<!--
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-10-12 18:00:45
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-10-12 18:17:26
 * @FilePath: \ustb-campus\frontend\src\components\SwiperComponent.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
  <div class="carousel-container">
    <van-swipe :autoplay="2000">
      <van-swipe-item v-for="(item, index) in images" :key="index">
        <img :src="item" alt="轮播图" class="carousel-image" />
      </van-swipe-item>
    </van-swipe>
  </div>


</template>

<script>
import { Swipe, SwipeItem } from 'vant';

export default {
  components: {
    [Swipe.name]: Swipe,
    [SwipeItem.name]: SwipeItem,
  },
  data() {
    return {
      images: [
        'https://www.ustb.edu.cn/images/content/2023-10/20231007103238584952.jpg',
        'https://www.ustb.edu.cn/images/content/2023-10/20231007110605606416.jpg',
        'https://www.ustb.edu.cn/images/content/2023-10/20231007103106249844.jpg',
        'https://news.ustb.edu.cn/__local/B/F7/CD/A317889687488C4515C5A8D44A1_D8C3ACF0_3D540A.jpg'
      ],
    };
  },
};
</script>


<style>
.carousel-image {
  width: 100%;
  height: auto;
}
.carousel-container {
  margin: 0 2%;
}
</style>