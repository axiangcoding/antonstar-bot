<template>
  <n-card>
    <template #header>
      <n-h1>关于我们</n-h1>
    </template>
    <n-space vertical style="text-align: left">
      <n-h2>项目介绍</n-h2>
      安东星是一个战争雷霆的资源网站，目标是聚合更多的玩家，建立一个完备的玩家社区。目前提供国际服玩家的资料查找，战绩查询的功能。
      <br/>
      后续，我们将陆续开放更多，更好玩的功能，进一步完善社区的建设。详情可见下面的更新计划，敬请期待
      <n-divider/>
      <n-h2>更新计划</n-h2>
      由于目前资源有限，因此更新计划也会跟着调整，但是大体的方向如下：
      <n-ol>
        <n-li>启用站内账号 —— 一切功能的前提是网站开始支持账号系统</n-li>
        <n-li>新增玩家标记 —— 被TK了？标记他，让其他用户给你评评理。开挂了？标记他，正义制裁后你会收到私信通知</n-li>
        <n-li>资源收集整合 —— 8111端口插件？皮肤？瞄具？不用费心去找了，这里都会有</n-li>
        <n-li>发布文章帖子 —— 攻略？教程？最新dev日志？新闻？都有的，开整！</n-li>
        <n-li>载具wiki —— 想看某个载具的参数，但是还要进游戏？不用那么麻烦，来这也能看啦</n-li>
        <n-li>商业化（靠充钱）or 社区维护（靠赞助） —— 虽然大家都知道运营要花钱（而且挺多），但是，你能再表演一下那个嘛.jpg</n-li>
      </n-ol>
      <n-divider/>
      <n-h2>赞助项目</n-h2>
      目前本项目，由StGGA联队提供测试工具人和启动资金支持，联队QQ群：662698286
      <br/>
      你可以通过下面的支付宝二维码向项目提供赞助，在此提前感谢慷慨解囊的你
      <n-image
          width="300"
          :src="qrAliPay"
          preview-disabled
      />
      <n-divider/>
      <n-h2>开发者</n-h2>
      目前，开发者一共有2个（其中1个是被骗来用爱发电的），画师1个（大触）。
      网站小老板联系方式如下：
      <div>B站：<a href="https://space.bilibili.com/8696650" target="_blank">摸鱼的33</a></div>
      <div>邮箱：<a href="mailto:axiangcoding@gmail.com" target="_blank">axiangcoding@gmail.com</a></div>
      <div>粉丝QQ群：689874918</div>
      <n-divider/>
      <n-h2>加入我们</n-h2>
      在目前这个阶段，我们很需要项目开发人员，比如：
      <n-ul>
        <n-li>会 vue3 + vite + ts 的前端大佬</n-li>
        <n-li>知道如何用 scrapy 从官网，wiki，thunder skill等站点爬取公开数据的爬虫大师</n-li>
        <n-li>搞 go + gin 的后端大哥</n-li>
        <n-li>会画图，有美感的大触</n-li>
      </n-ul>
      当然必须要说明的一点是，赞助的收入将全部投入服务器的花销中，所以有极大的概率你只能用爱发电...请谨慎考虑
      <n-divider/>
    </n-space>
  </n-card>
</template>

<script lang="ts" setup>
import qrAliPay from "@/assets/image/qr_alipay.jpg"
</script>

<style scoped>

</style>
