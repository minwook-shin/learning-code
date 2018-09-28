package main
 
import (
    "encoding/xml"
    "fmt"
)
 
type test struct {
    Title   string
    Text    string
    Index	int
}
 
func main() {
    t := test{"hello,world!", "text line", 0}
    b, _ := xml.Marshal(t)
    s := string(b)
	fmt.Println(s)
	
    d, _ := xml.Marshal(test{"hello,world!", "text line", 0})
    var t2 test
    xml.Unmarshal(d, &t2)
    fmt.Println(t2.Title,t2.Text, t2.Index)
}