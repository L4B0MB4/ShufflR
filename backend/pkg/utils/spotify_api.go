package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/L4B0MB4/Musicfriends/pkg/models"
	"github.com/rs/zerolog/log"
)

func SpotifyApiCall[T any](path string, accessToken string, method string, query url.Values, body []byte) *T {
	req := http.Request{}
	spotifyUrlStr, err := url.JoinPath("https://api.spotify.com/", path)
	if err != nil {
		log.Error().Err(err).Msg("Could not build the api url")
		return nil
	}
	spotifyUrl, err := url.Parse(spotifyUrlStr)
	if err != nil {
		log.Error().Err(err).Msg("Could not parse the api url")
		return nil
	}
	if query != nil {
		spotifyUrl.RawQuery = query.Encode()
	}
	req.URL = spotifyUrl
	req.Header = http.Header{}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Method = method
	if body != nil {
		req.Body = io.NopCloser(bytes.NewReader(body))
	}
	client := &http.Client{} // at some point improve instance numbers

	res, err := client.Do(&req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to do request to spotify")
		return nil
	}
	b, err := io.ReadAll(res.Body)
	//log.Debug().Str("body", string(b)).Msg("Body for call to " + spotifyUrl.String())
	if err != nil {
		log.Error().Err(err).Msg("Failed to do read spotify response body")
		return nil
	}
	if !(res.StatusCode >= 200 && res.StatusCode < 400) {
		return nil
	}
	var responseModel T
	json.Unmarshal(b, &responseModel)
	log.Info().Str("json-response", string(b)).Str("url", spotifyUrlStr).Msg("Response body of current request")
	return &responseModel
}

func GetAccessToken(code string, redirect_uri string, clientId string, clientSecret string) *models.TokenResponse {
	req := http.Request{}
	spotifyUrl, _ := url.Parse("https://accounts.spotify.com/api/token")
	req.URL = spotifyUrl
	form := url.Values{}
	form.Add("code", code)
	form.Add("redirect_uri", redirect_uri)
	form.Add("grant_type", "authorization_code")
	req.Body = io.NopCloser(bytes.NewBufferString(form.Encode()))
	req.Header = http.Header{}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(clientId+":"+clientSecret)))
	req.Method = "POST"
	client := &http.Client{}
	res, err := client.Do(&req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to do request to spotify")
		return nil
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error().Err(err).Msg("Failed to do read spotify response body")
		return nil
	}
	if !(res.StatusCode >= 200 && res.StatusCode < 400) {
		return nil
	}

	var tokenRes models.TokenResponse
	json.Unmarshal(b, &tokenRes)
	log.Info().Interface("token-response", tokenRes).Msg("Token response from current user who logged in")
	return &tokenRes
}
