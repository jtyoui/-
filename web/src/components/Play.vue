<template>
  <Header/>
  <el-row style="margin-top: 30px">
    <div class="video">
      <vue3-video-play v-bind="options"/>
      <div class="top">
        <el-icon :size="20">
          <share/>
        </el-icon>
        <span class="tittle">&emsp;我和僵尸有个约会1——马小玲爱上我</span>
        <span class="start">
        <el-icon :size="24"><star/></el-icon>
        <span> 20</span>
      </span>
      </div>
    </div>
    <div class="box">
      <el-menu :default-openeds="[0]">
        <el-sub-menu :index="0">
          <template #title>
            <el-icon>
              <Menu/>
            </el-icon>
            <span>我和僵尸有个约会一</span>
          </template>
          <div class="list" style="overflow: auto">
            <div v-for="i in 20" :key="i" class="infinite-list-item">
              <el-menu-item :index="i">
                <el-tooltip :content="'P'+i+'&emsp;我和僵尸有我和僵尸有个约会个约会'+i" placement="top" show-after="800">
                  <div class="name">
                    P{{ i }}&emsp;我和僵尸有我和僵尸有个约会个约会{{ i }}
                  </div>
                </el-tooltip>
                <div class="time">50:00</div>
              </el-menu-item>
            </div>
          </div>
        </el-sub-menu>
      </el-menu>
    </div>
  </el-row>
</template>

<script setup>
import Header from "./Header.vue"
import {reactive} from "vue";
import {Menu} from '@element-plus/icons'

const options = reactive({
  width: "984px",
  height: "700px",
  title: "我和僵尸有个约会开头",
  src: "https://jtyoui.oss-accelerate-overseas.aliyuncs.com/video/login.mp4",
  muted: false,
  webFullScreen: false,
  autoPlay: false,
  loop: false,
  control: true,
  volume: 0.3
})
</script>

<style scoped>
.video {
  margin-left: 20%;
}

.box {
  width: 300px;
  border: #dcd9d9 solid 2px;
  margin: 0 auto;
}

.tittle {
  font-size: 20px;
  color: red;
}

.start {
  float: right;
  margin-right: 100px;
  font-size: 25px;
}

.top {
  margin-top: 20px;
}

.time {
  width: 20px;
  margin-left: 30px;
}

.list {
  height: 700px;
  list-style: none;
  width: 100%;
}

.name {
  width: 200px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>