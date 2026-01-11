package utils

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"

	"gobee/ent/post"
	"gobee/internal/database"
)

// createAt time.Now().Unix()
func GenerateSlug(title string, createAt int64) (string, error) {

	hashInput := fmt.Sprintf("%s%d", title, createAt)
	hash := md5.Sum([]byte(hashInput))
	hashStr := hex.EncodeToString(hash[:])[:12]

	num, err := strconv.ParseInt(hashStr, 16, 64)
	if err != nil {
		return "", err
	}

	slug := strconv.FormatInt(num, 32)

	return slug, nil
}

func checkSlugExists(ctx context.Context, slug string, postID int) (bool, error) {
	client := database.DB
	count, err := client.Post.Query().
		Where(post.Slug(slug)).
		Where(post.IDNEQ(postID)).
		Count(ctx)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
