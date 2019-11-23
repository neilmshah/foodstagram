# Comments & Likes Services

For each picture, any signed up user can view all the comments for any picture and like or add additional comment on a picture.

## APIs for comments

### POST /comment/{photo_id}
```
request_body = {
		"User_id": <user_id>,
		"User_name": <user_name>,
		"Comment": <comment_on_picture>
	}
```
RESPONSES
* HTTP 201 -  Comment is successfully added into the collection along with a timestamp to sort comments later.
* HTTP 400 - Any invalid or missing field.

### GET /comment/{photo_id}/{user_id}
RESPONSE
```
response_body = {
		Liked: <1 or 0>,
		Comments: [
			{
				"User_id": <user_id>,
				"User_name": <user_name>,
				"Timestamp": <timestamp>
				"Comment": <comment_on_picture>
			},
			{
				"User_id": <user_id_2>,
				"User_name": <user_name_2>,
				"Timestamp": <timestamp_2>
				"Comment": <comment_on_picture_2>
			},
		]
	}
```

### DELETE /comment/{photo_id}
Users will be able to delete a comment they previously posted.
```
request_body = {
		"User_id": <user_id>,
		"User_name": <user_name>,
		"Comment": <comment_on_picture>
	}
```
RESPONSE
* HTTP 200 - If the service pulled the perivously appended comment successfully and decrements the count value
* HTTP 400 - If the database updates did not occur properly

### 
## APIs for likes

### POST /like/{photo_id}
```
request_body = {
		"User_id": <user_id>
	}
```
RESPONSES
* HTTP 201 - Like count is incremented by one, id of user who liked the photo is appended to the collection.
* HTTP 400 - Invalid photo ID, missing user name and improper update to the mongodb.

### GET /like/count/{photo_id}
Backup API to call if the the timeline service fails
```
response_body = {
		"Like_count": <number>
	}
```

### DELETE /like/count/{photo_id}
Disklike a previously liked photo.
```
request_body = {
		"User_id": <user_id>
	}
```
RESPONSES
* HTTP 200 - Sucessfully decremented like count and removed user ID from the collection for the given photo.
* HTTP 400 - If the database updates did not occur properly
