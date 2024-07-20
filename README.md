# X-Pay-Token util for Visa developer's API

## Description.

Some Visa Developer APIs require an API Key-Shared Secret Authentication, which Visa refers to as X-Pay Token. This library brings utils to create X-Pay tokens for your applications.

## Requirements

* Go 1.7 or later.
* Visa developer's API Key and Shared secret.

> [!TIP]
> See the [API Key – Shared Secret (X-Pay Token)](https://developer.visa.com/pages/working-with-visa-apis/x-pay-token) for generating both api keys and shared secrets from Visa developer's console.

> [!IMPORTANT]  
> This api key and shared secret should be kept secret on your server.

## Installation

To install the this library, please execute the following `go get` command.

```bash
go get vagrantvaquita.github.io/go/xpay
```


## Usage

Sample usage of the Directions API with an API key:

```go
package main

import (
    "net/url"

    "github.com/kr/pretty"
	"vagrantvaquita.github.io/go/xpay"
)

var (
	APIKey       = os.Getenv("VISA_API_KEY")
	SharedSecret = os.Getenv("VISA_SHARED_SECRET")
)


func main() {
	query := make(url.Values)
	query.Set("apiKey=", APIKey)

	tokenConfig := &TokenConfig{
		SharedSecret: SharedSecret,
		DateTime:     time.Now(),
		ResourcePath: "helloworld",
		QueryParams:  query,
		PostBody:     "",
	}


	token :=xpay.NewXPayToken(tokenConfig)

	pretty.Println(route)
}
```

> [!IMPORTANT]  
> For this example we set the API key and Shared Secret iside `VISA_API_KEY` and `VISA_SHARED_SECRET` environment variables.