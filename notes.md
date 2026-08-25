
A url shortener has 2 endpoints:
- `/shorten` POST endpoint that takes the OG url and returns the shortened url
`POST api/v1/data/shorten`
- `/redirect` GET endpoint that returns the OG url for HTTP redirect
`GET api/v1/shortUrl`

