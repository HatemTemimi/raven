package lib

import (
	"github.com/HatemTemimi/raven/pkg/lib/models"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

type Checker struct {
	Client *http.Client
}

func (c *Checker) CheckAgainstTarget(proxy models.Proxy, target string, data *[]models.Proxy, wg *sync.WaitGroup) {
	defer wg.Done()
	proxyURL, err := url.Parse("http://" + proxy.Ip + ":" + strconv.FormatInt(proxy.Port, 10))

	if err != nil {
		log.Fatal(err)
	}

	transport := &http.Transport{
		Proxy:                 http.ProxyURL(proxyURL),
		ResponseHeaderTimeout: 5 * time.Second,
	}

	c.Client.Transport = transport

	httpsPoke, _ := http.NewRequest("GET", "https://"+target, nil)

	httpsPoke.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36")

	httpsResp, httpsErr := c.Client.Do(httpsPoke)

	if httpsErr != nil {
		proxy.Status = "down"
		log.Println(proxyURL, "is Down against target: ", target)
	} else {
		if httpsResp != nil {
			proxy.Status = "up"
			log.Println(proxy, " is Up & Fresh against: ", target)
			*data = append(*data, proxy)
		}
	}
}

func (c *Checker) Check(proxies []models.Proxy, targets []string) []models.Proxy {
	var data []models.Proxy
	var wg sync.WaitGroup
	for _, t := range targets {
		for _, proxy := range proxies {
			wg.Add(1)
			go c.CheckAgainstTarget(proxy, t, &data, &wg)
		}
	}
	wg.Wait()
	return data
}
