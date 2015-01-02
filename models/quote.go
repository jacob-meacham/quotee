package models

import (
    "encoding/csv"
    "os"
    "math/rand"
    "errors"
    "log"
    "net/http"
    "strings"
    "io/ioutil"
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
    resp, err := http.Get("http://www.quotedb.com/quote/quote.php?action=random_quote")
    if err != nil {
        return Quote{}, err
    }

    defer resp.Body.Close()
    httpBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return Quote{}, err
    }
    quoteData := string(httpBody[:])

    // Parse the response to pull out the body and author
    quoteParts := strings.Split(quoteData, "document.write('")
    quoteBody := quoteParts[1]
    quoteBody = quoteBody[:strings.LastIndex(quoteBody, "<br>")]

    // TODO: Ugly
    // The author has the form of <i>More quotes from <a href="url">Author</a></i>
    // so we split on > and then strip. Lots of magic here
    quoteAuthorParts := strings.Split(quoteParts[2], ">")
    quoteAuthor := quoteAuthorParts[2]
    quoteAuthor = quoteAuthor[:strings.LastIndex(quoteAuthor, "</a")]

    return Quote{quoteBody, quoteAuthor}, nil
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