package reader

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Reader struct{}

func (r *Reader) ReadTxtfile(path string) ([]string, error) {

	var proxies []string

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		proxies = append(proxies, fscanner.Text())
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

func (r *Reader) ReadJsonFile(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var Decoder *json.Decoder = json.NewDecoder(file)
	if err != nil {
		log.Fatal(err)
	}

	var proxies []string

	err = Decoder.Decode(&proxies)
	if err != nil {
		log.Fatal(err)
	}

	for _, proxy := range proxies {
		proxies = append(proxies, proxy)
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

	var Decoder *json.Decoder = json.NewDecoder(file)
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
