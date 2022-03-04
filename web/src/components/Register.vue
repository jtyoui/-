<template>
  <div>
    <video autoplay loop class="video-bg" muted>
      <source src="https://jtyoui.oss-accelerate-overseas.aliyuncs.com/video/register.mp4" type="video/mp4">
    </video>

    <div class="layout">
      <div class="text color center">注 册</div>

      <div class="box">
        <el-container>
          <el-header class="color name">张伟 著</el-header>

          <el-container>
            <el-aside width="300px">

              <el-header class="color title">
                My date with a vampire.
              </el-header>

              <el-main class="color">
                <div class="center to-star">主演</div>
                <div class="move">
                  尹天照、万绮雯、杨恭如、陈启泰、
                  吴廷烨、张文慈、杜汶泽、张国权、
                  黄树棠、吕有慧、杨嘉诺、姜皓文、
                  冼灏英、程　东、炜　烈
                </div>
              </el-main>
            </el-aside>

            <el-main>
              <div class="right">
                <el-tooltip content="我们保证不会记录你的任何信息" placement="top">
                  <el-button size="medium" type="text" class="text">🌍 欢 迎 注 册</el-button>
                </el-tooltip>
              </div>

              <div class="input">
                <div class="distance">
                  <el-form :rules="ruleForm.rules" :model="ruleForm">
                    <el-form-item prop="name">
                      <el-input clearable v-model="ruleForm.name" placeholder="昵称" :prefix-icon="User" autofocus/>
                    </el-form-item>
                  </el-form>

                </div>

                <div class="distance">
                  <el-select v-model="names.gender" placeholder="性别" :suffix-icon="Male" class="lengthen">
                    <el-option v-for="item in options" :key="item" :label="item.label" :value="item.value"/>
                  </el-select>
                </div>

                <div class="distance">
                  <el-date-picker v-model="names.birth_date" type="date" placeholder="年龄" style="width: 100%;"/>
                </div>

                <div class="distance center">
                  <el-button type="primary" :icon="Stamp" :loading="loading" auto-insert-space @click="register">
                    点 击 注 册
                  </el-button>
                </div>
              </div>
            </el-main>
          </el-container>
        </el-container>
      </div>

    </div>

  </div>
</template>

<script setup>
import {Male, Stamp, User} from '@element-plus/icons'
import {reactive, ref} from "vue";
import axios from "axios";
import {useRouter} from "vue-router";

const names = reactive({
  "name": "",
  "gender": "",
  "birth_date": ""
})
const router = useRouter()

const loading = ref(false)

const options = ref([
  {
    value: '男',
    label: '男'
  },
  {
    value: '女',
    label: '女'
  }
])

const ruleForm = reactive({
  name: "",
  rules: {
    name: [{required: true, max: 10, min: 4, message: "昵称的长度为4到10位", trigger: 'blur'}]
  }
})

function register() {
  names.name = ruleForm.name
  axios.post("/api/user/register", names).then(resp => {
    if (resp.data.code !== 200) {
      alert(resp.data.msg)
    } else {
      alert("请记住你的唯一僵尸码：" + resp.data.data)
      router.push({path: "/login"})
    }
  })
}
</script>

<style scoped>
.layout {
  width: 50%;
  height: 500px;
  margin: 10% auto;
  background-color: rgba(27, 27, 65, 0.9);
}

.color {
  color: aliceblue;
}

.layout .text {
  font-size: 30px;
}

.box {
  width: 80%;
  margin: 0 auto;
}

.name {
  text-align: right;
}

.title {
  font-size: 23px;
  font-weight: bolder;
}

.to-star {
  font-size: 20px;
}

.right {
  text-align: center;
}

.right .text {
  font-size: 20px;
}

.distance {
  margin-top: 20px;
  width: 100%;
}

.input {
  margin: 0 auto;
  width: 70%;
}

.move {
  margin-top: 10%;
}
</style>