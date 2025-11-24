<template>
  <div class="settings-container p-6">
    <n-card title="系统设置" class="settings-card">
      <n-tabs type="line" animated>
        <!-- 基本设置 -->
        <n-tab-pane name="basic" tab="基本设置">
          <div class="tab-content">
            <basic-setting/>
          </div>
        </n-tab-pane>
        <!-- AI设置 -->
        <n-tab-pane name="ai" tab="AI设置">
          <div class="tab-content">
            <ai-setting />
          </div>
        </n-tab-pane>
        <!-- 邮件设置 -->
        <n-tab-pane name="email" tab="邮件设置">
          <div class="tab-content">
            <email-setting />
          </div>
        </n-tab-pane>
        <!-- 站点设置 -->
        <n-tab-pane name="site" tab="站点设置">
          <div class="tab-content">
            <site-setting />
          </div>
        </n-tab-pane>
        <!-- 安全设置 -->
        <n-tab-pane name="security" tab="安全设置">
          <div class="tab-content">
            <security-setting />
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
        <!-- SEO设置 -->
        <n-tab-pane name="seo" tab="SEO设置">
          <div class="tab-content">
            <seo-setting />
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
import BasicSetting from './components/basicSetting.vue'
import AiSetting from './components/aiSetting.vue'
import EmailSetting from './components/emailSetting.vue'
import siteSetting from './components/siteSetting.vue'
import SecuritySetting from './components/securitySetting.vue'
import paymentSetting from './components/paymentSetting.vue'
import notifySetting from './components/notifySetting.vue'
import seoSetting from './components/seoSetting.vue'
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
