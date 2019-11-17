# User Profile Service

# Api usage:

1. POST    /register 
body :
{
	"username":"amit3",
	"password":"abc",
	"firstname":"amit3",
	"lastname":"bharadia3"
}

2. POST   /login
{
	"username":"amit3",
	"password":"abc"
}

3. GET    /profile
param : Authorization = "auth_key"
