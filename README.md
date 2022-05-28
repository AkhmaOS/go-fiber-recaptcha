# Go fiber recaptcha


## Install
go get github.com/AkhmaOS/go-fiber-recaptcha

## Example


Set config
```

type Config struct {
ReCaptcha recaptcha.Config
}

func GetConfig() *Config {

    conf := &Config{}
    
    apiKey := os.Getenv("RECAPTCHA_API_KEY)
    conf.ReCaptcha.ApiKey = apiKey
    return conf
}

```


Fiber use case as middleware
```

conf := GetConfig()
app := fiber.New()
app.Use(recaptcha.New(conf.ReCaptcha))
```

Set Custom header
```
func GetConfig() *Config {

    conf := &Config{}
    
    apiKey := os.Getenv("RECAPTCHA_API_KEY)
    conf.ReCaptcha.ApiKey = apiKey
    
    conf.ReCaptcha.ReTokenHeader = "Custom-Header"
    return conf
}
```

Set Custom Scope
```
func GetConfig() *Config {

    conf := &Config{}
    
    apiKey := os.Getenv("RECAPTCHA_API_KEY)
    conf.ReCaptcha.ApiKey = apiKey
    
    conf.ReCaptcha.Scope = 0.8
    return conf
}
```


Set Custom Verify URL
```
func GetConfig() *Config {

    conf := &Config{}
    
    apiKey := os.Getenv("RECAPTCHA_API_KEY)
    conf.ReCaptcha.ApiKey = apiKey
    
    conf.ReCaptcha.VerifyUrl = "https://example.com/"
    return conf
}
```