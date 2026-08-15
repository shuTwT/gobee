<template>
  <div class="settings-container p-6">
    <n-card title="系统设置" class="settings-card">
      <n-tabs type="line" animated>
        <!-- 站点设置 -->
        <n-tab-pane name="site" tab="站点设置">
          <div class="tab-content">
            <site-setting />
          </div>
        </n-tab-pane>
        <!-- 支付设置 -->
        <n-tab-pane name="payment" tab="支付设置">
          <div class="tab-content">
            <payment-setting />
          </div>
        </n-tab-pane>
        <!-- 通知设置 -->
        <n-tab-pane name="notification" tab="通知设置">
          <div class="tab-content">
            <notify-setting />
          </div>
        </n-tab-pane>

        <!-- 备份设置 -->
        <n-tab-pane name="backup" tab="备份设置">
          <div class="tab-content">
            <backup-setting />
          </div>
        </n-tab-pane>

        <!-- 日志设置 -->
        <n-tab-pane name="logs" tab="日志设置">
          <div class="tab-content">
            <log-setting />
          </div>
        </n-tab-pane>
      </n-tabs>
    </n-card>
  </div>
</template>

<script setup lang="ts">
import { useMessage } from 'naive-ui'
import siteSetting from './components/siteSetting.vue'
import paymentSetting from './components/paymentSetting.vue'
import notifySetting from './components/notifySetting.vue'
import backupSetting from './components/backupSetting.vue'
import logSetting from './components/logSetting.vue'
import type { SettingsProps } from './utils/types'



const settings = ref<SettingsProps>({})

</script>

<style scoped>
.settings-container {
  background-color: #f5f7fa;
}

.settings-card {
  max-width: 1600px;
  margin: 0 auto;
}

.tab-content {
  padding: 24px 0;
}

.settings-form {
  max-width: 600px;
}

.slider-value {
  margin-left: 12px;
  min-width: 40px;
  text-align: center;
  font-weight: 500;
}

@media (max-width: 768px) {
  .settings-container {
    padding: 16px;
  }

  .tab-content {
    padding: 16px 0;
  }
}
</style>
