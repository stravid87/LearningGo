package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/submit", bar)
	http.ListenAndServe(":8080", nil)
}

type myClaims struct {
	jwt.StandardClaims
	Email string
}

const myKey = "I love thursday when it rains"

func getJWT(msg string) (string, error) {
	claims := myClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
		Email: msg,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	ss, err := token.SignedString([]byte(myKey))
	if err != nil {
		return "", fmt.Errorf("couldn't SignedSting %w", err)
	}
	return ss, nil
}

func bar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.FormValue("emailThing")
	log.Println("email")
	if email == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	ss, err := getJWT(email)
	if err != nil {
		http.Error(w, "couldn't getJWT", http.StatusInternalServerError)
		return
	}

	c := http.Cookie{
		Name:  "session",
		Value: ss,
	}

	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func foo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{}
	}

	ss := c.Value
	afterVerificationToken, err := jwt.ParseWithClaims(ss, &myClaims{},
		func(beforeVeritificationToken *jwt.Token) (interface{}, error) {
			if beforeVeritificationToken.Method.Alg() != jwt.SigningMethodES256.Alg() {
				return nil, fmt.Errorf("SOMEONE TRIED TO HACK changed sigining method")
			}
			return []byte(myKey), nil
		})

	// StandardClaims has the ...
	// Valid() error
	// ... method which means it implements the Claims interface ...
	//
	/*
		type Claims interface {
			Valid() error
		}
	*/
	//
	// ... when you ParseClaims as with "ParseWithClaims" ...
	// the Valid() method gets run
	// ... and if all is well, then returns no "error" and
	// type TOKEN which has a field VALID will be true

	isEqual := err == nil && afterVerificationToken.Valid

	message := "Not logged in"
	if isEqual {
		message = "Logged in"
		claims := afterVerificationToken.Claims.(*myClaims)
		fmt.Println(claims.Email)
		fmt.Println(claims.ExpiresAt)
	}

	html := `<!DOCTYPE html>
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
			<p>` + c.Value + `</p>
			<p>` + message + `</p>
			<form action="/submit" method="post">
				<input type="email" name="emailThing">
				<input type="submit">
			</form>
		</body>
	</html>`
	io.WriteString(w, html)
}
