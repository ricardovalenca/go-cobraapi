package request

type GitActionFetch struct {
	Name string `json:"name"`
}

type GitActionRun struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
