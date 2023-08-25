package client

import (
	"context"
	"github/erdemakyildiz/audit-log-client/internal/core/domain"
)

type UseCase interface {
	CreateAuditLog(ctx context.Context, log domain.AuditLog) (bool, error)
}
