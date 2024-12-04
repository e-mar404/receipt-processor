# Receipt Processor

## Prerequisites

Only prerequisite to run the docker container is to have docker-compose installed and docker daemon running.

I do have an `air.toml` file that I was using with [air](https://github.com/air-verse/air) for hot reload development, but air does not need to be installed.

## Running

1. Get the repo and cd into the dir that it creates.

```
git clone https://github.com/e-mar404/receipt-processor

cd receipt-processor
```

2. Running the following command should start the test suite and the docker instance

```
./run
```

If that script does not run you can copy this to run it manually 

Tests: 

```
go test ./test
```

Docker: 

```
docker-compose down
docker-compose up --build
```

*Note* : The application automatically starts on port 8080 and you can access it localhost ([localhost:8080](http://localhost:8080))

### Approach

Since this is a fairly light weight api with no need to render anything to html I decided to use [gin](https://github.com/gin-gonic/gin). This makes creating the api routs very simple. If we were needing to do more html rendering and even more flexibility maybe it might've been a good idea to look at using [echo](https://github.com/labstack/echo). Didn't thing there was a need to recreate a whole server protocol or http server since these technologies are very well documented, tested and reliable.

I knew I needed some testing since manually testing would have been very tedious and taken a really long time. At first I was testing only an overarching `CountPoints()` but then I saw that this was not going to be extensible and started to have some confusion as to why tests were failing so even if it seems over kill I added unit tests for each rule and if those pass then we can have confidence that that any combination of those rules will yield the correct point count
