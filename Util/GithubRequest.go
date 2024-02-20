package util

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type githubUser struct {
	Name string `json:"login"`
}

// supply the authentication token to get the github username
func GetGithubUsername(token string) (string, error) {

	headers := http.Header{}
	headers.Set("Accept", "application/vnd.github+json")
	headers.Set("User-Agent", "boonsboos-genderstockmarket")
	headers.Set("X-Github-Api-Version", "2022-11-28")
	headers.Set("Authorization", "Bearer "+token)

	link, _ := url.Parse("https://api.github.com/user")

	req := http.Request{
		Method: "GET",
		URL:    link,
		Header: headers,
	}

	resp, err := http.DefaultClient.Do(&req)
	if err != nil {
		return "", err
	}

	var user githubUser

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &user)
	if err != nil || user.Name == "" {
		return "", errors.New("name is null")
	}

	return user.Name, nil
}

type authResponse struct {
	AccessToken string `json:"access_token"`
}

func GetUserAccessToken(code string) (string, error) {

	headers := http.Header{}
	headers.Set("Accept", "application/vnd.github+json")
	headers.Set("User-Agent", "boonsboos-genderstockmarket")
	headers.Set("X-Github-Api-Version", "2022-11-28")

	link, _ := url.Parse(
		"https://github.com/login/oauth/access_token?" +
			"client_id=" + Options.GithubID +
			"&client_secret=" + Options.GithubToken +
			"&code=" + code,
	)

	req := http.Request{
		Method: "POST",
		URL:    link,
		Header: headers,
	}

	resp, err := http.DefaultClient.Do(&req)
	if err != nil {
		return "", err
	}

	var auth authResponse

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &auth)
	if err != nil || auth.AccessToken == "" {
		return "", errors.New("token is null")
	}

	return auth.AccessToken, nil
}
