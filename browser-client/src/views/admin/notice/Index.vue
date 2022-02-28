<template>
  <n-card size="small" :bordered="true" style="text-align: left">
    <n-space vertical>
      <n-h2 prefix="bar" align-text type="primary">修改全站公告</n-h2>
      <div id="vditor" style="text-align: left"></div>
      <n-space justify="end">
        <n-button type="error">重置</n-button>
        <n-button type="success" @click="getHTML">发布</n-button>
      </n-space>
    </n-space>

  </n-card>
</template>

<script lang="ts" setup>
// TODO 编辑全站公告
import {onMounted, ref} from "vue";
import Vditor from "vditor";
import {toolbarMini} from "@/util/vditor-utils";

const contentEditor = ref()


const getHTML = () => {
  console.log(contentEditor.value.getValue());
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
    }
    // after: () => {
    //   contentEditor.value.setValue('')
    // },
  })
})

</script>

<style scoped>

</style>
