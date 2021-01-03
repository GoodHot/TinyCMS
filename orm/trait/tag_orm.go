package trait

type Tag struct {
	BaseORM
	Title string `json:"title"`
	Path  string `json:"path"`
}
