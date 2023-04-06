package models

import "time"

type Session struct {
	UserID int64
}

type TokenResponse struct {
	Token   string `json:"token"`
	Expires time.Time
}
