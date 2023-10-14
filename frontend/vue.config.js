/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-10-08 17:35:10
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-10-10 13:22:04
 * @FilePath: \ustb-campus\frontend\vue.config.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

const { defineConfig } = require('@vue/cli-service')

module.exports = defineConfig({
  transpileDependencies: true,
  configureWebpack: {
    devServer: {
      proxy: {
        "/proxy_url": {
          target: 'http://ustb.campus.work.wobushidalao.top',
          ws: true,
          changeOrigin: true,
          pathRewrite: {
            '^/proxy_url': '/'
          }
        }
      }
    }
  }
})

// const { defineConfig } = require('@vue/cli-service')
// module.exports = defineConfig({
//   transpileDependencies: true,
//   devServer: {
//     proxy: {
//       "/proxy_url": {           // /proxy_url 这个用来和根路径 baseURL 进行匹配
//         // target: 'http://127.0.0.1:8000',  // 这个是填写跨域的请求域名+端口号，也就是要请求的URL(不包含URL路径)
//         target: 'http://ustb.campus.work.wobushidalao.top',  // 这个是填写跨域的请求域名+端口号，也就是要请求的URL(不包含URL路径)
//         ws:true,//代理websocked
//         changeOrigin: true,  // 是否允许跨域请求，在本地会创建一个虚拟服务端，然后发送请求的数据，并同时接收请求的数据，这样服务端和服务端进行数据的交互就不会有跨域问题
//         pathRewrite: {   // 路径重写
//           '^/proxy_url': '/' // 替换target中的请求地址，原请求为 http://127.0.0.1:8000/kuayu 实际请求为 http://127.0.0.1:8000/proxy_url/kuayu
//         }
//       }
//     }
//   }
// })
