package fixerio

import (
	"testing"
	"time"
)

func TestDefaultParameters(t *testing.T) {
	expected := "https://" + baseUrl + "/latest?base=EUR"
	actual := New().GetUrl()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestUnsecure(t *testing.T) {
	expected := "http://" + baseUrl + "/latest?base=EUR"

	f := New()
	f.Secure(false)
	actual := f.GetUrl()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestBase(t *testing.T) {
	expected := "https://" + baseUrl + "/latest?base=USD"

	f := New()
	f.Base(USD)
	actual := f.GetUrl()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestSymbols(t *testing.T) {
	expected := "https://" + baseUrl + "/latest?base=GBP&symbols=EUR,USD,AUD"

	f := New()
	f.Base(GBP)
	f.Symbols(EUR, USD, AUD)
	actual := f.GetUrl()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestSingleSymbol(t *testing.T) {
	expected := "https://" + baseUrl + "/latest?base=GBP&symbols=EUR"

	f := New()
	f.Base(GBP)
	f.Symbols(EUR)
	actual := f.GetUrl()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestEmptySymbols(t *testing.T) {
	expected := "https://" + baseUrl + "/latest?base=EUR"

	f := New()
	f.Symbols()
	actual := f.GetUrl()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestHistorical(t *testing.T) {
	expected := "https://" + baseUrl + "/2016-06-09?base=EUR"

	f := New()
	f.Historical(time.Date(2016, time.June, 9, 0, 0, 0, 0, time.UTC))
	actual := f.GetUrl()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestAllParameters(t *testing.T) {
	expected := "http://" + baseUrl + "/latest?base=USD&symbols=EUR,GBP"

	f := New()
	f.Base(USD)
	f.Symbols(EUR, GBP)
	f.Secure(false)
	actual := f.GetUrl()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}