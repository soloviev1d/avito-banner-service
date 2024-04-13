package cache

import "fmt"

func GenCacheKey(tag, feature int) string {
	return fmt.Sprintf("tag%dfeature%d", tag, feature)
}
