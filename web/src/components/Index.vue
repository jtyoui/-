<template>
  <Header/>
  <div class="box">
    <div class="title">
      <h1>精彩片段从这里开始 </h1>
    </div>
    <hr/>
    <div class="content">
      &emsp;&emsp;关于一个很久很久的都市传说，相传在公元2010年的某个下午，有一个小男孩去他父母打工的地方，偶然间在暄闹的菜市旁边的小黑屋中，
      发现了一碟DVD，上面满满写着一行大字，假装横做头依靠着一点小学文化隐隐约约听见八个大字《我和僵尸有个约会》，从此故事由此开头......
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
  src: "https://jtyoui.oss-accelerate-overseas.aliyuncs.com/video/login.mp4",
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
  margin: 30px auto 30px auto;
}

.content {
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