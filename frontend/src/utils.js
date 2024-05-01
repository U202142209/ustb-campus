/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-10-15 18:16:51
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-10-22 14:10:02
 * @FilePath: \ustb-campus\frontend\src\utils.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// utils.js

import CryptoJS from 'crypto-js';

// 加密删除帖子的id
export function jiama(s) {
  console.log("加密前的文字:" + s)
  // 设置密钥和初始向量
  var i = "1234512345123456";
  var c = "1234512345123456";
  // 将密钥和初始向量转换为Utf8格式
  i = CryptoJS.enc.Utf8.parse(i);
  c = CryptoJS.enc.Utf8.parse(c);
  // 构造待加密的字符串
  var r = '{"id":"' + s + '"}';
  // 使用AES算法进行加密
  var l = CryptoJS.AES.encrypt(
    r, i, {
    iv: c,
    mode: CryptoJS.mode.ECB,
    padding: CryptoJS.pad.Pkcs7
  });

  // 将加密后的结果转换为大写字符串
  l = l.ciphertext.toString().toUpperCase();
  // 输出加密结果
  console.log("加密后端文字:" + l);
  return l;
}


export function jiemi(s) {
  // 设置密钥和初始向量
  var i = "1234512345123456";
  var c = "1234512345123456";
  // 将密钥和初始向量转换为Utf8格式
  i = CryptoJS.enc.Utf8.parse(i);
  c = CryptoJS.enc.Utf8.parse(c);
  // 将加密后的结果转换为字节数组
  var ciphertext = CryptoJS.enc.Hex.parse(s);
  // 使用AES算法进行解密
  var decrypted = CryptoJS.AES.decrypt(
    { ciphertext: ciphertext },
    i,
    { iv: c, mode: CryptoJS.mode.ECB, padding: CryptoJS.pad.Pkcs7 }
  );
  // 将解密结果转换为字符串
  var plaintext = decrypted.toString(CryptoJS.enc.Utf8);
  // 输出解密结果
  console.log("解密后的文字:" + plaintext);
  return plaintext;
}

// 拆分图片链接
export function splitImageUrl(links) {
  if (!links) {
    return []
  }
  const regex = /(http:\/\/[^,]+\.jpg)/g;
  return links.match(regex);
}

// 加密发布帖子的内容
export function encroyptedContent(content, c_time = 123) {
  // 设置密钥和初始向量
  var i = "1234512345123456";
  var c = "1234512345123456";
  // 将密钥和初始向量转换为Utf8格式
  i = CryptoJS.enc.Utf8.parse(i);
  c = CryptoJS.enc.Utf8.parse(c);
  // 构造待加密的字符串
  var r = '{"content":"' + content + '","title":"' + content.replace(/\s+/g, "") + '","verify":"zzyq","c_time":"' + c_time + '"}'
  // 使用AES算法进行加密
  var l = CryptoJS.AES.encrypt(
    r, i,
    {
      iv: c,
      mode: CryptoJS.mode.ECB,
      padding: CryptoJS.pad.Pkcs7
    });

  // 将加密后的结果转换为大写字符串
  l = l.ciphertext.toString().toUpperCase();
  // 输出加密结果
  console.log("加密原文："+content)
  console.log("加密时间："+c_time)
  console.log("加密后端文字:" + l);
  return l;
}