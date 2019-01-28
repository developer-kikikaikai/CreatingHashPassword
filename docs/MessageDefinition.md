# HTTP API definition

All of request/response format  are JSON

## passphrase /api/passphrase

This is for generating passphrase.

### POST /api/passphrase

- Request format:
	```json
	{
		"title": "string of this passphrase's title",
		"keyphrase": "keyphrase string",
		"algorithm": "generating passphrase algorithm, write hash string",
		"seed": "seed value of string to generate passphrase",
		"length": length of generating passphrase
		"disable_symbol": length of generating passphrase, true or false
	}
	```

- Response:
	- HTTP 200 OK if server generate passphrase:
		```json
		{
			"result": "passphrase result string"
		}
		```
	- (Other status code if server failed to generate passphrase, not implement)

### Others

- GET/PUT/DELETE api/passphrase are Nothing

## passphraseInformation /api/passphraseInfo

This is for login user's passphrase information.  
User can store base of generating passphrase information.

### GET /api/passphraseInfo

- Request format:
	- User autorization is by Digest autorization.

- Response:
	- HTTP 200 OK if server generate passphrase:
		```json
		[
			{
				"title": "title"
				"algorithm": "generating passphrase algorithm, write hash string",
				"seed": "seed value of string to generate passphrase",
				"length": length of generating passphrase,
				"disable_symbol": length of generating passphrase, true or false
			},
			...
		]
		```
	- (Other status code if server failed to generate passphrase, not implement)

### POST /api/passphraseInfo

- Nothing (implement all in PUT)

### PUT /api/passphraseInfo

Create/Update passphraseInformation

- Request format:
	- body data is following:
		```json
		[
			{
				"title": "title"
				"algorithm": "generating passphrase algorithm, write hash string",
				"seed": "seed value of string to generate passphrase",
				"length": length of generating passphrase,
				"disable_symbol": length of generating passphrase, true or false
			},
			...
		]
		```
	- User autorization is by Digest autorization.

- Response:
	- HTTP 200 OK if server success to update (body data is NULL).
	- Other status code if server failed to generate passphrase, not implement
		- 400 Bad request
			- if there is no algorithm.

### DELETE /api/passphraseInfo

Delete passphraseInformation

- Request format:
	- body data is following:
		```json
		[
			"value of title",
			...
		]
		```
	- User autorization is by Digest autorization.

- Response:
	- HTTP 200 OK if server success to update:
		```json
		[
			"deleted title",
			...
		]
		```
	- Other status code if server failed to generate passphrase, not implement
		- 400 Bad request
			- if there is no data.

## UserAccount /api/userAccount

This is for generating/updating/deleting login user account.

### GET/POST /api/userAccount

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
       - Other status code if server failed to generate passphrase, not implement
               - 400 Bad request
                       - if there is already used username.


### PUT /api/userAccount

Update UserAccount

- Request format:
	- body data is following:
		```json
		{
			"passphrase": "passphrase string"
		}
		```
	- User autorization is by Digest autorization.

- Response:
	- HTTP 200 OK if server success to update (body data is NULL).
	- Other status code if server failed to generate passphrase, not implement
		- 400 Bad request
			- if there is no user.

### DELETE /api/userAccount

Delete UserAccount

- Request format:
	- no body data
	- User autorization is by Digest autorization.

- Response:
	- HTTP 200 OK if server success to update (body data is NULL).

- Note:
	- All of passphrase information which is created by the user are deleted.

## Logout

There is no "Login/Logout" design in HTTP digest authorization.
So user can use this api to remove login information from browser.

### GET /api/logout

- Request format:
	no body data

- Response
	- HTTP 401 Unauthorized
