package models

import (
    "fmt"
    "encoding/csv"
    "os"
    "math/rand"
    "errors"
    "net/http"
    "strings"
    "io/ioutil"

    "gopkg.in/jmcvetta/napping.v1"
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

func (source QuoteDBSource) String() string {
    return fmt.Sprintf("QuoteDBSource - Quotes from http://www.quotedb.com/quote/quote.php?action=random_quote, using %s categories.", source.Categories)
}

type TheySaidSoQuoteSource struct {
    Url string
    Categories []string
}

func (source TheySaidSoQuoteSource) String() string {
    return fmt.Sprintf("TheySaidSoQuoteSource - Quotes from %s, using %s categories.", source.Url, source.Categories)
}

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

type TheySaidSoQuote struct {
    Contents struct {
        Quote string `json:"quote"`
        Author string `json:"author"`
        Length string `json:"length"`
        Tags []string `json:"tags"`
        Category string `json:"category"`
    } `json:"contents"`
}

func (q TheySaidSoQuote) String() string {
    return fmt.Sprintf("%s - %s (%d)", q.Contents.Quote, q.Contents.Author, q.Contents.Length)
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

func (source FileQuoteSource) GetQuote() (Quote, error) {
    if len(source.quotes) == 0 {
        return Quote{}, errors.New("No quotes in this source")
    }

    return source.quotes[rand.Intn(len(source.quotes))], nil
}