package stucts

import "time"

type User struct {
	ID       int    `json:"user_id"`
	Name     string `json:"user_name"`
	Username string `json:"user_username"`
	Email    string `json:"user_email"`
	Avatar   string `json:"user_avatar"`
}

type Project struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	WebURL            string `json:"web_url"`
	GitSSHURL         string `json:"git_ssh_url"`
	GitHTTPURL        string `json:"git_http_url"`
	Namespace         string `json:"namespace"`
	VisibilityLevel   int    `json:"visibility_level"`
	PathWithNamespace string `json:"path_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
	Homepage          string `json:"homepage"`
	URL               string `json:"url"`
	SSHURL            string `json:"ssh_url"`
	HTTPURL           string `json:"http_url"`
}

type Commit struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	URL       string    `json:"url"`
	Author    User      `json:"author"`
	Added     []string  `json:"added"`
	Modified  []string  `json:"modified"`
	Removed   []string  `json:"removed"`
}

type Repository struct {
	Name            string `json:"name"`
	URL             string `json:"url"`
	Description     string `json:"description"`
	Homepage        string `json:"homepage"`
	GitHTTPURL      string `json:"git_http_url"`
	GitSSHURL       string `json:"git_ssh_url"`
	VisibilityLevel int    `json:"visibility_level"`
}

type GitLabWebhookPayload struct {
	ObjectKind        string     `json:"object_kind"`
	EventName         string     `json:"event_name"`
	Before            string     `json:"before"`
	After             string     `json:"after"`
	Ref               string     `json:"ref"`
	CheckoutSHA       string     `json:"checkout_sha"`
	Message           string     `json:"message"`
	User              User       `json:"user"`
	ProjectID         int        `json:"project_id"`
	Project           Project    `json:"project"`
	Commits           []Commit   `json:"commits"`
	TotalCommitsCount int        `json:"total_commits_count"`
	PushOptions       struct{}   `json:"push_options"`
	Repository        Repository `json:"repository"`
}

type Event struct {
	Type     string   `json:"type"`
	Commands []string `json:"commands"`
}

type Notification struct {
	Type string   `json:"type"`
	To   []string `json:"to"`
}

type Repo struct {
	Name          string         `json:"name"`
	Path          string         `json:"path"`
	Branch        string         `json:"branch"`
	Events        []Event        `json:"events"`
	Notifications []Notification `json:"notifications"`
}

type Config struct {
	Key   string `json:"key"`
	Repos []Repo `json:"repos"`
}
