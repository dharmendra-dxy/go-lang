package auth

import "fmt"

// PRACTICE: func name is written with Capital letter "L": means we are exporting it 
// DEFINING SCOPE: "L" capital means can be used in project, else can be used in same file only
func LoginWithCredentials(username string, password string){
	fmt.Println("USER LOGGED IN WITH CREDENTIALS:", username, password)
}