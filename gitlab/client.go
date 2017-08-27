package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	Config *Config
}

func (c *Client) Project(owner, name string) (*Project, error) {
	id := url.PathEscape(fmt.Sprintf("%s/%s", owner, name))
	uri := c.Config.URL + "/api/v4/projects/" + id
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("PRIVATE-Token", c.Config.PrivateToken)
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var p = &Project{}
	err = json.Unmarshal(body, p)

	return p, err
}

func (c *Client) CreateMergeRequest(p *Project, params map[string]interface{}) (*http.Response, error) {
	var uri = c.Config.URL + "/api/v4/projects/" + p.UrlString() + "/merge_requests"
	buf, err := getBuffer(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", uri, buf)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("PRIVATE-Token", c.Config.PrivateToken)
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func getBuffer(body interface{}) (*bytes.Buffer, error) {
	json, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(json)
	return buf, nil
}

func NewClient(config *Config) *Client {
	return &Client{
		Config: config,
	}
}
