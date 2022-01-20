package models

import "time"

type ShortLink struct {
	CreatedAt  time.Time  `firestore:"created_at"`
	ExpiredAt  *time.Time `firestore:"expired_at"`
	Name       string     `firestore:"name"`
	URL        string     `firestore:"url"`
	Short      string     `firestore:"short"`
	ClickCount int64      `firestore:"click_count"`
}
