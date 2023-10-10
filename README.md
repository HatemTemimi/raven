# Raven: HTTP Proxy Lister & Tester
### In the Box:
<ul>
<li> Get over 5000 http public proxies from multiples sources </li>
<li> Test the http proxies against custom targets </li>
<li> Export the http proxies to JSON formatted files </li>
<li> Proxies in the format <strong>IP:PORT</strong> </li>

</ul>


## Proxies in one command with docker
You can run the Raven as a container and it will print out the result proxies for you, it also comes with flags support to filter out the results.

### Setup and First Run
Pull and run the docker image
```bash
    docker run --rm  hatemtemimi94/raven
```
This is sample of what you should expect to receive:
```json
[
	160.3.168.70:8080
	81.103.105.130:8888
	146.59.14.159:80
	45.70.236.194:999
	103.155.54.38:83
	203.57.51.53:80
	47.254.198.237:3128
	54.92.199.26:80
	144.76.75.25:4444
	186.155.230.114:999
	190.61.32.168:6969
	94.198.40.18:80
	190.61.48.24:999
	181.176.221.151:9812
	193.151.130.114:8086
	95.56.254.139:3128
	213.6.170.17:80
	69.94.136.71:8443
	178.128.157.114:443
	41.128.148.78:1976
	45.162.135.201:999
	64.227.106.157:80
	...
]
```
### Flags
-fetch: fetching method, set to "all" for all proxies, "valid" for validation against target
-t [ target ]: target url against which the tests will be performed, **mandatory when using -fetch valid**
-o [output]: path to the file in which to write the results
-h [ help ]: prints help about the commands

#### Fetching untested public proxies
the fetch flag accepts these values: all | valid
```bash
	  docker run --rm  hatemtemimi94/raven -fetch all 
	  #will return all public proxies
```
#### Fetching tested public proxies against custom target url
with the fetch flag set to "valid", we have to specify the target url to test on, using the -t flag
```bash
	  docker run --rm  hatemtemimi94/raven -fetch valid -t www.example.com 
	  #will return all valid public proxies tested against the target url
	  #replace it with your value for example www.google.com
```
#### Exporting
Exporting to file is done via the flag -o
```bash
	  docker run --rm  hatemtemimi94/raven -fetch valid -t www.example.com -o proxies.json
	  #will return all valid public proxies tested against the target url
```

## Library usage for developement
## Add Raven to your project
### Requirements: Go, Git

Install the package
```console
    go get github.com/HatemTemimi/Raven
 ```

Create  *Raven* instance and call the *Init* function to set it up
```go
 import (
	"github.com/HatemTemimi/Raven/raven"
)

  raven  := raven.Raven{}
  raven.Init() //!!Mandatory to intialize http client
```
Now that the instance is setup and ready you can call various fetch functions:
### FetchAll()

 - returns a string array of all the proxies without testing any
 - returns an error in case it could not fetch the proxies

```go
	proxies, err := raven.FetchAll()
	log.Println(proxies)
```
### FetchValid(target string)
the target is the url you want to test against, for example: *www.scrapeme.live*
**use this format for the target url, no need to add the protocol to the url**
 - returns a string array of all the successfully tested proxies against given target
 - returns an error in case it could not fetch the proxies

```go
	proxies, err := raven.FetchValid("www.scrapeme.live")
	if err != nil {
		log.Println(err)
	}
	log.Println(proxies)
```
### FetchAllToStdOut()

 - prints all the proxies to stdout
 - returns an error in case it could not fetch the proxies

```go
	err  := raven.FetchAllToStdOut()
	if err !=  nil {
		log.Println(err)
	}
```
### FetchValidToStdOut(target string)

 - prints all the successfully tested proxies to stdout
 - returns an error in case it could not fetch the proxies

```go
	err  := raven.FetchValidToStdOut()
	if err !=  nil {
		log.Println(err)
	}
```
### FetchAllToJsonFile(filePath string) (error)

 -  writes to a json file all the proxies 
 - returns an error in case it could not fetch the proxies

```go
	err  := raven.FetchAllToJsonFile("./proxies.json")
	if err !=  nil {
		log.Println(err)
	}
```
### FetchValidToJsonFile(filePath string) (error)

 - writes to a json file all the successfully tested proxies 
 - returns an error in case it could not fetch the proxies

```go
	err  := raven.FetchValidToJsonFile("./proxies.json")
	if err !=  nil {
		log.Println(err)
	}
```

#### Notes
This is a project that i made in my freetime because i did not find a consistent package made to test proxies against a custom target, any PRs are welcome to improve it and I will be happy to improve it.
For now Raven:v1.1.0 only supports HTTP / HTTPS proxies, support for SOCKS might be considered in the future depending on how the package evolves, if it ever does.