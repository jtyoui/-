<template>
  <Header/>
  <el-container>
    <el-main>
      <div class="video">
        <vue3-video-play v-bind="options"/>
        <div class="top">
          <el-icon :size="20">
            <share/>
          </el-icon>
          <span class="tittle">&emsp;{{ options.title }}</span>
          <el-button icon="star" class="start" type="success">20</el-button>
        </div>
      </div>
    </el-main>

    <el-aside width="30%">
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
              <div v-for="(video,i) in getList()" :key="i" class="infinite-list-item">
                <el-menu-item :index="i" @click="play(video.id)">
                  <el-tooltip :content="video.name" placement="top" show-after="800">
                    <div class="name">
                      P{{ i + 1 }}&emsp;{{ video.name }}
                    </div>
                  </el-tooltip>
                  <div class="time">{{ video.time }}</div>
                </el-menu-item>
              </div>
            </div>
          </el-sub-menu>
        </el-menu>
      </div>
    </el-aside>
  </el-container>
</template>

<script setup>
import Header from "./Header.vue"
import {reactive} from "vue";
import {Menu} from '@element-plus/icons'

const options = reactive({
  width: "984px",
  height: "700px",
  title: "我和僵尸有个约会开头",
  src: "http://localhost:3030/hls/我和僵尸有个约会3/01.mp4/master.m3u8",
  type: "m3u8",
  muted: false,
  webFullScreen: false,
  autoPlay: false,
  loop: false,
  control: true,
  volume: 0.3
})

function getList() {
  const list = []
  for (let i = 0; i < 10; i++) {
    list.push({
      id: "1234",
      name: "我和僵尸有个约会第一部#第" + i + "集",
      time: "56:00"
    })
  }
  return list
}

function play(id) {
  options.title = ""
  options.src = ""
}

</script>

<style scoped>
.video {
  margin: 0 15%;
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