'use strict';
var $ = require('jquery');

module.exports = class APIResult {
	constructor(resolve, reject) {
		this.resolve = resolve;
		this.reject = reject;
	}
}
