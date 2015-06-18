package main

import "github.com/hgfischer/go-otp"

func main() {
	totp := &otp.TOTP{Secret: "your-secret", IsBase32Secret: true}
	token := totp.Get()
	println(token)
}