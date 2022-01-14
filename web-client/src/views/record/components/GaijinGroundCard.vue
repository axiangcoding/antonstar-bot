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
      <el-descriptions-item  width="50%" label="参战次数(陆战)">
        {{ data['参战次数(陆战)'] }}
      </el-descriptions-item>
      <el-descriptions-item  width="50%" label="参战次数(地面载具)">
        {{ data['参战次数(地面载具)'] }}
      </el-descriptions-item>
      <el-descriptions-item  width="50%" label="参战次数(坦克歼击车)">
        {{ data['参战次数(坦克歼击车)'] }}
      </el-descriptions-item>
      <el-descriptions-item  width="50%" label="参战次数(重型坦克)">
        {{ data['参战次数(重型坦克)'] }}
      </el-descriptions-item>
      <el-descriptions-item  width="50%" label="参战次数(自行防空炮)">
        {{ data['参战次数(自行防空炮)'] }}
      </el-descriptions-item>
      <el-descriptions-item  width="50%" label="游戏时长(陆战)">
        {{ data['游戏时长(陆战)'] }}
      </el-descriptions-item>
      <el-descriptions-item  width="50%" label="游戏时长(地面单位)">
        {{ data['游戏时长(地面单位)'] }}
      </el-descriptions-item>
      <el-descriptions-item  width="50%" label="游戏时长(坦克歼击车)">
        {{ data['游戏时长(坦克歼击车)'] }}
      </el-descriptions-item>
      <el-descriptions-item  width="50%" label="游戏时长(重型坦克)">
        {{ data['游戏时长(重型坦克)'] }}
      </el-descriptions-item>
      <el-descriptions-item  width="50%" label="游戏时长(自行防空炮)">
        {{ data['游戏时长(自行防空炮)'] }}
      </el-descriptions-item>
      <el-descriptions-item  width="50%" label="击毁目标总计">
        {{ data['击毁目标总计'] }}
      </el-descriptions-item>
      <el-descriptions-item  width="50%" label="空中单位摧毁数">
        {{ data['空中单位摧毁数'] }}
      </el-descriptions-item>
      <el-descriptions-item  width="50%" label="地面单位摧毁数">
        {{ data['地面单位摧毁数'] }}
      </el-descriptions-item>
      <el-descriptions-item  width="50%" label="水面单位摧毁数">
        {{ data['水面单位摧毁数'] }}
      </el-descriptions-item>
    </el-descriptions>
  </el-card>
</template>

<script>
export default {
  name: "GaijinGroundCard",
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
