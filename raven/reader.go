package raven

import (
	"bufio"
	"encoding/json"
	"log"
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

func (r *Reader) ReadJsonFile(path string)([]string, error){
	// open the file pointer
	file, err := os.Open(path)
	if err != nil {
		return nil,err
	}
	defer file.Close()

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
	return proxies, nil

}




