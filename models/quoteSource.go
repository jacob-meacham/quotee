package models

type Quote struct {
    Body string
    Author string
}

type QuoteSource interface {
    GetQuote() (Quote, error)
}