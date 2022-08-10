package model

type Article struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Likes    []int `json:"likes"`
	Dislikes []int `json:"dislikes"`
	Author   User   `json:"author"`
	Deleted  bool   `json:"is_deleted"`
}
