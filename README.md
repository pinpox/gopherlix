
[![Build Status](https://travis-ci.org/binaryplease/gopher-server.svg?branch=master)](https://travis-ci.org/binaryplease/gopher-server)
[![GoDoc](https://godoc.org/github.com/binaryplease/gopher-server?status.svg)](https://godoc.org/github.com/binaryplease/gopher-server)
[![Go Report Card](https://goreportcard.com/badge/github.com/binaryplease/gopher-server)](https://goreportcard.com/report/github.com/binaryplease/gopher-server)
[![codecov](https://codecov.io/gh/binaryplease/gopher-server/branch/master/graph/badge.svg)](https://codecov.io/gh/binaryplease/gopher-server)


# gopher-server

A Server for the Gopher protocol, written in the Go.

<p align="center">
  <img src="./logo.png">
</p>

## Usage

To start the server create a directory with you content. At the moment only
textfiles and directories are supported.

The client will be able to request paths. Text files will be rendered in the
client's browser and directories will return a listing of all files in them.

### Custom gophermaps

To show custom content instead of the generated directory listing for a
requested directory, place a file named `index.gph` in it. It will be shown
instead of the default listing.
