var api = require('./api');

(async() => {
  let jkm = await api.GetGjXgymjzxx('王*平', '身份证号码');
  console.log(jkm);
})();
