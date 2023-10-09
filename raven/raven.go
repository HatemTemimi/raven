package raven
import "net/http"


type Raven struct {
    scan Scanner
    check Checker
    writer Writer
}

func (r *Raven) Init(){
  client := http.Client{}
  r.scan.client = &client
  r.check.client = &client
}

func (r *Raven) FetchAll() ([]string, error){
  proxies, err := r.scan.ScanDefaultSources();
  if (err !=nil){
    return nil, err
  }
  return proxies, nil
}


func (r *Raven) FetchValid(target string) ([]string, error){
  proxies, err := r.scan.ScanDefaultSources();
  if (err !=nil){
    return nil, err
  }
  workingAgainst := r.check.Check(proxies, []string{target})
  return workingAgainst, nil
}


func (r *Raven) FetchAllToFile(target string) (error){
  proxies, err := r.scan.ScanDefaultSources();
  if (err !=nil){
    return err
  }
  err = r.writer.WriteToJsonFile(proxies)
  if err != nil {
    return err
  }
  return nil
}


func (r *Raven) FetchValidToFile(target string) (error){
  proxies, err := r.scan.ScanDefaultSources();
  if (err !=nil){
    return err
  }
  workingAgainst := r.check.Check(proxies, []string{target})
  err = r.writer.WriteToJsonFile(workingAgainst)
  if err != nil {
    return err
  }
  return nil
}
