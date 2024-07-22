package response

type GitActionResponse struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}
