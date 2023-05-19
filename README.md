# gRPC/Protobuf demo

## Overview

Demo App which contains simple example of gRPC/Protobuf server and client.

Includes:

* simple `.proto` file
* Makefile with all necessary targets
* simple RPC calls covered by UT
* simple storage implementation to store input data
* streaming examples, server-side and client-side
* optional package compression
* integration with swagger (standalone HTTP server)

## Build & run

Build with:

```bash
make
```

It will build the client and server executables, placed in the `bin` directory; and run the local tests.

You may use:

```bash
./bin/client
./bin/server
```

to run the client and server accordingly. The client shows streaming functionality.

## Integration with swagger

The server already exposes the swagger json file at `http://127.0.0.1:8081/swagger-ui`. So you may use an external swagger-ui tool like the following:

```bash
docker run --name sw --rm -p 80:8080 -e SWAGGER_JSON_URL=http://127.0.0.1:8081/swagger-ui swaggerapi/swagger-ui
```

After this, open your browser and go to `http://127.0.0.1` to see the available documentation generated from the `.proto` file.
