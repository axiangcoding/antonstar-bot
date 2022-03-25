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
        <n-popconfirm
            @positive-click="restoreOrigin"
            @negative-click=""
        >
          <template #trigger>
            <n-button type="error">重置</n-button>
          </template>
          是否重置到上次保存的内容？
        </n-popconfirm>

        <n-button type="success" @click="handleSubmit">发布</n-button>
      </n-space>
    </n-space>

  </n-card>
</template>

<script lang="ts" setup>
import {onMounted, ref} from "vue";
import Vditor from "vditor";
import {toolbarMini} from "@/util/vditor-utils";
import {getLastSiteNotice, NoticeForm, postSiteNotice} from "@/services/site-notice";
import {useStore} from "vuex";
import {useMessage} from "naive-ui";
import {v4 as uuid} from "uuid";

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

const restoreOrigin = () => {
  contentEditor.value.setValue(originText)
}

// 本地测试需要更改该地址，否则会有代理问题
const uploadUrl = "/api/v1/upload/picture"
let originText = ''
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
    upload: {
      url: uploadUrl,
      // linkToImgUrl: uploadUrl,
      accept: 'image/jpg, image/jpeg, image/png',
      multiple: false,
      max: 4 * 1024 * 1024,
      fieldName: 'file',
      filename(name: string) {
        return uuid() + name.substring(name.lastIndexOf("."))
      },
      headers: {
        'Authorization': store.state.auth
      },
      linkToImgCallback() {
        console.log("外部图片暂不支持转为内部图片")
      },
      format(files: File[], res) {
        let result: any = {
          msg: '',
          code: 0,
          data: {
            "errFiles": [],
            "succMap": {}
          }
        }
        let filename = files[0].name;
        let resJson = JSON.parse(res)
        if (resJson.code == 0) {
          result.data.succMap[filename] = resJson.data.url
        } else {
          result.data.errFiles.push(filename)
        }
        return JSON.stringify(result)
      },
      error(msg) {
        console.log("图片上传失败：" + msg)
      },
    },
    after: () => {
      getLastSiteNotice().then(res => {
        contentEditor.value.setValue(res.data.content)
        originText = res.data.content
        title.value = res.data.title
      })

    },
  })
})

</script>

<style scoped>

</style>
