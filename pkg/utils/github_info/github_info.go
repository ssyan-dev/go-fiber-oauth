package github_info

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

type UserEmail struct {
	Email    string `json:"email"`
	Primary  bool   `json:"primary"`
	Verified bool   `json:"verified"`
}

func GetUserInfo(ctx context.Context, cfg *oauth2.Config, token *oauth2.Token) (string, error) {
	client := cfg.Client(ctx, token)

	res, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		return "", fmt.Errorf("failed to get emails: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != fiber.StatusOK {
		return "", fmt.Errorf("error from github: %s", res.Status)
	}

	var emails []UserEmail
	if err := json.NewDecoder(res.Body).Decode(&emails); err != nil {
		return "", fmt.Errorf("failed to decode: %v", err)
	}

	for _, email := range emails {
		if email.Primary {
			return email.Email, nil
		}
	}

	return "", fmt.Errorf("primary email not found")
}
