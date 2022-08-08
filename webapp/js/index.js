queryObj = getAllQueryObject()
window.isNew = true
window.inSubmit = false

function submitCont() {
  if (inSubmit) return;
  inSubmit = true
  var _data = {
    id: parseIntMy($q('#id').value),
    user: $q('#user').value,
    cont: $q('#cont').value,
    time: isNew ? new Date().Format('yyyy-MM-dd') : $q('#time').value
  }
  if (!_data.user) return alert('请输入姓名!') || $q('#user').focus()
  if (!_data.cont) return alert('请输入内容!') || $q('#cont').focus()
  axios({
    url: '/api/report/save',
    method: 'post',
    data: _data
  }).then(function (res) {
    console.log(res);
    if (res.data && res.data.id && isNew) setStorage('report_today', res.data)
    inSubmit = false
    alert('提交成功!')
  }).catch(function (e) {
    inSubmit = false
    if (isNew) setStorage('report_today', _data)
    alert('提交失败, 请重试!')
  })
}

function resetCont() {
  $q('#cont').value = ''
}

window.onload = function () {
  var report = getStorage('report_today')
  if (report && report.user) {
    $q('#id').value = report.id
    $q('#user').value = report.user
    $q('#cont').value = report.cont
    if (report.time != new Date().Format('yyyy-MM-dd')) $q('#id').value = 0
  }

  if (queryObj.id) {
    axios({
      url: '/api/report/get?id=' + queryObj.id,
      method: 'get'
    }).then(function (res) {
      console.log(res);
      if (res.data && res.data.id) {
        isNew = false
        $q('#timewrap').style.display = 'block'
        $q('#id').value = res.data.id
        $q('#user').value = res.data.user
        $q('#cont').value = res.data.cont
        $q('#time').value = res.data.time
      }
    }).catch(function (e) {
      alert('查询失败, 请重试!')
    })
  }
}