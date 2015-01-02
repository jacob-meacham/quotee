package models

type Quote struct {
    body string
    author string
}

type QuoteSource interface {
    GetQuote() Quote
}

// Quotes from 
type QuoteDBSource struct {
    categories []string
}

type TheySaidSoQuoteSource struct {
    categories []string
}

type FileQuoteSource struct {

}

func (source *QuoteDBSource) GetQuote() Quote {
    return Quote{"QuoteDBSource", "bar"}
}

func (source *TheySaidSoQuoteSource) GetQuote() Quote {
    return Quote{"TheySaidSoQuoteSource", "bar"}
}

func (source *FileQuoteSource) GetQuote() Quote {
    return Quote{"FileQuoteSource", "bar"}
}