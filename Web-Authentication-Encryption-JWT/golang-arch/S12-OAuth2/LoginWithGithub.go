package main

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var githubConnections map[string]string

type githubResponse struct {
	Data struct {
		Viewer struct {
			ID string `json:"id"`
		} `json: "viewer"`
	} `json: "data"`
}

var githubOauthConfig = &oauth2.Config{
	ClientID:     "233ac783feea25fadb1f",
	ClientSecret: "7d357c9ddd38cd5fa821a9d60431287f4d5538a8",
	Endpoint:     github.Endpoint,
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/oauth/github", startGithubOauth)
	http.HandleFunc("/oauth2/receive", completeGithubOauth)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<title>Document</title>
			<meta name="description" content="">
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<link rel="stylesheet" href="">
		</head>
		<body>
			<form action="/oauth/github" method="POST">
				<input type="submit" value="Login with Github">
			</form>
		</body>
	</html>`)
}

func startGithubOauth(w http.ResponseWriter, r *http.Request) {
	redirectURL := githubOauthConfig.AuthCodeURL("0000")
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func completeGithubOauth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("something")
	code := r.FormValue("code")
	state := r.FormValue("state")

	if state != "0000" {
		http.Error(w, "State is incorrect", http.StatusBadRequest)
		return
	}

	token, err := githubOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Couldn't login", http.StatusInternalServerError)
		return
	}

	ts := githubOauthConfig.TokenSource(r.Context(), token)
	client := oauth2.NewClient(r.Context(), ts)

	requestBody := strings.NewReader(`{"query": "query {viewer {id}}"}`)
	resp, err := client.Post("https://api.github.com/graphql", "application/json", requestBody)
	if err != nil {
		http.Error(w, "Couldn't get user", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "couldn't read github information", http.StatusInternalServerError)
		return
	}

	log.Println(string(bs))

	fmt.Fprint(w, `<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<title>Document</title>
			<meta name="description" content="">
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<link rel="stylesheet" href="">
		</head>
		<body>
			<h1>This is code that coming from github: %s </h1>
		</body>
	</html>`,code)
	// var gr githubResponse
	// err = json.NewDecoder(resp.Body).Decode(&gr)
	// if err != nil {
	// 	http.Error(w, "GIthub invalid response", http.StatusInternalServerError)
	// 	return
	// }

	// githubID := gr.Data.Viewer.ID
	// userID, ok := githubConnections[githubID]
	// if !ok {
	// 	// New user create account
	// 	return
	// }
	// fmt.Println(userID)
}
