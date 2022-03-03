<template>
  <n-card size="small" :bordered="true" style="text-align: left">
    <n-space vertical>
      <n-h2 prefix="bar" align-text type="primary">修改全站公告</n-h2>
      <n-input-group>
        <n-input-group-label>公告标题</n-input-group-label>
        <n-input :style="{ width: '50%' }" placeholder="请输入标题" maxlength="30" show-count v-model:value="title">标题
        </n-input>
      </n-input-group>
      <div id="vditor" style="text-align: left"></div>
      <n-space justify="end">
        <n-button type="error">重置</n-button>
        <n-button type="success" @click="handleSubmit">发布</n-button>
      </n-space>
    </n-space>

  </n-card>
</template>

<script lang="ts" setup>
// TODO 编辑全站公告
import {onMounted, ref} from "vue";
import Vditor from "vditor";
import {toolbarMini} from "@/util/vditor-utils";
import {getLastSiteNotice, NoticeForm, postSiteNotice} from "@/services/site-notice";
import {useStore} from "vuex";
import {useMessage} from "naive-ui";

const contentEditor = ref()

const title = ref('')

const message = useMessage();
const store = useStore();
const handleSubmit = () => {
  let form: NoticeForm = {
    content: contentEditor.value.getValue(),
    title: title.value
  }

  postSiteNotice(store.state.auth, form).then((res: any) => {
    if (res.code === 0) {
      message.success('公告发布成功！')
    }

  })
}

onMounted(() => {
  contentEditor.value = new Vditor('vditor', {
    height: 800,
    toolbar: toolbarMini,
    toolbarConfig: {
      pin: true,
    },
    cache: {
      id: 'vditor-notice',
      enable: true,
    },
    counter: {
      enable: true,
      max: 500
    },
    after: () => {
      getLastSiteNotice().then(res => {
        contentEditor.value.setValue(res.data.content)
        title.value = res.data.title
      })

    },
  })
})

</script>

<style scoped>

</style>
