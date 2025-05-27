package models

import "time"

type CSPViolation struct {
	ID                 int       `json:"id" db:"id"`
	DocumentURI        string    `json:"document_uri" db:"document_uri"`
	Referrer           string    `json:"referrer" db:"referrer"`
	ViolatedDirective  string    `json:"violated_directive" db:"violated_directive"`
	EffectiveDirective string    `json:"effective_directive" db:"effective_directive"`
	OriginalPolicy     string    `json:"original_policy" db:"original_policy"`
	Disposition        string    `json:"disposition" db:"disposition"`
	BlockedURI         string    `json:"blocked_uri" db:"blocked_uri"`
	LineNumber         int       `json:"line_number" db:"line_number"`
	SourceFile         string    `json:"source_file" db:"source_file"`
	StatusCode         int       `json:"status_code" db:"status_code"`
	ScriptSample       string    `json:"script_sample" db:"script_sample"`
	UserAgent          string    `json:"user_agent" db:"user_agent"`
	IPAddress          string    `json:"ip_address" db:"ip_address"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
}

type CSPViolationInsertDB struct {
	DocumentURI        string `db:"document_uri"`
	Referrer           string `db:"referrer"`
	ViolatedDirective  string `db:"violated_directive"`
	EffectiveDirective string `db:"effective_directive"`
	OriginalPolicy     string `db:"original_policy"`
	Disposition        string `db:"disposition"`
	BlockedURI         string `db:"blocked_uri"`
	LineNumber         int    `db:"line_number"`
	SourceFile         string `db:"source_file"`
	StatusCode         int    `db:"status_code"`
	ScriptSample       string `db:"script_sample"`
	UserAgent          string `db:"user_agent"`
	IPAddress          string `db:"ip_address"`
}

type CSPViolationStats struct {
	TotalViolations       int            `json:"total_violations"`
	ViolationsByDirective map[string]int `json:"violations_by_directive"`
	ViolationsByURI       map[string]int `json:"violations_by_uri"`
	RecentViolations      []CSPViolation `json:"recent_violations"`
}
