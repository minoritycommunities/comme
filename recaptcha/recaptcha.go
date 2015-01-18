package recaptcha

// reCaptcha Documentation:
// https://developers.google.com/recaptcha/docs/verify

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var secret string

var recaptchaError = errors.New("Invalid ReCaptcha")

type recaptchaResult struct {
	Success bool   `json:"success"`
	Errors  []string `json:"error-codes"`
}

func Init(newSecret string) {
	secret = newSecret
}

func Check(response, remoteip string) error {
	apiEndpoint := "https://www.google.com/recaptcha/api/siteverify"
	params := fmt.Sprintf("?secret=%s&response=%s&remoteip=%s",
		secret,
		response,
		remoteip)
	resp, err := http.Get(apiEndpoint + params)
	defer resp.Body.Close()
	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		log.Println("reCaptcha result:", string(body))
		if err == nil {
			var data recaptchaResult
			json.Unmarshal(body, &data)
			if data.Success {
				return nil
			}
			if len(data.Errors) >= 1 {
				err = errors.New("reCaptcha error(s): " + fmt.Sprintf("%v", data.Errors))
			} else {
				err = recaptchaError
			}
			return err
		}
		return err
	}
	return err
}
