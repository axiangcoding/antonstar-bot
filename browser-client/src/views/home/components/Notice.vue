<template>
  <n-card :bordered="false" size="small">
    <template #header>
      <n-h2 prefix="bar" align-text type="primary">全站公告</n-h2>
    </template>
    <n-space vertical>
      <n-tag type="primary">由 {{ editorUserNick }} 于 {{ createAt }} 发布</n-tag>
      <div id="site-notice"></div>
    </n-space>
  </n-card>
</template>

<script lang="ts" setup>
import VditorPreview from 'vditor'
import {onMounted, ref} from "vue";
import {getLastSiteNotice} from "@/services/site-notice";
import {parseLocalTime} from "@/util/time";
import {userInfo} from "@/services/user";
import {useStore} from "vuex";

const editorUserNick = ref('')
const createAt = ref('')

const store = useStore();
onMounted(() => {
  getLastSiteNotice().then(res => {
    VditorPreview.preview(document.getElementById("site-notice") as HTMLDivElement, res.data.content)
    createAt.value = parseLocalTime(res.data.create_at)
    userInfo(store.state.auth, res.data.editor_user_id).then(res => {
      editorUserNick.value = res.data.nickname
    })
  })
})
</script>

<style scoped>

</style>
