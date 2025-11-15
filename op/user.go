package op

type UserInputParams struct {
	UserID      string         `json:"userID"`
	TaskID      string         `json:"taskID"`
	ThreadID    string         `json:"threadID"`
	Metadata    map[string]any `json:"metadata"`
	Content     string         `json:"content"`
	ContentType string         `json:"contentType"`
}
