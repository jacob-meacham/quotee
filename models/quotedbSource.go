package models

import (
    "fmt"
    "net/http"
    "strings"
    "io/ioutil"
)

// Quotes from 
type QuoteDBSource struct {
    Categories []string
}

func (source QuoteDBSource) String() string {
    return fmt.Sprintf("QuoteDBSource - Quotes from http://www.quotedb.com/quote/quote.php?action=random_quote, using %s categories.", source.Categories)
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
