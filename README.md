# Description

One of the best parts of Go is that compiled programs can be distributed and executed as a single dependency-free binary file. Unfortunately, programs requiring access to static files for thing like configuration, web server assets, or database migrations must be distributed alongside those static files, eliminating the single-file benefit. Fortunately, there's a way to bundle these assets inside the binary itself.

This post demonstrates how to embed static assets inside Go binaries using the pkger.

> Credit: https://codesalad.dev/blog/embed-static-files-in-go-binaries-with-pkger-15

# What is pkger?
pkger is a tool for embedding static files into Go binaries.

# Hands On
First thing that we need to do is generating pkged.go source file. In order to do that we need a "pkger" cli, so the following command will install this utility to our host
```bash
$ go get github.com/markbates/pkger/cmd/pkger
```

Then clone this repository and create pkged.go source file using the following command:
```bash
$ pkger
```

To check what's happening, we can also run pkger list to preview the assets that will be embedded.
```bash
$ pkger list
github.com/developer-guy/static-file-embedding-in-go-using-pkger
 > github.com/developer-guy/static-file-embedding-in-go-using-pkger:/hello-world.txt
```

Lets dockerize our application and build based on empty layer using special image called "scratch".
```bash
$ docker image build -t static-embed:v1 .
[+] Building 4.4s (14/14) FINISHED
 => [internal] load build definition from Dockerfile                                                                                                                                                        0.0s
 => => transferring dockerfile: 320B                                                                                                                                                                        0.0s
 => [internal] load .dockerignore                                                                                                                                                                           0.0s
 => => transferring context: 2B                                                                                                                                                                             0.0s
 => [internal] load metadata for docker.io/library/golang:1.15.7-alpine                                                                                                                                     1.8s
 => [auth] library/golang:pull token for registry-1.docker.io                                                                                                                                               0.0s
 => CACHED [stage-1 1/2] WORKDIR /app                                                                                                                                                                       0.0s
 => [stage-0 1/6] FROM docker.io/library/golang:1.15.7-alpine@sha256:dbda4e47937a3abb515c386d955002be5116d060c90d936127cc24ac439c815c                                                                       0.0s
 => [internal] load build context                                                                                                                                                                           0.0s
 => => transferring context: 29.98kB                                                                                                                                                                        0.0s
 => CACHED [stage-0 2/6] WORKDIR /app                                                                                                                                                                       0.0s
 => CACHED [stage-0 3/6] COPY go.mod go.sum ./                                                                                                                                                              0.0s
 => CACHED [stage-0 4/6] RUN go mod download                                                                                                                                                                0.0s
 => [stage-0 5/6] COPY ./ ./                                                                                                                                                                                0.0s
 => [stage-0 6/6] RUN go build -o static-embed                                                                                                                                                              2.4s
 => [stage-1 2/2] COPY --from=0 /app/static-embed ./                                                                                                                                                        0.0s
 => exporting to image                                                                                                                                                                                      0.0s
 => => exporting layers                                                                                                                                                                                     0.0s
 => => writing image sha256:013112fa5145b4f6e7cdebe7e5dd292d6ac81ff868ea58f84e36a346c23c66bc                                                                                                                0.0s
 => => naming to docker.io/library/static-embed:v1                                                                                                                                                          0.0s
```

Lets run and verify the output that should be match with the content of the file
```bash
$ docker container run static-embed:v1
Name:  hello-world.txt
Size:  12
Mode:  -rw-r--r--
ModTime:  2021-02-10 18:37:49.168973512 +0300 +0300
hello world
```

# References

* https://codesalad.dev/blog/embed-static-files-in-go-binaries-with-pkger-15
* https://github.com/markbates/pkger
* https://blog.gobuffalo.io/introducing-pkger-static-file-embedding-in-go-1ce76dc79c65
