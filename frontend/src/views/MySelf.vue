<!--
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-10-11 22:16:16
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-10-26 13:40:35
 * @FilePath: \frontend\src\views\MySelf.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
  <div class="user-profile">
    <div class="info">
      <van-image round src="https://img.yzcdn.cn/vant/cat.jpeg" />
      <h3 class="name">
        用户名
        <br />
        <van-tag size="mini">申请认证</van-tag>
      </h3>
    </div>
    <van-row>
      <van-col span="8">
        <p>0</p>
        <p>动态</p>
      </van-col>
      <van-col span="8">
        <p>0</p>
        <p>关注</p>
      </van-col>
      <van-col span="8">
        <p>0</p>
        <p>粉丝</p>
      </van-col>
    </van-row>
  </div>
  <van-row class="user-links">
    <van-col span="8">
      <van-icon name="newspaper-o" color="#7af" />我的作品
    </van-col>
    <van-col span="8">
      <van-icon name="star-o" color="#f00" />我的收藏
    </van-col>
    <van-col span="8">
      <van-icon name="tosend" color="#fa0" />阅读历史
    </van-col>
  </van-row>

  <van-cell-group class="user-group">
    <van-cell icon="edit" title="编辑资料"  is-link />
    <van-cell icon="chat-o" title="小智同学"  is-link />
    <van-cell icon="setting-o" title="系统设置" is-link />
    <van-cell icon="warning-o" title="退出登录" to="/404" is-link />
  </van-cell-group>
</template>
<script>
export default {

}

</script>

<style>
.user {
  &-profile {
    width: 100%;
    height: 150px;
    display: block;
    background: #3296fa;
    color: #fff;

    .info {
      display: flex;
      padding: 20px;
      align-items: center;

      .van-image {
        width: 64px;
        height: 64px;
      }

      .name {
        font-size: 16px;
        font-weight: normal;
        margin-left: 10px;
      }

      .van-tag {
        background: #fff;
        color: #3296fa;
      }
    }

    p {
      margin: 0;
      text-align: center;
    }
  }

  &-group {
    margin-bottom: 15px;
  }

  &-links {
    padding: 15px 0;
    font-size: 12px;
    text-align: center;
    background-color: #fff;

    .van-icon {
      display: block;
      font-size: 24px;
      padding-bottom: 5px;
    }
  }
}
</style>