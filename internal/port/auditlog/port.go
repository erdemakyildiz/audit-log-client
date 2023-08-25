package auditlog

import (
	"context"
	"github.com/erdemakyildiz/audit-log-client/internal"
)

type UseCase interface {
	PublishAuditLog(ctx context.Context, log internal.AuditLog) error
}
