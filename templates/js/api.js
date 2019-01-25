'use strict';
var $ = require('jquery');

module.exports = class API {

	constructor() {
		this._var_baseurl = '/api/'
	}

	_sendapi(url, request, apires) {
		fetch(url, request).then( res=> {
			if (res.ok) {
				//body is null=> call resolve with null
				if (Number.parseInt(res.headers.get("Content-Length")) === 0) {
					apires.resolve(null);
					return;
				}

				//otherwise => parse json
				res.json().then(result => {
					apires.resolve(result)
				});
			} else {
				apires.reject("Request failed: " + res.status)
			}
		});//fetch
	}

	generatePassphrase(title, keyphrase, algorithm, seed, length, disable_symbol, apires) {
		this._sendapi(
		//1st param
		this._var_baseurl + 'passphrase',
		//second param
		{
			method: "POST",
			headers: {
				'Accept': 'application/json',
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				'title': title,
				"keyphrase": keyphrase,
				"algorithm": algorithm,
				"seed": seed,
				"length": length,
				"disable_symbol": disable_symbol
			})
		},
		//3rd param
		apires);
	}

	savePassphraseInfo(title, algorithm, seed, length, disable_symbol, apires) {
		this._sendapi(
		//1st param
		this._var_baseurl + 'passphraseInfo',
		//second param
		{
			method: "PUT",
			credentials: "same-origin",
			headers: {
				'Accept': 'application/json',
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				'title': title,
				"algorithm": algorithm,
				"seed": seed,
				"length": length,
				"disable_symbol": disable_symbol
			})
		},
		//3rd param
		apires);
	}

	createUser(user, passphrase, apires) {
		this._sendapi(
		//1st param
		this._var_baseurl + 'userAccount',
		//second param
		{
			method: "POST",
			credentials: "same-origin",
			headers: {
				'Accept': 'application/json',
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				'username': user,
				"passphrase": passphrase
			})
		},
		//3rd param
		apires);
	}

	updateUser(passphrase, apires) {
		this._sendapi(
		//1st param
		this._var_baseurl + 'userAccount',
		//second param
		{
			method: "PUT",
			credentials: "same-origin",
			headers: {
				'Accept': 'application/json',
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				"passphrase": passphrase
			})
		},
		//3rd param
		apires);
	}

	deleteUser(apires) {
		this._sendapi(
		//1st param
		this._var_baseurl + 'userAccount',
		//second param
		{
			method: "DELETE",
			credentials: "same-origin",
			headers: {
				'Accept': 'application/json',
				'Content-Type': 'application/json'
			}
		},
		//3rd param
		apires);
	}

	getPassphraseInfo(apires) {
		this._sendapi(
		//1st param
		this._var_baseurl + 'passphraseInfo',
		//second param
		{
			method: "GET",
			credentials: "same-origin",
			headers: {
				'Accept': 'application/json',
				'Content-Type': 'application/json'
			}
		},
		//3rd param
		apires);
	}

	logout(apires) {
		this._sendapi(
		//1st param
		this._var_baseurl + 'logout',
		//second param
		{
			method: "GET",
			credentials: "same-origin",
			headers: {
				'Accept': 'application/json',
				'Content-Type': 'application/json'
			}
		},
		//3rd param
		apires);
	}
}
