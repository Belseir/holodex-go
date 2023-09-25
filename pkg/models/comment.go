package models

type Comment struct {
	CommentKey string `json:"comment_key"`
	VideoID    string `json:"video_id"`
	Message    string `json:"message"`
}
