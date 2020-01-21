package main

type MessagesResponse struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	CreateAt string `json:"createdAt"`
	Type     string `json:"type"`
}
