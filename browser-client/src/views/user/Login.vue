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
          <n-card size="small" bordered embedded style="min-width: 300px">
            <template #header><h3>登录</h3></template>
            <n-form
                :model="formValue"
                :rules="rules"
                :size="size"
                ref="formRef"
            >
              <n-form-item label="账号" path="username">
                <n-input v-model:value="formValue.username" placeholder="输入登录名或者邮箱"/>
              </n-form-item>
              <n-form-item label="密码" path="password">
                <n-input placeholder="输入密码" type="password"
                         show-password-on="click" v-model:value="formValue.password"/>
              </n-form-item>
              <n-form-item label="验证码" path="captcha">
                <n-space vertical>
                  <n-input v-model:value="formValue.captcha" placeholder="输入验证码"/>
                  <n-image width="240" height="80" preview-disabled src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"></n-image>
                </n-space>

              </n-form-item>
              <n-form-item style="display: flex; justify-content: flex-end;">
                <n-space>
                  <n-button type="success">没有账号？注册一个</n-button>
                  <n-button @click="handleValidateClick"  type="primary">登录</n-button>
                </n-space>
              </n-form-item>
            </n-form>
          </n-card>

        </n-gi>
        <n-gi/>
      </n-grid>

    </n-space>
  </n-card>
</template>

<script lang="ts" setup>
import {ref} from "vue";

const formValue = ref({
  username: '',
  password: '',
  captcha:''
})
const rules = ref()
const size = ref()
</script>

<style scoped>

</style>
