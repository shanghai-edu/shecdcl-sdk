var api = require('./api');

(async() => {
  let jkm = await api.GetGjXgfyHsjcsjfwjk('王*平', '身份证号码HERE');
  console.log(jkm);
})();

(async() => {
  let cyxx = await api.GetHscyxxcx('身份证件号码');
  console.log(cyxx);
})();