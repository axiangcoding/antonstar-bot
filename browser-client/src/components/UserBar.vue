<template>
  <n-dropdown :options="options" @select="handleSelect">
    <n-avatar

        round
        size="large"
        src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
    />
  </n-dropdown>
</template>

<script lang="ts" setup>
import {ref} from "vue";
import {renderIcon} from "@/util/naive";
import {Edit, SignOutAlt, User} from "@vicons/fa";
import {userLogout} from "@/services/user";
import {useStore} from "vuex";
import {useMessage} from "naive-ui";
import {useRouter} from "vue-router";

const options = ref([
  {
    label: '个人信息',
    key: 'profile',
    icon: renderIcon(User),
    disabled: true
  },
  {
    label: '编辑用户资料',
    key: 'editProfile',
    icon: renderIcon(Edit),
    disabled: true
  },
  {
    label: '退出登录',
    key: 'logout',
    icon: renderIcon(SignOutAlt)
  }
])

const message = useMessage();
const store = useStore();
const router = useRouter();
const handleSelect = (key: string) => {
  switch (key) {
    case "profile":
      break
    case "editProfile":
      break
    case "logout":
      userLogout(store.state.auth).then(res => {
        store.commit('logout')
        router.push({name: 'home'})
        message.success('你已成功退出登录，再见！')
      })
      break
    default:
      break
  }
}


</script>

<style scoped>

</style>
