package fixerio

import (
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
	"bytes"
	"errors"
)

type request struct {
	base     string
	protocol string
	date     string
	symbols  []string
}

type response struct {
	Base  string `json:"base"`
	Date  string `json:"date"`
	Rates rates `json:"rates"`
}

type rates map[string]float32

var baseUrl = "api.fixer.io"

// Initializes fixerio.
func New() *request {
	return &request{
		base:     EUR,
		protocol: "https",
		date:     "",
		symbols:  make([]string, 0),
	}
}

// Set base currency.
func (f *request) Base(currency string) {
	f.base = currency
}

// Make the connection secure or not by setting the
// secure argument to true or false.
func (f *request) Secure(secure bool) {
	if secure {
		f.protocol = "https"
	} else {
		f.protocol = "http"
	}
}

// List of currencies that should be returned.
func (f *request) Symbols(currencies ...string) {
	f.symbols = currencies
}

// Specify a date in the past to retrieve historical records.
func (f *request) Historical(date time.Time) {
	f.date = date.Format("2006-01-02")
}

// Retrieve the exchange rates.
func (f *request) GetRates() (rates, error) {
	url := f.buildUrl()
	body, err := f.makeRequest(url)

	if err != nil {
		return rates{}, err
	}

	response, err := f.parseJson(body)

	if err != nil {
		return rates{}, err
	}

	return response, nil
}

func (f *request) buildUrl() string {
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

func (f *request) makeRequest(url string) (string, error) {
	response, err := http.Get(url)

	if err != nil {
		return "", errors.New("Couldn't connect to server")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", errors.New("Couldn't parse response")
	}

	return string(body), nil
}

func (f *request) parseJson(body string) (rates, error) {
	var response response

	err := json.Unmarshal([]byte(body), &response)

	if err != nil {
		return rates{}, errors.New("Couldn't parse response")
	}

	return response.Rates, nil
}
