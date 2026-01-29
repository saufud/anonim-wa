package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type FonteSender struct {
	APIKey string
}

func NewFonte(apiKey string) *FonteSender {
	return &FonteSender{APIKey: apiKey}
}

func (f *FonteSender) Send(to string, message string) error {
	body := map[string]string{
		"target":      to,
		"message":     message,
		"countryCode": "62",
	}

	b, _ := json.Marshal(body)

	req, err := http.NewRequest(
		"POST",
		"https://api.fonnte.com/send",
		bytes.NewBuffer(b),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", f.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("FONNTE RESPONSE:", string(respBody))

	if resp.StatusCode >= 400 {
		return fmt.Errorf("fonnte error: %s", resp.Status)
	}

	return nil
}
