package goinsta

import (
	"net/http"
	"net/http/cookiejar"

	response "github.com/abramovic/goinsta/response"
)

type Informations struct {
	Username  string
	Password  string
	DeviceID  string
	UUID      string
	RankToken string
	Token     string
	PhoneID   string
}

type Instagram struct {
	Cookiejar *cookiejar.Jar
	InstaType
	Transport  http.Transport
	reqOptions *reqOptions
}

func (insta *Instagram) canChallenge() bool {
	if insta.Challenge == nil {
		return false
	}
	return true
}

//type ChallengeDelivery string

type ChallengeOptions struct {
	Path     string
	Delivery string
	Code     string
}

type TwoFactorOptions struct {
	ID   string
	Code string
}

func (c *ChallengeOptions) Choice() string {
	switch c.Delivery {
	// use the delivery method provided
	case GOINSTA_CHALLENGE_SMS, GOINSTA_CHALLENGE_EMAIL:
		return c.Delivery
	}
	// by default we will use email
	return GOINSTA_CHALLENGE_EMAIL
}

type InstaType struct {
	IsLoggedIn   bool
	ChallengeReq bool
	TwoFactorReq bool

	Informations Informations
	LoggedInUser response.User

	Proxy string

	Challenge *ChallengeOptions
	TwoFactor *TwoFactorOptions
}

type BackupType struct {
	Cookies []http.Cookie
	InstaType
}
