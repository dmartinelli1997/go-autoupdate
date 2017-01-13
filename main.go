package autoupdate

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
)

func Update(executavel, endereço string) error {
	//executavel para ser atualizado
	//endereço de onde esta o novo executavel da atualização
	var (
		Exec string
	)
	//verificar qual sistema operacional está sendo usado para uso de estenção caso necessario
	if runtime.GOOS == `windows` {
		Exec = executavel + ".exe"

	} else if runtime.GOOS == "linux" {
		Exec = executavel
	}
	// renomeamos o antigo executavel para que possamos substituir o mesmo por outro com o sistema em produção
	os.Rename(Exec, Exec+"_old.exe")
	out, err := os.Create(Exec)
	if err != nil {
		return err
	}
	defer out.Close()

	// coletamos o novo arquivo
	resp, err := http.Get(endereço)
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//
	err = ioutil.WriteFile(Exec, responseData, 0644)
	if err != nil {
		return err
	}
	// os.Exit(0) fara com que o serviço feche, quando ja estiver como um serviço no windows ele reiniciará sozinho.
	os.Exit(0)
	return nil
}
