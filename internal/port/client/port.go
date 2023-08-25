package client

import (
	"context"
	"github.com/erdemakyildiz/audit-log-client/internal"
)

type UseCase interface {
	CreateAuditLog(ctx context.Context, log internal.AuditLog) (bool, error)
}
