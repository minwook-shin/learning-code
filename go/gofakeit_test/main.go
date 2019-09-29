package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit"
)

func main() {
	fmt.Println("Person : ", gofakeit.Person())
	fmt.Println("Name : ", gofakeit.Name())
	fmt.Println("NamePrefix : ", gofakeit.NamePrefix())
	fmt.Println("NameSuffix : ", gofakeit.NameSuffix())
	fmt.Println("FirstName : ", gofakeit.FirstName())
	fmt.Println("LastName : ", gofakeit.LastName())
	fmt.Println("Gender : ", gofakeit.Gender())

	fmt.Println("Contact : ", gofakeit.Contact())
	fmt.Println("Email : ", gofakeit.Email())

	fmt.Println("Phone : ", gofakeit.Phone())
	fmt.Println("PhoneFormatted : ", gofakeit.PhoneFormatted())

	fmt.Println("Username : ", gofakeit.Username())
	fmt.Println("Password : ", gofakeit.Password(true, true, true, true, false, 16))

	fmt.Println("Address : ", gofakeit.Address())
	fmt.Println("City : ", gofakeit.City())
	fmt.Println("Country : ", gofakeit.Country())
	fmt.Println("CountryAbr : ", gofakeit.CountryAbr())
	fmt.Println("State : ", gofakeit.State())
	fmt.Println("StateAbr : ", gofakeit.StateAbr())
	fmt.Println("StatusCode : ", gofakeit.StatusCode())
	fmt.Println("Street : ", gofakeit.Street())
	fmt.Println("StreetName : ", gofakeit.StreetName())
	fmt.Println("StreetNumber : ", gofakeit.StreetNumber())
	fmt.Println("StreetPrefix : ", gofakeit.StreetPrefix())
	fmt.Println("StreetSuffix : ", gofakeit.StreetSuffix())
	fmt.Println("Zip : ", gofakeit.Zip())

	fmt.Println("Latitude : ", gofakeit.Latitude())
	fmt.Println("Longitude : ", gofakeit.Longitude())

	fmt.Println("UUID : ", gofakeit.UUID())

	fmt.Println("Color : ", gofakeit.Color())
	fmt.Println("HexColor : ", gofakeit.HexColor())
	fmt.Println("RGBColor : ", gofakeit.RGBColor())
	fmt.Println("SafeColor : ", gofakeit.SafeColor())

	fmt.Println("URL : ", gofakeit.URL())
	fmt.Println("ImageURL : ", gofakeit.ImageURL(320, 320))
	fmt.Println("DomainName : ", gofakeit.DomainName())
	fmt.Println("DomainSuffix : ", gofakeit.DomainSuffix())
	fmt.Println("IPv4Address : ", gofakeit.IPv4Address())
	fmt.Println("IPv6Address : ", gofakeit.IPv6Address())

	fmt.Println("SimpleStatusCode : ", gofakeit.SimpleStatusCode())
	fmt.Println("HTTPMethod : ", gofakeit.HTTPMethod())
	fmt.Println("UserAgent : ", gofakeit.UserAgent())
	fmt.Println("ChromeUserAgent : ", gofakeit.ChromeUserAgent())
	fmt.Println("FirefoxUserAgent : ", gofakeit.FirefoxUserAgent())
	fmt.Println("SafariUserAgent : ", gofakeit.SafariUserAgent())

	fmt.Println("Date : ", gofakeit.Date())
	fmt.Println("NanoSecond : ", gofakeit.NanoSecond())
	fmt.Println("Second : ", gofakeit.Second())
	fmt.Println("Minute : ", gofakeit.Minute())
	fmt.Println("Hour : ", gofakeit.Hour())
	fmt.Println("Month : ", gofakeit.Month())
	fmt.Println("Day : ", gofakeit.Day())
	fmt.Println("WeekDay : ", gofakeit.WeekDay())
	fmt.Println("Year : ", gofakeit.Year())

	fmt.Println("TimeZone : ", gofakeit.TimeZone())
	fmt.Println("TimeZoneAbv : ", gofakeit.TimeZoneAbv())
	fmt.Println("TimeZoneFull : ", gofakeit.TimeZoneFull())
	fmt.Println("TimeZoneOffset : ", gofakeit.TimeZoneOffset())

	fmt.Println("Price : ", gofakeit.Price(1000, 10000))
	fmt.Println("Currency : ", gofakeit.Currency())
	fmt.Println("CurrencyLong : ", gofakeit.CurrencyLong())
	fmt.Println("CurrencyShort : ", gofakeit.CurrencyShort())

	fmt.Println("Language : ", gofakeit.Language())
	fmt.Println("LanguageAbbreviation : ", gofakeit.LanguageAbbreviation())
	fmt.Println("ProgrammingLanguage : ", gofakeit.ProgrammingLanguage())

	fmt.Println("Extension : ", gofakeit.Extension())
	fmt.Println("MimeType : ", gofakeit.MimeType())
}
