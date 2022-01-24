'use strict';

// 国家信息
const gjxxTable = {
  "004": "阿富汗",
  "008": "阿尔巴尼亚",
  "012": "阿尔及利亚",
  "016": "美属萨摩亚",
  "020": "安道尔",
  "024": "安哥拉",
  "660": "安圭拉",
  "010": "南极洲",
  "028": "安提瓜和巴布达",
  "032": "阿根廷",
  "051": "亚美尼亚",
  "533": "阿鲁巴",
  "036": "澳大利亚",
  "040": "奥地利",
  "031": "阿塞拜疆",
  "044": "巴哈马",
  "048": "巴林",
  "050": "孟加拉国",
  "052": "巴巴多斯",
  "112": "白俄罗斯",
  "056": "比利时",
  "084": "伯利兹",
  "204": "贝宁",
  "060": "百慕大",
  "064": "不丹",
  "068": "玻利维亚",
  "070": "波黑",
  "072": "博茨瓦纳",
  "074": "布维岛",
  "076": "巴西",
  "086": "英属印度洋领土",
  "096": "文莱",
  "100": "保加利亚",
  "854": "布基纳法索",
  "108": "布隆迪",
  "116": "柬埔寨",
  "120": "喀麦隆",
  "124": "加拿大",
  "132": "佛得角",
  "136": "开曼群岛",
  "140": "中非",
  "148": "乍得",
  "152": "智利",
  "156": "中国",
  "162": "圣诞岛",
  "166": "科科斯（基林）群岛",
  "170": "哥伦比亚",
  "174": "科摩罗",
  "178": "刚果（布）",
  "180": "刚果（金）",
  "184": "库克群岛",
  "188": "哥斯达黎加",
  "384": "科特迪瓦",
  "191": "克罗地亚",
  "192": "古巴",
  "196": "塞浦路斯",
  "203": "捷克",
  "208": "丹麦",
  "262": "吉布提",
  "212": "多米尼克",
  "214": "多米尼加共和国",
  "626": "东帝汶",
  "218": "厄瓜多尔",
  "818": "埃及",
  "222": "萨尔瓦多",
  "226": "赤道几内亚",
  "232": "厄立特里亚",
  "233": "爱沙尼亚",
  "231": "埃塞俄比亚",
  "238": "福克兰群岛（马尔维纳斯）",
  "234": "法罗群岛",
  "242": "斐济",
  "246": "芬兰",
  "250": "法国",
  "254": "法属圭亚那",
  "258": "法属波利尼西亚",
  "260": "法属南部领土",
  "266": "加蓬",
  "270": "冈比亚",
  "268": "格鲁吉亚",
  "276": "德国",
  "288": "加纳",
  "292": "直布罗陀",
  "300": "希腊",
  "304": "格陵兰",
  "308": "格林纳达",
  "312": "瓜德罗普",
  "316": "美国（关岛）",
  "320": "危地马拉",
  "324": "几内亚",
  "624": "几内亚比绍",
  "328": "圭亚那",
  "332": "海地",
  "334": "赫德岛和麦克唐纳岛",
  "340": "洪都拉斯",
  "348": "匈牙利",
  "352": "冰岛",
  "356": "印度",
  "360": "印度尼西亚",
  "364": "伊朗",
  "368": "伊拉克",
  "372": "爱尔兰",
  "376": "以色列",
  "380": "意大利",
  "388": "牙买加",
  "392": "日本",
  "400": "约旦",
  "398": "哈萨克斯坦",
  "404": "肯尼亚",
  "296": "基里巴斯",
  "408": "朝鲜",
  "410": "韩国",
  "414": "科威特",
  "417": "吉尔吉斯斯坦",
  "418": "老挝",
  "428": "拉脱维亚",
  "422": "黎巴嫩",
  "426": "莱索托",
  "430": "利比里亚",
  "434": "利比亚",
  "438": "列支敦士登",
  "440": "立陶宛",
  "442": "卢森堡",
  "807": "北马其顿",
  "450": "马达加斯加",
  "454": "马拉维",
  "458": "马来西亚",
  "462": "马尔代夫",
  "466": "马里",
  "470": "马耳他",
  "584": "马绍尔群岛",
  "474": "马提尼克",
  "478": "毛里塔尼亚",
  "480": "毛里求斯",
  "175": "马约特",
  "484": "墨西哥",
  "583": "密克罗尼西亚联邦",
  "498": "摩尔多瓦",
  "492": "摩纳哥",
  "496": "蒙古",
  "500": "蒙特塞拉特",
  "504": "摩洛哥",
  "508": "莫桑比克",
  "104": "缅甸",
  "516": "纳米比亚",
  "520": "瑙鲁",
  "524": "尼泊尔",
  "528": "荷兰",
  "530": "荷属安的列斯",
  "540": "新喀里多尼亚",
  "554": "新西兰",
  "558": "尼加拉瓜",
  "562": "尼日尔",
  "566": "尼日利亚",
  "570": "纽埃",
  "574": "诺福克岛",
  "580": "北马里亚纳",
  "578": "挪威",
  "512": "阿曼",
  "586": "巴基斯坦",
  "585": "帕劳",
  "275": "巴勒斯坦",
  "591": "巴拿马",
  "598": "巴布亚新几内亚",
  "600": "巴拉圭",
  "604": "秘鲁",
  "608": "菲律宾",
  "612": "皮特凯恩群岛",
  "616": "波兰",
  "620": "葡萄牙",
  "630": "波多黎各",
  "634": "卡塔尔",
  "638": "留尼汪",
  "642": "罗马尼亚",
  "643": "俄罗斯联邦",
  "646": "卢旺达",
  "654": "圣赫勒拿",
  "659": "圣基茨和尼维斯",
  "662": "圣卢西亚",
  "666": "圣皮埃尔和密克隆",
  "670": "圣文森特和格林纳丁斯",
  "882": "萨摩亚",
  "674": "圣马力诺",
  "678": "圣多美和普林西比",
  "682": "沙特阿拉伯",
  "686": "塞内加尔",
  "690": "塞舌尔",
  "694": "塞拉利昂",
  "702": "新加坡",
  "703": "斯洛伐克",
  "705": "斯洛文尼亚",
  "090": "所罗门群岛",
  "706": "索马里",
  "710": "南非",
  "239": "南乔治亚岛和南桑德韦奇岛",
  "724": "西班牙",
  "144": "斯里兰卡",
  "736": "苏丹",
  "740": "苏里南",
  "744": "斯瓦尔巴群岛",
  "748": "斯威士兰",
  "752": "瑞典",
  "756": "瑞士",
  "760": "叙利亚",
  "762": "塔吉克斯坦",
  "834": "坦桑尼亚",
  "764": "泰国",
  "768": "多哥",
  "772": "托克劳",
  "776": "汤加",
  "780": "特立尼达和多巴哥",
  "788": "突尼斯",
  "792": "土耳其",
  "795": "土库曼斯坦",
  "796": "特克斯科斯群岛",
  "798": "图瓦卢",
  "800": "乌干达",
  "804": "乌克兰",
  "784": "阿联酋",
  "826": "英国",
  "840": "美国",
  "581": "美国本土外小岛屿",
  "858": "乌拉圭",
  "860": "乌兹别克斯坦",
  "548": "瓦努阿图",
  "336": "梵蒂冈",
  "862": "委内瑞拉",
  "704": "越南",
  "092": "英属维尔京群岛",
  "850": "美属维尔京群岛",
  "876": "瓦利斯和富图纳",
  "732": "西撒哈拉",
  "887": "也门",
  "891": "前南斯拉夫",
  "894": "赞比亚",
  "716": "津巴布韦",
  "688": "塞尔维亚",
  "499": "黑山共和国",
  "728": "南苏丹"
};

