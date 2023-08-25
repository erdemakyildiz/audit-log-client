package client

import (
	"audit-log-client/internal/core/domain"
	"audit-log-client/internal/port/auditlog"
	"context"
	"errors"
	"fmt"
	"github.com/nats-io/nats.go"
	"sync"
)

var lock = &sync.Mutex{}

var clientInstance *UseCase

type UseCase struct {
	auditLog auditlog.UseCase
}

func New(nc *nats.Conn) *UseCase {
	return &UseCase{auditLog: nil}
}

func getInstance() *UseCase {
	if clientInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if clientInstance == nil {
			clientInstance = &UseCase{}
		}
	}

	return clientInstance
}

func (uc *UseCase) CreateAuditLog(ctx context.Context, log domain.AuditLog) (bool, error) {
	if log.Data.Event.EventCategory == "" {
		return false, errors.New("event category can not be empty")
	}

	if log.Data.Event.Name == "" {
		return false, errors.New("event name can not be empty")
	}

	err := uc.auditLog.PublishAuditLog(ctx, log)
	if err != nil {
		return false, fmt.Errorf("error sending audit log: %w", err)
	}

	return true, nil
}
