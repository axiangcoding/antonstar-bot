<template>
  <n-card :bordered="false">
    <template #header>
      <h1>欢迎来到安东星！</h1>
    </template>
    <n-space vertical style="text-align: left">
      <!--<n-alert title="本站须知" type="warning">请阅读以下说明</n-alert>-->
      <n-grid cols="1 768:3 1200:3 1920:3">
        <n-gi/>
        <n-gi>
          <n-card size="small" bordered embedded>
            <template #header><h3>登录</h3></template>
            <n-form
                :model="formValue"
                :rules="rules"
                :size="size"
                ref="formRef"
            >
              <n-form-item label="账号" path="username">
                <n-input v-model:value="formValue.username" placeholder="输入用户名"/>
              </n-form-item>
              <n-form-item label="密码" path="password">
                <n-input placeholder="输入密码" type="password"
                         show-password-on="click" v-model:value="formValue.password"/>
              </n-form-item>
              <n-form-item label="验证码" required path="captcha_val">
                <n-input v-model:value="formValue.captchaVal" placeholder="输入验证码"/>
              </n-form-item>
              <n-image width="240" height="80" preview-disabled
                       @click="refreshCaptcha"
                       class="img-captcha"
                       :src="prefix+captchaFile+'?'+randomStr"></n-image>
              <n-form-item style="display: flex; justify-content: flex-end;">
                <n-space>
                  <n-button type="success" text @click="router.push({'name':'register'})">没有账号？注册一个</n-button>
                  <n-button @click="" type="primary">登录</n-button>
                </n-space>
              </n-form-item>
            </n-form>
            <!--{{ formValue }}-->
          </n-card>

        </n-gi>
        <n-gi/>
      </n-grid>

    </n-space>
  </n-card>
</template>

<script lang="ts" setup>
import {onMounted, ref} from "vue";
import http from "@/services/request";
import {useRouter} from "vue-router";
import {getRegex} from "@/util/validation";
import {captcha} from "@/services/user";

onMounted(() => {
  generateCaptcha()
})

const formValue = ref({
  username: '',
  password: '',
  captchaVal: '',
})

const captchaId = ref('')

const router = useRouter()
const rules = {
  username: [
    {
      required: true,
      message: '请输入用户名',
      trigger: 'blur'
    },
    {
      pattern: getRegex('username'),
      message: '字母，数字，下划线组成，长度在5-16位之间',
      trigger: 'blur'
    }
  ],
  password: [
    {
      required: true,
      message: '请输入密码',
      trigger: 'blur'
    },
    {
      pattern: getRegex('password'),
      message: '密码长度应该在8-16位之间',
      trigger: 'blur'
    }
  ],
  captchaVal: {
    required: true,
    message: '请输入验证码',
    trigger: 'blur'
  }
}
const size = ref()
const captchaFile = ref()
const randomStr = ref(0)

const prefix = import.meta.env.VITE_APP_REQUEST_BASE_URL + 'v1/captcha/'
const generateCaptcha = () => {
  captcha('', false).then(res => {
    captchaFile.value = res.data.id + "." + res.data.ext
    captchaId.value = res.data.id
  })
}

const refreshCaptcha = () => {
  captcha(captchaFile.value, true).then(res => {
    randomStr.value = new Date().getTime()
  })
}
</script>

<style lang="scss" scoped>
.img-captcha {
  &:hover {
    cursor: pointer;
  }
}
</style>
