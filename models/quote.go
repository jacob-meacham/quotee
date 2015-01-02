package models

import (
    "encoding/csv"
    "os"
    "math/rand"
    "errors"
    "log"
)

type Quote struct {
    Body string
    Author string
}

type QuoteSource interface {
    GetQuote() (Quote, error)
}

// Quotes from 
type QuoteDBSource struct {
    Categories []string
}

type TheySaidSoQuoteSource struct {
    Categories []string
}

type FileQuoteSource struct {
    quotes []Quote
}

func CreateFileQuoteSource(filename string) (source FileQuoteSource, err error) {
    f, err := os.Open(filename)
    if err != nil {
        return
    }
    defer f.Close()

    reader := csv.NewReader(f)
    records, err := reader.ReadAll()
    if (err != nil) {
        return
    }

    for _, record := range records {
        source.quotes = append(source.quotes, Quote{record[0],record[1]})
    }
    log.Printf("Source from %s contains %d quotes", filename, len(source.quotes))

    return
}

func (source QuoteDBSource) GetQuote() (Quote, error) {
    return Quote{"QuoteDBSource", "bar"}, nil
}

func (source TheySaidSoQuoteSource) GetQuote() (Quote, error) {
    return Quote{"TheySaidSoQuoteSource", "bar"}, nil
}

func (source FileQuoteSource) GetQuote() (Quote, error) {
    if len(source.quotes) == 0 {
        return Quote{}, errors.New("No quotes in this source")
    }

    return source.quotes[rand.Intn(len(source.quotes))], nil
}