package reader

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/HatemTemimi/raven/pkg/lib/models"
	"github.com/HatemTemimi/raven/pkg/lib/utils"
	"log"
	"os"
)

//this class reads proxies from txt or json file
// and returns an array of proxies on success

type Reader struct{}

func (r *Reader) ReadTxtfile(path string) ([]models.Proxy, error) {

	var proxies []models.Proxy

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		proxy, err := utils.ParseProxyFromAddress(fscanner.Text())
		if err != nil {
			proxies = append(proxies, *proxy)
		}
	}

	return proxies, nil
}

func (r *Reader) ReadTxtfileToStdOut(path string) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		fmt.Println(fscanner.Text())
	}
	return nil
}

func (r *Reader) ReadJsonFile(path string) ([]models.Proxy, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var Decoder = json.NewDecoder(file)
	if err != nil {
		return nil, err
	}

	var decoded []string
	var proxies []models.Proxy

	err = Decoder.Decode(&decoded)
	if err != nil {
		return nil, err
	}

	for _, proxy := range decoded {
		parsed, err := utils.ParseProxyFromAddress(proxy)
		if err != nil {
			proxies = append(proxies, *parsed)
		}
	}

	return proxies, nil
}

func (r *Reader) ReadJsonFileToStdOut(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var Decoder = json.NewDecoder(file)
	if err != nil {
		log.Fatal(err)
	}

	var proxies []string

	err = Decoder.Decode(&proxies)
	if err != nil {
		log.Fatal(err)
	}

	for _, proxy := range proxies {
		log.Println(proxy)
	}

	return nil
}
