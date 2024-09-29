package main

import (

	"time"

	"github.com/google/uuid"
	"github.com/rss-aggregator/internal/database"
)

type User struct {
	ID   uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Name   string  `json:"name"`
	APIKey  string `json:"api_key"`
}


type Feed struct {
	ID   uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Name   string  `json:"name"`
	Url  string `json:"url"`
	UserID uuid.UUID `json:"user_id"`
}

type FeedFollow struct {
	ID   uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	UserID uuid.UUID `json:"user_id"`
	FeedID uuid.UUID `json:"feed_id"`
}

func databaseUsertoUser(dbUser database.User) User {
     return User{
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
		APIKey: dbUser.ApiKey,
	 }

}


func databaseFeedtoFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID: dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name: dbFeed.Name,
		Url: dbFeed.Url,
		UserID: dbFeed.UserID,
	}
}


func databaseFeedstoFeeds(dbFeed []database.Feed) []Feed {
	feeds := []Feed{}

	for _, dbFeed := range dbFeed {
		feeds = append(feeds, databaseFeedtoFeed(dbFeed))
	}

	return feeds
}

func databaseFeedFollowtoFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID: dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID: dbFeedFollow.UserID,
		FeedID: dbFeedFollow.FeedID,
	}
}



func databaseFeedFollowstoFeedFollows(dbFeedFollow []database.FeedFollow) []FeedFollow {
	

	feed_follows := []FeedFollow{}

	for _, dbFeeddbFeedFollow := range dbFeedFollow {
		feed_follows = append(feed_follows, databaseFeedFollowtoFeedFollow(dbFeeddbFeedFollow))
	}

	return feed_follows
}

type Post struct {
	ID          uuid.UUID      `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Title       string         `json:"title"`
	Description *string        `json:"description"`
	PublishedAt time.Time      `json:"published_at"`
	Url         string         `json:"url"`
	FeedID      uuid.UUID      `json:"feed_id"`
}