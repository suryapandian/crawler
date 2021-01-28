### A simple crawler in Go!

The crawler uses `net/http` package to parse a web page and fetch the sitemap.
When the application is started the crawler is initalized with the web page to be crawled.
The web page could be passed to the application by setting the environment variable `URL_TO_CRAWL`, if no environment variable is set then the application crawls the default url("https://www.cuvva.com/")

The crawler maintains two channels, one to read the urls to be crawled and another to process the sitemaps. When the application starts the url is pushed to the go channel.
The crawler also runs a pool of go routines to keep reading from the urlsChannel. The url that is pushed is picked by one of the go channel and it is parsed to get the sitemap. The sitemap is then pushed to the sitemaps channel which is then processed by another go routine.

A cache of the urls that is parsed is maintained to avoid parsing the same web page again.
Context cancellation is used to gracefully close the crawler on sigterm.

## Extra details about the application
- The application has unit tests for the scraper.
- The application also has github actions that run on every push to ensure that the build passes with the code change
- The application also has a pre-commit hook that makes sure the code is formatted and the tests pass before getting committed

## Running the application

Using pre-installed go:

```
go run main.go
```

Using docker:

```
 docker build -t crawler .  

 docker run -i -t "crawler"       

 ```

## Pre-commit 
To ensure a clean code.

   Install pre-commit [https://pre-commit.com/#install]
   Install .pre-commit-config.yaml as a pre-commit hook

	pre-commit install

    Go static analysis tools run automatically on pre-commit. Run checks manually if needed using

	pre-commit run --all-files
