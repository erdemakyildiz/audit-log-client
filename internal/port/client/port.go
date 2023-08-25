package client

import (
	"audit-log-client/internal/core/domain"
	"context"
)

type UseCase interface {
	CreateAuditLog(ctx context.Context, log domain.AuditLog) (bool, error)
}
