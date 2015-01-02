package models

import (
    "fmt"
    "math/rand"
    "errors"

    "gopkg.in/jmcvetta/napping.v1"
)

type TheySaidSoQuoteSource struct {
    Url string
    Categories []string
}

func (source TheySaidSoQuoteSource) String() string {
    return fmt.Sprintf("TheySaidSoQuoteSource - Quotes from %s, using %s categories.", source.Url, source.Categories)
}

type TheySaidSoQuote struct {
    Contents struct {
        Quote string `json:"quote"`
        Author string `json:"author"`
        Length string `json:"length"`
        Tags []string `json:"tags"`
        Category string `json:"category"`
    } `json:"contents"`
}

func (source TheySaidSoQuoteSource) GetQuote() (Quote, error) {
    params := napping.Params {
    "category": source.Categories[rand.Intn(len(source.Categories))],
    }
    result := TheySaidSoQuote{}
    resp, err := napping.Get(source.Url, &params, &result, nil)
    if err != nil {
        return Quote{}, err
    }
    if resp.Status() != 200 {
        return Quote{}, errors.New("Received failed response from theysaidso") 
    }

    return Quote{result.Contents.Quote, result.Contents.Author}, nil
}