/*
	Package Fixerio provides a simple interface to the
	fixer.io API, a service for currency exchange rates.
 */
package fixerio

import (
	"strings"
	"net/http"
	"encoding/json"
	"time"
	"bytes"
	"errors"
)

// Holds the request parameters.
type Request struct {
	base     string
	protocol string
	date     string
	symbols  []string
}

// JSON response object.
type Response struct {
	Base  string `json:"base"`
	Date  string `json:"date"`
	Rates rates `json:"rates"`
}

type rates map[string]float32

var baseUrl = "api.fixer.io"

// Initializes fixerio.
func New() *Request {
	return &Request{
		base:     EUR,
		protocol: "https",
		date:     "",
		symbols:  make([]string, 0),
	}
}

// Sets base currency.
func (f *Request) Base(currency string) {
	f.base = currency
}

// Make the connection secure or not by setting the
// secure argument to true or false.
func (f *Request) Secure(secure bool) {
	if secure {
		f.protocol = "https"
	} else {
		f.protocol = "http"
	}
}

// List of currencies that should be returned.
func (f *Request) Symbols(currencies ...string) {
	f.symbols = currencies
}

// Specify a date in the past to retrieve historical records.
func (f *Request) Historical(date time.Time) {
	f.date = date.Format("2006-01-02")
}

// Retrieve the exchange rates.
func (f *Request) GetRates() (rates, error) {
	url := f.GetUrl()
	response, err := f.makeRequest(url)

	if err != nil {
		return rates{}, err
	}

	return response, nil
}

// Formats the URL correctly for the API Request.
func (f *Request) GetUrl() string {
	var url bytes.Buffer

	url.WriteString(f.protocol)
	url.WriteString("://")
	url.WriteString(baseUrl)
	url.WriteString("/")

	if f.date == "" {
		url.WriteString("latest")
	} else {
		url.WriteString(f.date)
	}

	url.WriteString("?base=")
	url.WriteString(string(f.base))

	if len(f.symbols) >= 1 {
		url.WriteString("&symbols=")
		url.WriteString(strings.Join(f.symbols, ","))
	}

	return url.String()
}

func (f *Request) makeRequest(url string) (rates, error) {
	var response Response
	body, err := http.Get(url)

	if err != nil {
		return rates{}, errors.New("Couldn't connect to server")
	}

	defer body.Body.Close()

	err = json.NewDecoder(body.Body).Decode(&response)

	if err != nil {
		return rates{}, errors.New("Couldn't parse Response")
	}

	return response.Rates, nil
}