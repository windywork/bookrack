function namespaceChange(obj) {
  const namespace = $(obj).val()
  $.cookie('namespace', namespace)
  window.location.reload()
}

function searchKeyChange(obj) {
  const searchKey = $(obj).val()
  $.cookie('searchKey', searchKey)
  console.log('hello')
}

function initCurrentNav() {
  $('li', '.main-header').removeClass('uk-active')
  const aList = $('li>a', '.main-header')
  for (let item of aList) {
    if ($(item).attr('href') === location.pathname) {
      $(item).parent().addClass('uk-active')
    }
  }
}

function initNsSelect() {
  const namespace = $.cookie('namespace')
  $('option', '.namespace-select').attr("selected", false)
  $('option[value=' + namespace + ']', '.namespace-select').attr("selected", true)
}

window.TTYD_URL = '/ttyd'

function entryTerminal() {
  const el = '#modal-terminal'
  $('iframe', el).attr('src', window.TTYD_URL)
  UIkit.modal(el).show()
}

function writeKubeConfig(el) {
  if (!el) {
    el = '#modal-reset-config'
  }
  UIkit.modal.prompt('请输入 [ OK ] 确认重写: ', '').then(function (val) {
    if (val === 'OK') {
      $.ajax({
        url: '/write-kube-config',
        type: 'POST',
        dataType: 'json',
        data: $('textarea', el).val(),
        success:function(req){
          if (!req.code) {
            location.href = '/'
          } else {
            UIkit.modal.alert('重写错误:' + req.message)
          }
        },
        error: function(err){
          UIkit.modal.alert('重写错误:' + err.responseText)
        }
      })
    } else if (val != null) {
      UIkit.modal.alert('输入错误, 无法重写!')
    }
  })
}

function kubeDeletePod() {
  UIkit.modal.prompt('请输入 [ DELETE ] 确认删除: ', '').then(function (val) {
    if (val === 'DELETE') {
      console.log('Delete Pod:', val)
    } else if (val != null){
      UIkit.modal.alert('输入错误,无法删除!')
    }
  })
}
function kubeCreateNamespace() {
  UIkit.modal.prompt('请输入创建的空间名称: ', '').then(function (val) {
      if (val != null && val != '') {
        UIkit.modal.alert('创建成功!')
      }
  })
}
function kubeDeleteNamespace() {
  UIkit.modal.prompt('请输入 [ DELETE ] 确认删除: ', '').then(function (val) {
    if (val === 'DELETE') {
      console.log('Delete Namespace:', val)
    } else if (val != null){
      UIkit.modal.alert('输入错误,无法删除!')
    }
  })
}
function kubeLogs(name, ns) {
  var el = '#modal-view-log'
  var cmd = 'kubectl logs --tail=800 -f -n ' + ns + ' ' + name
  var url = window.TTYD_URL + '?arg=' + $.base64.encode(cmd)
  $('iframe', el).attr('src', url)
  $('.uk-modal-header>.uk-modal-title', el).text('查看日志:' + name)
  UIkit.modal(el).show()
}

function kubeExec(name, container, ns) {
  var el = '#modal-entry-container'
  var cmd = 'kubectl exec -it -n ' + ns + ' ' + name + ' sh'
  var url = window.TTYD_URL + '?arg=' + $.base64.encode(cmd)
  $('iframe', el).attr('src', url)
  $('.uk-modal-header>.uk-modal-title', el).text('执行命令:' + name)
  UIkit.modal(el).show()
}



jQuery(() => {
  initCurrentNav()
  initNsSelect()
  UIkit.util.on('#modal-view-log,#modal-terminal,#modal-entry-container', 'hide', function (e) {
      $('iframe', e.target).attr('src', '')
  })
})