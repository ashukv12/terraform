package zoom

import (
	// "encoding/json"
	"encoding/json"
	// "errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	// "strings"
	"time"
)

type Client struct {
	authToken  string
	httpClient *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		authToken:  token,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}


func (c *Client) doRequest(req *http.Request, method string) ([]byte, error) {
	req.Header.Set("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MjAxMDExOTUsImlhdCI6MTYxOTQ5NjQwN30.P85E91mA-T_-tISlKA7GX0XRcD6bheVArg-spWLwSTk")
	fmt.Println(method)
	if method == "POST" || method == "PATCH" {
		req.Header.Add("content-type", "application/json")
	}
	// fmt.Println("do ",c.authToken)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if method == "POST" {
		if res.StatusCode != 201 {
			return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
		}
	} else if method == "DELETE" {
		if res.StatusCode != 204 {
			return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
		}
	} else if method == "PATCH" {
		if res.StatusCode != 204 {
			return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
		}
	} else if res.StatusCode != 200 {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
	return body, err
}

func (c *Client) GetUserData(UserId string) (*Users, error) {
	// time.Sleep(10 * time.Second)
	URL := "https://api.zoom.us/v2/users/" + UserId
	req, err := http.NewRequest("GET", URL, nil)
	fmt.Println("data user")
	if err != nil {
		return nil, err
	}

	r, err := c.doRequest(req, "GET")
	if err != nil {
		return nil, err
	}
	user := Users{}
	err = json.Unmarshal(r, &user)
	if err != nil {
		return nil, err
	}
	fmt.Println("Get Data", UserId)
	return &user, nil
}

func (c *Client) CreateUser(userCreateInfo UserCreate, UserId string) (*UserCreate, error) {
	fmt.Println("createfun in ", UserId)
	client := &http.Client{Timeout: 10 * time.Second}
	URL := "https://api.zoom.us/v2/users/" + UserId
	r, _ := http.NewRequest("GET", URL, nil)
	// c.authToken= "Bearer "+ c.authToken
	r.Header.Add("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MjAxMDExOTUsImlhdCI6MTYxOTQ5NjQwN30.P85E91mA-T_-tISlKA7GX0XRcD6bheVArg-spWLwSTk")

	// fmt.Println("createfun")
	// fmt.Println(c.authToken)
	re, er := client.Do(r)
	if er != nil {
		return nil, fmt.Errorf("status: %d", re.StatusCode)
	}
	defer re.Body.Close()
	b, e := ioutil.ReadAll(re.Body)
	if e != nil {
		return nil, fmt.Errorf("status: %d body %v", re.StatusCode, b)
	}
	if re.StatusCode != 404 {
		if re.StatusCode == 200 {
			return nil, fmt.Errorf("status: %d  User already exist %v  ", re.StatusCode, UserId)
		}
	} else if re.StatusCode == 404 {

		reqb, err := json.Marshal(userCreateInfo)
		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest("POST", "https://api.zoom.us/v2/users", strings.NewReader(string(reqb)))
		time.Sleep(10 * time.Second)
		if err != nil {
			return nil, err
		}
		body, err := c.doRequest(req, "POST")

		if err != nil {
			return nil, err
		}
		user := UserCreate{}
		err = json.Unmarshal(body, &user)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, nil

}

func (c *Client) GetCreatedUserData(UserId string) (*UserCreateInfo, error) {

	URL := "https://api.zoom.us/v2/users/" + UserId
	req, err := http.NewRequest("GET", URL, nil)
	time.Sleep(10 * time.Second)
	// fmt.Println("user get")
	// fmt.Println("user get ",UserId)
	if err != nil {
		return nil, err
	}

	r, err := c.doRequest(req, "GET")
	if err != nil {
		return nil, err
	}

	user := UserCreateInfo{}
	err = json.Unmarshal(r, &user)
	if err != nil {
		return nil, err
	}
	// fmt.Println("user get final", UserId)
	// fmt.Println(user.Email)
	return &user, nil
}
func (c *Client) UpdateUser(UserId string, userUpdateInfo UserCreateInfo) (*UserCreateInfo, error) {
	reqb, err := json.Marshal(userUpdateInfo)
	if err != nil {
		return nil, err
	}
	URL := "https://api.zoom.us/v2/users/" + UserId

	req, err := http.NewRequest("PATCH", URL, strings.NewReader(string(reqb)))
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req, "PATCH")
	if err != nil {
		return nil, err
	}
	user := UserCreateInfo{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}
	// fmt.Println("update func")
	return &user, nil
}
func (c *Client) DeleteUser(UserId string) error {
	time.Sleep(10 * time.Second)
	URL := "https://api.zoom.us/v2/users/" + UserId
	req, err := http.NewRequest("DELETE", URL, nil)
	fmt.Println("delete func init")
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: 10 * time.Second}
	req.Header.Add("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MjAxMDExOTUsImlhdCI6MTYxOTQ5NjQwN30.P85E91mA-T_-tISlKA7GX0XRcD6bheVArg-spWLwSTk")
	// fmt.Println(c.authToken)
	re, err := client.Do(req)
	fmt.Println(re.StatusCode)
	if re.StatusCode != 204 {
		if re.StatusCode != 429 {
			return fmt.Errorf("status: %v", re.StatusCode)
		}
	}
	if err != nil {
		return fmt.Errorf("status: %v", err)
	}

	return nil
}
