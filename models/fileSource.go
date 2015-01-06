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
    Quotes []Quote
}

func (source FileQuoteSource) String() string {
    return fmt.Sprintf("FileQuoteSource - %d quotes loaded from %s", len(source.Quotes), source.filename)
}

func CreateFileQuoteSource(filename string) (source FileQuoteSource, err error) {
    f, err := os.Open(filename)
    if err != nil {
        return
    }
    defer f.Close()

    reader := csv.NewReader(f)
    reader.Comma = '|'
    reader.TrimLeadingSpace = true
    reader.FieldsPerRecord = 2
    records, err := reader.ReadAll()
    if (err != nil) {
        return
    }

    source.filename = filename
    for _, record := range records {
        source.Quotes = append(source.Quotes, Quote{record[0],record[1]})
    }

    return
}

func (source FileQuoteSource) GetQuote() (Quote, error) {
    if len(source.Quotes) == 0 {
        return Quote{}, errors.New("No quotes in this source")
    }

    return source.Quotes[rand.Intn(len(source.Quotes))], nil
}
