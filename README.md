# Raven: HTTP Proxy Lister & Tester
### In the Box:
<ul>
<li> Get over 5000 http public proxies from multiples sources </li>
<li> Test the http proxies against custom targets </li>
<li> Export the http proxies to JSON formatted files </li>
</ul>

## Add Raven to your project
### Requirements: Go, Git

Install the package
```console
    go get github.com/HatemTemimi/Raven
 ```

## Library Usage

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


## Command line Usage: Under construction

  
TBC
