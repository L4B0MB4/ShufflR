package manager

import (
	"net/url"

	"github.com/L4B0MB4/Musicfriends/pkg/database"
	"github.com/L4B0MB4/Musicfriends/pkg/models"
	"github.com/L4B0MB4/Musicfriends/pkg/utils"
)

type PersonalInfoManager struct {
	db *database.DatabaseConnection
}

func (m *PersonalInfoManager) SetUp(db *database.DatabaseConnection) {
	m.db = db
}

func (m PersonalInfoManager) UpsertProfile(user *models.CurrentUserProfile) {
	res := database.GetUserProfile(m.db, user.ID)
	if res == nil {
		database.InsertUserProfile(m.db, user)
	}
}

func (m PersonalInfoManager) GetUserProfile(userId string) *models.CurrentUserProfile {
	res := database.GetUserProfile(m.db, userId)
	if res == nil {
		return &models.CurrentUserProfile{}
	}
	return res
}

func (m PersonalInfoManager) GetOrReadTopTracks(userContext *models.UserContext) *models.TopTracksResponse {

	topTracks := database.GetTopTracks(m.db, userContext.ID)
	if topTracks == nil {

		query := url.Values{"time_range": []string{"short_term"}, "limit": []string{"50"}}
		topTracks = utils.SpotifyApiCall[models.TopTracksResponse]("/v1/me/top/tracks", userContext.AccessToken, "GET", query, nil)
		//database.SaveTopTracks(m.db, userContext.ID, topTracks)
	}
	return topTracks
}
