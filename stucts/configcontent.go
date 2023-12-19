package stucts

type Event struct {
	Type     string   `json:"type"`
	Commands []string `json:"commands"`
}

type Repo struct {
	Name   string  `json:"name"`
	Path   string  `json:"path"`
	Branch string  `json:"branch"`
	Events []Event `json:"events"`
}

type Config struct {
	Key           string `json:"key"`
	Repos         []Repo `json:"repos"`
	Email         string `json:"email"`
	EmailPassword string `json:"email_password"`
	SMTPHost      string `json:"smtp_host"`
	SMTPPort      string `json:"smtp_port"`
}
