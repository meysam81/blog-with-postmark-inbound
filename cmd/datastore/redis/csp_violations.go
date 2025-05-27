package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/meysam81/tarzan/cmd/models"
)

func (r *redisClient) InsertCSPViolation(ctx context.Context, violation *models.CSPViolationInsertDB) error {
	violationData := models.CSPViolation{
		DocumentURI:        violation.DocumentURI,
		Referrer:           violation.Referrer,
		ViolatedDirective:  violation.ViolatedDirective,
		EffectiveDirective: violation.EffectiveDirective,
		OriginalPolicy:     violation.OriginalPolicy,
		Disposition:        violation.Disposition,
		BlockedURI:         violation.BlockedURI,
		LineNumber:         violation.LineNumber,
		SourceFile:         violation.SourceFile,
		StatusCode:         violation.StatusCode,
		ScriptSample:       violation.ScriptSample,
		UserAgent:          violation.UserAgent,
		IPAddress:          violation.IPAddress,
		CreatedAt:          time.Now(),
	}

	data, err := json.Marshal(violationData)
	if err != nil {
		return err
	}

	score := float64(time.Now().Unix())

	return r.r.ZAdd(ctx, r.cspCiolationsKey, struct {
		Score  float64
		Member interface{}
	}{Score: score, Member: data}).Err()
}
