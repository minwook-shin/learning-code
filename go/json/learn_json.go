package main
 
import (
    "encoding/json"
    "fmt"
)
 
// test struct
type test struct {
    Title   string
    Text    string
    Index   int 
}

// test2 struct
type test2 struct {
    Title   string `json:"title"`
    Text    string `json:"txt"`
    Index   int `json:"idx"`
}
 
func main() {
    s := test{"hello", "hello world, go!", 0}
    b, _ := json.Marshal(s)
    str := string(b)
    fmt.Println(str)

    s2 := test2{"hello", "hello world, go!", 0}
    b2, _ := json.Marshal(s2)
    str2 := string(b2)
    fmt.Println(str2)

    b3, _ := json.Marshal(test{"title", "hello world! go!",0})
    var str3 test
    json.Unmarshal(b3, &str)
    fmt.Println(str3.Title, str3.Text, str3.Index)
}