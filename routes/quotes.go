package routes

import (
    //"fmt"
    "net/http"
    //"strconv"
    "math/rand"

    "github.com/go-martini/martini"
    "github.com/jacob-meacham/quotee/models"
)

var quoteSources map[string]models.QuoteSource
var quoteKeys []string

func SetQuoteSources(sources map[string]models.QuoteSource) {
    quoteSources = sources
    for k := range quoteSources {
        quoteKeys = append(quoteKeys, k)
    }
}

func GetQuote(r *http.Request, enc Encoder, parms martini.Params) (int, string) {
    source := parms["source"]

    quoteSource, ok := quoteSources[source]
    if ok == false {
        // If there isn't any source specified, just use a random source
        quoteSource = quoteSources[quoteKeys[rand.Intn(len(quoteKeys))]]
    }

    quote, err := quoteSource.GetQuote()
    if err != nil {
        panic(err)
    }
    return http.StatusOK, Must(enc.Encode(quote))
}