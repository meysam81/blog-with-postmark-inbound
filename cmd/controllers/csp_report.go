package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

type CSPReport struct {
	CSPReport struct {
		DocumentURI        string `json:"document-uri"`
		Referrer           string `json:"referrer"`
		ViolatedDirective  string `json:"violated-directive"`
		EffectiveDirective string `json:"effective-directive"`
		OriginalPolicy     string `json:"original-policy"`
		Disposition        string `json:"disposition"`
		BlockedURI         string `json:"blocked-uri"`
		LineNumber         int    `json:"line-number"`
		SourceFile         string `json:"source-file"`
		StatusCode         int    `json:"status-code"`
		ScriptSample       string `json:"script-sample"`
	} `json:"csp-report"`
}

func (a *AppState) CSPViolationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var report CSPReport
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&report); err != nil {
		log.Printf("Error decoding CSP report: %v", err)
		http.Error(w, "Invalid report format", http.StatusBadRequest)
		return
	}

	log.Printf("CSP Violation: Blocked URI: %s, Violated Directive: %s, Document URI: %s",
		report.CSPReport.BlockedURI,
		report.CSPReport.ViolatedDirective,
		report.CSPReport.DocumentURI)

	w.WriteHeader(http.StatusNoContent)
}
