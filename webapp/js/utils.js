window.baseURL = window && window.location ? 
                          `${window.location.protocol}//${window.location.hostname}:3457`
                          : 'http://localhost:3457'
window.request = axios.create({
  baseURL: baseURL,
  timeout: 1000,
  headers: {'X-Custom-Header': 'foobar'}
});
request.interceptors.response.use(function (response) {
  // 对响应数据做点什么
  // if (app) app.$message.success('操作成功!')
  console.log(response);
  return response;
}, function (error) {
  // 对响应错误做点什么
  if (app) app.$message.error('操作失败, 请重试!')
  return Promise.reject(error);
});

/************** utils  **********************/
Date.prototype.Format = function (fmt) {  
  var o = {  
      "M+": this.getMonth() + 1,
      "d+": this.getDate(),
      "H+": this.getHours(),
      "m+": this.getMinutes(),
      "s+": this.getSeconds(),
      "q+": Math.floor((this.getMonth() + 3) / 3),
      "S": this.getMilliseconds()
  };  
  if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));  
  for (var k in o)  
  if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));  
  return fmt;  
} 

function $q(sel) {
  return document.querySelector(sel)
}

function $qa(sel) {
  return document.querySelectorAll(sel)
}

function getType(value) {
  return Object.prototype.toString.call(value).toLowerCase()
}

function isString(value) {
  return getType(value) === '[object string]'
}

function setStorage(key, value) {
  if (window && window.localStorage) {
    if (isString(value)) return window.localStorage.setItem(key, value)
    window.localStorage.setItem(key, JSON.stringify(value))
  } else {
    console.log('window.localStorage not supported...');
  }
}

function getStorage(key) {
  if (window && window.localStorage) {
    let value = window.localStorage.getItem(key)
    if (!value) return value
    try {
      return JSON.parse(value)
    } catch (e) {
      return value
    }
  } else {
    console.log('window.localStorage not supported...');
  }
}

function getValueType(value) {
	return Object.prototype.toString.call(value).slice(8, -1);
}

function parseFloatMy(value) {
	if (getValueType(value) == 'Undefined' || value == '' || isNaN(value)) return 0;
	return parseFloat(value);
}

function parseIntMy(value) {
	if (getValueType(value) == 'Undefined' || value == '' || isNaN(value)) return 0;
	return parseInt(value);
}

/**
 * @param {string} url
 * @returns {Object}
 */
function getAllQueryObject(url) {
  var _url = url ? url : window.location.href
  let search = _url.substring(_url.indexOf('?') + 1)
  search = search.replace(/#/g, '?')
  const obj = {}
  const reg = /([^?&=]+)=([^?&=]*)/g
  search.replace(reg, (rs, $1, $2) => {
    const name = decodeURIComponent($1)
    let val = decodeURIComponent($2)
    val = String(val)
    obj[name] = val
    return rs
  })
  return obj
}

function $copy(text, container) { 
  return new Promise(function (resolve, reject) {
    var fakeElement = document.createElement('button')
    var clipboard = new ClipboardJS(fakeElement, {
      text: function () { return text + '' },
      action: function () { return 'copy' },
      // container: typeof container === 'object' ? container : document.body
    })
    clipboard.on('success', function (e) {
      clipboard.destroy()
      resolve(e)
    })
    clipboard.on('error', function (e) {
      clipboard.destroy()
      reject(e)
    })
    document.body.appendChild(fakeElement)
    fakeElement.click()
    document.body.removeChild(fakeElement)
    setTimeout(() => {
    }, 100);
  })
}

function bindListener(ele, ename, keyCode, callback) {
  if (!ele) return;
  let eles = [].concat(ele)
  eles.forEach(function(item) {
    item.addEventListener(ename, function (e) {
      // console.log(e);
      var currKey=0, e=e||event||window.event;
      currKey = e.keyCode||e.which||e.charCode;
      if (currKey == keyCode || ename.indexOf('key') == -1) callback && callback()
    })
  })
}