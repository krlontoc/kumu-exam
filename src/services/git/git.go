package git

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	hlpr "kumu-exam/src/helpers"
)

type ListForm struct {
	Users []string `json:"users"`
}

type Request struct {
	HttpClient *http.Client
	EndPoint   string
	Method     string
}

type ResponseBody struct {
	Name        string `json:"name,omitempty"`
	Login       string `json:"login"`
	Company     string `json:"company,omitempty"`
	Followers   int    `json:"followers,omitempty"`
	PublicRepos int    `json:"public_repos,omitempty"`
	Messege     string `json:"message,omitempty"`
}

func GetUsers(param ListForm) ([]ResponseBody, error) {
	// container of user name to be used for data sorting
	arrUserName := []string{}
	queriedUser := map[string]interface{}{}

	data := map[string]ResponseBody{}
	req := Request{
		HttpClient: &http.Client{Timeout: 30 * time.Second},
		Method:     "GET",
	}

	wg := new(sync.WaitGroup)

	// to ensure that the waitgroup will only do 10 concurrent processes
	guard := make(chan struct{}, 10)
	baseURL := "https://api.github.com/users/"
	for _, user := range param.Users {
		// bypass duplicates
		if queriedUser[user] != nil {
			continue
		}

		// set map key to check for duplicate
		queriedUser[user] = struct{}{}

		// check cache before hitting git API
		if item, found := hlpr.GetFromCache(user); found {
			value := item.(ResponseBody)

			// replace spaces to dash(-) for sorting purposes
			infoKeys := strings.ToLower(strings.Replace(value.Name, " ", "-", -1))
			arrUserName = append(arrUserName, infoKeys)
			data[infoKeys] = value
			continue
		}

		wg.Add(1)
		guard <- struct{}{}
		go func(username string) {
			defer func() {
				<-guard
			}()
			defer wg.Done()

			addToCache := true
			req.EndPoint = baseURL + username
			resp, err := req.consume()
			if err != nil {
				// incase of http request error, add the error message on the user info
				// also, don't add the user info to cache
				resp.Messege = err.Error()
				addToCache = false
			}

			// for instances that user doesn't exist, set back the login
			if resp.Login == "" {
				resp.Login = username
			}

			// set name to user that doesn't set their profile name
			if resp.Name == "" {
				resp.Name = resp.Login
			}

			// replace spaces to dash(-) for sorting purposes
			infoKeys := strings.ToLower(strings.Replace(resp.Name, " ", "-", -1))
			arrUserName = append(arrUserName, infoKeys)
			data[infoKeys] = *resp

			// added conditional addToCache check for http request error instances
			if addToCache {
				// add user info to cache
				itemExp := (2 * time.Minute)
				if err := hlpr.AddToCache(username, *resp, &itemExp); err != nil {
					log.Printf("[ERROR] - Unable to cache [%s] info, Error: %s", username, err.Error())
				}
			}

		}(user)
	}

	wg.Wait()

	// sort response data
	arrInfo := []ResponseBody{}
	sort.Strings(arrUserName)
	for _, sUserName := range arrUserName {
		arrInfo = append(arrInfo, data[sUserName])
	}

	return arrInfo, nil
}

func (r *Request) consume() (*ResponseBody, error) {
	if r.HttpClient == nil {
		r.HttpClient = &http.Client{Timeout: 30 * time.Second}
	}

	req, err := http.NewRequest(r.Method, r.EndPoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := r.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	respBody := &ResponseBody{}
	err = json.Unmarshal([]byte(string(body)), &respBody)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
