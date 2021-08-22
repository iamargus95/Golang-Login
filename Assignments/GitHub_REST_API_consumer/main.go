package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

var (
	tpl          *template.Template
	formUsername string
)

func main() {
	fmt.Println("Starting listener on http://localhost:8080")
	http.HandleFunc("/", getUserName)
	http.ListenAndServe(":8080", nil)
}

type Userinfo struct {
	Login               string
	Id                  int
	Node_id             string
	Avatar_url          string
	Gravatar_id         string
	Url                 string
	Html_url            string
	Followers_url       string
	Following_url       string
	Gists_url           string
	Starred_url         string
	Subscriptions_url   string
	Organizations_url   string
	Repos_url           string
	Events_url          string
	Received_events_url string
	Type                string
	Site_admin          bool
	Name                string
	Company             string
	Blog                string
	Location            string
	Email               string
	Hireable            bool
	Bio                 string
	Twitter_username    string
	Public_repos        int
	Public_gists        int
	Followers           int
	Following           int
	Created_at          string
	Updated_at          string
}

func getUserName(w http.ResponseWriter, r *http.Request) {

	tpl = template.Must(template.ParseGlob("./index.html"))
	tpl.ExecuteTemplate(w, "index.html", nil)

	formUsername = r.FormValue("username")

	url := "https://api.github.com/users/" + formUsername

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var users Userinfo
	json.Unmarshal([]byte(body), &users)
	// fmt.Println(users)
	fmt.Printf("Name: %s\n\nLogin: %s\n\nURL: %s\n\nBio: %s\nPublic Repositories: %d\n\nFollowers: %d\n\nFollowing: %d\n\n", users.Name, users.Login, users.Url, users.Bio, users.Public_repos, users.Followers, users.Following)
}
