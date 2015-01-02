package routes

import (
    //"fmt"
    "net/http"
    //"strconv"

    //"github.com/go-martini/martini"
    "github.com/jacob-meacham/quotee/models"
)

var quoteSources *map[string]models.QuoteSource

func SetQuoteSources(sources *map[string]models.QuoteSource) {
    quoteSources = sources
}

func GetQuote(r *http.Request, enc Encoder) string {
    return "foo"
}