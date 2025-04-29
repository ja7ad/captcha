package captcha_test

import (
	"context"
	"fmt"
	"github.com/ja7ad/captcha"
	"log"
)

func ExampleNew() {
	cloudflare, err := captcha.New(captcha.CloudflareProvider, "secret")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("provider: ", cloudflare.Provider())

	valid, err := cloudflare.Validate(context.Background(), "challenge", "ip")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(valid)
}
