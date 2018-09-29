package main
 
import (
    "log"
    "os"
)

func test() {
    log.Print("Test")
}
 
var logging *log.Logger
 
func main() {
	test()

    logging = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	logging.Println("include info")
	

	logging.SetFlags(0)
	logging.Println("remove time")

	logging.SetPrefix("prefix: ")
	logging.Println("remove time and set prefix")
}
 
