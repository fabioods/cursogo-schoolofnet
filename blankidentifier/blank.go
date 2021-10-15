package blankidentifier

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Blank() {
	resp, _ := http.Get("https://api.github.com/users/fabioods")
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("\n\n%s", body)
}
