package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"go-OAuth/config"
	jwt "go-OAuth/pkg/jwt"
	"go-OAuth/pkg/res"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

const redirectURL = "http://localhost:9090/api/ouath/google/login"
const scope = "https://www.googleapis.com/auth/userinfo.profile"

// GoogleAccsess GoogleAccsess
func GoogleAccsess(c *gin.Context) {
	res.Success(c, gin.H{
		"url": oauthURL(),
	})
}

func oauthURL() string {
	u := "https://accounts.google.com/o/oauth2/v2/auth?client_id=%s&response_type=code&scope=%s&redirect_uri=%s"

	return fmt.Sprintf(u, config.Val.GoogleClientID, scope, redirectURL)
}

// GoogleLogin GoogleLogin
func GoogleLogin(c *gin.Context) {
	code := c.Query("code")
	token, err := accessToken(code)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Debug("accessToken error")
		c.Redirect(http.StatusFound, "/")
		return
	}

	id, name, err := getGoogleUserInfo(token)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Debug("getGoogleUserInfo error")
		c.Redirect(http.StatusFound, "/")
		return
	}

	log.Infof("id: %v, name: %v", id, name)

	jwtToken, err := jwt.GenerateToken(id, name)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Debug("GenerateToken error")
		c.Redirect(http.StatusFound, "/")
		return
	}

	c.SetCookie(jwt.Key, jwtToken, config.Val.JWTTokenLife, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/")
}

func accessToken(code string) (token string, err error) {
	u := "https://www.googleapis.com/oauth2/v4/token"

	data := url.Values{"code": {code}, "client_id": {config.Val.GoogleClientID}, "client_secret": {config.Val.GoogleSecretKey}, "grant_type": {"authorization_code"}, "redirect_uri": {redirectURL}}
	body := strings.NewReader(data.Encode())

	resp, err := http.Post(u, "application/x-www-form-urlencoded", body)
	if err != nil {
		return token, err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return token, err
	}

	token = gjson.GetBytes(b, "access_token").String()
	return token, nil
}

func getGoogleUserInfo(token string) (id, name string, err error) {
	u := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", token)
	resp, err := http.Get(u)
	if err != nil {
		return id, name, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return id, name, err
	}

	name = gjson.GetBytes(body, "name").String()
	id = gjson.GetBytes(body, "id").String()

	return id, name, nil
}
