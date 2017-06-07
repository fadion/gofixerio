# Thin wrapper for Fixer.io

A thin wrapper in Go for [Fixer.io](http://www.fixer.io), a service for foreign exchange rates and currency conversion. It provides a few methods to easily construct the url, makes the api call and gives back the response.

## Installation

As for any other package, you can use `go get`:

```
$ go get github.com/fadion/gofixerio
```

For better package management however, I'd recommend [glide](https://github.com/Masterminds/glide) or the still in alpha [dep](https://github.com/golang/dep).

## Usage

First, let's import the package:

```go
import "github.com/fadion/gofixerio"
```

Let's see an exhaustive example with all the parameters:

```go
exchange := gofixerio.New()

exchange.Base(gofixerio.EUR)
exchange.Symbols(gofixerio.USD, gofixerio.AUD)
exchange.Secure(true)
exchange.Historical(time.Date(2016, time.December, 15, 0, 0, 0, 0, time.UTC))

if rates, err := exchange.GetRates(); err == nil {
    fmt.Println(rates[gofixerio.USD])
}
```

Every parameter can be omitted as the package provides some sensible defaults. The base currency is `EUR`, makes a secure connection by default and returns all the supported currencies.

## Response

The response is a simple `map[string]float32` with currencies as keys and ratios as values. For a request like the following:

```go
exchange := gofixerio.New()
exchange.Symbols(gofixerio.USD, gofixerio.GBP)

rates, _ := exchange.GetRates()
fmt.Println(rates)
```

the response will be the map:

```go
map[USD:1.1188 GBP:0.87093]
```

which you can access with the keys as strings or using the currency constants:

```go
rates["USD"];
rates[gofixerio.GBP];
```
