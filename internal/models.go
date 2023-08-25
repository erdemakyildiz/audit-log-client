package internal

import "time"

type AuditLog struct {
	Data struct {
		Date           time.Time `json:"date"`
		Actor          Actor     `json:"actor"`
		Event          Event     `json:"event"`
		Description    string    `json:"description"`
		OrganizationID string    `json:"organization_id"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type Actor struct {
	Username  string `json:"username"`
	UserAgent string `json:"user_agent"`
	IP        string `json:"ip"`
}

type Event struct {
	EventCategory string `json:"event_category"`
	Name          string `json:"name"`
}
