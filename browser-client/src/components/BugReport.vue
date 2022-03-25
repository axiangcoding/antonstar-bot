<template>
  <n-card
      style="width: 600px"
      title="问题反馈"
      :bordered="false"
      size="huge"
      role="dialog"
      aria-modal="true"
  >
    <template #header-extra>
      <n-text italic>一起让安东星变得更好！</n-text>
    </template>
    <n-form
        :model="formValue"
        :rules="rules"
        :size="size"
        ref="formRef"
    >
      <n-form-item label="问题类型" required path="bugType">
        <n-select v-model:value="formValue.bugType"
                  placeholder="请选择你要报告的问题类型"
                  :options="bugOptions"/>

      </n-form-item>

      <n-form-item label="简要标题" required path="title">
        <n-input placeholder="简单描述下你要报告的问题" maxlength="30" show-count
                 v-model:value="formValue.title"/>
      </n-form-item>
      <n-form-item label="详细描述" path="content">
        <n-input v-model:value="formValue.content" :autosize="{
                    minRows: 4,
                    maxRows: 8
                  }" maxlength="2000"
                 placeholder="如果可以，请详细描述下这个问题是如何出现的"
                 type="textarea"/>
      </n-form-item>

      <n-form-item style="display: flex; justify-content: flex-end;">
        <n-space>
          <n-checkbox v-model:checked="formValue.anonymous" :disabled="!store.state.login" text type="success">
            匿名反馈
          </n-checkbox>
          <n-button @click="handlerClick" type="primary">反馈问题</n-button>
        </n-space>
      </n-form-item>
    </n-form>
    <!--<template #footer>-->
    <!--  <n-space justify="end">-->
    <!--    <n-button></n-button>-->
    <!--  </n-space>-->
    <!--</template>-->
  </n-card>
</template>

<script lang="ts" setup>


import {ref} from "vue";
import {useStore} from "vuex";

const message = useMessage();
import {getRegex} from "@/util/validation";
import {userInfo, userLogin} from "@/services/user";
import {BugReportForm, postBugReport} from "@/services/bug-report";
import router from "@/router";
import {useMessage} from "naive-ui";

const store = useStore()
const size = ref()
const formRef = ref()
const formValue = ref({
  bugType: '',
  title: '',
  content: '',
  anonymous: !store.state.login,
})

const rules = {
  bugType: [
    {
      required: true,
      message: '请选择BUG类型',
      trigger: 'blur'
    },
  ],
  title: [
    {
      required: true,
      message: '请输入简要的描述，30字以内',
      max: 30,
      trigger: 'blur'
    },
  ],
}

const bugOptions = [
  {
    label: '功能异常',
    value: '1',
  },
  {
    label: '界面或展示',
    value: '2',
  },
  {
    label: '账号及战绩',
    value: '3',
  },
  {
    label: '意见或建议',
    value: '4'
  },
  {
    label: '其他',
    value: '5'
  }
]

const handlerClick = (e: Event) => {
  e.preventDefault()
  formRef.value.validate((errors: any) => {
    if (!errors) {
          let reqForm: BugReportForm = {
            type: formValue.value.bugType,
            title: formValue.value.title,
            content: formValue.value.content,
            anonymous: formValue.value.anonymous
          }
          postBugReport(store.state.auth, reqForm).then(res => {
            message.success('您的反馈已记录，请耐心等待处理哦')
          })
        }
      }
  )
}
</script>

<style scoped>

</style>
