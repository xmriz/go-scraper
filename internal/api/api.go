package api

import (
	"encoding/json"
	"fmt"

	"github.com/xmriz/go-scraper/internal/model"
	"github.com/xmriz/go-scraper/internal/utils"
)

const baseURL = "https://dummyapi.io/data/v1"

// FetchUsers retrieves a list of users from the API
func FetchUsers(page int, appID string) ([]model.UserPreview, error) {
	url := fmt.Sprintf("%s/user?page=%d", baseURL, page)
	resp, err := utils.MakeRequest(url, appID)
	if err != nil {
		return nil, err
	}

	var apiResponse model.APIResponseUser
	if err := json.Unmarshal(resp, &apiResponse); err != nil {
		return nil, err
	}

	return apiResponse.Data, nil
}

// FetchUserDetail retrieves detailed information about a specific user
func FetchUserDetail(userID string, appID string) (model.UserFull, error) {
	url := fmt.Sprintf("%s/user/%s", baseURL, userID)
	resp, err := utils.MakeRequest(url, appID)
	if err != nil {
		return model.UserFull{}, err
	}

	var userDetail model.UserFull
	if err := json.Unmarshal(resp, &userDetail); err != nil {
		return model.UserFull{}, err
	}

	return userDetail, nil
}

// FetchPosts retrieves a list of posts from the API
func FetchPosts(page int, appID string) ([]model.Post, error) {
	url := fmt.Sprintf("%s/post?page=%d", baseURL, page)
	resp, err := utils.MakeRequest(url, appID)
	if err != nil {
		return nil, err
	}

	var apiResponse model.APIResponsePost
	if err := json.Unmarshal(resp, &apiResponse); err != nil {
		return nil, err
	}

	return apiResponse.Data, nil
}
