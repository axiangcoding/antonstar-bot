<template>
  <n-card :bordered="false" id="body-register">
    <template #header>
      <h1>欢迎来到安东星！</h1>
    </template>
    <n-space vertical style="text-align: left">
      <!--<n-alert title="注册须知" type="warning">请阅读以下说明</n-alert>-->
      <n-card size="small" bordered embedded class="card-register">
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
          <n-image v-if="captchaFile!==undefined" width="240" height="80" preview-disabled
                   @click="refreshCaptcha"
                   class="img-captcha"
                   :src="prefix+captchaFile+'?'+randomStr"
          ></n-image>
          <n-form-item label="用户协议" required path="agreeLicense">
            <n-checkbox v-model:checked="formValue.agreeLicense">同意行为准则</n-checkbox>
            <n-button text type="info" @click="showUA=true">《安东星行为准则》</n-button>
          </n-form-item>
          <n-form-item style="display: flex; justify-content: flex-end;">
            <n-space>
              <n-button type="success" text @click="router.push({'name':'login'})">已有账号？点我登录</n-button>
              <n-button @click="handlerClick" type="primary">注 册</n-button>
            </n-space>
          </n-form-item>
        </n-form>
      </n-card>
    </n-space>
    <UserAgreementCard :show="showUA" @update-show="updateShowUA"/>
    <RegSuccessCard :show="showRegSuccess"/>
  </n-card>
</template>

<script lang="ts" setup>
import {onMounted, ref} from "vue";
import {onBeforeRouteLeave, useRouter} from "vue-router";
import {getRegex} from "@/util/validation";
import {captcha, CaptchaForm} from "@/services/captcha";
import {userRegister, RegForm, userValueExist} from "@/services/user";
import {useMessage} from "naive-ui";
import RegSuccessCard from "@/views/user/components/RegSuccessCard.vue";
import UserAgreementCard from "@/views/user/components/UserAgreementCard.vue";

onMounted(() => {
  generateCaptcha()
})

const formRef = ref()
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

const showUA = ref(false)
const updateShowUA = (show: boolean) => {
  showUA.value = show
}

const message = useMessage()
const showRegSuccess = ref(false)
const handlerClick = (e: Event) => {
  e.preventDefault()
  formRef.value.validate((errors: any) => {
    if (!errors) {
      let user: RegForm = {
        email: formValue.value.email,
        invitedCode: formValue.value.inviteCode,
        password: formValue.value.password,
        username: formValue.value.username
      }
      let cap: CaptchaForm = {
        captchaId: captchaId.value,
        captchaVal: formValue.value.captchaVal

      }
      userRegister(user, cap)
          .then((res: any) => {
            if (res.code === 0) {
              showRegSuccess.value = true
              countdown()
            } else if (res.code === 11004) {
              message.warning('验证码不正确，请重新输入！')
            } else {
              message.warning('注册失败！')
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

const countdown = () => {
  setTimeout(() => {
    showRegSuccess.value = false
    router.push({name: 'login'})
  }, 3000)
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
      pattern: getRegex('email'),
      message: '请输入正确的邮箱',
      trigger: 'blur'
    },
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
    type: 'enum',
    enum: [true],
    message: '请阅读并同意行为准则',
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

onBeforeRouteLeave((to, from) => {
  const answer = window.confirm(
      '你确定要离开这个页面吗？未保存的更改将会被丢弃'
  )
  // 取消导航并停留在同一页面上
  if (!answer) return false
})
</script>

<style lang="scss" scoped>
.img-captcha {
  &:hover {
    cursor: pointer;
  }
}

//#body-register {
//  background: url("@/assets/image/bg-reg.jpg") no-repeat;
//  background-size: 100% 100%;
//  min-height: calc(100vh - var(--footer-height) - var(--header-height));
//}

.card-register {
  max-width: 550px;
  margin: 0 auto;
}
</style>
