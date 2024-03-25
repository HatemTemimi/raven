package writer

import (
	"encoding/json"
	"fmt"
	"os"
)

type Writer struct{}

func (w *Writer) WriteToJsonFile(proxies []string, path string) error {

	proxyJson, err := json.Marshal(proxies)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, proxyJson, 0660)
	if err != nil {
		return err
	}

	return nil
}

func (w *Writer) WriteToTxtFile(proxies []string, path string) error {

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	for _, value := range proxies {
		_, err := fmt.Fprintln(f, value)
		if err != nil {
			return err
		}
	}

	return nil

}

func (w *Writer) WriteToStdout(proxies []string) {

	for _, value := range proxies {
		fmt.Println(value)
	}

}
