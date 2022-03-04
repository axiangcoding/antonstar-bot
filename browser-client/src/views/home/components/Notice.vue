<template>
  <n-card :bordered="false" size="small">
    <template #header>
      <n-h2 prefix="bar" align-text type="primary">全站公告</n-h2>
    </template>
    <n-card embedded>
      <template #header>
        <n-space vertical>
          <n-h1>{{ title }}</n-h1>
          <n-tag type="info">
            <template #avatar>
              <n-avatar
              >
                <n-icon>
                  <UserRegular/>
                </n-icon>
              </n-avatar>
            </template>
            由 {{ editorUserNick }} 于 {{ createAt }} 发布
          </n-tag>
        </n-space>
      </template>
      <div id="site-notice"></div>
    </n-card>

  </n-card>
</template>

<script lang="ts" setup>
import VditorPreview from 'vditor'
import {onMounted, ref} from "vue";
import {getLastSiteNotice} from "@/services/site-notice";
import {parseLocalTime} from "@/util/time";
import {userInfo} from "@/services/user";
import {useStore} from "vuex";
import {UserRegular} from "@vicons/fa";

const editorUserNick = ref('')
const title = ref('')
const createAt = ref('')

const store = useStore();
onMounted(() => {
  getLastSiteNotice().then(res => {
    VditorPreview.preview(document.getElementById("site-notice") as HTMLDivElement, res.data.content)
    title.value = res.data.title
    createAt.value = parseLocalTime(res.data.create_at)
    userInfo(store.state.auth, res.data.editor_user_id).then(res => {
      editorUserNick.value = res.data.nickname
    })
  })
})
</script>

<style scoped>

</style>
