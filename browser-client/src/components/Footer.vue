<template>
  <n-layout-footer>
    <n-space vertical>
      <n-text strong>安东星 <n-tag type="error" size="small">{{ store.state.systemInfo.version }} beta</n-tag></n-text>
      <div>
        <n-button quaternary @click="router.push({name:'about'})">关于我们</n-button>
        <n-divider vertical/>
        <n-button quaternary>问题反馈</n-button>
        <n-divider vertical/>
        <n-button quaternary>赞助名单</n-button>
        <n-divider vertical/>
        <n-button quaternary @click="showModal=true">商业合作</n-button>
      </div>
      <n-text italic>Power by 安东星 All Rights Reserved.</n-text>
    </n-space>
    <n-modal v-model:show="showModal">
      <n-card
          style="width: 600px;"
          title="商业合作"
          :bordered="false"
          size="huge"
          role="dialog"
          aria-modal="true"
      >
        <template #header-extra></template>
        <div>邮箱：<a href="mailto:axiangcoding@gmail.com" target="_blank">axiangcoding@gmail.com</a></div>
        <!--<template #footer> 人也是要恰饭的嘛</template>-->
      </n-card>
    </n-modal>
  </n-layout-footer>
</template>

<script lang="ts" setup>
import {useRouter} from "vue-router";
import {ref} from "vue";
import {useStore} from "vuex";

const router = useRouter()
const store = useStore();
const showModal = ref(false)
</script>


<style lang="scss" scoped>
.n-layout-footer {
  min-height: var(--footer-height);
  overflow: hidden;
}
</style>
