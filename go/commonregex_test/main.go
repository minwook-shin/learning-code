package main

import (
	"fmt"

	"github.com/mingrammer/commonregex"
)

func main() {
	textString := `Jan 9th 2012 5:00PM 4:00 $9,000 test@gmail.com www.example.com 127.0.0.1 8A-99-D9-96-89-5B https://github.com/mingrammer/commonregex.git (519)-236-2723x341 #ffffff`

	date := commonregex.Date(textString)
	fmt.Println(date)

	time := commonregex.Time(textString)
	fmt.Println(time)

	price := commonregex.Prices(textString)
	fmt.Println(price)

	email := commonregex.Emails(textString)
	fmt.Println(email)

	link := commonregex.Links(textString)
	fmt.Println(link)

	git := commonregex.GitRepos(textString)
	fmt.Println(git)

	phone := commonregex.PhonesWithExts(textString)
	fmt.Println(phone)

	ip := commonregex.IPs(textString)
	fmt.Println(ip)

	mac := commonregex.MACAddresses(textString)
	fmt.Println(mac)
}
