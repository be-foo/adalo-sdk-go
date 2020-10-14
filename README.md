# Adalo SDK Go

This is an unofficial SDK for the use with Go. It provides a thin wrapper for the endpoints
provided by [The Adalo API](https://help.adalo.com/integrations/the-adalo-api).

[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/be-foo/adalo-sdk-go)
[![Build Status](https://github.com/be-foo/adalo-sdk-go/workflows/test/badge.svg)](https://github.com/be-foo/adalo-sdk-go/actions)
[![Coverage Status](https://coveralls.io/repos/github/be-foo/adalo-sdk-go/badge.svg?branch=develop)](https://coveralls.io/github/be-foo/adalo-sdk-go?branch=develop)
[![Go Report Card](https://goreportcard.com/badge/github.com/be-foo/adalo-sdk-go)](https://goreportcard.com/report/github.com/be-foo/adalo-sdk-go)
![License](https://img.shields.io/github/license/be-foo/adalo-sdk-go)
## Installation

Make sure your project is using Go Modules (it will have a `go.mod` file in its
root if it already is):

``` sh
go mod init
```

Then, reference stripe-go in a Go program with `import`:

``` go
import "github.com/be-foo/adalo-sdk-go"
```

Run any of the normal `go` commands (`build`/`install`/`test`). The Go
toolchain will resolve and fetch the stripe-go module automatically.

Alternatively, you can also explicitly `go get` the package into a project:

``` sh
go get -u github.com/be-foo/adalo-sdk-go
```

## Using the SDK

Before using any of the SDK functions, make sure you have configured the credentials for your Adalo app.
The keys will be used globally as parameters to your requests.

``` go
import "github.com/be-foo/adalo-sdk-go"

func main() {
    adalo.ApiKey = "<YOUR-API-KEY>"
    adalo.AppKey = "<YOUR-APP-KEY>"
}
```

### Collection API

The API enables you to run basic CRUD operations on your Adalo collections.
See below for some examples on how to use it with this SDK.

**First initialize your collection**
``` go
personCollection := adalo.NewCollection("<ID-OF-PERSON-COLLECTION>")
```

**Get All Items**
``` go
var persons []interface{} // result will be bind to this variable

err := personCollection.All(&persons)

if err != nil {
    panic(err)
}
```

**Get Item by ID**
``` go
var person interface{} // result will be bind to this variable

err := personCollection.Get(1, &person)

if err != nil {
    panic(err)
}
```

**Insert Item**
``` go
var person interface{} // will bind created item to this variable

err := personCollection.Insert(struct{
    Name    string
    Age     int
}{
    Name:   "John",
    Age:    21,
}, &person)

if err != nil {
    panic(err)
}
```

**Update Item**
``` go
var person interface{} // will bind updated item to this variable

err := personCollection.Update(1, struct{
    Name    string
    Age     int
}{
    Name:   "John",
    Age:    21,
}, &person)

if err != nil {
    panic(err)
}
```

#### Additional Notes

Because each collection has different fields, the SDK must work with the `interface{}` type.
For type safety and a over-all better developer experience, we suggest you implement a wrapper/ superset
of the `Collection` type, where you declare types according to your specific collections and then
create wrapper functions for the functions of `Collection`, using your type.

You can see a full example of how this can look like in [example/person.go](./example/person.go).
