<p align="center">
<h1 align="center">HttpClient</h1>
<p align="center">A lightweight HTTP client package for Golang</p>
</p>

<p align="center">
<p align="center"><a href="https://codecov.io/gh/calvin-puram/httpclient" >
 <img src="https://codecov.io/gh/calvin-puram/httpclient/graph/badge.svg?token=PS9FK3IHFA"/>
 </a></p>
</p>



## Installation

```bash
# Go Modules
require https://github.com/calvin-puram/httpclient/ghttp
```

## Usage

```go
// Import ghttp into your code.
import "github.com/calvin-puram/httpclient/ghttp"

// Configure the client
// First you need to configure and build the client as you need

reqHeaders := make(http.Header)
reqHeaders.Set("headers-key", "headers-value")

// setting global headers use:
SetHeaders(reqHeaders)
// setting user agent use:
SetUserAgent("user-agent")
// set max ideal connection per host use:
SetMaxIdleConns(5)
//specifies the amount of time to wait for a server's response headers after fully writing the request use:
SetResponseHeaderTimeout(5 * time.Second)
// default false: disable client request timeout use:
SetDisableTimeouts(false)


client := ghttp.NewBuilder().SetHeaders(reqHeaders).SetUserAgent("user-agent").
SetMaxIdleConns(5).SetResponseHeaderTimeout(5 * time.Second)SetDisableTimeouts(false).
Build() // build the client


```

## Making HTTP calls
Take a look at all of the [examples folder](./examples/) for more information on how to make http calls

