<template>
  <n-card :bordered="false" id="body-login">
    <template #header>
      <h1>欢迎来到安东星！</h1>
    </template>
    <n-space vertical style="text-align: left">
      <n-card size="small" bordered embedded class="card-login">
        <template #header><h3>登录</h3></template>
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
          <n-form-item label="验证码" required path="captchaVal">
            <n-input v-model:value="formValue.captchaVal" placeholder="输入验证码"/>
          </n-form-item>
          <n-image v-if="captchaFile!==undefined" width="240" height="80" preview-disabled
                   @click="refreshCaptcha"
                   class="img-captcha"
                   :src="prefix+captchaFile+'?'+randomStr"
          ></n-image>
          <n-form-item style="display: flex; justify-content: flex-end;">
            <n-space>
              <n-button type="success" text @click="router.push({'name':'register'})">没有账号？点我注册一个</n-button>
              <n-button @click="handlerClick" type="primary">登 录</n-button>
            </n-space>
          </n-form-item>
        </n-form>
      </n-card>
    </n-space>
  </n-card>
</template>

<script lang="ts" setup>
import {onMounted, ref} from "vue";
import {useRouter} from "vue-router";
import {captcha, CaptchaForm} from "@/services/captcha";
import {LoginForm, RegForm, userLogin} from "@/services/user";
import {useMessage} from "naive-ui";
import {getRegex} from "@/util/validation";
import {useStore} from "vuex";

onMounted(() => {
  generateCaptcha()
})

const formRef = ref()
const formValue = ref({
  username: '',
  password: '',
  captchaVal: '',
})
const captchaId = ref()

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
  ],
  password: [
    {
      required: true,
      message: '请输入密码',
      trigger: 'blur'
    },
  ],
  captchaVal: {
    required: true,
    message: '请输入验证码',
    trigger: 'blur'
  }
}

const store = useStore();
const message = useMessage()
const handlerClick = (e: Event) => {
  e.preventDefault()
  formRef.value.validate((errors: any) => {
    if (!errors) {
      let user: LoginForm = {
        password: formValue.value.password,
        username: formValue.value.username
      }
      let cap: CaptchaForm = {
        captchaId: captchaId.value,
        captchaVal: formValue.value.captchaVal

      }
      userLogin(store.state.auth, user, cap)
          .then((res: any) => {
            if (res.code === 0) {
              message.success(`欢迎回来， ${formValue.value.username}！`)
              store.commit('setAuth', res.data.Authorization)
              store.commit('setLogin', true)
              router.back()
            } else if (res.code === 11004) {
              message.warning('验证码不正确，请重新输入！')
            } else {
              message.error('登录失败！用户名或密码错误！')
            }
            generateCaptcha()
          })
          .catch(err => {
            generateCaptcha()
          })
    } else {

    }
  })
}

const router = useRouter()

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

//#body-login {
//  background: url("@/assets/image/bg-login.jpg") no-repeat;
//  background-size: 100% 100%;
//  min-height: calc(100vh - var(--footer-height) - var(--header-height));
//}

.card-login {

  max-width: 550px;
  margin: 0 auto;
}
</style>
