package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MetaSender struct {
	Token   string
	PhoneID string
}

func NewMeta(token, phoneID string) *MetaSender {
	return &MetaSender{Token: token, PhoneID: phoneID}
}

func (m *MetaSender) Send(to, message string) error {
	url := fmt.Sprintf(
		"https://graph.facebook.com/v19.0/%s/messages",
		m.PhoneID,
	)

	payload := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":   to,
		"type": "text",
		"text": map[string]string{"body": message},
	}

	b, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Authorization", "Bearer "+m.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("meta error: %s", resp.Status)
	}
	return nil
}
