package audit

import (
	"context"
	"fmt"
	"github.com/erdemakyildiz/audit-log-client/internal/core/domain"
	"github.com/erdemakyildiz/audit-log-client/internal/usecase/auditlog"
	"github.com/nats-io/nats.go"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

type Client struct {
	useCase auditlog.UseCase
}

var client *Client

func GetInstance(nc *nats.Conn) *Client {
	if client == nil {
		lock.Lock()
		defer lock.Unlock()
		if client == nil {
			useCase := auditlog.New(nc)
			client = &Client{useCase: *useCase}
		}
	}

	return client
}

func (c *Client) CreateAuditLog(organizationID, username, userAgent, ip, eventName, eventCategory, eventDescription string) error {
	log := domain.AuditLog{
		Data: domain.Data{
			Date: time.Now(),
			Actor: domain.Actor{
				Username:  username,
				UserAgent: userAgent,
				IP:        ip,
			},
			Event: domain.Event{
				EventCategory: eventCategory,
				Name:          eventName,
			},
			Description:    eventDescription,
			OrganizationID: organizationID,
		},
	}

	err := c.useCase.PublishAuditLog(context.Background(), log)
	if err != nil {
		return fmt.Errorf("error creating audit log: %w", err)
	}

	return nil
}
