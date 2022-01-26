var api = require('./api');

(async() => {
  let jkm = await api.GetGjXgymjzxx('王玉平', '身份证号码');
  console.log(jkm);
})();
