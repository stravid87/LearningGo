// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password1234`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	newpassword := `password1234`
	err = bcrypt.CompareHashAndPassword(bs, []byte(newpassword))
	if err != nil {
		fmt.Println("You are not Log In")
		return
	}
	fmt.Println("You are logged IN")

}
