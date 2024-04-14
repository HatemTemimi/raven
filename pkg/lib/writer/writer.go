package writer

import (
	"encoding/json"
	"fmt"
	"github.com/HatemTemimi/raven/pkg/lib/models"
	"github.com/HatemTemimi/raven/pkg/lib/utils"
	"os"
)

type Writer struct{}

func (w *Writer) WriteToJsonFile(proxies []models.Proxy, path string) error {

	formatted := utils.ProxiesToArray(proxies)

	proxyJson, err := json.Marshal(formatted)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, proxyJson, 0660)
	if err != nil {
		return err
	}

	return nil
}

func (w *Writer) WriteToTxtFile(proxies []models.Proxy, path string) error {

	formatted := utils.ProxiesToArray(proxies)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	for _, value := range formatted {
		_, err := fmt.Fprintln(f, value)
		if err != nil {
			return err
		}
	}

	return nil

}

func (w *Writer) WriteToStdout(proxies []models.Proxy) {
	formatted := utils.ProxiesToArray(proxies)
	for _, value := range formatted {
		fmt.Println(value)
	}
}
