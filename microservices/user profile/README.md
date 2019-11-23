# User Profile Service

## APIs for comments

### POST /register
Create a new user and assign a UUID
```
request_body={
		"Username": <user name provided during signup>,
		"Password": <password corresponding to username>,
		"Firstname": <first name of user>,
		"Lastname": <last name of user>
	}
```
RESPONSE
* HTTP 200 - Sucessfully created a new user profile and saved it in database
* HTTP 400 - Invalid entries or database updated failed.

### POST /login
Authenticate user credentials
```
request_body={
		"Username": <user name provided during signup>,
		"Password": <password corresponding to username>
	}
```
RESPONSE
* HTTP 200 - Sucessfully authenticated user details and returns token to handle session for user
* HTTP 404 - Error reading data or invalid credentials


### GET /profile
Get user information after sucessful login to populate the user information in homepage.
```
response_body={
		"UserID": <unique id assigned to user by the system>,
		"Username": <user name provided during signup>,
		"Firstname": <first name of user>,
		"Lastname": <last name of user>
	}
```