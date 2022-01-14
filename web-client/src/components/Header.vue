<template>
  <el-header class="header" height="60px">
    <el-row justify="space-between" style="height: 100%">
      <el-col :xs="4" :sm="4" :md="4" :lg="6" :xl="4">
        <div>
          <el-image
              @click="pushRouter('Home')"
              :src="url"
              fit="scale-down"
              class="img-logo"></el-image>
          <el-tag
              class="tag-version hidden-md-and-down"
              type="danger"
              effect="dark"
          >{{ version }}
          </el-tag>
        </div>
      </el-col>

      <el-col :xs="18" :sm="18" :md="18" :lg="16" :xl="18">
        <el-menu
            style="border: 0"
            :default-active="activeIndex"
            mode="horizontal"
            @select="handleSelect"
        >
          <el-menu-item v-for="item in titleList" :index="item.id" :disabled="item.disabled">
            <span style="font-size: var(--el-font-size-medium);font-weight: bold">{{ item.name }}</span>
          </el-menu-item>
        </el-menu>
      </el-col>
      <el-col :xs="2" :sm="2" :md="2" :lg="2" :xl="2" style="text-align: right">
        <el-dropdown>
          <el-avatar size="default"
                     src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
          ></el-avatar>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item>Action 1</el-dropdown-item>
              <el-dropdown-item>Action 2</el-dropdown-item>
              <el-dropdown-item>Action 3</el-dropdown-item>
              <el-dropdown-item>Action 4</el-dropdown-item>
              <el-dropdown-item>Action 5</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </el-col>
    </el-row>
  </el-header>
</template>

<script lang="ts" setup>
import {computed, ref} from 'vue'
import logo from "../assets/logo/logo.png"
import {useRoute, useRouter} from "vue-router";

const route = useRoute()
const activeIndex = computed(() => {
  return route.name
})

const titleList = [
  {id: 'Record', name: '战绩查询', disabled: false},
  {id: '4', name: '资源下载', disabled: true},
  {id: '6', name: '封神榜', disabled: true},
  {id: 'About', name: '关于我们', disabled: true},
]

const version = ref('Alpha 测试版')

const url = ref(logo)


const router = useRouter()
const handleSelect = (key: string, keyPath: string[]) => {
  pushRouter(key)
}

const pushRouter = (name: string) => {
  router.push({name: name})
}

</script>
<style lang="css" scoped>

.header {
  text-align: left;

  border-bottom: 1px #ccc solid;
  background-color: #ffffff;
}

.tag-version {
  vertical-align: top;
}

.img-logo {
  width: auto;
  height: 60px;
}

.img-logo:hover {
  cursor: pointer;
}
</style>
