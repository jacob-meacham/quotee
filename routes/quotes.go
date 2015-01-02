package routes

import (
    //"fmt"
    "net/http"
    //"strconv"

    "github.com/go-martini/martini"
    "github.com/jacob-meacham/quotee/models"
)

var quoteSources []models.QuoteSourceEntry

func SetQuoteSources(sources []models.QuoteSourceEntry) {
    quoteSources = sources
}

func GetQuote(r *http.Request, enc Encoder, parms martini.Params) (int, string) {
    source := parms["source"]

    // If there isn't any source specified, just use a random source

    return http.StatusOK, source
}