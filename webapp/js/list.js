window.onload = function () {
  $q('#time').value = new Date().Format('yyyy-MM-dd');
  if (getStorage('search_time')) $q('#time').value = getStorage('search_time')
  search()

  if (window.VConsole) new VConsole()
}

function search() {
  setStorage('search_time', $q('#time').value)
  let time = $q('#time').value
  let user = $q('#user').value
  let cont = $q('#cont').value
  let limit = $q('#limit').value
  axios({
    url: `/api/report/list?time=${time}&user=${user}&cont=${cont}&limit=${limit}`,
    method: 'get'
  }).then(function (res) {
    if (res && res.data.length) {
      window.nowReports = res.data
      $q('#reports').innerHTML = res.data.map(function (item, index) {
        return `<li><div onclick="chooseThis(this)"><input type="checkbox" name="chooseReport" value="${item.id}" />${item.time.substr(5)}: <b>${item.user}</b>:  ${item.cont} <a href="./?id=${item.id}">修改</a> <a href="#" onclick="del(${item.id})">删除</a></div></li>`
      }).join('')
    } else {
      $q('#reports').innerHTML = ''
    }
  })
}

function resetSearch() {
  $q('#time').value = '';
  $q('#user').value = '';
  $q('#cont').value = '';
  $q('#limit').value = 50;
}

function today() {
  $q('#time').value = new Date().Format('yyyy-MM-dd');
  $q('#limit').value = 50;
}

function chooseThis(obj) {
  var e = event || window.event
  if (e.target ==  obj.children[0]) return;
  obj.children[0].checked = !obj.children[0].checked
}

function selectAll() {
  var els = $qa("input[name='chooseReport']")
  var resu = !els[0].checked
  Array.from(els).map(function(el) {
    el.checked = resu
  })
}

function edit(id) {
  console.log(id);
}
function del(id) {
  console.log(id);
  var e = event || window.event
  e.preventDefault()
  e.stopPropagation()
  var sign = window.prompt('确定删除吗？','是')
  if (sign && sign != '否') {
    axios({
      url: `/api/report/del?id=${id}`,
      method: 'delete'
    }).then(function () {
      window.location.reload()
    })
  }
}

function copyReport() {
  var choosed = $qa("input[name='chooseReport']:checked")
  if (choosed.length && nowReports && nowReports.length) {
    let count = 0
    let txt = Array.from(choosed).map(function (el, ind) {
      let one = nowReports.find(r => r.id == el.value)
      return one.cont.split(/(\r|\n)/ig).filter(i => !!i.trim()).map(o => {
        count++
        return `${count}、${o}[${one.user}]`
      }).join('\r\n')
      // return `${count}、${one.cont} [${one.user}]`
    }).join('\r\n')
    console.log(txt);
    $copy(txt).catch(() => alert('复制失败!'))
  }
}

bindListener(Array.from($qa('.enter')), 'keyup', 13, search)