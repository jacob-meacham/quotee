quotee
======
[![Build Status](https://drone.io/github.com/jacob-meacham/quotee/status.png)](https://drone.io/github.com/jacob-meacham/quotee/latest)
[![Coverage Status](https://coveralls.io/repos/jacob-meacham/quotee/badge.png?branch=develop)](https://coveralls.io/r/jacob-meacham/quotee?branch=develop)

Quotee is a simple web app using Martini + angularJS. Show quotes from a variety of sources.

Usage
=====
Quotee is intended for use with a high-definition media panel. With that in mind, the quotee front-end is designed to be simple and elegant. You can see an example at [https://quotee.herokuapp.com/](https://quotee.herokuapp.com/).

The front-end also takes a few parameters:

* style: If set to inverse, forces the inverse theme. *[https://quotee.herokuapp.com/?style=inverse](https://quotee.herokuapp.com/?style=inverse)*
* autoplay: If set to true, auto-refreshes the quote every 5 minutes. *[https://quotee.herokuapp.com/?autoplay=true](https://quotee.herokuapp.com/?autoplay=true)*

Sources
=======
Quotee exposes an extremely simple REST API and can pull quotes from a variety of sources. To get a quote from a random source, you can use:

```
GET /api/quote/
```

Each source is exposed via its name on the api

```
GET /api/quote/{sourceName}
```

Quotes are returned as simply as possible:

```
{ 
    "body" : "body of quote",
    "author" : "author of quote"
}
```

Any source can also render a quote as xml - just add .xml to the path

```
GET /api/quote.xml

<result>
    <Quote>
    <body>Body of Quote</body>
    <author>Author of Quote</author>
    </Quote>
</result>
```

By default, Quotee comes with 3 standard sources:

* File: Loads quotes from a csv file in the form of body,author
* QuoteDB: Returns quotes from quotedb.com
* TheySaidSo: Returns quotes from the theysaidso.com Quote of the Day, using a specified set of categories.

### Adding New Sources
Adding new sources is easy. A source must implement the QuoteSource interface

```
type QuoteSource interface {
    GetQuote() (Quote, error)
}
```
*See models/fileSource.go for an example*

Once you've implemented a new source, add it to the source map in server.go

```
sourceMap["mynewsource"] = MyNewSource{}
```

The source will be added to the list of sources when using the /api/quote endpoint. It will also be available directly at /api/quote/mynewsource.