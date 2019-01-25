const API = require('./api.js');
const APIResult = require('./api_result.js');
const Element = require('./view.js');
var api = new API();
var element = new Element();

var $ = require('jquery');

/*****************************/
//function for generate passphrase
//result: response json format
function generatePassphrase_resolved(result) {
	document.getElementById('result_In_Generated_passphrase').value = result.result
}
function generatePassphrase_reject(result) {
	alert(result)
}
//generate passphrase event
$(document).on('click', '[id="SubmitGeneratePassphrase"]', function(){
	var title = document.getElementById('title_In_PassphraseInformation').value
	var keyphrase = document.getElementById('keyphrase_In_PassphraseInformation').value
	var algorithmSelect = document.getElementById('algorithmSelect_In_PassphraseInformation').value
	var extraInfo = document.getElementById('extraInfo_In_PassphraseInformation').value
	var length = Number(document.getElementById('maxLength_In_PassphraseInformation').value)
	var disableSymbol = ($("[name=UseSymbol_In_PassphraseInformation]").prop("checked") != true)
	api.generatePassphrase(title, keyphrase, algorithmSelect, extraInfo, length, disableSymbol, new APIResult(generatePassphrase_resolved, generatePassphrase_reject));
});

/*****************************/
//function for save setting
//result: response json format
function savePassphraseInfo_resolved(result) {
	alert("Success to save passphrase setting")
}

function savePassphraseInfo_reject(result) {
	alert(result)
}
//save setting event
$(document).on('click', '[id="SubmitSaveSetting"]', function(){
	var title = document.getElementById('title_In_PassphraseInformation').value
	var algorithmSelect = document.getElementById('algorithmSelect_In_PassphraseInformation').value
	var extraInfo = document.getElementById('extraInfo_In_PassphraseInformation').value
	var length = Number(document.getElementById('maxLength_In_PassphraseInformation').value)
	var disableSymbol = ($("[name=UseSymbol_In_PassphraseInformation]").prop("checked") != true)
	api.savePassphraseInfo(title, algorithmSelect, extraInfo, length, disableSymbol, new APIResult(savePassphraseInfo_resolved, savePassphraseInfo_reject));
});

/*****************************/
//function for save setting
//result: response json format
function createUser_resolved(result) {
	alert("Success to save passphrase setting")
}
function createUser_reject(result) {
	alert(result)
}
//Create user event
$(document).on('click', '[id="SubmitCreateUser_In_CreateUser"]', function(){
	var user = document.getElementById('Username_In_CreateUser').value
	var pass = document.getElementById('LoginPassphrase_In_CreateUser').value
	api.createUser(user, pass, new APIResult(createUser_resolved, createUser_reject))
});

/*****************************/
//Update user event
function updateUser_resolved(result) {
	alert("Success to update passphrase setting")
}
function updateUser_reject(result) {
	alert(result)
}
//Create user event
$(document).on('click', '[id="SubmitUpdateUser_In_UpdateUser"]', function(){
	var pass = document.getElementById('LoginPassphrase_In_UpdateUser').value
	api.updateUser(pass, new APIResult(updateUser_resolved, updateUser_reject))
});

/*****************************/
//Delete user event
function deleteUser_resolved(result) {
	alert("Success to delete passphrase setting")
}
function deleteUser_reject(result) {
	alert(result)
}
//Create user event
$(document).on('click', '[id="SubmitDeleteUser_In_UpdateUser"]', function(){
	api.deleteUser(new APIResult(deleteUser_resolved, deleteUser_reject))
});

/*****************************/
//Get setting event
function getPassphraseInfo_resolved(result) {
	$("#user_passphrase_settings_In_PassphraseSettings").empty()

	if(!Object.keys(result).length) {
		return
	}

	result.forEach(function(item, index, array) {
		title_id = item.title + '_In_PassphraseSetting'
		html = '<tr><td algorithm="'+item.algorithm+'" extra="' +item.seed+ '" maxlength="' + item.length + '" disable_symbol="' + item.disable_symbol + '" id="' + title_id + '"> '+ item.title + '</td>'
		$("#user_passphrase_settings_In_PassphraseSettings").append(html)

		//add event
		$(document).off('click', '[id="' + title_id + '"]')
		$(document).on('click', '[id="' + title_id + '"]', function(event) {
			title = event.currentTarget.innerText
			document.getElementById('title_In_PassphraseInformation').value = title
			title_element = document.getElementById(title + '_In_PassphraseSetting')
			document.getElementById('algorithmSelect_In_PassphraseInformation').value = title_element.getAttribute('algorithm')
			document.getElementById('extraInfo_In_PassphraseInformation').value = title_element.getAttribute('extra')
			document.getElementById('maxLength_In_PassphraseInformation').value = title_element.getAttribute('maxlength')
			document.getElementById('UseSymbol_In_PassphraseInformation').checked = (title_element.getAttribute('disable_symbol') != 'true')
		});
	})
}
function getPassphraseInfo_reject(result) {
	alert(result)
}
//Get setting event
$(document).on('click', '[id="SubmitGetSetting_In_PassphraseSettings"]', function(){
	api.getPassphraseInfo(new APIResult(getPassphraseInfo_resolved, getPassphraseInfo_reject));
});

/*****************************/
//Get setting event
function logout_resolved(result) {
	alert("Success to update passphrase setting")
}
function logout_reject(result) {
}
//Get setting event
$(document).on('click', '[id="SubmitLogoutUser_In_UpdateUser"]', function(){
	api.logout(new APIResult(logout_resolved, logout_reject));
});

/*****************************/
//Set select event
function getMaxLength(algorithm) {
	switch(algorithm) {
		case "sha256":
		case "sha512_256":
		case "black2s_256":
		case "sha3_256":
			return [64,32,16];
		case "sha384":
		case "sha3_384":
		case "black2s_384":
			return [96,48,32,16];
		case "sha512":
		case "sha3_512":
		case "black2s_512":
			return [128,64,32,16];
	}
}
function updateMaxLengthField() {
	var algorithmSelect = document.getElementById('algorithmSelect_In_PassphraseInformation').value
	$("#maxLength_In_PassphraseInformation").empty()
	getMaxLength(algorithmSelect).forEach(function(length) {
		var html = "<option>" + length + "</option>"
		$("#maxLength_In_PassphraseInformation").append(html)
	})
}
$(document).on('change', '[id="algorithmSelect_In_PassphraseInformation"]', updateMaxLengthField)
window.onload = function() {
	updateMaxLengthField();
}
