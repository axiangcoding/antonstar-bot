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

const options = ref([
  {
    label: '个人信息',
    key: 'profile',
    icon: renderIcon(User)
  },
  {
    label: '编辑用户资料',
    key: 'editProfile',
    icon: renderIcon(Edit)
  },
  {
    label: '退出登录',
    key: 'logout',
    icon: renderIcon(SignOutAlt)
  }
])

const message = useMessage();
const store = useStore();
const handleSelect = (key: string) => {
  switch (key) {
    case "profile":
      break
    case "editProfile":
      break
    case "logout":
      userLogout(store.state.auth).then(res => {
        store.commit('setAuth', '')
        store.commit('setLogin', false)
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
