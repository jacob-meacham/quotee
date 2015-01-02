package models

type Quote struct {
    Body string `json:"body" xml:"body"`
    Author string `json:"author" xml:"author"`
}

type QuoteSource interface {
    GetQuote() (Quote, error)
}