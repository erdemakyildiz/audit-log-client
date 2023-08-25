package auditlog

import (
	"context"
	"github.com/erdemakyildiz/audit-log-client/internal/core/domain"
)

type UseCase interface {
	PublishAuditLog(ctx context.Context, log domain.AuditLog) error
}
