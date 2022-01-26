var api = require('./api');

(async ()=>{
let token = await api.getAccessToken();

console.log(token);
})();
