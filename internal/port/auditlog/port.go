package auditlog

import (
	"context"
	domain "github.com/erdemakyildiz/audit-log-client"
)

type UseCase interface {
	PublishAuditLog(ctx context.Context, log domain.AuditLog) error
}
