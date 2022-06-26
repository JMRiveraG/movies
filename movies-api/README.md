# movies-api

`movies-api` is an API that returns information on movies, based on a text search. Currently it uses [omdbapi.com](https://omdbapi.com) to do so.

## Working with movies-api

`movies-api` requires go 1.13 to be built.

Here are some basic go commands you can use to work with movies-api.

* To build: `go build .`
* To run (after building): `./movies-api`
* To execute tests (in all packages): `go test -v ./...`

NOTE: The API by default listens on port 5432; this is set in the main function.
