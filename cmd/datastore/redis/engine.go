package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/meysam81/tarzan/cmd/models"
	"github.com/redis/go-redis/v9"
)

func (r *redisClient) ListPosts(ctx context.Context, transformers ...func(*models.Post)) (*[]models.Post, error) {
	postIDs, err := r.r.ZRevRange(ctx, "posts:timeline", 0, -1).Result()
	if err != nil {
		return nil, err
	}

	var posts []models.Post
	for _, postID := range postIDs {
		postKey := fmt.Sprintf("post:%s", postID)
		postData, err := r.r.HGetAll(ctx, postKey).Result()
		if err != nil {
			continue
		}

		if len(postData) == 0 {
			continue
		}

		var p models.Post
		if idStr, ok := postData["id"]; ok {
			if idInt, err := strconv.Atoi(idStr); err == nil {
				p.ID = idInt
			}
		}
		p.Title = postData["title"]
		p.Content = postData["content"]
		p.AuthorEmail = postData["author_email"]
		p.AuthorName = postData["author_name"]

		if createdAtStr, exists := postData["created_at"]; exists {
			if createdAt, err := time.Parse(time.RFC3339, createdAtStr); err == nil {
				p.CreatedAt = createdAt.Format("2026-01-02")
			}
		}

		for _, t := range transformers {
			t(&p)
		}
		posts = append(posts, p)
	}

	return &posts, nil
}

func (r *redisClient) InsertEmail(ctx context.Context, email *models.EmailInsertDB) error {
	postID, err := r.r.Incr(ctx, "posts:counter").Result()
	if err != nil {
		return err
	}

	postIDStr := strconv.FormatInt(postID, 10)
	postKey := fmt.Sprintf("post:%s", postIDStr)
	now := time.Now()

	postData := map[string]interface{}{
		"id":           postIDStr,
		"title":        email.Subject,
		"content":      email.Content,
		"author_email": email.From,
		"author_name":  email.FromName,
		"created_at":   now.Format(time.RFC3339),
	}

	pipe := r.r.TxPipeline()
	pipe.HMSet(ctx, postKey, postData)
	pipe.ZAdd(ctx, "posts:timeline", redis.Z{
		Score:  float64(now.Unix()),
		Member: postIDStr,
	})

	_, err = pipe.Exec(ctx)
	return err
}

func (r *redisClient) FindAuthorizedSenderByEmail(ctx context.Context, email string) (*models.AuthorizedSender, error) {
	senderKey := fmt.Sprintf("authorized_sender:%s", email)
	senderData, err := r.r.HGetAll(ctx, senderKey).Result()
	if err != nil {
		return nil, err
	}

	if len(senderData) == 0 {
		return nil, fmt.Errorf("authorized sender not found")
	}

	var sender models.AuthorizedSender
	if idStr, ok := senderData["id"]; ok {
		if idInt, err := strconv.Atoi(idStr); err == nil {
			sender.ID = idInt
		}
	}
	sender.Email = senderData["email"]

	if createdAtStr, exists := senderData["created_at"]; exists {
		if createdAt, err := time.Parse(time.RFC3339, createdAtStr); err == nil {
			sender.CreatedAt = createdAt.Format("2026-01-02")
		}
	}

	return &sender, nil
}

func (r *redisClient) Close() error {
	return r.r.Close()
}
