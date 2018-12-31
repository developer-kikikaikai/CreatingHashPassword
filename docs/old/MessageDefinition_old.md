## message definition

All of request/response format  are JSON

### Login

#### Base information:

This is for checking user.

- url
	- api/login

#### Base information:

### Login

- Get request format:

	{  
	  "client_keycode":value of client  
	}

	- client_keycode is a value of under 30 bit value for Diffie-Hellman key exchange algorithm.

	```
	Diffie-Hellman key exchange
	1. Set the p, which is a prime number (normaly this value is defined in [RFC3526](https://tools.ietf.org/html/rfc3526)).
	2. And set a number g in Z/pZ (={n in Z|n = m mod p for all m in Z}).
	3. Let a( and b) in {n in Z | 0 <= n <= p-2} are a value as Client( and server) side private key.
	4. At that time, Client and Server are exchange following values:
	Client:
		A = g^a mod p
	Server:
		B = g^b mod p
	5. Then Client and Server can get the same value A^b mod p = B^a mod p = g^(a*b) mod p. And it's not easy to get others.
	So only Client and Server can know same key from shared value A and B.
	```

	In this system, define the value p is the Mersenne prime "2^31 - 1".
	(it's better to use a prime which bit length is over 2048, please see [RFC3526](https://tools.ietf.org/html/rfc3526) for more secure key exchange).

	- User is checked by using Digest autorization.

    - After using it, all of POST data is encrypted by g^(a*b) mod p key value.

- Get request format:

- request body:
	- api/login
	- api/login
