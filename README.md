# RESTful API with Go and Gin

Example of RESTful API written in Go and
[Gin framework](https://gin-gonic.com/).

This example project simply followed the official tutorial
[Tutorial: Developing a RESTful API with Go and Gin](https://golang.org/doc/tutorial/web-service-gin).

It may be slightly different from the original (more comments etc.).

---

## How to get and run this example

<pre>
  $ git clone https://github.com/rideee/restful-api-with-go-and-gin
  $ cd restful-api-with-go-and-gin
  $ go mod tidy
  $ go run .
</pre>

### Now you are ready to do some HTTP requests, you can use tools like

- [Postman](https://www.postman.com/)
- [Insomnia](https://insomnia.rest/)
- [curl](https://curl.se/) - like in official tutorial
- [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) -
  Visual Studio Code Extension

## REST_Requests.http file

_This file is used by
[REST Client (Visual Studio Code extension)](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)._

If you are using Visual Studio Code, you can install extension called
[REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)
which give you possibility to make HTTP requests from your editor.

Some examples requests are in REST_Requests.http file.

---

# Troubleshooting

## **[Error]** go: modules disabled by GO111MODULE=off (...)

Error occurs when using 'go mod tidy':

<pre>
  $ go mod tidy
  go: modules disabled by GO111MODULE=off; see 'go help modules'
</pre>

It means that your **GO111MODULE**
[environment variable](https://en.wikipedia.org/wiki/Environment_variable) is
set to **off**.

You can check this using:

<pre>
  $ go env
  GO111MODULE="off"
  GOARCH="amd64"
  GOBIN=""
  ...
</pre>

### What is GO111MODULE?

> From **Go version 1.11**, the way to deal with modules was **revamped**.
>
> Beforehand, it was required that you place all your Go code under a $GOPATH.
> Which many Go developers didn’t much like, as it seemed chaotic at best.
>
> Also, having a package management/module management solution is so much more
> productive, and manageable.
>
> [Source](https://ao.ms/solved-go-mod-init-modules-disabled-by-go111moduleoff/)

<pre>
  $ go help modules
  Modules are how Go manages dependencies.

  A module is a collection of packages that are released, versioned, and
  distributed together. Modules may be downloaded directly from version control
  repositories or from module proxy servers.

  For a series of tutorials on modules, see
  https://golang.org/doc/tutorial/create-module.

  For a detailed reference on modules, see https://golang.org/ref/mod.

  By default, the go command may download modules from https://proxy.golang.org.
  It may authenticate modules using the checksum database at
  https://sum.golang.org. Both services are operated by the Go team at Google.
  The privacy policies for these services are available at
  https://proxy.golang.org/privacy and https://sum.golang.org/privacy,
  respectively.

  The go command's download behavior may be configured using GOPROXY, GOSUMDB,
  GOPRIVATE, and other environment variables. See 'go help environment'
  and https://golang.org/ref/mod#private-module-privacy for more information.
</pre>

- [Go to -> Official - Series of tutorials on modules](https://golang.org/doc/tutorial/create-module)
- [Go to -> Official - Detailed reference on modules](https://golang.org/ref/mod)
- [Go to -> Official - Private Module Privacy](https://golang.org/ref/mod#private-module-privacy)

## Solution

The easiest way to fix the problem is to set **GO111MODULE** environment
variable to **auto**, using **go env** command (like below), then verify also by
using **go env**:

<pre>
  $ go env -w GO111MODULE=auto
  $ go env
  GO111MODULE="auto"
  GOARCH="amd64"
  GOBIN=""
  ...
</pre>
