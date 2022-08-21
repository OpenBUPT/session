package bupt_session

import (
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"strings"
)

const loginUrl = "https://auth.bupt.edu.cn/authserver/login"

func Login(client *http.Client, username, password string) (*http.Client, error) {

	if client.Jar == nil {
		jar, err := cookiejar.New(nil)
		if err != nil {
			return client, err
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
	resp, err := client.Get(loginUrl)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`<input name="execution" value="[0-9a-zA-Z-=_]+"`)
	findExecution := re.FindAll(body, -1)
	execution := strings.Split(string(findExecution[0]), "\"")[3]
	loginMap["execution"] = []string{execution}

	_, err = client.PostForm(loginUrl, loginMap)
	return client, err
}
