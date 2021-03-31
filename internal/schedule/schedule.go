package schedule

import (
	"fmt"
	"share/internal/database"
	"time"
)

func init() {
	visitStatistics()
}

func visitStatistics() {
	if database.Redis != nil {
		go func() {
			for {
				now := time.Now()
				next := now.Add(time.Hour * 24)
				next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
				t := time.NewTimer(next.Sub(now))
				<-t.C
				// statistics
				// uv
				database.Redis.PFMerge("uv_total", "uv_total", fmt.Sprintf("uv:%s", now.Format("20060102")))
			}
		}()
	}
}
