var config = require('./config');
var GW = require('../shecgw');

module.exports = new GW(config.appId, config.appSecret, config.apiGwEndPoint, config.debug);
