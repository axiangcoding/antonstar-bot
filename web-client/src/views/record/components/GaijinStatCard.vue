<template>
  <el-card shadow="hover" :body-style="{'padding':0}">
    <template #header>
      <div class="card-header">
                <span style="font-size: 20px;font-weight: bolder">{{
                    title === 'ab' ? '街机娱乐' : title === 'rb' ? '历史性能' : '全真模拟'
                  }}</span>
        <el-button type="info" size="small" disabled>导出为图片</el-button>
      </div>
    </template>
    <el-descriptions
        :column="1"
        border
    >
      <el-descriptions-item width="50%" label="胜利场次">
        {{ data['胜利场次'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="任务总数">
        {{ data['任务总数'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="作战胜率">
        {{ data['作战胜率'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="阵亡数">
        {{ data['阵亡数'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="银狮获得数">
        {{ data['银狮获得数'] }}
      </el-descriptions-item>
      <el-descriptions-item width="50%" label="游戏时间">
        {{ data['游戏时间'] }}
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
  name: "GaijinStatCard",
  props: {
    data: {},
    title: ''
  }
}
</script>

<style lang="scss" scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
