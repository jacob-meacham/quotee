package models

import (
    "fmt"
    "encoding/csv"
    "os"
    "math/rand"
    "errors"
)

type FileQuoteSource struct {
    filename string
    quotes []Quote
}

func (source FileQuoteSource) String() string {
    return fmt.Sprintf("FileQuoteSource - %d quotes loaded from %s", len(source.quotes), source.filename)
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

    source.filename = filename
    for _, record := range records {
        source.quotes = append(source.quotes, Quote{record[0],record[1]})
    }

    return
}

func (source FileQuoteSource) GetQuote() (Quote, error) {
    if len(source.quotes) == 0 {
        return Quote{}, errors.New("No quotes in this source")
    }

    return source.quotes[rand.Intn(len(source.quotes))], nil
}
