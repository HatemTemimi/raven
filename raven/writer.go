package raven

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
	defer f.Close()

	for _, value := range proxies {
		fmt.Fprintln(f, value)  // print values to f, one per line
	}

	return nil

}

func (w *Writer) WriteToStdout(proxies []string) {

	for _, value := range proxies {
		fmt.Println(value)  
	}

}
