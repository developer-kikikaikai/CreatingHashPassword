# HTTP API definition

All of request/response format  are JSON

## Password /api/password

This is for generating password.

### GET /api/password

- Request format:
	```json
	{
		"title": "string of this password's title",
		"keyphrase": "keyphrase string",
		"algorithm": "generating password algorithm, write hash string",
		"seed": "seed value of string to generate password"
	}
	```

- Response:
	- HTTP 200 OK if server generate password:
		```json
		{
			"result": "password result string"
		}
		```
	- (Other status code if server failed to generate password, not implement)

### Others

- POST/PUT/DELETE api/password are Nothing

## PasswordInformation /api/passwordInfo

This is for login user's password information.  
User can store base of generating password information.

### GET /api/passwordInfo

- Request format:
	- body data is NULL
	- User autorization is by Digest autorization.

- Response:
	- HTTP 200 OK if server generate password:
		```json
		{
			"value of title": {
				"algorithm": "generating password algorithm, write hash string",
				"seed": "seed value of string to generate password"
			},
			...
		}
		```
	- (Other status code if server failed to generate password, not implement)

### POST/PUT /api/passwordInfo

Create/Update PasswordInformation

- Request format:
	- body data is following:
		```json
		[
			"value of title": {
				"algorithm": "generating password algorithm, write hash string",
				"seed": "seed value of string to generate password"
			},
			...
		]
		```
	- User autorization is by Digest autorization.

- Response:
	- HTTP 200 OK if server success to update (body data is NULL).
	- Other status code if server failed to generate password, not implement
		- 400 Bad request
			- if PUT request and there is no same title's data, and
			- if POST reqest and there is already created same title's data.

### DELETE /api/passwordInfo

Delete PasswordInformation

- Request format:
	- body data is following:
		```json
		{
			"value of title",
			...
		}
		```
	- User autorization is by Digest autorization.

- Response:
	- HTTP 200 OK if server success to update (body data is NULL).
	- Other status code if server failed to generate password, not implement
		- 400 Bad request
			- if there is no data.

## UserAccount /api/userAccount

This is for generating/updating/deleting login user account.

### GET /api/userAccount

- Nothing

### POST /api/userAccount

Create UserAccount

- Request format:
	- body data is following:
		```json
		{
			"username": "user name information",
			"passphrase": "passphrase string"
		}
		```

- Response:
	- HTTP 200 OK if server success to update (body data is NULL).
	- Other status code if server failed to generate password, not implement
		- 400 Bad request
			- if there is already used username.

### PUT /api/userAccount

Update UserAccount

- Request format:
	- body data is following:
		```json
		{
			"username": "user name information",
			"passphrase": "passphrase string"
		}
		```
	- User autorization is by Digest autorization.

- Response:
	- HTTP 200 OK if server success to update (body data is NULL).
	- Other status code if server failed to generate password, not implement
		- 400 Bad request
			- if there is no user.

### DELETE /api/userAccount

Delete UserAccount

- Request format:
	- body data is following:
		```json
		{
			"username": "user name information",
		}
		```
	- User autorization is by Digest autorization.

- Response:
	- HTTP 200 OK if server success to update (body data is NULL).
	- Other status code if server failed to generate password, not implement
		- 400 Bad request
			- if there is no user.

- Note:
	- All of password information which is created by the user are deleted.