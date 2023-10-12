package raven

import (
	"bufio"
	"os"
)


type Reader struct {}

func (r *Reader) ReadTxtfile(path string) ([]string, error){

	var proxies []string

	file, err := os.Open(path)
	if err != nil {
		return nil,err
	}

	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		proxies = append(proxies, fscanner.Text())
	}
	return proxies, nil
}




