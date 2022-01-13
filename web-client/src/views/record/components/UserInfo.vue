<template>
  <div>
    <el-space v-if="!gaijinInfo.found">
      <h1>未找到来自官网的该账号信息！</h1>
    </el-space>

    <el-space v-else style="width: 100%" direction="vertical" fill>
      <el-badge value="已被封禁！" :hidden="!gaijinInfo.banned">
        <el-avatar shape="square" :size="size"
                   src="https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png"></el-avatar>
      </el-badge>
      <el-link :href="'https://warthunder.com/'+gaijinInfo.clan_url" target="_blank"
               style="font-size: 20px;color: goldenrod"
               :underline="false">{{ gaijinInfo.clan }}
      </el-link>
      <div class="text-title-1">【{{ gaijinInfo.title }}】{{ gaijinInfo.nick }}</div>
      <el-tag size="large" color="goldenrod" type="info" class="tag-level">
        等级 {{ gaijinInfo.level }}
      </el-tag>
      <div class="text-title-2">注册于 {{ gaijinInfo.register_date }}</div>
      <div>
        <el-space>
          <el-tag size="medium" type="info">本数据获取于 {{ gaijinInfo.updated_at }}</el-tag>
          <el-link type="primary" :icon="Refresh">数据过时？点我刷新</el-link>
        </el-space>
      </div>
      <el-space :size="size" spacer="|">
        <el-button size="small" type="success" plain disabled>这是我</el-button>
        <el-button size="small" type="danger" plain disabled>举报为疑似外挂</el-button>
      </el-space>
      {{ gaijinInfo }}
      <el-row :gutter="10">
        <el-col :span="8">
          <el-card shadow="hover">
            <template #header>
              <div class="card-header">
                <span style="font-size: 20px;font-weight: bolder">街机模式</span>
                <el-button type="primary" size="medium" round disabled>导出为图片</el-button>
              </div>
            </template>
            <el-descriptions
                title="总览"
                :column="1"
                border
            >
              <template #extra>

              </template>
              <el-descriptions-item label="胜利场次">
                placeholder
              </el-descriptions-item>
              <el-descriptions-item label="任务总数">
                placeholder
              </el-descriptions-item>
              <el-descriptions-item label="作战胜率">
                placeholder
              </el-descriptions-item>
              <el-descriptions-item label="阵亡数">
                placeholder
              </el-descriptions-item>
              <el-descriptions-item label="银狮获得数">
                placeholder
              </el-descriptions-item>
              <el-descriptions-item label="游戏时间">
                placeholder
              </el-descriptions-item>
              <el-descriptions-item label="空中单位摧毁数">
                placeholder
              </el-descriptions-item>
              <el-descriptions-item label="地面单位摧毁数">
                placeholder
              </el-descriptions-item>
              <el-descriptions-item label="水面单位摧毁数">
                placeholder
              </el-descriptions-item>
            </el-descriptions>
          </el-card>

        </el-col>
        <el-col :span="8">

        </el-col>
        <el-col :span="8">

        </el-col>
      </el-row>
      <el-row :gutter="10">
        <!--  <el-col :span="8">-->
        <!--    <el-card>-->
        <!--      <el-descriptions-->
        <!--          class="margin-top"-->
        <!--          title="官网"-->
        <!--          :column="1"-->
        <!--          border-->
        <!--      >-->
        <!--        <template #extra>-->
        <!--          <el-button type="primary">Operation</el-button>-->
        <!--        </template>-->
        <!--        <el-descriptions-item>-->
        <!--          <template #label>-->
        <!--            <div class="cell-item">-->
        <!--              <el-icon :style="iconStyle">-->
        <!--                <user/>-->
        <!--              </el-icon>-->
        <!--              Username-->
        <!--            </div>-->
        <!--          </template>-->
        <!--          kooriookami-->
        <!--        </el-descriptions-item>-->
        <!--        <el-descriptions-item>-->
        <!--          <template #label>-->
        <!--            <div class="cell-item">-->
        <!--              <el-icon :style="iconStyle">-->
        <!--                <iphone/>-->
        <!--              </el-icon>-->
        <!--              Telephone-->
        <!--            </div>-->
        <!--          </template>-->
        <!--          18100000000-->
        <!--        </el-descriptions-item>-->
        <!--        <el-descriptions-item>-->
        <!--          <template #label>-->
        <!--            <div class="cell-item">-->
        <!--              <el-icon :style="iconStyle">-->
        <!--                <location/>-->
        <!--              </el-icon>-->
        <!--              Place-->
        <!--            </div>-->
        <!--          </template>-->
        <!--          Suzhou-->
        <!--        </el-descriptions-item>-->
        <!--        <el-descriptions-item>-->
        <!--          <template #label>-->
        <!--            <div class="cell-item">-->
        <!--              <el-icon :style="iconStyle">-->
        <!--                <tickets/>-->
        <!--              </el-icon>-->
        <!--              Remarks-->
        <!--            </div>-->
        <!--          </template>-->
        <!--          <el-tag size="small">School</el-tag>-->
        <!--        </el-descriptions-item>-->
        <!--        <el-descriptions-item>-->
        <!--          <template #label>-->
        <!--            <div class="cell-item">-->
        <!--              <el-icon :style="iconStyle">-->
        <!--                <office-building/>-->
        <!--              </el-icon>-->
        <!--              Address-->
        <!--            </div>-->
        <!--          </template>-->
        <!--          No.1188, Wuzhong Avenue, Wuzhong District, Suzhou, Jiangsu Province-->
        <!--        </el-descriptions-item>-->
        <!--      </el-descriptions>-->
        <!--    </el-card>-->

        <!--  </el-col>-->
        <!--  <el-col :span="8">-->
        <!--    <el-card>-->
        <!--      <el-descriptions-->
        <!--          class="margin-top"-->
        <!--          title="官网"-->
        <!--          :column="1"-->
        <!--          border-->
        <!--      >-->
        <!--        <template #extra>-->
        <!--          <el-button type="primary">Operation</el-button>-->
        <!--        </template>-->
        <!--        <el-descriptions-item>-->
        <!--          <template #label>-->
        <!--            <div class="cell-item">-->
        <!--              <el-icon :style="iconStyle">-->
        <!--                <user/>-->
        <!--              </el-icon>-->
        <!--              Username-->
        <!--            </div>-->
        <!--          </template>-->
        <!--          kooriookami-->
        <!--        </el-descriptions-item>-->
        <!--        <el-descriptions-item>-->
        <!--          <template #label>-->
        <!--            <div class="cell-item">-->
        <!--              <el-icon :style="iconStyle">-->
        <!--                <iphone/>-->
        <!--              </el-icon>-->
        <!--              Telephone-->
        <!--            </div>-->
        <!--          </template>-->
        <!--          18100000000-->
        <!--        </el-descriptions-item>-->
        <!--        <el-descriptions-item>-->
        <!--          <template #label>-->
        <!--            <div class="cell-item">-->
        <!--              <el-icon :style="iconStyle">-->
        <!--                <location/>-->
        <!--              </el-icon>-->
        <!--              Place-->
        <!--            </div>-->
        <!--          </template>-->
        <!--          Suzhou-->
        <!--        </el-descriptions-item>-->
        <!--        <el-descriptions-item>-->
        <!--          <template #label>-->
        <!--            <div class="cell-item">-->
        <!--              <el-icon :style="iconStyle">-->
        <!--                <tickets/>-->
        <!--              </el-icon>-->
        <!--              Remarks-->
        <!--            </div>-->
        <!--          </template>-->
        <!--          <el-tag size="small">School</el-tag>-->
        <!--        </el-descriptions-item>-->
        <!--        <el-descriptions-item>-->
        <!--          <template #label>-->
        <!--            <div class="cell-item">-->
        <!--              <el-icon :style="iconStyle">-->
        <!--                <office-building/>-->
        <!--              </el-icon>-->
        <!--              Address-->
        <!--            </div>-->
        <!--          </template>-->
        <!--          No.1188, Wuzhong Avenue, Wuzhong District, Suzhou, Jiangsu Province-->
        <!--        </el-descriptions-item>-->
        <!--      </el-descriptions>-->
        <!--    </el-card>-->
        <!--  </el-col>-->
        <!--  <el-col :span="8">-->
        <!--    <el-card>-->
        <!--      <el-descriptions-->
        <!--          class="margin-top"-->
        <!--          title="官网"-->
        <!--          :column="1"-->
        <!--          border-->
        <!--      >-->
        <!--        <template #extra>-->
        <!--          <el-button type="primary">Operation</el-button>-->
        <!--        </template>-->
        <!--        <el-descriptions-item>-->
        <!--          <template #label>-->
        <!--            <div class="cell-item">-->
        <!--              <el-icon :style="iconStyle">-->
        <!--                <user/>-->
        <!--              </el-icon>-->
        <!--              Username-->
        <!--            </div>-->
        <!--          </template>-->
        <!--          kooriookami-->
        <!--        </el-descriptions-item>-->
        <!--        <el-descriptions-item>-->
        <!--          <template #label>-->
        <!--            <div class="cell-item">-->
        <!--              <el-icon :style="iconStyle">-->
        <!--                <iphone/>-->
        <!--              </el-icon>-->
        <!--              Telephone-->
        <!--            </div>-->
        <!--          </template>-->
        <!--          18100000000-->
        <!--        </el-descriptions-item>-->
        <!--        <el-descriptions-item>-->
        <!--          <template #label>-->
        <!--            <div class="cell-item">-->
        <!--              <el-icon :style="iconStyle">-->
        <!--                <location/>-->
        <!--              </el-icon>-->
        <!--              Place-->
        <!--            </div>-->
        <!--          </template>-->
        <!--          Suzhou-->
        <!--        </el-descriptions-item>-->
        <!--        <el-descriptions-item>-->
        <!--          <template #label>-->
        <!--            <div class="cell-item">-->
        <!--              <el-icon :style="iconStyle">-->
        <!--                <tickets/>-->
        <!--              </el-icon>-->
        <!--              Remarks-->
        <!--            </div>-->
        <!--          </template>-->
        <!--          <el-tag size="small">School</el-tag>-->
        <!--        </el-descriptions-item>-->
        <!--        <el-descriptions-item>-->
        <!--          <template #label>-->
        <!--            <div class="cell-item">-->
        <!--              <el-icon :style="iconStyle">-->
        <!--                <office-building/>-->
        <!--              </el-icon>-->
        <!--              Address-->
        <!--            </div>-->
        <!--          </template>-->
        <!--          No.1188, Wuzhong Avenue, Wuzhong District, Suzhou, Jiangsu Province-->
        <!--        </el-descriptions-item>-->
        <!--      </el-descriptions>-->
        <!--    </el-card>-->
        <!--  </el-col>-->
      </el-row>
    </el-space>
  </div>


</template>

<script setup lang="ts">
import {onMounted, ref} from 'vue'
import {
  User,
  Iphone,
  Location,
  Refresh,
  Tickets,
  OfficeBuilding,
} from '@element-plus/icons-vue'
import service from "../../../util/request";

const size = ref('')


const gaijinInfo = ref({})
const thunderskillInfo = ref({})
const mockGetInfo = () => {
  service.get('v1/war_thunder/userinfo/query',
      {
        params: {
          "query_id": '770e56e4-4338-408d-bbd5-694f898ae94d'
        }
      }).then(res => {
    gaijinInfo.value = res.data['gaijin']
    thunderskillInfo.value = res.data['thunder_skill']
  })
}

onMounted(mockGetInfo)

</script>

<style lang="scss" scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.tag-level {
  font-size: 20px;
  font-weight: bolder;
  color: gold;
}

.text-title-1 {
  font-size: 20px;
  font-weight: bolder;
}

.text-title-2 {
  font-size: 18px;
  font-weight: bolder;
}
</style>
