
[![Build Status](https://travis-ci.org/binaryplease/gopherlix.svg?branch=master)](https://travis-ci.org/binaryplease/gopherlix)
[![GoDoc](https://godoc.org/github.com/binaryplease/gopherlix?status.svg)](https://godoc.org/github.com/binaryplease/gopherlix)
[![Go Report Card](https://goreportcard.com/badge/github.com/binaryplease/gopherlix)](https://goreportcard.com/report/github.com/binaryplease/gopherlix)
[![codecov](https://codecov.io/gh/binaryplease/gopherlix/branch/master/graph/badge.svg)](https://codecov.io/gh/binaryplease/gopherlix)


# gopherlix

A Server for the Gopher protocol, written in the Go.

<p align="center">
  <img src="./logo.png">
</p>

## Usage

### Quickstart

#### Installation

Just clone the repository and run `go build` in it.

```sh
git clone https://github.com/binaryplease/gopherlix.git
cd gopherlix
go build
```

#### Basic configuration

```ini
[paths]
content = "data/content"
templates = "data/templates"

[server]
port = "8000"
domain = "localhost"
host = "0.0.0.0"
```

The configuration file (`config.ini`) should be self explanatory. The server
includes some sample pages and should run with it out of the box on port `8000`.
You may want to change the port to gophers default port `70`

#### Run the server

Gopherlix will look for the configuration file in the same folder as the server
per defalt.

To start it just run `./gopherlix` and start browsing `goopher://localhost:8000`


### Features

#### Adding content

You can add content in the `content` directory configured in `config.ini`.
The `templates` directory includes templates that can be rendered in inside
other files. The `header.gph` and `footer.gph` will be added automatically to
all pages.


#### Directory listings

A request like `gopher://localhost/something` will look for
`data/content/something` and check if it's a file or a directory.
For a file, the contents of it will be returned. If the requested path is a
directory, it will try to find a file called `index.gph` in it and display it.

If no `index.gph` is found in this directory, it will generate a listing with
links to all the directory contents.

#### Templating

Gopherlix leverages the power of golangs
[templates](https://golang.org/pkg/text/template/) to generate content
dynamically. This allows to include ariables in the form of `{{.Variablename}}`
in any page that will be replaced accordingly. At the moment the following
variables are suported (more to come in future updates):

| Variable        | Description                    |
|-----------------|--------------------------------|
| {{.Directory}}  | Path of the current request    |
| {{.ServerName}} | Shows the server's domain name |

# Contributing

Pull-request, issues, feature-requst and contributions of any kind are *very*
welcome!
