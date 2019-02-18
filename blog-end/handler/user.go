package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"publisher-end/database"
	"publisher-end/utils"
)

func getAccessToken(code string) (*database.User, error) {
	remote := "https://github.com/login/oauth/access_token"
	params := url.Values{
		"client_id": {"b35c9c1d9c1649803683"},
		"client_secret": {"8f14354bc4e4b11c3b7ca932a8dc794a5f7da584"},
		"code": {code},
	}
	res, err := http.PostForm(remote, params)
	defer res.Body.Close()
	if err != nil {
		fmt.Println("getAccessToken - %v", err)
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("getAccessToken - %v", err)
		return nil, err
	}
	kws := utils.QueryStringToMap(string((data)))
	access_token, exists := kws["access_token"]
	if !exists {
		fmt.Println("获取access_token失败")
		return nil, errors.New("access_token not exists")
	}
	res, err = http.Get("https://api.github.com/user?access_token=" + access_token)
	if err != nil {
		fmt.Println("getAccessToken - %v", err)
		return nil, err
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("getAccessToken - %v", err)
		return nil, err
	}
	var objMap map[string]*json.RawMessage
	err = json.Unmarshal(data, &objMap)
	if err != nil {
		fmt.Println("parseJson - %v", err)
		return nil, err
	}
	var  name, avatar string
	var id int
	err = json.Unmarshal(*objMap["login"], &name)
	if err != nil {
		fmt.Println("parseJson - %v", err)
		return nil, err
	}
	err = json.Unmarshal(*objMap["avatar_url"], &avatar)
	if err != nil {
		fmt.Println("parseJson - %v", err)
		return nil, err
	}
	err = json.Unmarshal(*objMap["id"], &id)
	if err != nil {
		fmt.Println("parseJson - %v", err)
		return nil, err
	}
	fmt.Println("id = ", id)
	err = database.SaveUser(fmt.Sprintf("%d", id), avatar, name)
	if err != nil {
		fmt.Println("SaveUser - %v", err)
		return nil, err
	}
	return &database.User{
		ID: fmt.Sprintf("%d", id),
		Avatar: avatar,
		Name: name,
	}, nil
}

func GetUserInfo(c *gin.Context) {
	code := c.PostForm("code")
	user, err := getAccessToken(code)
	if err != nil {
		fmt.Println("GetUserInfo - %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "发生错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
