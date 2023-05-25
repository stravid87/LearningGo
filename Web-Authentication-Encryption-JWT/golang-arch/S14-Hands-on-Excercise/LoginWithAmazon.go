package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/amazon"
)

type user struct {
	password []byte
	First    string
}

var oauth = &oauth2.Config{
	ClientID:     "amzn1.application-oa2-client.40c73345c7bf49aeba2c54c2557310b2",
	ClientSecret: "amzn1.oa2-cs.v1.c6f039950cf22a84a46703622fd0765f937a41129bad5761dcadfaa905db6617",
	Endpoint:     amazon.Endpoint,
	RedirectURL:  "http://localhost:8080/oauth/amazon/receive",
	Scopes:       []string{"postal_code"},
}

// key is email, value password
var db = map[string]user{
	"test@example.com": user{
		First: "testFirstName",
	},
}

var sessions = map[string]string{}

var oauthExp = map[string]time.Time{}

var oauthConnections = map[string]string{}

var key = []byte("my secret key 007 james bond rule")

type customClaims struct {
	jwt.StandardClaims
	SID string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/oauth/amazon/login", oAmazonLogin)
	http.HandleFunc("/oauth/amazon/receive", oAmazonReceive)
	http.HandleFunc("/partial-register", partialRegister)
	http.HandleFunc("/oauth/amazon/register", oAmazonRegister)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("sessionID")
	if err != nil {
		c = &http.Cookie{
			Name:  "sessionID",
			Value: "",
		}
	}

	sID, err := parseToken(c.Value)
	if err != nil {
		log.Println("index parseToken", err)
	}

	var e string
	if sID != "" {
		e = sessions[sID]
	}

	var f string
	if user, ok := db[e]; ok {
		f = user.First
	}

	errMsg := r.FormValue("msg")
	fmt.Fprintf(w, `<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<title></title>
			<meta name="description" content="">
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<link rel="stylesheet" href="">
		</head>
		<body>
			<h1>YOU HAVE SESSION, HERE IS YOUR NAME: %s </h1>
			<h1>YOU HAVE SESSION, HERE IS YOUR EMAIL: %s </h1>
			<h1>IF THERE IS ANY MESSAGE FOR YOU, HERE IT IS: %s </h1>
			<h1>LOG IN with AMAZON</h1>
			<form action="/oauth/amazon/login" method="POST">
				<input type="submit" value="Login with Amazon">
			</form>
			<h1>LOGOUT</h1>
			<form action="/logout" method="POST">
				<input type="submit" value="Login with Amazon">
			</form>
		</body>
	</html>`, e, f, errMsg)
}

func createToken(sid string) (string, error) {
	cc := customClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
		SID: sid,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cc)
	st, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("couldn't sign token in createToken %w", err)
	}
	return st, nil
}

func parseToken(ss string) (string, error) {
	token, err := jwt.ParseWithClaims(ss, &customClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("parseWithClaims couldn't different algorithm")
		}
		return key, nil
	})

	if err != nil {
		return "", fmt.Errorf("couldn't ParseWithClaims in parseToken %w", err)
	}

	if !token.Valid {
		return "", fmt.Errorf("token not valid in parseToken")
	}

	return token.Claims.(*customClaims).SID, nil
}

func logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	c, err := r.Cookie("sessionID")
	if err != nil {
		c = &http.Cookie{
			Name:  "sessionID",
			Value: "",
		}
	}

	sID, err := parseToken(c.Value)
	if err != nil {
		log.Println("index parseToken", err)
	}

	delete(sessions, sID)

	c.MaxAge = -1

	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func createSession(e string, w http.ResponseWriter) error {
	sUUID := uuid.New().String()
	sessions[sUUID] = e
	token, err := createToken(sUUID)
	if err != nil {
		return fmt.Errorf("Couln't createToken un Create session %w", err)
	}

	c := http.Cookie{
		Name:  "sessionID",
		Value: token,
		Path: "/",
	}

	http.SetCookie(w, &c)
	return nil
}

func oAmazonLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	id := uuid.New().String()
	oauthExp[id] = time.Now().Add(time.Hour)

	// here we redirect to the Amazon at the AuthURL
	http.Redirect(w, r, oauth.AuthCodeURL(id), http.StatusSeeOther)
}

