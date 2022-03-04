module.exports = {
  appId: process.env.SHECGW_APPID || '',
  appSecret: process.env.SHECGW_APPSECRET || '',
  apiGwEndPoint: process.env.SHECGW_APIGWENDPOINT || 'https://apigw.shec.edu.cn',
  debug: process.env.DEBUG || false,
  unitId : process.env.UNITID || '',
};
