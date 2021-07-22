# go-url-shortener

go-url-shortener allows you to create short URLs to redirect to any fully-qualified URL, and retrieve the referrals by simply querying the shortened URL.

* Shortens any URL to a 12-characdter re-direct.
* Retrieves the fully-qualified URL for any short URL.
* Accessible via RESTful interface or with provided CLI tool.

## Building & Running

For the API:

```bash
cd ./cmd && go build -o <name> && ./<name>
```

For the CLI:

```bash
cd .cmd/cli && go build -o <name> && ./<name> <command> <params>
```

## API Interface

### Shorten

Accessed via:

```
POST /shorten/
```

#### Request 

```json
{
    "url": "https://www.lush.com/uk/en/p/good-karma-everybody-needs-some-shower-gel",
}
```

#### Response

```json
{
    "longURL": "https://www.lush.com/uk/en/p/good-karma-everybody-needs-some-shower-gel",
    "shortURL": "https://df.dv/9UhhRFxVDElPV06C",
}
```

### Longen

Accessed via:

```
GET /
```

#### Request

```
GET https://df.dv/9UhhRFxVDElPV06C
```

#### Response

```json
{
    "longURL": "https://www.lush.com/uk/en/p/good-karma-everybody-needs-some-shower-gel",
    "shortURL": "https://df.dv/9UhhRFxVDElPV06C",
}
```

## CLI Commands

The CLI tool supports both shortening and longening of URLs, accessible by running it in one of the two following ways.

### Shorten

#### Request

```bash
./cli shorten -shorten-URL=https://www.lush.com/uk/en/p/good-karma-everybody-needs-some-shower-gel
```

#### Response

```bash
https://df.dv/9UhhRFxVDElPV06C
```

### Longen

#### Request

```bash
./cli longen -longen-URL=kSDOejD9yPsuKRuO
```

#### Response

```bash
https://www.lush.com/uk/en/p/good-karma-everybody-needs-some-shower-gel
```

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