<template>
  <n-card :bordered="false">
    <template #header>
      <h1>欢迎来到安东星！</h1>
    </template>
    <n-space vertical style="text-align: left">
      <n-alert title="注册须知" type="warning">请阅读以下说明</n-alert>
      <n-grid cols="1 768:3 1200:3 1920:3">
        <n-gi/>
        <n-gi>
          <n-card size="small" bordered embedded>
            <template #header><h3>注册</h3></template>
            <n-form
                :model="formValue"
                :rules="rules"
                :size="size"
                ref="formRef"
            >
              <n-form-item label="登录账号" required path="username">
                <n-input v-model:value="formValue.username" placeholder="输入登录名"/>
              </n-form-item>

              <n-form-item label="密码" required path="password">
                <n-input placeholder="输入密码" type="password"
                         show-password-on="click" v-model:value="formValue.password"/>
              </n-form-item>
              <n-form-item label="重复密码" required path="password">
                <n-input placeholder="请再次输入密码" type="password"
                         show-password-on="click" v-model:value="formValue.secPassword"/>
              </n-form-item>
              <n-form-item label="邮箱" path="email">
                <n-input v-model:value="formValue.email" placeholder="输入邮箱"/>
              </n-form-item>
              <n-form-item label="邀请码" path="inviteCode">
                <n-input v-model:value="formValue.inviteCode" placeholder="输入邀请码"/>
              </n-form-item>
              <n-form-item label="验证码" required path="captcha">
                <n-space vertical>
                  <n-input v-model:value="formValue.captcha" placeholder="输入验证码"/>
                  <n-image width="240" height="80" preview-disabled
                           @click="refreshCaptcha"
                           class="img-captcha"
                           :src="prefix+captchaFile+'?'+randomStr"></n-image>
                </n-space>
              </n-form-item>
              <n-form-item label="用户协议" required path="captcha">
                <n-checkbox v-model:checked="formValue.agreeLicense">同意用户协议</n-checkbox>
                <n-button text type="info">用户协议</n-button>
              </n-form-item>
              <n-form-item style="display: flex; justify-content: flex-end;">
                <n-space>
                  <n-button type="success" text @click="router.push({'name':'login'})">已有账号？点我登录</n-button>
                  <n-button @click="handleValidateClick" type="primary">注册</n-button>
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

onMounted(() => {
  generateCaptcha()
})

const formValue = ref({
  username: '',
  email: '',
  password: '',
  secPassword: '',
  inviteCode: '',
  captcha_id: '',
  captcha_val: '',
  agreeLicense: false
})

const router = useRouter()
const rules = ref()
const size = ref()
const captchaFile = ref()
const randomStr = ref(0)

const prefix = import.meta.env.VITE_APP_REQUEST_BASE_URL + 'v1/captcha/'
const generateCaptcha = () => {
  http.get('/v1/captcha').then(res => {
    captchaFile.value = res.data.id + "." + res.data.ext
    formValue.value.captcha_id = res.data.id
  })
}

const refreshCaptcha = () => {
  http.get('/v1/captcha/' + captchaFile.value, {
    params: {
      reload: true
    }
  }).then(res => {
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
