# gotest

[![Build Status](https://travis-ci.org/jasonlam604/gotest.svg?branch=master)](https://travis-ci.org/jasonlam604/gotest)



## Overview

Go Skelton project that supports Travis CI, Unit Testing and Glide

## Prerequisite 

Assumption you have of installed GoLang and set the GOPATH, if not please refer to the [GoLang.org](https://golang.org/) site

As well, you will need to install [Glide](https://glide.sh/] the package manager).

I personally used brew:

```bash
brew install glide
```

## Quick Start

Assuming you have all the prerequisite software installed, first step is to install dependencies using Glide

Note because the *[vendor](https://golang.org/cmd/go/#hdr-Vendor_Directories)* requirement is to reside under source, the Glide configuration files reside in *src* directory, namely the *glide.yaml* file.


```bash
> cd src
> glide install
```

In this instance of *glide install* the following dependencies are installed:

* StringUtil (for simple us of 3rd party library to reverse a string, see main.go)
* Testify (used for unit testing)

Go back to the root of GOPATH (your project root)

Run the usual [Go commands](https://golang.org/cmd/go/)

*Build*
```bash
go build github.com/jasonlam604/gotest
```
*Install*
```bash
go install github.com/jasonlam604/gotest
```

*Test*
```bash
go test github.com/jasonlam604/gotest
```

*Test with Verbose*
```bash
go test -v github.com/jasonlam604/gotest
```

## Adding more dependencies

Instead of the usual go get:

```bash
go get github.com/go-sql-driver/mysql
```

Use Glide in place of go (reminder you do need to be in the /src directory to do this)
```bash
glide get github.com/go-sql-driver/mysql
```

*Run the Program* 
```bash
~/bin/gotest
```

The output should be the reverse of "hello":

*Run the Program* 
```bash
olleh
```

## What You Need to Change

You will need to update the folder to match your repository.

Current folder src layout:

```bash
~/src/github.com/jasonlam604/gotest/
```

Example if you were to host it on GitHub:
```bash
~/src/github.com/[your-github-username]/[your-github-repo-name]/
```
