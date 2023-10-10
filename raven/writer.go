package raven

import (
	"encoding/json"
	"fmt"
	"os"
)

type Writer struct {}

func (w *Writer)WriteToJsonFile(proxies []string, path string) error{
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

func (w *Writer)WriteToStdout(proxies []string) {
  for _, p := range proxies{
    fmt.Println(p)
  }
}
