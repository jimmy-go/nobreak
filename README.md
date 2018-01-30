## NoBreak cache server for third parties services with intermitent uptimes.

<p align="center">
    <img src="logo.png" width="200" />
</p>

[![License MIT](https://img.shields.io/npm/l/express.svg)](http://opensource.org/licenses/MIT)
[![Build Status](https://travis-ci.org/jimmy-go/nobreak.svg?branch=master)](https://travis-ci.org/jimmy-go/nobreak)
[![Go Report Card](https://goreportcard.com/badge/github.com/jimmy-go/nobreak)](https://goreportcard.com/report/github.com/jimmy-go/nobreak)
[![GoDoc](http://godoc.org/github.com/jimmy-go/nobreak?status.png)](http://godoc.org/github.com/jimmy-go/nobreak)
[![Coverage Status](https://coveralls.io/repos/github/jimmy-go/nobreak/badge.svg?branch=master)](https://coveralls.io/github/jimmy-go/nobreak?branch=master)

### Install:
```
go get gopkg.in/jimmy-go/nobreak.v0
```

### Usage:

```
nobreak -config=$PWD/_examples/youtube.yml
```

#### Config file contents:

```
# Target host.
host: https://www.youtube.com

# Listen port.
port: 9090

# Port for admin configuration dashboard (WIP).
admin_port: 8383

# Http client timeout in milliseconds.
timeout: 2000

# Auto flag for automatic cache return.
auto: true

# Database connection url for sqlite3. If empty in memory will be used.
database: ':memory:'

# Enable tls server.
tls_enabled: false

# TLS cert pem file.
tls_cert: 'none'

# TLS key pem file.
tls_key: 'none'
```

### License:

The MIT License (MIT)

Copyright (c) 2018 Angel del Castillo

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

