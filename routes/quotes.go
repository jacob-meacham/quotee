package routes

import (
    "net/http"
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

func GetTest(r *http.Request, parms martini.Params) (int, string) {
    return http.StatusOK, `{"success":{"total":1},"contents":{"id":"UA6fPz652xt_AJzwP1ULiweF","quote":"Women want men, careers, money, children, friends, luxury, comfort, independence, freedom, respect, love, and a three-dollar pantyhose that won't run.","author":"Phyllis Diller","length":"150","tags":["funny","humor","women"],"category":"funny"}}`
}