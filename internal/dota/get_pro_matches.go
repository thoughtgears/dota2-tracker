package dota

import (
	"encoding/json"
	"net/http"

	"github.com/thoughtgears/dota2-tracker/models"
)

func (c *Client) GetProMatches() []models.ProMatch {
	req, _ := http.NewRequest("GET", c.baseUrl+"/proMatches", nil)
	res, err := c.client.Do(req)
	if err != nil {
		return nil
	}

	var proMatches []models.ProMatch
	if err := json.NewDecoder(res.Body).Decode(&proMatches); err != nil {
		return nil
	}

	return proMatches
}
