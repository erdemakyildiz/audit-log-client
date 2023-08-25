package client

import (
	"context"
	domain "github.com/erdemakyildiz/audit-log-client"
)

type UseCase interface {
	CreateAuditLog(ctx context.Context, log domain.AuditLog) (bool, error)
}
