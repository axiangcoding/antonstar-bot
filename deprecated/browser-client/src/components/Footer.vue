<template>
  <n-layout-footer>
    <n-space vertical>
      <n-text strong>安东星 <n-tag type="error" size="small">{{ store.state.systemInfo.version }} beta</n-tag></n-text>
      <div>
        <n-button quaternary @click="router.push({name:'about'})">关于我们</n-button>
        <n-divider vertical/>
        <n-button quaternary @click="showBugReport=true">问题反馈</n-button>
        <n-divider vertical/>
        <n-button quaternary>赞助名单</n-button>
        <n-divider vertical/>
        <n-button quaternary @click="showModal=true">商业合作</n-button>
      </div>
      <n-text italic>Power by 安东星 All Rights Reserved.</n-text>
    </n-space>
    <BusinessWork v-model:show="showModal"/>
    <BugReport v-model:show="showBugReport"/>

  </n-layout-footer>
</template>

<script lang="ts" setup>
import {useRouter} from "vue-router";
import {ref} from "vue";
import {useStore} from "vuex";
import BugReport from "@/components/BugReport.vue";
import BusinessWork from "@/components/BusinessWork.vue";

const router = useRouter()
const store = useStore();
const showModal = ref(false)
const showBugReport = ref(false)


</script>


<style lang="scss" scoped>
.n-layout-footer {
  min-height: var(--footer-height);
  overflow: hidden;
}
</style>
