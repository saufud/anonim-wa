package config

import "os"

type Config struct {
	Port        string
	Provider    string // fonte | meta
	FonteAPIKey string

	MetaToken   string
	MetaPhoneID string

	Watermark string
}

func Load() *Config {
	return &Config{
		Port:        get("PORT", "8080"),
		Provider:    get("PROVIDER", "fonte"),
		FonteAPIKey: os.Getenv("FONTE_API_KEY"),
		MetaToken:   os.Getenv("WHATSAPP_TOKEN"),
		MetaPhoneID: os.Getenv("WHATSAPP_PHONE_ID"),
		Watermark:   get("WATERMARK", "webpesananonim"),
	}
}

func get(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
