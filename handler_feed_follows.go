package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User){
	
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	

	params := parameters{}
	err := decoder.Decode(&params)

	fmt.Println(params.FeedID)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return 
	}

	feed_follow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID: params.FeedID,
		UserID: user.ID,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Coudnt create feed follow: %v", err))
		return 
	}
	
	respondWithJSON(w, 201, databaseFeedFollowtoFeedFollow(feed_follow))
}





