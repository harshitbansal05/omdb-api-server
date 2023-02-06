## omdb-api-server

The server can be begun by running the command: `go run main.go`. Before that, a `.env` file must be placed in the root directory with the OMDB url and access keys. A `.sample-env` is also present in the root directory for reference.

To run the tests, the command `go test -v ./...` can be used. It requires a `.env` file to be placed in the controllers directory as well.  