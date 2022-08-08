function getOption(id) {
  return new Promise((resolve, reject) => {
    request.get('/api/tasklog/list?limit=500' + (id && id > 0 ? ('&taskId=' + id) : '')).then(res => {
      if (res.data.length) {
        let logs = res.data
        let legends = {}
        let x = []
        logs.map(item => {
          if (!legends[item.taskId]) legends[item.taskId] = item.name
          let time = item.createdAt.substring(4,19).replace(/[-]/ig, '').replace('T', '_')
          item.time = time
          if (x.indexOf(time) === -1) x.push(time)
        })
        series = Object.keys(legends).map(key => {
          return {
            name: legends[key],
            type: 'line',
            connectNulls: true,
            data: x.map(item => {
              let one = logs.find(log => log.taskId == key && log.time == item)
              return one && one.duration ? one.duration : null
            })
          }
        })
        console.log(legends);
        console.log(x);
        // https://echarts.apache.org/examples/zh/editor.html?c=line-stack
        const option = {
          title: {
            text: '任务明细记录'
          },
          tooltip: {
            trigger: 'axis'
          },
          legend: {
            selector: true,
            data: Object.values(legends)
          },
          grid: {
            left: '3%',
            right: '4%',
            bottom: '3%',
            containLabel: true
          },
          toolbox: {
            feature: {
              saveAsImage: {}
            }
          },
          xAxis: {
            type: 'category',
            boundaryGap: false,
            data: x.reverse() //['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
          },
          yAxis: {
            type: 'value',
            // max: 1000,
            axisLine: {
              show: true
            }
          },
          series: series
        }
    
        resolve(option)
        
      } else {
        resolve({})
      }
    })
  });
}