package main

type image struct {
	Id				string
	Description 	string
	Url				string
	UserId			string
	UserName		string
	Timestamp		int64
	LikeCount		int64
	CommentCount	int64
}

type count struct {
	Id				string
	Num				int64
}

type sns struct {
	Type				string
	MessageId			string
	TopicArn			string
	Message				string
	Timestamp			string
	SignatureVersion 	string
	Signature			string
	SigningCertURL		string
	SubscribeURL		string
	UnsubscribeURL		string
}