// 证件类型
const zjlxTable = {
  "01": "居民身份证",
  "02": "居民户口薄",
  "03": "护照",
  "04": "军官证",
  "05": "驾驶证",
  "06": "港澳居民来往内地通行证",
  "07": "台湾居民来往内地通行证",
}

// 疫苗名称
const ymmcTable = {
  "5601": "新型冠状病毒灭活疫苗（Vero细胞）",
  "5602": "重组新型冠状病毒疫苗（腺病毒载体）",
  "5603": "重组新型冠状病毒疫苗（CHO细胞）",
}

// 疫苗生产企业代码表
const ymqyTable = {
  "02": "北京生物",
  "10": "武汉生物",
  "11": "成都生物",
  "12": "医科院生物",
  "13": "兰州生物",
  "14": "长春生物",
  "66": "科兴（大连）",
  "68": "北京智飞绿竹",
  "80": "北京科兴中维",
  "81": "康希诺生物",
  "82": "安徽智飞",
}

// 行政区划代码表
const xzqhTable = {
  "11": "北京市",
  "12": "天津市",
  "13": "河北省",
  "14": "山西省",
  "15": "内蒙古自治区",
  "21": "辽宁省",
  "22": "吉林省",
  "23": "黑龙江省",
  "31": "上海市",
  "32": "江苏省",
  "33": "浙江省",
  "34": "安徽省",
  "35": "福建省",
  "36": "江西省",
  "37": "山东省",
  "41": "河南省",
  "42": "湖北省",
  "43": "湖南省",
  "44": "广东省",
  "45": "广西壮族自治区",
  "46": "海南省",
  "50": "重庆市",
  "51": "四川省",
  "52": "贵州省",
  "53": "云南省",
  "54": "西藏自治区",
  "61": "陕西省",
  "62": "甘肃省",
  "63": "青海省",
  "64": "宁夏回族自治区",
  "65": "新疆维吾尔自治区",
  "71": "台湾省",
  "81": "香港特别行政区",
  "82": "澳门特别行政区",
  "66": "新疆兵团",

}

module.exports.getGjxx = function (code) {
  if (gjxxTable.hasOwnProperty(code)) {
    return gjxxTable[code];
  }

  return "N/A (" + code + ")";
}

module.exports.getZjlx = function (code) {
  if (zjlxTable.hasOwnProperty(code)) {
    return zjlxTable[code];
  }

  return 'N/A (' + code + ')';
}

module.exports.getYmmc = function (code) {
  if (ymmcTable.hasOwnProperty(code)) {
    return ymmcTable[code];
  }

  return 'N/A (' + code + ')';
}

module.exports.getYmqy = function (code) {
  if (ymqyTable.hasOwnProperty(code)) {
    return ymqyTable[code];
  }

  return 'N/A (' + code + ')';
}

/**
 * 获取行政区划名称
 * @param {string} code
 * @returns 
 */
module.exports.getXzqh = function (code) {
  if (xzqhTable.hasOwnProperty(code)) {
    return xzqhTable[code];
  }

  return 'N/A (' + code + ')';
}