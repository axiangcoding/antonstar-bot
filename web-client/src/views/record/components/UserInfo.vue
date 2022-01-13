<template>
  <div>
    <el-space style="width: 100%" direction="vertical" fill>
      <el-badge value="已被封禁！" :hidden="!gaijinInfo.banned">
        <el-avatar shape="square" :size="size"
                   src="https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png"></el-avatar>
      </el-badge>
      <el-link :href="'https://warthunder.com/'+gaijinInfo.clan_url" target="_blank"
               :underline="false" class="link-clan">{{ gaijinInfo.clan }}
      </el-link>
      <div class="text-title-1">【{{ gaijinInfo.title }}】{{ gaijinInfo.nick }}</div>
      <el-tag size="large" color="goldenrod" type="info" class="tag-level">
        等级 {{ gaijinInfo.level }}
      </el-tag>
      <div class="text-title-2">注册于 {{ gaijinInfo.register_date }}</div>
      <div>
        <el-space>
          <el-tag size="medium" type="info">本数据获取于 {{ gaijinInfo.updated_at }}</el-tag>
          <el-link type="primary" :icon="Refresh" @click="refreshInfoQueries(gaijinInfo.nick)">数据过时？点我刷新</el-link>
        </el-space>
      </div>
      <el-row :gutter="10">
        <el-col :span="8" v-for="(index,key) in gaijinInfo.user_stat" :key="key">
          <GaijinStatCard :data="index" :title="key"></GaijinStatCard>
        </el-col>
      </el-row>

      <!--{{ gaijinInfo.user_rate }}-->
      <el-tabs  type="card" >
        <el-tab-pane label="空军">
          <el-row :gutter="10">
            <el-col :span="8" v-for="(index,key) in gaijinInfo.user_rate.aviation" :key="key">
              <GaijinRateAviationCard :data="index" :title="key"></GaijinRateAviationCard>
            </el-col>
          </el-row>
        </el-tab-pane>

        <el-tab-pane label="陆军">
          <el-row :gutter="10">
            <el-col :span="8" v-for="(index,key) in gaijinInfo.user_rate.ground_vehicles" :key="key">
              <GaijinGroundCard :data="index" :title="key"></GaijinGroundCard>
            </el-col>
          </el-row>
        </el-tab-pane>
        <el-tab-pane label="海军">
          <el-row :gutter="10">
            <el-col :span="8" v-for="(index,key) in gaijinInfo.user_rate.fleet" :key="key">
              <GaijinFleetCard :data="index" :title="key"></GaijinFleetCard>
            </el-col>
          </el-row>
        </el-tab-pane>

      </el-tabs>


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

<script lang="ts">
export default {
  props: {
    thunderskillInfo: {
      'found': false,
    },
    gaijinInfo: {'found': false,}
  }
}
</script>

<script setup lang="ts">
import {onMounted, ref} from 'vue'
import {
  Refresh,
} from '@element-plus/icons-vue'
import GaijinStatCard from "./GaijinStatCard.vue";
import GaijinRateAviationCard from "./GaijinRateAviationCard.vue";
import GaijinGroundCard from "./GaijinGroundCard.vue";
import GaijinFleetCard from "./GaijinFleetCard.vue";
import service from "../../../util/request";
import {ElMessage} from "element-plus";

const size = ref('')

const refreshInfoQueries = (nick: string) => {
  service.post('v1/war_thunder/userinfo/refresh',
      {}, {
        params: {
          "nickname": nick
        }
      }).then(res => {
    if (res.data['refresh'] === true) {
      ElMessage({
        message: '已发送查询请求',
        type: 'success',
      })
    } else {
      ElMessage({
        type: 'warning',
        message: "同一个玩家24小时内仅能查询一次！"
      })
    }
  })
}

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

:deep(.el-card__header) {
  background-color: #f4f4f6;
}

.link-clan {
  font-size: 20px;
  color: goldenrod;
}
</style>
