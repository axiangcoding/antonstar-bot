<template>
  <n-card size="small" :bordered="false" style="text-align: left">
    <n-space vertical>
      <n-h2 prefix="bar" align-text type="primary">用户管理</n-h2>
      <n-data-table
          ref="table"
          :bordered="false"
          :columns="columns"
          :data="data"
          :pagination="pagination"
      />
    </n-space>
  </n-card>
</template>

<script lang="ts" setup>
import {h, onMounted, reactive, ref} from "vue";
import {NButton} from "naive-ui";

const data = Array.apply(null, {length: 46}).map((_, index) => ({
  key: index,
  user_id: `id ${index}`,
  user_name: `user ${index}`,
  nick_name: `London ${index}`
}))


const columns = [
  {
    title: '用户ID',
    key: 'user_id',
  },
  {
    title: '用户名',
    key: "user_name"
  },
  {
    title: '用户昵称',
    key: 'nick_name',
  },
  {
    title: '邮箱',
    key: 'email'
  },
  {
    title: '用户角色',
    key: 'roles',
  },
  {
    title: '创建时间',
    key: 'created_at'
  },
  {
    title: 'Action',
    key: 'actions',
    render(row: any) {
      return h(
          NButton,
          {
            strong: true,
            type: 'primary',
            tertiary: true,
            size: 'small',
            onClick: () => {
              console.log("click")
            }
          },
          {default: () => '操作'}
      )
    }
  }
]

const paginationReactive = reactive({
  page: 2,
  pageSize: 5,
  pageCount: 200,
  showSizePicker: true,
  pageSizes: [3, 5, 7],
  onChange: (page: number) => {
    paginationReactive.page = page
  },
  onUpdatePageSize: (pageSize: number) => {
    paginationReactive.pageSize = pageSize
    paginationReactive.page = 1
  },
  prefix({itemCount}: any) {
    return `总共有 ${itemCount} 个注册用户`
  }
})

const pagination = paginationReactive
</script>


<style scoped>

</style>
