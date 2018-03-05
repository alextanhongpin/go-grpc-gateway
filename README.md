# go-grpc-gateway


## Installation
```
$ brew install protobuf
$ go get -u -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
$ go get -u -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
$ go get -u -v github.com/golang/protobuf/protoc-gen-go
```

## Run

Steps to run:

1. Generate gRPC stub with `build.sh` (in `go`).
2. Implement your service in grpc, if you need to compile to any language other than go.
3. Generate reverse-proxy with `build-reverse-proxy.sh`
4. Optionally generate swagger endpoint `build-swagger.sh`

Your service should be up and running at `localhost:9090`.

It has four endpoints:

### The `POST` service EchoPost endpoint. 

Verify it by making a `POST` call to the endpoint:

```bash
$ curl -X POST -H "Content-Type: application/json" http://localhost:8080/v1/example/echo_post -d '{"value":"100"}'
```

Output:

```
{"value":"100"}
```

### The `GET` Echo endpoint:
```bash
$ curl -X POST -i http://localhost:8080/v1/example/echo/abcde
```

Output:
```
HTTP/1.1 200 OK
Content-Type: application/json
Grpc-Metadata-Bar: bar1
Grpc-Metadata-Content-Type: application/grpc
Grpc-Metadata-Foo: foo1
Trailer: Grpc-Trailer-Foo
Trailer: Grpc-Trailer-Bar
Date: Mon, 05 Mar 2018 14:30:03 GMT
Transfer-Encoding: chunked

{"id":"abcde"}
```

### The `POST` Echo endpoint:

```
$ curl -i http://localhost:8080/v1/example/echo/abcde/200
```

Output:

```
HTTP/1.1 200 OK
Content-Type: application/json
Grpc-Metadata-Bar: bar1
Grpc-Metadata-Content-Type: application/grpc
Grpc-Metadata-Foo: foo1
Trailer: Grpc-Trailer-Foo
Trailer: Grpc-Trailer-Bar
Date: Mon, 05 Mar 2018 14:30:25 GMT
Transfer-Encoding: chunked

{"id":"abcde","num":"200"}
```

### The swagger endpoint:

```bash
$ curl http://localhost:8080/swagger/hello.swagger.json
```

Output:

```{
  "swagger": "2.0",
  "info": {
    "title": "hello/hello.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/example/echo/{id}": {
      "post": {
        "operationId": "Echo",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/helloSimpleMessage"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "HelloService"
        ]
      }
    },
    "/v1/example/echo/{id}/{num}": {
      "get": {
        "operationId": "Echo2",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/helloSimpleMessage"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "num",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          }
        ],
        "tags": [
          "HelloService"
        ]
      }
    },
    "/v1/example/echo_post": {
      "post": {
        "operationId": "EchoPost",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/helloStringMessage"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/helloStringMessage"
            }
          }
        ],
        "tags": [
          "HelloService"
        ]
      }
    }
  },
  "definitions": {
    "helloSimpleMessage": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "num": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "helloStringMessage": {
      "type": "object",
      "properties": {
        "value": {
          "type": "string"
        }
      }
    }
  }
}
```
