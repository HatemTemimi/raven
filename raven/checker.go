package raven

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

type Checker struct {
	client *http.Client
}

func (c *Checker) CheckAgainstTarget(proxy string, target string, data *[]string, wg *sync.WaitGroup) {
	defer wg.Done()
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		log.Fatal(err)
	}

	transport := &http.Transport{
		Proxy:                 http.ProxyURL(proxyURL),
		ResponseHeaderTimeout: 30 * time.Second,
	}

	c.client.Transport = transport

	httpsPoke, _ := http.NewRequest("GET", "https://"+target, nil)

	httpsPoke.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36")

	httpsResp, httpsErr := c.client.Do(httpsPoke)

	if httpsErr != nil {
		log.Println(proxy, "is Down against: ", target)
	} else {
		if httpsResp != nil {
			log.Println(proxy, " is Up & Fresh against: ", target)
			*data = append(*data, proxy)
			proxyJson, _ := json.Marshal(*data)
			os.WriteFile("proxies.json", proxyJson, 0660)
		}
	}
}

func (c *Checker) Check(proxies []string, targets []string) []string {
	var data []string
	var wg sync.WaitGroup
	for _, t := range targets {
		for _, e := range proxies {
			wg.Add(1)
			go c.CheckAgainstTarget("http://"+e, t, &data, &wg)
		}
	}
	wg.Wait()
	return data

}
