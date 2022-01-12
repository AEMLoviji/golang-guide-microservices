package github

// https://docs.github.com/en/rest/reference/repos#create-a-repository-for-the-authenticated-user--parameters
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

type CreateRepoResponse struct {
	Id          int64           `json:"id"`
	Name        string          `json:"name"`
	FullName    string          `json:"full_name"`
	Owner       RepoOwner       `json:"owner"`
	Premissions RepoPremissions `json:"premissions"`
}

type RepoOwner struct {
	Id      int64  `json:"id"`
	Login   string `json:"login"`
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
}

type RepoPremissions struct {
	IsAdmin int64  `json:"is_admin"`
	HasPull string `json:"has_pull"`
	HasPush string `json:"has_push"`
}
