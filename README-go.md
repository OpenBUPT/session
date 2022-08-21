# BUPT Session

北邮统一登录网关 Session。用于需要登录的网络请求。

## Usage

```golang
package main

import (
	"fmt"
	"net/http"

	bupt_session "github.com/OpenBUPT/session/bupt_session_go"
)

func main() {
	client := &http.Client{}
	client, err := bupt_session.Login(client, "username", "password")
	if err != nil {
		fmt.Println(err)
	}
}
```