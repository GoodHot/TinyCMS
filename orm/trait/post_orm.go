package trait

type Post struct {
	BaseORM
	Title     string `json:"title"`
	Content   string `json:"content"`
	Tags      string `json:"tags"`
	ChannelID int    `json:"channel_id"`
}
