package main

import (
	"flag"
	"os"
)

func initConfig() *Config {
	cfg := &Config{
		Port: 1188,
	}

	// Port flag
	flag.IntVar(&cfg.Port, "port", cfg.Port, "set up the port to listen on")
	flag.IntVar(&cfg.Port, "p", cfg.Port, "set up the port to listen on")

	// DL Session flag
	flag.StringVar(&cfg.DlSession, "s", "", "set the dl-session for /v1/translate endpoint")
	if cfg.DlSession == "" {
		if dlSession, ok := os.LookupEnv("DL_SESSION"); ok {
			cfg.DlSession = dlSession
		}
	}

	// Access token flag
	flag.StringVar(&cfg.Token, "token", "", "set the access token for /translate endpoint")
	if cfg.Token == "" {
		if token, ok := os.LookupEnv("TOKEN"); ok {
			cfg.Token = token
		}
	}

	// DeepL Official Authentication key flag
	flag.StringVar(&cfg.AuthKey, "authkey", "", "The authentication key for DeepL API")
	if cfg.AuthKey == "" {
		if authKey, ok := os.LookupEnv("AUTHKEY"); ok {
			cfg.AuthKey = authKey
		}
	}

	// HTTP Proxy flag
	flag.StringVar(&cfg.Proxy, "proxy", "", "set the proxy URL for HTTP requests")
	if cfg.Proxy == "" {
		if proxy, ok := os.LookupEnv("PROXY"); ok {
			cfg.Proxy = proxy
		}
	}

	flag.Parse()
	return cfg
}
