package main

type likecount struct{
	Count int64
}

type like struct{
	User_id string
}

type user_comment struct{
	User_id 	string
	Timestamp 	int64
	Comment 	string
}

type comment_list struct{
	Comments [] string
}

type data_struct struct{
	Photo_id 		string
	Like_count		int64
	Comment_count	int64
	Likes			[]string
	Comments 	[]user_comment
}

type modal_struct struct{
	Liked		bool
	Comments 	[]user_comment
}
