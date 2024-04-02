# Web Scrapper GoLang 

*Source : [zerotomastery.io/blog/golang-practice-projects/](https://zerotomastery.io/blog/golang-practice-projects/)*

A web scraper is a great project for getting started with Go because it has a small number of steps. Even better still, Go has packages available to help you with these steps!

This web scraper is a terminal (command line) application. It must connect to a web resource, download specified data, and output it to the terminal.

[You can grab the package for this here.](https://pkg.go.dev/github.com/gocolly/colly/v2#section-readme)

Then, once you’ve built the scraper and got it working, feel free to level up your skills by implementing more features.

URL for test : `https://go-colly.org/`

Use this command to *install packages*:
```
go mod download
```

Use this command to *build and run* project :
```
go build . 
./web-scrapper
```

## Level up even further!
**Scrape more data:**
- Generating paginated URLs based on the input URL, such as<br> example.com/blog/1, example.com/blog/2, etc
- Following links

**Add CLI arguments and flags to allow the user to customize the application:**
- Specify the URLs to scrape
- Use custom CSS selectors
- Enable auto-pagination based on an input pattern
- Enable link following using a CSS selector

**Implement [rate limiting](https://en.wikipedia.org/wiki/Rate_limiting) per domain so the user doesn't get blocked if they make lots of requests**

**[Cache](https://en.wikipedia.org/wiki/Cache_(computing)) pages so subsequent runs don't need to download the same page again:**
- Download a fresh copy if the cached page is over 1 day old
- Add a CLI flag to clear the cache
- Add a CLI flag to ignore the cache and download the page anyway
- Save the new page in the cache

## Real life example
* Cryto coin martket cap : [Click here](https://go-colly.org/docs/examples/cryptocoinmarketcap/)
* Instagram Images : [Click here](https://go-colly.org/docs/examples/instagram/)
* Shopify Sitemap : [Click here](https://go-colly.org/docs/examples/shopify_sitemap/)

### HAPPY CODING !!!