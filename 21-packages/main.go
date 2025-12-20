package main

import (
	"fmt"

	"github.com/dharmendra-dxy/go-lang.git/21-packages/auth"
	"github.com/dharmendra-dxy/go-lang.git/21-packages/user"
)

// PACKAGES: can be used to re-use the code
func main(){
	auth.LoginWithCredentials("dharmendra", "pass@1234")

	session:=auth.GetSession()
	fmt.Println("session:", session);

	user:= user.User{
		Email: "email@example.com",
		Name: "Test User",
	}

	fmt.Println("USER DETAILS:", user.Email, user.Name)
}
