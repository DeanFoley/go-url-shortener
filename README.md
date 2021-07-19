# Go URL Shortener

## Stated Objectives

Write an API for a URL shortener which satisfies the following behaviour:

* Accepts a URL to be shortened
* Generates a short URL for the original URL
* Accepts a short URL and redirects the user to the original URL

### Bonus

Comes with a CLI which can be used to call your service.

## Things We'd Like To See

* Sound design approach; not over-complicated or over-engineered.
* Code that's easy to read; not "clever".
* Sensible tests in place.

## Benchmarks

### Handlers

| handler name | ns/op | bytes/op | allocs/op |
|----|----|----|----|
| shortenURLHandler | 259277 | 18716 | 136 |
| retrieveURLHandler | 289681 | 18245 | 131 |

### App

| app name | ns/op | bytes/op | allocs/op |
|----|----|----|----|
| stripURL | 2854 | 105 | 0 |
| urlShortener | 705.7 | 16 | 1 |
| validateShortURL | 226.2 | 0 | 0 |