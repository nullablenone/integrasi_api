package utils

import "fmt"

func GetUserCacheKey(id string) string {
	return fmt.Sprintf("user:%s", id)
}
