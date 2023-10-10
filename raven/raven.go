package raven
import "net/http"


type Raven struct {
    Scanner Scanner
    Checker Checker
    Writer Writer
}

func (r *Raven) Init(){
  client := http.Client{}
  r.Scanner.client = &client
  r.Checker.client = &client
}

func (r *Raven) FetchAll() ([]string, error){
  proxies, err := r.Scanner.ScanDefaultSources();
  if (err !=nil){
    return nil, err
  }
  return proxies, nil
}


func (r *Raven) FetchValid(target string) ([]string, error){
  proxies, err := r.Scanner.ScanDefaultSources();
  if (err !=nil){
    return nil, err
  }
  workingAgainst := r.Checker.Check(proxies, []string{target})
  return workingAgainst, nil
}


func (r *Raven) FetchAllToStdOut() (error){
  proxies, err := r.Scanner.ScanDefaultSources();
  if (err !=nil){
    return  err
  }
  r.Writer.WriteToStdout(proxies)
  return nil
}


func (r *Raven) FetchValidToStdOut(target string) (error){
  proxies, err := r.Scanner.ScanDefaultSources();
  if (err !=nil){
    return  err
  }
  workingAgainst := r.Checker.Check(proxies, []string{target})
  r.Writer.WriteToStdout(workingAgainst)
  return nil
}


func (r *Raven) FetchAllToFile() (error){
  proxies, err := r.Scanner.ScanDefaultSources();
  if (err !=nil){
    return err
  }
  err = r.Writer.WriteToJsonFile(proxies)
  if err != nil {
    return err
  }
  return nil
}


func (r *Raven) FetchValidToFile(target string) (error){
  proxies, err := r.Scanner.ScanDefaultSources();
  if (err !=nil){
    return err
  }
  workingAgainst := r.Checker.Check(proxies, []string{target})
  err = r.Writer.WriteToJsonFile(workingAgainst)
  if err != nil {
    return err
  }
  return nil
}
