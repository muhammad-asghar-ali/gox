package models

import "time"

type (
	UrlStore struct {
		ShortUrl       string    `bson:"short_url" json:"short_url"`
		LongUrl        string    `bson:"long_url" json:"long_url"`
		ExpriationDate time.Time `bson:"expiration_date" json:"expiration_date"`
		UserID         string    `bson:"user_id" json:"user_id"`
	}

	UrlStoreRepository struct{}
)

func (u *UrlStoreRepository) Insert(us *UrlStore) error {
	return nil
}
