package disbotlist

import (
	"io/ioutil"
	"net/http"
	"time"
	"fmt"
)


type DisBotListClient struct {
    Token  string
}

func (client DisBotListClient) ServerCountPost(servers string)  string {

	req, err := http.NewRequest("POST", "https://disbots.xyz/api/bots/stats",  nil)

	if err != nil {
		return ""
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", client.Token)
	req.Header.Set("serverCount", servers)

	r, err := httpClient.Do(req)


	if err != nil {
		return ""
	}

  defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return ""
	}

	return string(body)



}

func (client DisBotListClient) HasVoted(id string) string {

  req, err := http.NewRequest("GET", fmt.Sprintf("https://disbots.xyz/api/bots/check/%s", id),  nil)

	if err != nil {
		return ""
	}

	req.Header.Set("Authorization", client.Token)
	req.Header.Set("Content-Type", "application/json")

	r, err := httpClient.Do(req)


  if err != nil {
      return ""
  }

  defer r.Body.Close()


	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return ""
	}

	return string(body)

}

func (client DisBotListClient) Search(id string) string {

	req, err := http.NewRequest("GET", fmt.Sprintf("https://disbots.xyz/api/bots/%s", id),  nil)

	if err != nil {
		return ""
	}

	req.Header.Set("Content-Type", "application/json")

	r, err := httpClient.Do(req)


  if err != nil {
      return ""
  }

  defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return ""
	}

	return string(body)

}



var httpClient = &http.Client{Timeout: 10 * time.Second}
