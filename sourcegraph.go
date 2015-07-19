package main
import (
  "encoding/json"
  "flag"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
)
func main() {
  login := flag.String("login", "", "GitHub login of user")
  flag.Parse()
  if *login == "" {
    log.Fatal("must specify login")
  }
  log.Println("Looking up GitHub user: ", *login)
    resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s", *login))
    if err != nil {
      log.Fatal(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatal(err)
    }
    var userdata map[string]interface{}
    err = json.Unmarshal(body, &userdata)
    if err != nil {
      log.Fatal(err)
    }
    log.Println("User's full name: ", userdata["name"])
    log.Println("User's company: ", userdata["company"])
}
