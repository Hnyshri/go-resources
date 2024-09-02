package main

import (
	"fmt"

	auth "github.com/Hnyshri/go-resources/packages"
	user "github.com/Hnyshri/go-resources/packages/user"
)

func packageResource() {
	fmt.Println("check Scope function")
	auth.AuthLoginWithCredentials("shiryansh", "ceuih")
	session := auth.GetSession()
	fmt.Println(session)

	userData := user.User{
		Email: "shriyiansh@gmai",
		Name:  "shiryanh",
	}
	fmt.Println(userData)

}
