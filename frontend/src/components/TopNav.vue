<!--
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-10-10 11:23:34
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-10-12 18:19:35
 * @FilePath: \ustb-campus\frontend\src\components\TopNav.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
  <div>
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
      <div class="container">
        <a class="text-primary navbar-brand" href="#">
          <router-link to="/">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"
              class="bi bi-backpack3-fill" viewBox="0 0 16 16">
              <path d="M5 10v3h6v-3h-1v.5a.5.5 0 0 1-1 0V10H5Z" />
              <path
                d="M6 2v.341a5.978 5.978 0 0 0-1.308.653l-.416-1.247a1 1 0 0 0-1.749-.284l-.77 1.027a1 1 0 0 0-.149.917l.803 2.407A5.985 5.985 0 0 0 2 8v5.5A2.5 2.5 0 0 0 4.5 16h7a2.5 2.5 0 0 0 2.5-2.5V8c0-.771-.146-1.509-.41-2.186l.801-2.407a1 1 0 0 0-.148-.917l-.77-1.027a1 1 0 0 0-1.75.284l-.415 1.247A5.976 5.976 0 0 0 10 2.34V2a2 2 0 1 0-4 0Zm1 0a1 1 0 0 1 2 0v.083a6.04 6.04 0 0 0-2 0V2Zm5.941 2.595a6.036 6.036 0 0 0-.8-.937l.531-1.595.77 1.027-.5 1.505ZM3.86 3.658a6.031 6.031 0 0 0-.8.937L2.557 3.09l.77-1.027.532 1.595Zm.18 3.772a4 4 0 0 1 7.92 0 .5.5 0 1 1-.99.142 3 3 0 0 0-5.94 0 .5.5 0 1 1-.99-.142ZM4 9.5a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5h-7a.5.5 0 0 1-.5-.5v-4Z" />
            </svg>
            贝壳校园墙网页版
          </router-link>
        </a>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav"
          aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarNav">
          <ul class="navbar-nav colorblack">
            <li class="nav-item">
              <a class="nav-link" href="#">
                <router-link to="/home">首页</router-link>
              </a>
            </li>
            <li class="nav-item dropdown">
              <a class="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-bs-toggle="dropdown"
                aria-expanded="false">
                关于
              </a>
              <ul class="dropdown-menu" aria-labelledby="navbarDropdown">
                <li>
                  <a target="_blank" href="https://blog.csdn.net/V123456789987654" class="dropdown-item">作者的博客</a>
                </li>
                <li>
                  <a target="_blank" href="https://github.com/U202142209" class="dropdown-item">GitHub</a>
                </li>
                <a target="_blank" href="http://blog.wobushidalao.top/" class="dropdown-item">我的个人网站</a>
              </ul>
            </li>
            <!-- <li class="nav-item">
              <a class="nav-link" href="#">
                <router-link to="/login">登录/注册</router-link>
              </a>
            </li> -->
          </ul>
        </div>
      </div>
    </nav>
  </div>
</template>