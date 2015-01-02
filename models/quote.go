package models

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

}

func (source QuoteDBSource) GetQuote() (Quote, error) {
    return Quote{"QuoteDBSource", "bar"}, nil
}

func (source TheySaidSoQuoteSource) GetQuote() (Quote, error) {
    return Quote{"TheySaidSoQuoteSource", "bar"}, nil
}

func (source FileQuoteSource) GetQuote() (Quote, error) {
    return Quote{"FileQuoteSource", "bar"}, nil
}