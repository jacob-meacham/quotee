package models_test

import (
	. "github.com/jacob-meacham/quotee/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

    "io"
    "net/http"
    "net/http/httptest"
)

func ServerMock(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, `{"success":{"total":1},"contents":{"id":"UA6fPz652xt_AJzwP1ULiweF","quote":"Behind every successful man is a woman, behind her is his wife.","author":"Groucho Marx","length":"150","tags":["funny","humor","women"],"category":"funny"}}`)
}

var _ = Describe("Models", func() {
    Describe("TheySaidSoSource", func() {
        It("has a list of categories", func() {
            categories := []string{"funny", "life", "inspire", "love"}
            source := TheySaidSoQuoteSource{Url: "http://api.theysaidso.com/qod.json", Categories: categories}

            Expect(source.Categories).To(Equal(categories))
        })
        
        Context("When the URL is well-formed", func() {
            It("returns a quote", func() {
                testServer := httptest.NewServer(http.HandlerFunc(ServerMock))
                defer testServer.Close()

                source := TheySaidSoQuoteSource{Url: testServer.URL + "/api/quote/theysaidso/static", Categories: []string{"funny", "life", "inspire", "love"}}
                Expect(source.String()).To(Equal("TheySaidSoQuoteSource - Quotes from " + testServer.URL + "/api/quote/theysaidso/static, using [funny life inspire love] categories."))
                quote, err := source.GetQuote()
                Expect(err).ToNot(HaveOccurred())
                Expect(quote.Body).To(Equal("Behind every successful man is a woman, behind her is his wife."))
                Expect(quote.Author).To(Equal("Groucho Marx"))
            })
        })

        Context("When the URL is not correct", func() {
            It("returns an error", func() {
                source := TheySaidSoQuoteSource{Url: "http://notreal.com", Categories: []string{"funny", "life", "inspire", "love"}}
                _, err := source.GetQuote()
                Expect(err).To(HaveOccurred())
            })
        })
    })

    Describe("QuoteDB", func() {
        Context("When the URL is well-formed", func() {
            It("returns a quote", func() {
                source := QuoteDBSource{Url: "http://www.quotedb.com/quote/quote.php?action=random_quote", Categories: []string{"funny", "life", "inspire", "love"}}
                Expect(source.String()).To(Equal("QuoteDBSource - Quotes from http://www.quotedb.com/quote/quote.php?action=random_quote, using [funny life inspire love] categories."))
                quote, err := source.GetQuote()
                Expect(err).ToNot(HaveOccurred())
                Expect(quote.Body).ToNot(BeNil())
                Expect(quote.Author).ToNot(BeNil())
            })
        })

        Context("When the URL is not correct", func() {
            It("returns an error", func() {
                source := QuoteDBSource{Url: "http://------.com/notreal", Categories: []string{"funny", "life", "inspire", "love"}}
                _, err := source.GetQuote()
                Expect(err).To(HaveOccurred())
            })
        })
    })

    Describe("FileSource", func() {
        Context("When the CSV is well-formed", func() {
            It("has a list of quotes", func() {
                source, err := CreateFileQuoteSource("./test/good_multi.csv")
                Expect(err).ToNot(HaveOccurred())
                Expect(len(source.Quotes)).To(Equal(2))
                Expect(source.Quotes[1].Body).To(Equal("Im a second test"))
                Expect(source.String()).To(Equal("FileQuoteSource - 2 quotes loaded from ./test/good_multi.csv"))
            })

            It("returns a quote from the list", func() {
                source, err := CreateFileQuoteSource("./test/good_single.csv")
                Expect(err).ToNot(HaveOccurred())

                quote, err := source.GetQuote()
                Expect(err).ToNot(HaveOccurred())
                Expect(quote.Body).To(Equal("I'm a test"))
                Expect(quote.Author).To(Equal("Tester"))
            })
        })

        Context("When the CSV can't be found", func() {
            It("returns an error", func() {
                _, err := CreateFileQuoteSource("notexistent")
                Expect(err).To(HaveOccurred())
            })
        })

        Context("When the CSV is malformed", func() {
            It("returns an error", func() {
                _, err := CreateFileQuoteSource("./test/bad.csv")
                Expect(err).To(HaveOccurred())
            })
        })

        Context("When the CSV is empty", func() {
            It("returns an error", func() {
                source := FileQuoteSource{}

                _, err := source.GetQuote()
                Expect(err).To(HaveOccurred())
            })
        })
    })
})
