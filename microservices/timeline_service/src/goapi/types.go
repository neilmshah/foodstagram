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


/*

{
  "Type" : "Notification",
  "MessageId" : "c164eb91-ded0-5737-88df-bdde1db5f796",
  "TopicArn" : "arn:aws:sns:us-east-1:650635599638:image",
  "Message" : "{\"Id\":\"b7d96d1f-7eef-48f3-bcca-ea5c6cb5f62d\",\"Description\":\"Food image creation2\",\"Url\":\"https://imageservicebucket.s3.amazonaws.com/images/b7d96d1f-7eef-48f3-bcca-ea5c6cb5f62d\",\"UserId\":\"1\",\"UserName\":\"Priyal Agrawal\",\"Timestamp\":\"2019-11-20 22:41:39.428767 +0000 UTC\"}",
  "Timestamp" : "2019-11-20T22:41:41.980Z",
  "SignatureVersion" : "1",
  "Signature" : "gBZ03pJdxZTfc/5kvnN45BVRlZaE7KkBCqTq33ckrZmG50j38aXmp6dmLDOib3xReB3P12P8+gKJe9JjjHbC91rLg264M+vQ4da4C5FUroujXVxQBKrp5L5goA8J2DqFpHtlh18X5z1A3KK2mJ9tBwO3/U2ldA5yV8Lsz/z8fps6zrlpkjPiONSTttL/2Z4yzKWt6ii9azLuAjDuecODbuVrMp+DaReYBhzZO0zkC257aO3WoBbWneO6yv+Ccd4D4nt9GnB2xLTJMljmYigMn2MnsBF7ZuIMtfIJa2n137jOzY3XcdZeh2YwcZFq0vnN9daKkGrTs9dieiMdYXbyYA==",
  "SigningCertURL" : "https://sns.us-east-1.amazonaws.com/SimpleNotificationService-6aad65c2f9911b05cd53efda11f913f9.pem",
  "UnsubscribeURL" : "https://sns.us-east-1.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:us-east-1:650635599638:image:7340d41f-9c8a-4618-8978-985c8a48cdac"
}
*/