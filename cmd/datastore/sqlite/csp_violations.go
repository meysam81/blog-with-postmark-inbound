package sqlite

import (
	"context"

	"github.com/meysam81/tarzan/cmd/models"
)

func (s *Sqlite) InsertCSPViolation(ctx context.Context, violation *models.CSPViolationInsertDB) error {
	query := `
		INSERT INTO csp_violations (
			document_uri, referrer, violated_directive, effective_directive,
			original_policy, disposition, blocked_uri, line_number,
			source_file, status_code, script_sample, user_agent, ip_address
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := s.DB.ExecContext(ctx, query,
		violation.DocumentURI, violation.Referrer, violation.ViolatedDirective,
		violation.EffectiveDirective, violation.OriginalPolicy, violation.Disposition,
		violation.BlockedURI, violation.LineNumber, violation.SourceFile,
		violation.StatusCode, violation.ScriptSample, violation.UserAgent, violation.IPAddress,
	)

	return err
}
