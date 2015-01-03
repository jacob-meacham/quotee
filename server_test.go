package main_test

import (
	. "github.com/jacob-meacham/quotee"
    . "github.com/jacob-meacham/quotee/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
    
    "github.com/go-martini/martini"
    "encoding/json"
    "encoding/xml"
    "net/http"
    "net/http/httptest"
)

func mapFromJSON(data []byte) map[string]interface{} {
    var result interface{}
    json.Unmarshal(data, &result)
    return result.(map[string]interface{})
}

func quoteFromXML(data []byte) Quote {
    var result Quote
    err := xml.Unmarshal(data, &result)
    Expect(err).ToNot(HaveOccurred())
    return result
}

var _ = Describe("Server", func() {
    var server *martini.Martini
    var request *http.Request
    var recorder *httptest.ResponseRecorder

    BeforeEach(func() {
        server = NewServer()
        recorder = httptest.NewRecorder()
    })

    Describe("GET /api/quote", func() {
        BeforeEach(func() {
            request, _ = http.NewRequest("GET", "/api/quote", nil)
        })

        Context("When there are quote sources", func() {
            It("returns a status code of 200", func() {
                server.ServeHTTP(recorder, request)
                Expect(recorder.Code).To(Equal(200))
            })

            It("returns a quote", func() {
                server.ServeHTTP(recorder, request)
                quoteJson := mapFromJSON(recorder.Body.Bytes())
                Expect(quoteJson["body"]).ToNot(BeNil())
                Expect(quoteJson["author"]).ToNot(BeNil())
            })
        })

        Context("When asked for a particular content type", func() {
            It("returns a json quote", func() {
                request, _ = http.NewRequest("GET", "/api/quote.json", nil)
                server.ServeHTTP(recorder, request)
                quoteJson := mapFromJSON(recorder.Body.Bytes())
                Expect(quoteJson["body"]).ToNot(BeNil())
                Expect(quoteJson["author"]).ToNot(BeNil())
            })

            It("returns an xml quote", func() {
                request, _ = http.NewRequest("GET", "/api/quote.xml", nil)
                server.ServeHTTP(recorder, request)
                quote := quoteFromXML(recorder.Body.Bytes())
                Expect(quote.Body).ToNot(BeNil())
                Expect(quote.Author).ToNot(BeNil())
            })
        })
    })

    Describe("GET /api/quote/file", func() {
        BeforeEach(func() {
            request, _ = http.NewRequest("GET", "/api/quote/file", nil)
        })

        Context("When there are quote sources", func() {
            It("returns a status code of 200", func() {
                server.ServeHTTP(recorder, request)
                Expect(recorder.Code).To(Equal(200))
            })

            It("returns a quote", func() {
                server.ServeHTTP(recorder, request)
                quoteJson := mapFromJSON(recorder.Body.Bytes())
                Expect(quoteJson["body"]).ToNot(BeNil())
                Expect(quoteJson["author"]).ToNot(BeNil())
            })
        })
    })

    Describe("GET /api/quote/theysaidso", func() {
        BeforeEach(func() {
            request, _ = http.NewRequest("GET", "/api/quote/theysaidso", nil)
        })

        Context("When there are quote sources", func() {
            It("returns a status code of 200", func() {
                server.ServeHTTP(recorder, request)
                Expect(recorder.Code).To(Equal(200))
            })

            It("returns a quote", func() {
                server.ServeHTTP(recorder, request)
                quoteJson := mapFromJSON(recorder.Body.Bytes())
                Expect(quoteJson["body"]).ToNot(BeNil())
                Expect(quoteJson["author"]).ToNot(BeNil())
            })
        })
    })

    Describe("GET /api/quote/quotedb", func() {
        BeforeEach(func() {
            request, _ = http.NewRequest("GET", "/api/quote/quotedb", nil)
        })

        Context("When there are quote sources", func() {
            It("returns a status code of 200", func() {
                server.ServeHTTP(recorder, request)
                Expect(recorder.Code).To(Equal(200))
            })

            It("returns a quote", func() {
                server.ServeHTTP(recorder, request)
                quoteJson := mapFromJSON(recorder.Body.Bytes())
                Expect(quoteJson["body"]).ToNot(BeNil())
                Expect(quoteJson["author"]).ToNot(BeNil())
            })
        })
    })
})
