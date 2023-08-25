package auditlog

import (
	"context"
	"encoding/json"
	"fmt"
	domain "github.com/erdemakyildiz/audit-log-client"
	"github.com/nats-io/nats.go"
)

type UseCase struct {
	nc *nats.Conn
}

func New(nc *nats.Conn) *UseCase {
	return &UseCase{nc: nc}
}

func (uc *UseCase) PublishAuditLog(ctx context.Context, log domain.AuditLog) error {
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
