package library

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	channelAccessToken string
	request            *http.Request
}

func NewClient(token string) *Client {
	return &Client{
		channelAccessToken: token,
	}
}

func (c *Client) SetGetRequest(url string) *Client {
	c.request, _ = http.NewRequest(
		"GET",
		url,
		nil,
	)
	return c
}

func (c *Client) SetAuthorization() *Client {
	token := fmt.Sprintf("Bearer %s", c.channelAccessToken)
	c.request.Header.Set("Authorization", token)
	return c
}

func (c *Client) Send(r interface{}) error {
	client := &http.Client{}
	response, err := client.Do(c.request)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &r); err != nil {
		return err
	}
	return nil
}
