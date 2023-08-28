package pkg

import (
	"context"
	"fmt"
	"log"

	"gogenggo/modules/secret"
)

func (p *PkgModules) SendMessage(ctx context.Context, chatFromID int64, message string) error {
	url := fmt.Sprintf("%s/sendMessage", secret.SecretObjects.Telegram.Host)
	headers := map[string]any{
		"Content-Type": "application/json",
	}
	params := map[string]any{
		"chat_id": chatFromID,
		"text":    message,
	}

	if _, err := p.Post(ctx, url, headers, params); err != nil {
		log.Println("[Pkg - SendMessage] Error doing post request, err:", err)
		return err
	}

	return nil
}

func (p *PkgModules) GetUpdates(ctx context.Context, offset, limit int64, timeout int, allowUpdates []string) ([]byte, error) {
	url := fmt.Sprintf("%s/getUpdates", secret.SecretObjects.Telegram.Host)
	headers := map[string]any{
		"Content-Type": "application/json",
	}
	params := map[string]any{
		"offset":  offset,
		"limit":   limit,
		"timeout": timeout,
		// "allowed_updates": allowUpdates,
	}

	bodyResponse, err := p.Post(ctx, url, headers, params)
	if err != nil {
		log.Println("[Pkg - GetUpdates] Error doing post request, err:", err)
		return nil, err
	}

	return bodyResponse, nil
}
