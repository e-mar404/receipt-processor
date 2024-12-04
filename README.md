# Receipt Processor

## Prerequisites

Only prerequisite to run the docker container is to have docker-compose installed and docker daemon running.

I do have an `air.toml` file that I was using with [air]() for hot reload development, but air does not need to be installed.

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
