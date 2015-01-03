package main

import (
    "net/http"
    "regexp"
    "strings"
    "time"
    "math/rand"
    "log"

    "github.com/go-martini/martini"
    "github.com/jacob-meacham/quotee/models"
    "github.com/jacob-meacham/quotee/routes"
)

func NewServer() *martini.Martini {
    m := martini.New()
    
    // Setup middleware
    m.Use(martini.Recovery())
    m.Use(martini.Logger())
    m.Use(martini.Static("public"))
    m.Use(MapEncoder)
    // Setup routes
    r := martini.NewRouter()

    sourceMap := getQuoteSources()
    logSources(sourceMap)
    routes.SetQuoteSources(sourceMap)

    r.Get("/api/quote", routes.GetQuote)
    r.Get("/api/quote/:source", routes.GetQuote)
    r.Get("/api/quote/theysaidso/static", routes.GetTest)
    
    // Add the router action
    m.Action(r.Handle)
    return m
}

func getQuoteSources() map[string]models.QuoteSource {
    fileQuoteSource, err := models.CreateFileQuoteSource("data/quotes.csv")
    if err != nil {
        panic(err)
    }

    return map[string]models.QuoteSource{
        "file": fileQuoteSource,
        //"theysaidso": models.TheySaidSoQuoteSource{Url: "http://api.theysaidso.com/qod.json", Categories: []string{"funny", "life", "inspire", "love"}},
        "theysaidso": models.TheySaidSoQuoteSource{Url: "http://localhost:3000/api/quote/theysaidso/static", Categories: []string{"funny", "life", "inspire", "love"}},
        "quotedb": models.QuoteDBSource{Url: "http://www.quotedb.com/quote/quote.php?action=random_quote"},
    }
}

func logSources(sources map[string]models.QuoteSource) {
    log.Print("Registered Sources:")
    for name, v := range sources {
        log.Printf("api/quote/%s: %s", name, v)
    }
}

// The regex to check for the requested format (allows an optional trailing
// slash).
var rxExt = regexp.MustCompile(`(\.(?:xml|text|json))\/?$`)

// MapEncoder intercepts the request's URL, detects the requested format,
// and injects the correct encoder dependency for this request. It rewrites
// the URL to remove the format extension, so that routes can be defined
// without it.
func MapEncoder(c martini.Context, w http.ResponseWriter, r *http.Request) {
    // Get the format extension
    matches := rxExt.FindStringSubmatch(r.URL.Path)
    ft := ".json"
    if len(matches) > 1 {
        // Rewrite the URL without the format extension
        l := len(r.URL.Path) - len(matches[1])
        if strings.HasSuffix(r.URL.Path, "/") {
            l--
        }
        r.URL.Path = r.URL.Path[:l]
        ft = matches[1]
    }
    // Inject the requested encoder
    switch ft {
    case ".xml":
        c.MapTo(routes.XmlEncoder{}, (*routes.Encoder)(nil))
        w.Header().Set("Content-Type", "application/xml")
    default:
        c.MapTo(routes.JsonEncoder{}, (*routes.Encoder)(nil))
        w.Header().Set("Content-Type", "application/json")
    }
}

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    server := NewServer()
    server.Run()
}
