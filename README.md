# Introduction

Simple HTTP Service returns JSON:
```
{
  "id": "1",
  "message": "Hello world"
}
```

# Run tests

```
docker run -it -v $PWD:/go/src/app -w /go/src/app -e CGO_ENABLED=0 golang:alpine go test .
```
