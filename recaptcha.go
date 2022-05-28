package recaptcha

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// scoreMinValue - minimal value of score
// 0.0 - is bot / 0.9 - is human. Optimal score value is 0.5
const scoreMinValue = 0.5

// Response - recaptcha response
type Response struct {
	Success    bool     `json:"success"`
	Score      float32  `json:"score"`
	Timestamp  string   `json:"challenge_ts"`
	Hostname   string   `json:"hostname"`
	ErrorCodes []string `json:"error-codes,omitempty"`
}

const (
	// VerifyUrl - Recaptcha verify url.
	// https://developers.google.com/recaptcha/docs/verify
	VerifyUrl = "https://www.google.com/recaptcha/api/siteverify"

	// ReTokenHeader - recaptcha token http header
	ReTokenHeader = "token"
)

type Config struct {
	ApiKey        string
	VerifyUrl     string
	ReTokenHeader string
	Scope         float32
}

var defaultConfig = Config{
	VerifyUrl:     VerifyUrl,
	ReTokenHeader: ReTokenHeader,
	Scope:         scoreMinValue,
}

func New(config Config) fiber.Handler {

	cfg := defaultConfig

	cfg.ApiKey = config.ApiKey

	if config.ReTokenHeader != "" {
		cfg.ReTokenHeader = config.ReTokenHeader
	}

	if config.VerifyUrl != "" {
		cfg.ReTokenHeader = config.ReTokenHeader
	}

	if config.Scope != 0 {
		cfg.Scope = config.Scope
	}

	return func(c *fiber.Ctx) error {

		token := c.Get(cfg.ReTokenHeader)

		if strings.TrimSpace(token) == "" {
			c.Status(fiber.StatusForbidden)
			return c.JSON(fmt.Sprintf("Header: %s was not set", cfg.ReTokenHeader))
		}

		resp, err := makeRequest(token, cfg.ApiKey, cfg.VerifyUrl)

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		if resp.Success && resp.Score > cfg.Scope {
			return c.Next()
		}

		c.Status(fiber.StatusForbidden)
		return c.JSON("Invalid key/timeout/duplicate")
	}
}

// makeRequest - request to validate google service
func makeRequest(token, apiKey, verifyUrl string) (Response, error) {
	var r Response

	res, err := http.PostForm(verifyUrl, url.Values{
		"secret":   {apiKey},
		"response": {token},
	})

	if err != nil {
		return r, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return r, err
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}
