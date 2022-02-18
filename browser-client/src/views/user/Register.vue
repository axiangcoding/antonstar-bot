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
              <n-form-item label="重复密码" required path="secPassword">
                <n-input placeholder="请再次输入密码" type="password"
                         show-password-on="click" v-model:value="formValue.secPassword"/>
              </n-form-item>
              <n-form-item label="邮箱" path="email">
                <n-input v-model:value="formValue.email" placeholder="输入邮箱"/>
              </n-form-item>
              <n-form-item label="邀请码" path="inviteCode">
                <n-input v-model:value="formValue.inviteCode" placeholder="输入邀请码"/>
              </n-form-item>
              <n-form-item label="验证码" required path="captchaVal">
                <n-input v-model:value="formValue.captchaVal" placeholder="输入验证码"/>
              </n-form-item>
              <n-image width="240" height="80" preview-disabled
                       @click="refreshCaptcha"
                       class="img-captcha"
                       :src="prefix+captchaFile+'?'+randomStr"></n-image>
              <n-form-item label="用户协议" required path="agreeLicense">
                <n-checkbox v-model:checked="formValue.agreeLicense">同意行为准则</n-checkbox>
                <n-button text type="info">《安东星行为准则》</n-button>
              </n-form-item>
              <n-form-item style="display: flex; justify-content: flex-end;">
                <n-space>
                  <n-button type="success" text @click="router.push({'name':'login'})">已有账号？点我登录</n-button>
                  <n-button @click="" type="primary">注册</n-button>
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
import {useRouter} from "vue-router";
import {getRegex} from "@/util/validation";
import {captcha, userValueExist} from "@/services/user";

onMounted(() => {
  generateCaptcha()
})

const formValue = ref({
  username: '',
  email: '',
  password: '',
  secPassword: '',
  inviteCode: '',
  captchaVal: '',
  agreeLicense: false
})

const captchaId = ref()

// 校验二次密码是否相同
const validatePasswordSame = (rule: any, value: string) => {
  return value === formValue.value.password
}

const validatePasswordStartWith = (rule: any, value: string) => {
  return (
      formValue.value.password &&
      formValue.value.password.startsWith(value) &&
      formValue.value.password.length >= value.length
  )
}

const validateEmailExist = async (rule: any, value: string, callback: any) => {
  if (value === '') {
    return
  }
  await userValueExist('email', value).then(res => {
    if (res.data.exists) {
      callback(new Error('邮箱已存在！'))
    }
  }).catch(err => {
    callback(new Error('重复邮箱检测失败，请稍后重试'))
  })
  callback()
}

const validateUsernameExist = async (rule: any, value: string, callback: any) => {
  if (value === '') {
    return
  }
  await userValueExist('username', value).then(res => {
    if (res.data.exists) {
      callback(new Error('用户名已存在！'))
    }
  })
  callback()
}


const router = useRouter()
const rules = {
  username: [
    {
      required: true,
      message: '请输入登录用户名',
      trigger: 'blur'
    },
    {
      pattern: getRegex('username'),
      message: '字母，数字，下划线组成，长度在5-16位之间',
      trigger: 'blur'
    },
    {
      message: '用户名已经存在，请更换！',
      validator: validateUsernameExist,
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
  secPassword: [
    {
      required: true,
      message: '请再次输入密码',
      trigger: ['input', 'blur']
    },
    {
      validator: validatePasswordStartWith,
      message: '两次密码输入不一致',
      trigger: 'input'
    },
    {
      validator: validatePasswordSame,
      message: '两次密码输入不一致',
      trigger: ['blur', 'password-input']
    }
  ],
  email: [
    {
      validator: validateEmailExist,
      trigger: 'blur'
    }
  ],
  captchaVal: {
    required: true,
    message: '请输入验证码',
    trigger: 'blur'
  },
  agreeLicense: {
    required: true,
    message: '请阅读并同意安东星行为准则',
    trigger: 'blur'
  }
}
const size = ref()
const randomStr = ref(0)

const captchaFile = ref()
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
