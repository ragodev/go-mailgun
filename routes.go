package mailgun

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type Route struct {
	Id          string   `json:"id"`
	Priority    int      `json:"priority"`
	Description string   `json:"description"`
	Expression  string   `json:"expression"`
	Actions     []string `json:"actions"`
}

func (c *Client) Routes(limit, skip int) (total int, res []Route, err error) {
	v := url.Values{}
	v.Set("limit", strconv.Itoa(limit))
	v.Set("skip", strconv.Itoa(skip))
	body, err := c.api("GET", "/routes", v)
	if err != nil {
		return
	}

	var j struct {
		Total int     `json:"total_count"`
		Items []Route `json:"items"`
	}

	err = json.Unmarshal(body, &j)
	if err != nil {
		return
	}
	total, res = j.Total, j.Items
	return

}

func (c *Client) Get(routeId string) (r Route, err error) {
	rsp, err := c.api("GET", "/routes/"+routeId, nil)
	if err != nil {
		return
	}
	var res struct {
		Message string `json:"message"`
		R       Route  `json:"route"`
	}
	err = json.Unmarshal(rsp, &res)
	r = res.R

	return
}

func (c *Client) Create(r *Route) (routeId string, err error) {
	v := url.Values{}

	v.Set("priority", strconv.Itoa(r.Priority))
	v.Set("description", r.Description)
	v.Set("expression", r.Expression)

	for _, a := range r.Actions {
		v.Add("action", a)
	}

	rsp, err := c.api("POST", "/routes", v)
	if err != nil {
		return
	}
	var res struct {
		Message string `json:"message"`
		R       Route  `json:"route"`
	}
	err = json.Unmarshal(rsp, &res)
	routeId = res.R.Id
	return
}

func (c *Client) Update(r *Route) (routeId string, err error) {
	v := url.Values{}

	v.Set("priority", strconv.Itoa(r.Priority))
	v.Set("description", r.Description)
	v.Set("expression", r.Expression)

	for _, a := range r.Actions {
		v.Add("action", a)
	}

	rsp, err := c.api("PUT", "/routes/"+r.Id, v)
	if err != nil {
		return
	}
	var res struct {
		Message string `json:"message"`
		R       Route  `json:"route"`
	}
	err = json.Unmarshal(rsp, &res)
	routeId = res.R.Id

	return
}

func (c *Client) Delete(r *Route) (err error) {
	rsp, err := c.api("DELETE", "/routes/"+r.Id, nil)
	if err != nil {
		return
	}
	var res struct {
		Message string `json:"message"`
	}
	err = json.Unmarshal(rsp, &res)
	return
}
