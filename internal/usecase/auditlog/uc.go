package auditlog

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/erdemakyildiz/audit-log-client/internal/core/domain"
	"github.com/nats-io/nats.go"
)

type UseCase struct {
	nc *nats.Conn
}

func New(nc *nats.Conn) *UseCase {
	return &UseCase{nc: nc}
}

func (uc *UseCase) PublishAuditLog(ctx context.Context, log domain.AuditLog) error {
	if log.Data.Event.EventCategory == "" {
		return errors.New("event category can not be empty")
	}

	if log.Data.Event.Name == "" {
		return errors.New("event name can not be empty")
	}

	auditLog, err := json.Marshal(log)
	if err != nil {
		return fmt.Errorf("error marshal audit log: %w", err)
	}

	err = uc.nc.Publish("AUDIT_LOG.create-log", auditLog)
	if err != nil {
		return fmt.Errorf("error publish audit log: %w", err)
	}

	return nil
}
