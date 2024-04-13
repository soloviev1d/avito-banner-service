package cache

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/soloviev1d/avito-banner-service/internal/database"
	"github.com/soloviev1d/avito-banner-service/internal/structs"
)

var BannerCache = struct {
	mu   sync.Mutex
	Data map[string]*structs.UniqueBanner
}{
	Data: make(map[string]*structs.UniqueBanner),
}

// token -> access level
var userCache map[string]int = make(map[string]int)

func AssembleCacheKey(tag, feature int) string {
	return fmt.Sprintf("tag%dfeature%d", tag, feature)
}

func StartCaching() {
	banners, err := database.GetAllBanners()
	if err != nil {
		log.Println("Failed to get banners:", err)
	}
	ticker := time.NewTicker(time.Minute * 5)
	go func() {
		for {
			<-ticker.C
			BannerCache.mu.Lock()
			for _, b := range banners {
				BannerCache.Data[AssembleCacheKey(b.TagId, b.FeatureId)] = b
			}
			BannerCache.mu.Unlock()
		}
	}()
}

func UserAccessLevel(token string) int {
	var (
		v  int
		ok bool
	)
	if v, ok = userCache[token]; ok {
		return v
	} else {
		v = database.GetUserAccess(token)
		if v > 0 {
			userCache[token] = v
		}
	}

	return v
}
