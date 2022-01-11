<template>
  <div class="about">
    <el-card>
      <template #header>
        <h1>关于我们</h1>
      </template>
      <el-collapse v-model="activeNames" @change="handleChange">
        <el-collapse-item title="项目介绍" name="0">
          <div class="text-about">
            这个网站设立的初衷是，建立一个完整的战争雷霆资源集散论坛，收集攻略，教程，皮肤，游戏数据，玩家战绩，8111插件等资源，以供玩家使用
            <br/>
            开发者白天里是个大数据工程师，一到夜晚就化身战争雷霆玩家，同时也是B站的一个小UP主。这一来二去，就知道玩家的需求了，知道需求了，那就做点什么吧，于是安东星就出来了
          </div>
        </el-collapse-item>


        <el-collapse-item title="意见反馈" name="1">
          <div class="text-about">
            意见反馈可以让网站随时调整功能的优先级，确保用户呼声最高的功能优先被开发。
            <br/>
            因此，如果你对网站有不同的看法想法，有合理的意见和建议，可以在B站私信
            <a href="https://space.bilibili.com/8696650" target="_blank">@摸鱼的33</a> 进行反馈，当然，最好顺手点个关注啦
          </div>
        </el-collapse-item>
        <el-collapse-item title="作者的闲话" name="2">
          <div class="text-about">
            虽然是关于我们，但是实际上就我一个在开发...网站开发完全是出于热情，以及空闲时间想做点什么。简单来说是：一合计，能干，就干了
            <br/>
            不过相信你们也看得出来，我想把他做成战争雷霆的一个资源集散论坛，但这个路很长很长
            <br/>
            用爱发电能不能做成？做成了能不能长久运营？这些都是问题。合理的运营模式会在开发的过程中不断探索，开源？商业？社区？这些犹未可知。
            <br/>
            不过，在测试版早期，我可以明确说，安东星是免费使用的。当然，服务器资源有限，有些功能会设限制，所以你用起来不一定会很舒服
          </div>
        </el-collapse-item>
        <el-collapse-item title="赞助" name="3">
          <div class="text-about">
            尽管看起来这个网站不咋地，但是是要钱的...服务器的支出就是个大头。因此，在目前还未决定运营方向时，如果负有余力的你可以提供赞助，那是最好不过的了
            <br/>
            <el-image style="width: 240px; height: 440px" :src="payCode" fit="scale-down">
            </el-image>
          </div>
        </el-collapse-item>
        <el-collapse-item title="加入我们" name="4">
          <div class="text-about">
            你刚好懂vue3+vite+element-plus的前端架构？
            <br/>
            你刚好玩go+gin玩得很六？
            <br/>
            你十分凑巧知道如何用scrapy从更多的数据源那获取数据？
            <br/>
            你会按照设计画出漂亮的图片？
            <br/>
            那还等什么，B站私信
            <a href="https://space.bilibili.com/8696650" target="_blank">@摸鱼的33</a> 详细描述你的技能！
            <br/>
            当然，现阶段，就算你成功加入开发组，也没有工资和补贴（悲
          </div>

        </el-collapse-item>
      </el-collapse>

    </el-card>
  </div>
</template>

<script lang="ts" setup>
import {useStore} from "vuex";
import {ref} from "vue";
import qrcode from '../../assets/image/1641908835717.jpg'

const store = useStore()
const activeNames = ref(['0', '1', '2', '3', '4'])
const handleChange = (val: string[]) => {
  console.log(val)
}


const payCode = ref(qrcode)

</script>

<style lang="scss" scoped>
.about {

}

:deep(.el-collapse-item__header) {
  font-size: 18px;
  font-weight: bolder;
}

.text-about {
  font-size: 16px;
}

</style>
