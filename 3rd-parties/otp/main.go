package main

import (
	"fmt"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"

	"encoding/base32"
	"encoding/base64"
	"time"
)

// Demo function, not used in main
// Generates Passcode using a UTF-8 (not base32) secret and custom parameters
func GeneratePassCode(utf8string string) string {
	secret := base32.StdEncoding.EncodeToString([]byte(utf8string))
	passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    30,
		Skew:      0,
		Digits:    otp.DigitsEight,
		Algorithm: otp.AlgorithmSHA512,
	})
	if err != nil {
		panic(err)
	}
	return passcode
}

func main() {
	s := "secret"
	code := GeneratePassCode(s)
	fmt.Println(code)

	basicAuthRaw := fmt.Sprintf("userid:%s", code)
	fmt.Println(basicAuthRaw)

	encoded := base64.StdEncoding.EncodeToString([]byte(basicAuthRaw))
	fmt.Println(encoded)
}
