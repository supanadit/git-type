# Gity

[![Build Status](https://travis-ci.com/supanadit/git-type.svg?branch=master)](https://travis-ci.com/supanadit/git-type)
[![Go Report Card](https://goreportcard.com/badge/github.com/supanadit/git-type)](https://goreportcard.com/report/github.com/supanadit/git-type)

Check type of Git url protocol simply whether is SSH, HTTP or HTTPS with Zero Dependencies

```golang
type Type struct {
	Url            string
	Type           string
	RepositoryName string
}
```

## Installation

```shell script
go get -u -v github.com/supanadit/gity
```

## How To Use
```golang
import "github.com/supanadit/gity"
```

### For SSH
```golang
gity, err := gity.Check("git@github.com:supanadit/jwt-go.git")
if err != nil {
    panic(err)
}
fmt.Println(gity.IsHTTPORS()) // false
fmt.Println(gity.IsHTTP()) // false
fmt.Println(gity.IsHTTPS()) // false
fmt.Println(gity.IsSSH()) // true
```

### For HTTP / HTTPS
```golang
type, err := gity.Check("https://github.com/supanadit/jwt-go.git")
if err != nil {
    panic(err)
}
fmt.Println(type.IsHTTPORS()) // true
fmt.Println(type.IsHTTP()) // false
fmt.Println(type.IsHTTPS()) // true
fmt.Println(type.IsSSH()) // false
```