func oAmazonReceive(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state == "" {
		msg := url.QueryEscape("state was empty in oAmazonReceive")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	code := r.FormValue("code")
	if code == "" {
		msg := url.QueryEscape("code was empty in oAmazonReceive")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	expT := oauthExp[state]
	if time.Now().After(expT) {
		msg := url.QueryEscape("oauth took too long time.Now.after")
		http.Redirect(w, r, "/msg?="+msg, http.StatusSeeOther)
		return
	}

	t, err := oauth.Exchange(r.Context(), code)
	if err != nil {
		msg := url.QueryEscape("couldn't do auth exchange" + err.Error())
		http.Redirect(w, r, "/msg?="+msg, http.StatusSeeOther)
		return
	}

	ts := oauth.TokenSource(r.Context(), t)
	c := oauth2.NewClient(r.Context(), ts)

	resp, err := c.Get("https://api.amazon.com/user/profile")
	if err != nil {
		msg := url.QueryEscape("couldn't get at amazon: " + err.Error())
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		msg := url.QueryEscape("not 200 response")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	// fmt.Fprint(w, string(bs))
	// io.WriteString(w, string(bs))

	
	type profileResponse struct {
		Email  string `json:"email"`
		Name   string `json:"name"`
		UserID string `json:"user_id"`
	}

	var pr profileResponse

	err = json.NewDecoder(resp.Body).Decode(&pr)
	if err != nil {
		msg := url.QueryEscape("not able to decode json response")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	eml, ok := oauthConnections[pr.UserID]

	if !ok {
		// not regisstred at out site yet with amazon
		st, err := createToken(pr.UserID)
		if err != nil {
			log.Println("couldn't createToken in oAmazonReceive", err)
			msg := url.QueryEscape("our server disn't get enough lunch and is not wokring")
			http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
			return
		}

		uv := url.Values{}
		uv.Add("sst", st)	
		uv.Add("name", pr.Name)	
		uv.Add("email", pr.Email)	

		http.Redirect(w, r, "/partial-register?"+uv.Encode(), http.StatusSeeOther)
		return
	}

	err = createSession(eml, w)
	if err != nil {
		log.Println("couldn't createSession  in oAmazonReceive", err)
		msg := url.QueryEscape("our server disn't get enough lunch and is not wokring")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	msg := url.QueryEscape("you logged in " + eml)
	http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
}

func partialRegister(w http.ResponseWriter, r *http.Request) {
	sst := r.FormValue("sst")
	name := r.FormValue("name")
	email := r.FormValue("email")

	if sst == "" {
		log.Println("couldn't get sst in partialRegister")
		msg := url.QueryEscape("our server disn't get enough lunch and is not wokring")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	fmt.Fprintf(w, `<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<title></title>
			<meta name="description" content="">
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<link rel="stylesheet" href="">
		</head>
		<body>
			<form action="/oauth/amazon/register" method="POST">

				<label for="firstName">First</label>
				<input type="text" name="first" id="Firstname" value="%s">

				<label for="Email">Email</label>
				<input type="text" name="email" id="Email" value="%s">

				<input type="hidden" name="oauthID" value="%s">

				<input type="submit">
			</form>
		</body>
	</html>`, name, email, sst)
}

func oAmazonRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		msg := url.QueryEscape("your method was not post")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	f := r.FormValue("first")
	e := r.FormValue("email")
	oauthID := r.FormValue("oauthID")

	if f == "" {
		msg := url.QueryEscape("your method needs not to be empty")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	if e == "" {
		msg := url.QueryEscape("your method needs not to be empty")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	if oauthID == "" {
		log.Println("oauthID came through as empty at oAmazonRegister - MAYBE BAN BOT PRANKSTAR IP ADDRESS")
		msg := url.QueryEscape("your oauthID needs not to be empty")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	amazonUID, err := parseToken(oauthID)
	if err != nil {
		log.Println("ParseToekn at oAmazonRegister didn't parse")
		msg := url.QueryEscape("there was an issue. send us money so we can fix it")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	db[e] = user{
		First: f,
	}

	oauthConnections[amazonUID] = e

	err = createSession(e, w)
	if err != nil {
		log.Println("couldn't CreateSession in oAmazonRegister", err)
		msg := url.QueryEscape("there was an issue. send us money so we can fix it")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}