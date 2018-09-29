package main
 
import (
    "fmt"
    "io/ioutil"
    "net/http"
)
 
func main() {
    page, _ := http.Get("http://info.cern.ch/")
    data, _ := ioutil.ReadAll(page.Body)
	fmt.Println(string(data))
	page.Body.Close()

    req, _ := http.NewRequest("GET", "http://info.cern.ch/", nil)
    client := &http.Client{}
    resp, _ := client.Do(req)
    bytes, _ := ioutil.ReadAll(resp.Body)
    str := string(bytes)
	fmt.Println(str)
	resp.Body.Close()
}