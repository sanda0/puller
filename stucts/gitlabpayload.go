package stucts

import "time"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar_url"`
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

type MergeParams struct {
	ForceRemoveSourceBranch  string `json:"force_remove_source_branch"`
	ShouldRemoveSourceBranch bool   `json:"should_remove_source_branch"`
}

type SourceInfo struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	WebURL            string `json:"web_url"`
	AvatarURL         string `json:"avatar_url"`
	GitSSHURL         string `json:"git_ssh_url"`
	GitHTTPURL        string `json:"git_http_url"`
	Namespace         string `json:"namespace"`
	VisibilityLevel   int    `json:"visibility_level"`
	PathWithNamespace string `json:"path_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
	CIConfigPath      string `json:"ci_config_path"`
	Homepage          string `json:"homepage"`
	URL               string `json:"url"`
	SSHURL            string `json:"ssh_url"`
	HTTPURL           string `json:"http_url"`
}

type CommitInfo struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	Title     string    `json:"title"`
	Timestamp time.Time `json:"timestamp"`
	URL       string    `json:"url"`
	Author    struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"author"`
}

type TargetInfo struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	WebURL            string `json:"web_url"`
	AvatarURL         string `json:"avatar_url"`
	GitSSHURL         string `json:"git_ssh_url"`
	GitHTTPURL        string `json:"git_http_url"`
	Namespace         string `json:"namespace"`
	VisibilityLevel   int    `json:"visibility_level"`
	PathWithNamespace string `json:"path_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
	CIConfigPath      string `json:"ci_config_path"`
	Homepage          string `json:"homepage"`
	URL               string `json:"url"`
	SSHURL            string `json:"ssh_url"`
	HTTPURL           string `json:"http_url"`
}

type LastCommitInfo struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	Title     string    `json:"title"`
	Timestamp time.Time `json:"timestamp"`
	URL       string    `json:"url"`
	Author    struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"author"`
}

type ObjectAttributes struct {
	AssigneeID                  interface{}    `json:"assignee_id"`
	AuthorID                    int            `json:"author_id"`
	CreatedAt                   string         `json:"created_at"`
	Description                 string         `json:"description"`
	Draft                       bool           `json:"draft"`
	HeadPipelineID              interface{}    `json:"head_pipeline_id"`
	ID                          int            `json:"id"`
	IID                         int            `json:"iid"`
	LastEditedAt                interface{}    `json:"last_edited_at"`
	LastEditedByID              interface{}    `json:"last_edited_by_id"`
	MergeCommitSHA              string         `json:"merge_commit_sha"`
	MergeError                  interface{}    `json:"merge_error"`
	MergeParams                 MergeParams    `json:"merge_params"`
	MergeStatus                 string         `json:"merge_status"`
	MergeUserID                 interface{}    `json:"merge_user_id"`
	MergeWhenPipelineSucceeds   bool           `json:"merge_when_pipeline_succeeds"`
	MilestoneID                 interface{}    `json:"milestone_id"`
	SourceBranch                string         `json:"source_branch"`
	SourceProjectID             int            `json:"source_project_id"`
	StateID                     int            `json:"state_id"`
	TargetBranch                string         `json:"target_branch"`
	TargetProjectID             int            `json:"target_project_id"`
	TimeEstimate                int            `json:"time_estimate"`
	Title                       string         `json:"title"`
	UpdatedAt                   string         `json:"updated_at"`
	UpdatedByID                 interface{}    `json:"updated_by_id"`
	URL                         string         `json:"url"`
	Source                      SourceInfo     `json:"source"`
	Target                      TargetInfo     `json:"target"`
	LastCommit                  LastCommitInfo `json:"last_commit"`
	WorkInProgress              bool           `json:"work_in_progress"`
	TotalTimeSpent              int            `json:"total_time_spent"`
	TimeChange                  int            `json:"time_change"`
	HumanTotalTimeSpent         interface{}    `json:"human_total_time_spent"`
	HumanTimeChange             interface{}    `json:"human_time_change"`
	HumanTimeEstimate           interface{}    `json:"human_time_estimate"`
	AssigneeIDs                 []interface{}  `json:"assignee_ids"`
	ReviewerIDs                 []interface{}  `json:"reviewer_ids"`
	Labels                      []interface{}  `json:"labels"`
	State                       string         `json:"state"`
	BlockingDiscussionsResolved bool           `json:"blocking_discussions_resolved"`
	FirstContribution           bool           `json:"first_contribution"`
	DetailedMergeStatus         string         `json:"detailed_merge_status"`
}

type GitLabWebhookPayload struct {
	ObjectKind        string           `json:"object_kind"`
	EventName         string           `json:"event_name"`
	Before            string           `json:"before"`
	After             string           `json:"after"`
	Ref               string           `json:"ref"`
	CheckoutSHA       string           `json:"checkout_sha"`
	Message           string           `json:"message"`
	User              User             `json:"user"`
	ProjectID         int              `json:"project_id"`
	Project           Project          `json:"project"`
	Commits           []Commit         `json:"commits"`
	TotalCommitsCount int              `json:"total_commits_count"`
	PushOptions       struct{}         `json:"push_options"`
	Repository        Repository       `json:"repository"`
	ObjectAttributes  ObjectAttributes `json:"object_attributes"`
	UserID            int              `json:"user_id"`
	UserName          string           `json:"user_name"`
	UserUsername      string           `json:"user_username"`
	UserEmail         string           `json:"user_email"`
	UserAvatar        string           `json:"user_avatar"`
}
