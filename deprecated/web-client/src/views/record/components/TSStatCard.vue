<template>
  <el-card>
    <template #header>
      <div class="card-header">
        <span style="font-size: 16px;font-weight: bolder">
        ThunderSkill效率值
           <el-tooltip
               class="box-item"
               effect="dark"
               content="这个效率值来自Thunder Skill，代表的是你的综合作战水平"
               placement="top-start"
           >
        <el-icon size="14"><QuestionFilled/></el-icon>
      </el-tooltip>

      </span>

      </div>
    </template>
    <div v-if="data.found">
      <el-row>

        <el-col :xs="24" :sm="8" :md="8" :lg="8" :xl="8">
          <el-progress type="dashboard" :percentage="data['a']['kpd']" :color="customColors"
                       :width="cycleWidth"
                       :stroke-width="strokeWidth"
          >
            <template #default="{ percentage }">
              <span class="percentage-value">{{ percentage }}%</span>
              <span class="percentage-value-2">{{ handlePercentText(data['a']['kpd']) }}</span>
              <span class="percentage-label">街机娱乐</span>
            </template>
          </el-progress>
        </el-col>
        <el-col :xs="24" :sm="8" :md="8" :lg="8" :xl="8">
          <el-progress type="dashboard" :percentage="data['r']['kpd']" :color="customColors"
                       :width="cycleWidth"
                       :stroke-width="strokeWidth">
            <template #default="{ percentage }">
              <span class="percentage-value">{{ percentage }}%</span>
              <span class="percentage-value-2">{{ handlePercentText(data['r']['kpd']) }}</span>
              <span class="percentage-label">历史性能</span>
            </template>
          </el-progress>
        </el-col>
        <el-col :xs="24" :sm="8" :md="8" :lg="8" :xl="8">
          <el-progress type="dashboard" :percentage="data['s']['kpd']" :color="customColors" :width="cycleWidth"
                       :stroke-width="strokeWidth">
            <template #default="{ percentage }">
              <span class="percentage-value">{{ percentage }}%</span>
              <span class="percentage-value-2">{{ handlePercentText(data['s']['kpd']) }}</span>
              <span class="percentage-label">全真模拟</span>
            </template>
          </el-progress>
        </el-col>
      </el-row>
    </div>
    <div v-else>
      <h2>未找到你在Thunder skill的数据</h2>
      <el-link type="primary">如何在ThunderSkill中查看自己的战绩？</el-link>
    </div>
  </el-card>
</template>

<script lang="ts">
export default {
  name: "TSStatCare.vue",
  props: {
    data: {}
  }
}
</script>

<script lang="ts" setup>
import {ref} from "vue";
import {QuestionFilled} from "@element-plus/icons-vue"

const customColors = [
  {color: '#ffffff', percentage: 0, text: '在？'},
  {color: '#5cb87a', percentage: 10, text: '战雷之拉'},
  {color: '#1989fa', percentage: 20, text: '战雷之屑'},
  {color: '#6f7ad3', percentage: 50, text: '战雷之虫'},
  {color: '#e4a13d', percentage: 70, text: '战雷之人'},
  {color: '#f56c6c', percentage: 80, text: '战雷之王'},
  {color: '#8B0000', percentage: 90, text: '战雷皇帝'},
  {color: '#000000', percentage: 100, text: '神中神'},
]


const cycleWidth = ref(180)
const strokeWidth = ref(12)

const handlePercentText = (percentage: number) => {
  for (let customColor of customColors) {
    if (percentage < customColor.percentage) {
      return customColor.text
    }
  }
  return customColors[customColors.length - 1].text
}
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.percentage-value {
  display: block;
  margin-top: 10px;
  font-size: 28px;
}

.percentage-label {
  display: block;
  margin-top: 10px;
  font-size: 12px;
}

.percentage-value-2 {
  display: block;
  margin-top: 10px;
  font-size: 20px;
}

</style>
