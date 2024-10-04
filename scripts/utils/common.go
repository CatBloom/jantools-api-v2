package utils

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

// uuidを作成し、ハイフンを取り除く関数
func GenerateUUIDWithoutHyphens() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(uuid.String(), "-", ""), nil
}

// 現在時間をjstで取得する関数
func NowJST() time.Time {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	return time.Now().In(jst)
}
