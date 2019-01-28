package main

import (
	"fmt"

	"github.com/abramovic/goinsta"
	"github.com/abramovic/goinsta/_examples/shared"
	"github.com/abramovic/goinsta/store"
)

func main() {

	var (
		insta *goinsta.Instagram

		// instagram username/password credentials
		username string = "USERNAME"
		password string = "PASSWORD"

		// 2FA credentials.
		identifier string = "IDENTIFIER_CODE" // provided by insta.Login() if 2FA is required
		authCode   string = "123 456"         // provided by user for 2FA
	)

	//	fmt.Println("Try to login a user. This will fail if a user has 2FA enabled")
	insta = loginUser(username, password)

	// demo trying to login a user via 2FA
	insta = loginUser2FA(username, password, identifier, authCode)
	// print the exported string to console
	// that way we can test using
	exportedString, err := store.ExportString(insta)
	shared.FatalOnErr(err)

	// now demo the login using the exported cookie string
	loginUsingCookieString(exportedString)

	// print the cookie string to console
	fmt.Println("\n\n\nCookie String: ")
	fmt.Println(exportedString)
}

func loginUser(username, password string) *goinsta.Instagram {
	insta := goinsta.New(username, password, nil)
	// try to login the user
	identifier_code, err := insta.Login()
	if identifier_code != "" {
		fmt.Println("\n\n")
		fmt.Println("two factor authorization is required. you will need to use this code when the user provides their 6 digit code: ")
		fmt.Println(identifier_code)
		fmt.Println("\n\n")
		return nil
	}
	shared.FatalOnErr(err)
	return insta
}

func loginUser2FA(username, password, identifier, code string) *goinsta.Instagram {
	// try to login the user via username/password
	insta := goinsta.New(username, password, nil)

	err := insta.TwoFactorLogin(identifier, code)
	shared.FatalOnErr(err)

	return insta
}

func loginUsingCookieString(cookieString string) {
	insta, err := store.ImportString(cookieString)
	shared.FatalOnErr(err)

	instaPage, err := insta.GetUserByUsername("thedodo")
	fmt.Println("GetUserByUsername", err)
	shared.FatalOnErr(err)
	fmt.Printf("%+v\n", instaPage)
}
