package main

type image struct {
	Id				string
	Description 	string
	Url				string
	UserId			string
	UserName		string
	Timestamp		string
	LikeCount		int64
	CommentCount	int64
}

type count struct {
	Num				int64
}

type count2 struct {
	Num				string
}
