# go-grpc-gateway


Steps to run:

1. Generate gRPC stub with `build.sh`
2. Implement your service in grpc with `build-service.sh`
3. Generate reverse-proxy with `build-reverse-proxy.sh`
4. Optionally generate swagger endpoint `build-swagger.sh`

Your service should be up and running at `localhost:9090`.

It has four endpoints:

1. The service EchoPost endpoint. Verify it by making a `POST` call to the endpoint:

```bash
$ curl -X POST -H "Content-Type: application/json" http://localhost:8080/v1/example/echo_post -d '{"value":"100"}'
```

2. The service Echo endpoint has a `GET` and `POST` endpoint:
```bash
$ curl -X POST -i http://localhost:8080/v1/example/echo/abcde

$ curl -i http://localhost:8080/v1/example/echo/abcde/200
```

3. The swagger endpoint:

```bash
$ curl http://localhost:8080/swagger/hello.swagger.json exposes the swagger endpoint
```
