package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"strings"
)

func main() {
	client := http.Client{}
	login(&client, "username", "password")
}

func login(client *http.Client, username, password string) {
	urls := make(map[string]string, 0)
	urls["check"] = "https://auth.bupt.edu.cn/authserver/login"
	
	if client.Jar == nil {
		jar, err := cookiejar.New(nil)
		if err != nil {
			panic("Jar init error")
		}
		client.Jar = jar
	}

	loginMap := make(map[string][]string)
	loginMap["submit"] = []string{"LOGIN"}
	loginMap["type"] = []string{"username_password"}
	loginMap["username"] = []string{username}
	loginMap["password"] = []string{password}
	loginMap["_eventId"] = []string{"submit"}

	// 获取登录用的 execution 参数
	resp, err := client.Get(urls["check"])
	body, err := ioutil.ReadAll(resp.Body)
	re := regexp.MustCompile(`<input name="execution" value="[0-9a-zA-Z-=_]+"`)
	findExecution := re.FindAll(body, -1)
	execution := strings.Split(string(findExecution[0]), "\"")[3]
	loginMap["execution"] = []string{execution}

	_, err = client.PostForm(urls["check"], loginMap)
	if err != nil {
		fmt.Println(err)
		panic("check error")
	}
}
	
