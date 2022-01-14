<template>
  <el-card shadow="hover" :body-style="{'padding':0}">
    <template #header>
      <div class="card-header">
                <span style="font-size: 20px;font-weight: bolder">{{
                    title === 'ab' ? '街机娱乐' : title === 'rb' ? '历史性能' : '全真模拟'
                  }}</span>
        <!--<el-button type="info" size="small" disabled>导出为图片</el-button>-->
      </div>
    </template>
    <el-descriptions
        :column="1"
        border
    >
      <el-descriptions-item width="50%" label="参战次数(海军)">
        {{ data['参战次数(海军)'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="参战次数(舰船)">
        {{ data['参战次数(舰船)'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="参战次数(鱼雷艇)">
        {{ data['参战次数(鱼雷艇)'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="参战次数(炮艇)">
        {{ data['参战次数(炮艇)'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="参战次数(鱼雷炮艇)">
        {{ data['参战次数(鱼雷炮艇)'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="参战次数(猎潜艇)">
        {{ data['参战次数(猎潜艇)'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="参战次数(海军驳渡船)">
        {{ data['参战次数(海军驳渡船)'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="游戏时长(海战)">
        {{ data['游戏时长(海战)'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="游戏时长(船舰)">
        {{ data['游戏时长(船舰)'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="游戏时长(鱼雷艇)">
        {{ data['游戏时长(鱼雷艇)'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="游戏时长(炮艇)">
        {{ data['游戏时长(炮艇)'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="游戏时长(鱼雷炮艇)">
        {{ data['游戏时长(鱼雷炮艇)'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="游戏时长(猎潜艇)">
        {{ data['游戏时长(猎潜艇)'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="游戏时长(驱逐舰)">
        {{ data['游戏时长(驱逐舰)'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="游戏时长(海军驳渡船)">
        {{ data['游戏时长(海军驳渡船)'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="击毁目标总计">
        {{ data['击毁目标总计'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="空中单位摧毁数">
        {{ data['空中单位摧毁数'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="地面单位摧毁数">
        {{ data['地面单位摧毁数'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="水面单位摧毁数">
        {{ data['水面单位摧毁数'] }}
      </el-descriptions-item>
    </el-descriptions>
  </el-card>
</template>

<script>
export default {
  name: "GaijinFleetCard",
  props: {
    data: {},
    title: ''
  }
}
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
