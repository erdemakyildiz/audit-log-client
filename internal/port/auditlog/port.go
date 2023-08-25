package auditlog

import (
	"audit-log-client/internal/core/domain"
	"context"
)

type UseCase interface {
	PublishAuditLog(ctx context.Context, log domain.AuditLog) error
}
