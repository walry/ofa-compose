import axios from 'axios'
import { Message } from 'element-ui'

var showLoginMessageBox = false
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8'
// 创建axios实例
// let hrefs = []
// if (window.location.href.indexOf("index.html") != -1) {
//   hrefs = window.location.href.split('index.html')
// } else {
//   hrefs = window.location.href.split('#')
// }
// let baseURL = hrefs.length > 0 ? hrefs[0] : window.location.href
// baseURL + 'index.php/' 默认请求地址
// process.env.BASE_API 自定义请求地址

window.BASE_URL = '/'

const service = axios.create({
  baseURL: window.BASE_URL, // api 的 base_url
  timeout: 180000 // 请求超时时间
})
let silentError = false // 静默错误信息，由上一层手动处理

// request拦截器
service.interceptors.request.use(
  config => {
    if (config.silentError !== undefined) {
      silentError = config.silentError
      delete config.silentError
    }
    return config
  },
  error => {
    // Do something with request error
    return Promise.reject(error)
  }
)

// response 拦截器
service.interceptors.response.use(
  response => {
    /**
     * code为非200是抛错
     */
    const res = response.data
    if (response.status === 200 && response.config.responseType === 'blob') { // 文件类型特殊处理
      return response
    } else if (res.code !== 0) {
      let errMsg = ''
      // 101	登录已失效 102	没有权限 103	账号已被删除或禁用
      if (res.code === 101) {
        if (!showLoginMessageBox) {
          showLoginMessageBox = true
        }
      } else if (res.code === 402) {
        console.log('err')
      } else {
        errMsg = res.error || res.message || res.msg
        if (errMsg) {
          if (!silentError) {
            Message({
              message: errMsg,
              type: 'error'
            })
          }
        }
      }
      return Promise.reject(errMsg)
    } else {
      return res
    }
  },
  error => {
    if (!silentError) {
      Message({
        message: '网络请求失败，请稍候再试. ' + error.toString(),
        type: 'error'
      })
      console.error(error)
    }
    return Promise.reject(error)
  }
)

export default service