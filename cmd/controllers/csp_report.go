package controllers

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/meysam81/tarzan/cmd/models"
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

	userAgent := r.Header.Get("User-Agent")
	ipAddress := getClientIP(r)

	violationRecord := &models.CSPViolationInsertDB{
		DocumentURI:        report.CSPReport.DocumentURI,
		Referrer:           report.CSPReport.Referrer,
		ViolatedDirective:  report.CSPReport.ViolatedDirective,
		EffectiveDirective: report.CSPReport.EffectiveDirective,
		OriginalPolicy:     report.CSPReport.OriginalPolicy,
		Disposition:        report.CSPReport.Disposition,
		BlockedURI:         report.CSPReport.BlockedURI,
		LineNumber:         report.CSPReport.LineNumber,
		SourceFile:         report.CSPReport.SourceFile,
		StatusCode:         report.CSPReport.StatusCode,
		ScriptSample:       report.CSPReport.ScriptSample,
		UserAgent:          userAgent,
		IPAddress:          ipAddress,
	}

	if err := a.DS.InsertCSPViolation(r.Context(), violationRecord); err != nil {
		log.Printf("Error storing CSP violation: %v", err)
	}

	log.Printf("CSP Violation: Blocked URI: %s, Violated Directive: %s, Document URI: %s, IP: %s",
		report.CSPReport.BlockedURI,
		report.CSPReport.ViolatedDirective,
		report.CSPReport.DocumentURI,
		ipAddress)

	w.WriteHeader(http.StatusNoContent)
}

func getClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		ips := strings.Split(xForwardedFor, ",")
		return strings.TrimSpace(ips[0])
	}

	xRealIP := r.Header.Get("X-Real-IP")
	if xRealIP != "" {
		return xRealIP
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}
