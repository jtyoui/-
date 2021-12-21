<template>
  <Header/>
  <div class="box">
    <h1>精彩片段从这里开始 </h1>
    <hr/>

    <div class="content">
      &emsp;虽然和你相识的日子不长，不过发生的事情实在是太多了。在这段时间里，有一件事我可以肯定，原来时间是毫无意义的，
      就算让你长生不死，又能怎样？寥寥寂寞一生，这比死更令人难受，反而短暂的相遇，更加令人死而无憾。
      让我们一起忘记这个世界和时间的存在，忘记生离死别，甚至是生命的长短。我们要做的就是珍惜现在，
      就算下一秒就是世界末日，只要这一秒没有到，这就不是事实，事实只有一个，就是况天佑不能没有马小玲......
    </div>
  </div>
  <div>
    <div class="video">
      <vue3-video-play v-bind="options"/>
    </div>
    <div class="arrow">
      <el-icon :size="100">
        <ArrowRightBold @click="video"/>
      </el-icon>
    </div>
  </div>

</template>

<script setup>
import Header from "./Header.vue"

import {reactive} from "vue";
import {ArrowRightBold} from '@element-plus/icons'
import {useRouter} from 'vue-router'

const router = useRouter()
const options = reactive({
  width: "100%",
  height: "600px",
  title: "我和僵尸有个约会开头",
  src: "http://localhost:3030/hls/制作特辑.mp4/master.m3u8",
  type: "m3u8",
  muted: true,
  webFullScreen: false,
  autoPlay: true,
  loop: false,
  control: true,
  volume: 0.3
})

function video() {
  router.push({path: "/video"})
}
</script>

<style scoped>
.box {
  width: 50%;
  text-align: center;
  margin: 10px auto;
}

.content {
  text-align: left;
  margin-top: 20px;
  color: #cccccc;
}

.video {
  width: 60%;
  float: left;
  margin-left: 20%;
}

.arrow {
  margin-top: 170px;
  width: 10%;
  float: right;
  cursor: pointer;
}

</style>