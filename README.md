# 🧩 Captcha – A Unified CAPTCHA Verification Library for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/ja7ad/captcha.svg)](https://pkg.go.dev/github.com/ja7ad/captcha)

`captcha` is a modular, provider-agnostic CAPTCHA verification library for Go that makes it easy to validate challenge responses from **Cloudflare Turnstile**, **Google reCAPTCHA**, and **hCaptcha** using a unified interface.

It also includes optional **Protobuf support** for gRPC-based services.

---

## ✨ Features

- ✅ Unified interface for multiple CAPTCHA providers
- 🔐 Secure server-side verification of challenge tokens
- 📦 Simple API for Go web and microservice backends
- 🔁 Optional [Protocol Buffers](https://developers.google.com/protocol-buffers) integration for gRPC use cases
- 📚 Clean, testable, idiomatic Go

---

## 📦 Installation

```bash
go get github.com/ja7ad/captcha
```

If you're using Protobuf/gRPC support:

```bash
go get github.com/ja7ad/captcha/proto
```

---


## 🔐 Supported Providers

| Provider        | Enum                         | Notes                           |
|-----------------|------------------------------|----------------------------------|
| Cloudflare      | `captcha.ProviderCloudflare` | Fast and privacy-respecting     |
| Google reCAPTCHA| `captcha.ProviderGoogle`     | Widely supported                 |
| hCaptcha        | `captcha.ProviderHcaptcha`   | Strong bot protection            |


---

## 🚀 Quick Start

```go
package main

import (
	"context"
	"fmt"
	"github.com/ja7ad/captcha"
)

func main() {
	c, err := captcha.New(captcha.ProviderCloudflare, "your-secret-key")
	if err != nil {
		panic(err)
	}

	success, err := c.Validate(context.Background(), "challenge-token-from-client")
	if err != nil {
		fmt.Println("Verification error:", err)
		return
	}

	if success {
		fmt.Println("CAPTCHA verified ✅")
	} else {
		fmt.Println("CAPTCHA failed ❌")
	}
}
```

---

## 🔄 With Full Challenge Response

```go
challenge, err := c.ValidateWithResponse(ctx, "challenge-token")
if err != nil {
	log.Fatal(err)
}

if challenge.Success {
	// do something
}
```

---

## 🤝 Contributing

Pull requests, issues, and discussions are welcome!  
Feel free to fork and enhance the library for your use case.
