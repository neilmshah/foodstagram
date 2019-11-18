package main

type likecount struct{
	Count int64
}

type like struct{
	User string
}

type comment struct{
	User string
	Comment string
}

type comment_list struct{
	Comments [] string
}

type data_struct struct{
	Photo_id 	string
	Likes		int64
	Comment_list 	[]string
}
