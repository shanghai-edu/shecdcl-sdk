var api = require('./api');

(async() => {
  let jkm = await api.GetSjSsmjm('王玉平', '身份证号码HERE');
  console.log(jkm);
})();
