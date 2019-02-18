package database

type Post struct {
	ID string `json:"id"`
	Title string `json:"title"`
	HtmlStr string `json:"html_str"`
	RawStr string `json:"raw_str"`
	CreateTime int `json:"create_time"`
}

type User struct {
	ID string `json:"id"`
	Avatar string `json:"avatar"`
	Name string `json:"name"`
}

type Comment struct {
	Post_ID string `json:"post_id"`
	Comment_ID int `json:"comment_id"`
	Commenter_ID string `json:"commenter_id"`
	Parent_ID int `json:"parent_id"`
	Create_time int `json:"create_time"`
	Replyer_id string `json:"replyer_id"`
	Replyer_name string `json:"replyer_name"`
	Content string `json:"content"`
	Commenter_avatar string `json:"commenter_avatar"`
	Commenter_name string `json:"commenter_name"`
	Sub_comments []Comment `json:"sub_comments"`
}
