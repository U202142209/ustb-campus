<template>
  <div class="sdjflasjlk">
    <div class="title">
      <div class="tac titlecontainer">欢迎使用贝壳校园墙网页版</div>
    </div>
    <div id="apphomt">
      <canvas ref="canvas"></canvas>
    </div>
    <div id="footer">
      <div id="footcontent">
        powered by <a target="_blank" id="alink" href="https://space.bilibili.com/661047235">@我不是大佬</a>
      </div>
    </div>
  </div>
</template>

<script>
import * as THREE from 'three';

export default {
  mounted() {
    this.init();
    this.animate();
  },
  methods: {
    init() {
      // 创建场景
      this.scene = new THREE.Scene();

      // 创建相机
      this.camera = new THREE.PerspectiveCamera(
        60, window.innerWidth / window.innerHeight,
        0.5, 1000
      );
      this.camera.position.z = 5;

      // 创建渲染器
      this.renderer = new THREE.WebGLRenderer({ canvas: this.$refs.canvas });
      this.renderer.setSize(window.innerWidth, window.innerHeight);

      // 创建球体
      const geometry = new THREE.SphereGeometry(1, 62, 62);
      const material = new THREE.MeshBasicMaterial({ color: 0x11a213 });
      this.sphere = new THREE.Mesh(geometry, material);
      this.scene.add(this.sphere);

      // 创建网络状球体
      const networkGeometry = new THREE.SphereGeometry(1.5, 64, 64);
      const networkMaterial = new THREE.MeshBasicMaterial({
        color: 0x5CACEE, wireframe: true,
      });
      this.networkSphere = new THREE.Mesh(networkGeometry, networkMaterial);
      this.scene.add(this.networkSphere);
    },
    animate() {
      requestAnimationFrame(this.animate);
      this.sphere.rotation.x += 0.01;
      this.sphere.rotation.y += 0.01;
      this.networkSphere.rotation.x -= 0.005;
      this.networkSphere.rotation.y -= 0.005;
      this.renderer.render(this.scene, this.camera);
    },
  },
};
</script>

<style>
#apphomt {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: rgb(30, 9, 9);
}

.sdjflasjlk {
  position: relative;
}

.title,
#footer {
  color: white;
  background-color: black;
  font-style: 楷体;
  position: absolute;
  left: 0;
  width: 100%;
}

.title {
  font-size: larger;
  top: 10%;
}

#footer {
  bottom: 10%;
}

.top-nav,#alink {
  color: white;
  background-color: black;
}
</style>