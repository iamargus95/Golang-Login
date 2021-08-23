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
	reposLoop    int
)

type Owner struct {
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
}

type Licence struct {
	Key     string
	Name    string
	Spdx_id string
	Url     string
	Node_id string
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

type ReposInfoJson struct {
	Id                int
	Node_id           string
	Name              string
	Full_name         string
	Private           bool
	Owner             Owner
	Html_url          string
	Description       string
	Fork              bool
	Url               string
	Forks_url         string
	Keys_url          string
	Collaborators_url string
	Teams_url         string
	Hooks_url         string
	Issue_events_url  string
	Events_url        string
	Assignees_url     string
	Branches_url      string
	Tags_url          string
	Blobs_url         string
	Git_tags_url      string
	Git_refs_url      string
	Trees_url         string
	Statuses_url      string
	Languages_url     string
	Stargazers_url    string
	Contributors_url  string
	Subscribers_url   string
	Subscription_url  string
	Commits_url       string
	Git_commits_url   string
	Comments_url      string
	Issue_comment_url string
	Contents_url      string
	Compare_url       string
	Merges_url        string
	Archive_url       string
	Downloads_url     string
	Issues_url        string
	Pulls_url         string
	Milestones_url    string
	Notifications_url string
	Labels_url        string
	Releases_url      string
	Deployments_url   string
	Created_at        string
	Updated_at        string
	Pushed_at         string
	Git_url           string
	Ssh_url           string
	Clone_url         string
	Svn_url           string
	Homepage          string
	Size              int
	Stargazers_count  int
	Watchers_count    int
	Language          string
	Has_issues        bool
	Has_projects      bool
	Has_downloads     bool
	Has_wiki          bool
	Has_pages         bool
	Forks_count       int
	Mirror_url        string
	Archived          bool
	Disabled          bool
	Open_issues_count int
	License           Licence
	Forks             int
	Open_issues       int
	Watchers          int
	Default_branch    string
}

type ReposInfoArray []ReposInfoJson

func main() {
	fmt.Println("\nStarting listener on http://localhost:8080")
	http.HandleFunc("/", getInput)
	http.ListenAndServe(":8080", nil)
}

func getInput(w http.ResponseWriter, r *http.Request) {

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

	bodyJson, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var users Userinfo
	json.Unmarshal([]byte(bodyJson), &users)
	// fmt.Println(users)
	reposLoop = users.Public_repos

	if users.Name != "" {
		fmt.Printf("\nName: %s\n\nLogin: %s\n\nBio: %s\n\nPublic Repositories: %d\n\nFollowers: %d\n\nFollowing: %d\n\n", users.Name, users.Login, users.Bio, users.Public_repos, users.Followers, users.Following)
		getRepos()
	}
}

func getRepos() {

	url := "http://api.github.com/users/" + formUsername + "/repos"

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

	bodyJson, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var reposArray ReposInfoArray
	json.Unmarshal([]byte(bodyJson), &reposArray)
	for i := 0; i < reposLoop; i++ {
		fmt.Printf("\nRepository No %d: %v", i+1, reposArray[i].Name)
		fmt.Printf("\nAvailable at : %v\n", reposArray[i].Html_url)
	}
}
