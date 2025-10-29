<template>
  <div class="settings-container">
    <n-card title="系统设置" class="settings-card">
      <n-tabs type="line" animated>
        <!-- 基本设置 -->
        <n-tab-pane name="basic" tab="基本设置">
          <div class="tab-content">
            <n-form
              ref="basicFormRef"
              :model="basicForm"
              label-placement="left"
              label-width="auto"
              require-mark-placement="right-hanging"
              class="settings-form"
            >
              <n-form-item label="网站名称" path="siteName">
                <n-input v-model:value="basicForm.siteName" placeholder="请输入网站名称" />
              </n-form-item>
              <n-form-item label="网站描述" path="siteDescription">
                <n-input
                  v-model:value="basicForm.siteDescription"
                  type="textarea"
                  placeholder="请输入网站描述"
                />
              </n-form-item>
              <n-form-item label="网站Logo" path="siteLogo">
                <n-input v-model:value="basicForm.siteLogo" placeholder="请输入Logo地址" />
              </n-form-item>
              <n-form-item label="网站图标" path="siteFavicon">
                <n-input v-model:value="basicForm.siteFavicon" placeholder="请输入图标地址" />
              </n-form-item>
              <n-form-item label="关键词" path="keywords">
                <n-input v-model:value="basicForm.keywords" placeholder="请输入关键词，用逗号分隔" />
              </n-form-item>
              <n-form-item label="作者" path="author">
                <n-input v-model:value="basicForm.author" placeholder="请输入作者名称" />
              </n-form-item>
              <n-form-item label="语言" path="language">
                <n-select v-model:value="basicForm.language" :options="languageOptions" />
              </n-form-item>
              <n-form-item label="时区" path="timezone">
                <n-select v-model:value="basicForm.timezone" :options="timezoneOptions" />
              </n-form-item>
              <n-form-item label="日期格式" path="dateFormat">
                <n-input v-model:value="basicForm.dateFormat" placeholder="例如: YYYY-MM-DD" />
              </n-form-item>
              <n-form-item label="时间格式" path="timeFormat">
                <n-input v-model:value="basicForm.timeFormat" placeholder="例如: HH:mm:ss" />
              </n-form-item>
              <n-form-item>
                <n-button type="primary" @click="saveBasicSettings" :loading="basicLoading">
                  保存基本设置
                </n-button>
              </n-form-item>
            </n-form>
          </div>
        </n-tab-pane>

        <!-- AI设置 -->
        <n-tab-pane name="ai" tab="AI设置">
          <div class="tab-content">
            <n-form
              ref="aiFormRef"
              :model="aiForm"
              label-placement="left"
              label-width="auto"
              require-mark-placement="right-hanging"
              class="settings-form"
            >
              <n-form-item label="OpenAI API Key" path="openaiApiKey">
                <n-input
                  v-model:value="aiForm.openaiApiKey"
                  type="password"
                  placeholder="请输入OpenAI API密钥"
                  show-password-on="mousedown"
                />
              </n-form-item>
              <n-form-item label="OpenAI API地址" path="openaiApiUrl">
                <n-input v-model:value="aiForm.openaiApiUrl" placeholder="https://api.openai.com/v1" />
              </n-form-item>
              <n-form-item label="AI模型" path="aiModel">
                <n-select v-model:value="aiForm.aiModel" :options="aiModelOptions" />
              </n-form-item>
              <n-form-item label="AI温度" path="aiTemperature">
                <n-slider v-model:value="aiForm.aiTemperature" :min="0" :max="2" :step="0.1" />
                <span class="slider-value">{{ aiForm.aiTemperature }}</span>
              </n-form-item>
              <n-form-item label="AI最大令牌数" path="aiMaxTokens">
                <n-input-number v-model:value="aiForm.aiMaxTokens" :min="1" :max="8192" />
              </n-form-item>
              <n-form-item label="AI Top P" path="aiTopP">
                <n-slider v-model:value="aiForm.aiTopP" :min="0" :max="1" :step="0.1" />
                <span class="slider-value">{{ aiForm.aiTopP }}</span>
              </n-form-item>
              <n-form-item label="AI频率惩罚" path="aiFrequencyPenalty">
                <n-slider v-model:value="aiForm.aiFrequencyPenalty" :min="-2" :max="2" :step="0.1" />
                <span class="slider-value">{{ aiForm.aiFrequencyPenalty }}</span>
              </n-form-item>
              <n-form-item label="AI存在惩罚" path="aiPresencePenalty">
                <n-slider v-model:value="aiForm.aiPresencePenalty" :min="-2" :max="2" :step="0.1" />
                <span class="slider-value">{{ aiForm.aiPresencePenalty }}</span>
              </n-form-item>
              <n-form-item>
                <n-button type="primary" @click="saveAISettings" :loading="aiLoading">
                  保存AI设置
                </n-button>
                <n-button @click="testAIConnection" style="margin-left: 12px;">
                  测试AI连接
                </n-button>
              </n-form-item>
            </n-form>
          </div>
        </n-tab-pane>

        <!-- 邮件设置 -->
        <n-tab-pane name="email" tab="邮件设置">
          <div class="tab-content">
            <n-form
              ref="emailFormRef"
              :model="emailForm"
              label-placement="left"
              label-width="auto"
              require-mark-placement="right-hanging"
              class="settings-form"
            >
              <n-form-item label="SMTP服务器" path="smtpHost">
                <n-input v-model:value="emailForm.smtpHost" placeholder="例如: smtp.gmail.com" />
              </n-form-item>
              <n-form-item label="SMTP端口" path="smtpPort">
                <n-input-number v-model:value="emailForm.smtpPort" :min="1" :max="65535" />
              </n-form-item>
              <n-form-item label="SMTP用户名" path="smtpUsername">
                <n-input v-model:value="emailForm.smtpUsername" placeholder="请输入SMTP用户名" />
              </n-form-item>
              <n-form-item label="SMTP密码" path="smtpPassword">
                <n-input
                  v-model:value="emailForm.smtpPassword"
                  type="password"
                  placeholder="请输入SMTP密码"
                  show-password-on="mousedown"
                />
              </n-form-item>
              <n-form-item label="加密方式" path="smtpEncryption">
                <n-select v-model:value="emailForm.smtpEncryption" :options="encryptionOptions" />
              </n-form-item>
              <n-form-item label="发件人邮箱" path="senderEmail">
                <n-input v-model:value="emailForm.senderEmail" placeholder="请输入发件人邮箱地址" />
              </n-form-item>
              <n-form-item label="发件人名称" path="senderName">
                <n-input v-model:value="emailForm.senderName" placeholder="请输入发件人名称" />
              </n-form-item>
              <n-form-item>
                <n-button type="primary" @click="saveEmailSettings" :loading="emailLoading">
                  保存邮件设置
                </n-button>
                <n-button @click="testEmailConnection" style="margin-left: 12px;">
                  测试邮件连接
                </n-button>
              </n-form-item>
            </n-form>
          </div>
        </n-tab-pane>

        <!-- 站点设置 -->
        <n-tab-pane name="site" tab="站点设置">
          <div class="tab-content">
            <n-form
              ref="siteFormRef"
              :model="siteForm"
              label-placement="left"
              label-width="auto"
              require-mark-placement="right-hanging"
              class="settings-form"
            >
              <n-form-item label="维护模式" path="maintenanceMode">
                <n-switch v-model:value="siteForm.maintenanceMode" />
              </n-form-item>
              <n-form-item label="允许注册" path="allowRegistration">
                <n-switch v-model:value="siteForm.allowRegistration" />
              </n-form-item>
              <n-form-item label="邮箱验证" path="emailVerification">
                <n-switch v-model:value="siteForm.emailVerification" />
              </n-form-item>
              <n-form-item label="评论审核" path="commentModeration">
                <n-switch v-model:value="siteForm.commentModeration" />
              </n-form-item>
              <n-form-item label="文件上传限制" path="uploadLimit">
                <n-input-number v-model:value="siteForm.uploadLimit" :min="1" :max="100">
                  <template #suffix>MB</template>
                </n-input-number>
              </n-form-item>
              <n-form-item label="允许的文件类型" path="allowedFileTypes">
                <n-dynamic-tags v-model:value="siteForm.allowedFileTypes" />
              </n-form-item>
              <n-form-item label="启用缓存" path="enableCache">
                <n-switch v-model:value="siteForm.enableCache" />
              </n-form-item>
              <n-form-item label="缓存时间" path="cacheTime">
                <n-input-number v-model:value="siteForm.cacheTime" :min="1">
                  <template #suffix>分钟</template>
                </n-input-number>
              </n-form-item>
              <n-form-item label="启用SSL" path="enableSSL">
                <n-switch v-model:value="siteForm.enableSSL" />
              </n-form-item>
              <n-form-item label="启用CDN" path="enableCDN">
                <n-switch v-model:value="siteForm.enableCDN" />
              </n-form-item>
              <n-form-item label="CDN地址" path="cdnUrl">
                <n-input v-model:value="siteForm.cdnUrl" placeholder="请输入CDN地址" />
              </n-form-item>
              <n-form-item>
                <n-button type="primary" @click="saveSiteSettings" :loading="siteLoading">
                  保存站点设置
                </n-button>
              </n-form-item>
            </n-form>
          </div>
        </n-tab-pane>

        <!-- 安全设置 -->
        <n-tab-pane name="security" tab="安全设置">
          <div class="tab-content">
            <n-form
              ref="securityFormRef"
              :model="securityForm"
              label-placement="left"
              label-width="auto"
              require-mark-placement="right-hanging"
              class="settings-form"
            >
              <n-form-item label="启用两步验证" path="enableTwoFactor">
                <n-switch v-model:value="securityForm.enableTwoFactor" />
              </n-form-item>
              <n-form-item label="登录失败重试次数" path="maxLoginAttempts">
                <n-input-number v-model:value="securityForm.maxLoginAttempts" :min="3" :max="10" />
              </n-form-item>
              <n-form-item label="账户锁定时间" path="lockoutDuration">
                <n-input-number v-model:value="securityForm.lockoutDuration" :min="5" :max="60">
                  <template #suffix>分钟</template>
                </n-input-number>
              </n-form-item>
              <n-form-item label="会话超时时间" path="sessionTimeout">
                <n-input-number v-model:value="securityForm.sessionTimeout" :min="30" :max="1440">
                  <template #suffix>分钟</template>
                </n-input-number>
              </n-form-item>
              <n-form-item label="密码最小长度" path="minPasswordLength">
                <n-input-number v-model:value="securityForm.minPasswordLength" :min="6" :max="32" />
              </n-form-item>
              <n-form-item label="密码复杂度要求" path="passwordComplexity">
                <n-checkbox-group v-model:value="securityForm.passwordComplexity">
                  <n-space vertical>
                    <n-checkbox value="uppercase">包含大写字母</n-checkbox>
                    <n-checkbox value="lowercase">包含小写字母</n-checkbox>
                    <n-checkbox value="number">包含数字</n-checkbox>
                    <n-checkbox value="special">包含特殊字符</n-checkbox>
                  </n-space>
                </n-checkbox-group>
              </n-form-item>
              <n-form-item label="API访问限制" path="apiRateLimit">
                <n-input-number v-model:value="securityForm.apiRateLimit" :min="10" :max="1000">
                  <template #suffix>次/分钟</template>
                </n-input-number>
              </n-form-item>
              <n-form-item label="IP白名单" path="ipWhitelist">
                <n-dynamic-tags v-model:value="securityForm.ipWhitelist" />
              </n-form-item>
              <n-form-item label="IP黑名单" path="ipBlacklist">
                <n-dynamic-tags v-model:value="securityForm.ipBlacklist" />
              </n-form-item>
              <n-form-item>
                <n-button type="primary" @click="saveSecuritySettings" :loading="securityLoading">
                  保存安全设置
                </n-button>
              </n-form-item>
            </n-form>
          </div>
        </n-tab-pane>

        <!-- 支付设置 -->
        <n-tab-pane name="payment" tab="支付设置">
          <div class="tab-content">
            <n-form
              ref="paymentFormRef"
              :model="paymentForm"
              label-placement="left"
              label-width="auto"
              require-mark-placement="right-hanging"
              class="settings-form"
            >
              <n-form-item label="启用支付宝" path="enableAlipay">
                <n-switch v-model:value="paymentForm.enableAlipay" />
              </n-form-item>
              <n-form-item label="支付宝应用ID" path="alipayAppId">
                <n-input v-model:value="paymentForm.alipayAppId" placeholder="请输入支付宝应用ID" />
              </n-form-item>
              <n-form-item label="支付宝私钥" path="alipayPrivateKey">
                <n-input
                  v-model:value="paymentForm.alipayPrivateKey"
                  type="textarea"
                  placeholder="请输入支付宝私钥"
                  :rows="4"
                />
              </n-form-item>
              <n-form-item label="支付宝公钥" path="alipayPublicKey">
                <n-input
                  v-model:value="paymentForm.alipayPublicKey"
                  type="textarea"
                  placeholder="请输入支付宝公钥"
                  :rows="4"
                />
              </n-form-item>
              <n-divider />
              <n-form-item label="启用微信支付" path="enableWechatPay">
                <n-switch v-model:value="paymentForm.enableWechatPay" />
              </n-form-item>
              <n-form-item label="微信商户号" path="wechatMchId">
                <n-input v-model:value="paymentForm.wechatMchId" placeholder="请输入微信商户号" />
              </n-form-item>
              <n-form-item label="微信API密钥" path="wechatApiKey">
                <n-input v-model:value="paymentForm.wechatApiKey" placeholder="请输入微信API密钥" />
              </n-form-item>
              <n-form-item label="微信APP ID" path="wechatAppId">
                <n-input v-model:value="paymentForm.wechatAppId" placeholder="请输入微信APP ID" />
              </n-form-item>
              <n-divider />
              <n-form-item label="支付回调域名" path="paymentNotifyUrl">
                <n-input v-model:value="paymentForm.paymentNotifyUrl" placeholder="https://yourdomain.com" />
              </n-form-item>
              <n-form-item label="订单超时时间" path="orderTimeout">
                <n-input-number v-model:value="paymentForm.orderTimeout" :min="5" :max="1440">
                  <template #suffix>分钟</template>
                </n-input-number>
              </n-form-item>
              <n-form-item label="退款审核" path="refundReview">
                <n-switch v-model:value="paymentForm.refundReview" />
              </n-form-item>
              <n-form-item>
                <n-button type="primary" @click="savePaymentSettings" :loading="paymentLoading">
                  保存支付设置
                </n-button>
                <n-button @click="testPaymentConnection" style="margin-left: 12px;">
                  测试支付连接
                </n-button>
              </n-form-item>
            </n-form>
          </div>
        </n-tab-pane>

        <!-- 通知设置 -->
        <n-tab-pane name="notification" tab="通知设置">
          <div class="tab-content">
            <n-form
              ref="notificationFormRef"
              :model="notificationForm"
              label-placement="left"
              label-width="auto"
              require-mark-placement="right-hanging"
              class="settings-form"
            >
              <n-form-item label="启用邮件通知" path="enableEmailNotification">
                <n-switch v-model:value="notificationForm.enableEmailNotification" />
              </n-form-item>
              <n-form-item label="启用短信通知" path="enableSmsNotification">
                <n-switch v-model:value="notificationForm.enableSmsNotification" />
              </n-form-item>
              <n-form-item label="启用站内通知" path="enableInAppNotification">
                <n-switch v-model:value="notificationForm.enableInAppNotification" />
              </n-form-item>
              <n-divider>用户相关通知</n-divider>
              <n-form-item label="新用户注册通知" path="notifyNewUserRegistration">
                <n-switch v-model:value="notificationForm.notifyNewUserRegistration" />
              </n-form-item>
              <n-form-item label="用户登录异常通知" path="notifyLoginAbnormal">
                <n-switch v-model:value="notificationForm.notifyLoginAbnormal" />
              </n-form-item>
              <n-form-item label="密码修改通知" path="notifyPasswordChange">
                <n-switch v-model:value="notificationForm.notifyPasswordChange" />
              </n-form-item>
              <n-divider>订单相关通知</n-divider>
              <n-form-item label="新订单通知" path="notifyNewOrder">
                <n-switch v-model:value="notificationForm.notifyNewOrder" />
              </n-form-item>
              <n-form-item label="订单支付成功通知" path="notifyOrderPaid">
                <n-switch v-model:value="notificationForm.notifyOrderPaid" />
              </n-form-item>
              <n-form-item label="订单发货通知" path="notifyOrderShipped">
                <n-switch v-model:value="notificationForm.notifyOrderShipped" />
              </n-form-item>
              <n-form-item label="订单完成通知" path="notifyOrderCompleted">
                <n-switch v-model:value="notificationForm.notifyOrderCompleted" />
              </n-form-item>
              <n-divider>系统相关通知</n-divider>
              <n-form-item label="系统异常通知" path="notifySystemError">
                <n-switch v-model:value="notificationForm.notifySystemError" />
              </n-form-item>
              <n-form-item label="系统维护通知" path="notifyMaintenance">
                <n-switch v-model:value="notificationForm.notifyMaintenance" />
              </n-form-item>
              <n-form-item label="通知接收邮箱" path="notificationEmail">
                <n-input v-model:value="notificationForm.notificationEmail" placeholder="管理员接收通知的邮箱" />
              </n-form-item>
              <n-form-item>
                <n-button type="primary" @click="saveNotificationSettings" :loading="notificationLoading">
                  保存通知设置
                </n-button>
                <n-button @click="testNotification" style="margin-left: 12px;">
                  测试通知
                </n-button>
              </n-form-item>
            </n-form>
          </div>
        </n-tab-pane>

        <!-- SEO设置 -->
        <n-tab-pane name="seo" tab="SEO设置">
          <div class="tab-content">
            <n-form
              ref="seoFormRef"
              :model="seoForm"
              label-placement="left"
              label-width="auto"
              require-mark-placement="right-hanging"
              class="settings-form"
            >
              <n-form-item label="启用SEO优化" path="enableSEO">
                <n-switch v-model:value="seoForm.enableSEO" />
              </n-form-item>
              <n-form-item label="网站标题后缀" path="siteTitleSuffix">
                <n-input v-model:value="seoForm.siteTitleSuffix" placeholder="例如: - 我的网站" />
              </n-form-item>
              <n-form-item label="网站关键词" path="siteKeywords">
                <n-input v-model:value="seoForm.siteKeywords" placeholder="请输入网站关键词，用逗号分隔" />
              </n-form-item>
              <n-form-item label="网站描述" path="siteDescription">
                <n-input
                  v-model:value="seoForm.siteDescription"
                  type="textarea"
                  placeholder="请输入网站描述"
                  :rows="3"
                />
              </n-form-item>
              <n-form-item label="作者信息" path="siteAuthor">
                <n-input v-model:value="seoForm.siteAuthor" placeholder="请输入网站作者" />
              </n-form-item>
              <n-form-item label="版权信息" path="siteCopyright">
                <n-input v-model:value="seoForm.siteCopyright" placeholder="请输入版权信息" />
              </n-form-item>
              <n-form-item label="ICP备案号" path="siteIcp">
                <n-input v-model:value="seoForm.siteIcp" placeholder="请输入ICP备案号" />
              </n-form-item>
              <n-form-item label="公安备案号" path="sitePoliceIcp">
                <n-input v-model:value="seoForm.sitePoliceIcp" placeholder="请输入公安备案号" />
              </n-form-item>
              <n-form-item label="统计代码" path="analyticsCode">
                <n-input
                  v-model:value="seoForm.analyticsCode"
                  type="textarea"
                  placeholder="请输入百度统计、Google Analytics等统计代码"
                  :rows="4"
                />
              </n-form-item>
              <n-form-item label=" robots.txt 内容" path="robotsContent">
                <n-input
                  v-model:value="seoForm.robotsContent"
                  type="textarea"
                  placeholder="请输入robots.txt内容"
                  :rows="6"
                />
              </n-form-item>
              <n-form-item label="Sitemap自动生成" path="autoGenerateSitemap">
                <n-switch v-model:value="seoForm.autoGenerateSitemap" />
              </n-form-item>
              <n-form-item label="URL伪静态" path="urlRewrite">
                <n-switch v-model:value="seoForm.urlRewrite" />
              </n-form-item>
              <n-form-item label="404页面自定义内容" path="custom404Content">
                <n-input
                  v-model:value="seoForm.custom404Content"
                  type="textarea"
                  placeholder="请输入自定义404页面内容"
                  :rows="4"
                />
              </n-form-item>
              <n-form-item>
                <n-button type="primary" @click="saveSEOSettings" :loading="seoLoading">
                  保存SEO设置
                </n-button>
                <n-button @click="generateSitemap" style="margin-left: 12px;">
                  生成Sitemap
                </n-button>
              </n-form-item>
            </n-form>
          </div>
        </n-tab-pane>

        <!-- 备份设置 -->
        <n-tab-pane name="backup" tab="备份设置">
          <div class="tab-content">
            <n-form
              ref="backupFormRef"
              :model="backupForm"
              label-placement="left"
              label-width="auto"
              require-mark-placement="right-hanging"
              class="settings-form"
            >
              <n-form-item label="启用自动备份" path="enableAutoBackup">
                <n-switch v-model:value="backupForm.enableAutoBackup" />
              </n-form-item>
              <n-form-item label="备份频率" path="backupFrequency">
                <n-select v-model:value="backupForm.backupFrequency" :options="backupFrequencyOptions" />
              </n-form-item>
              <n-form-item label="备份保留天数" path="backupRetentionDays">
                <n-input-number v-model:value="backupForm.backupRetentionDays" :min="1" :max="365">
                  <template #suffix>天</template>
                </n-input-number>
              </n-form-item>
              <n-form-item label="备份存储位置" path="backupStorageLocation">
                <n-select v-model:value="backupForm.backupStorageLocation" :options="backupStorageOptions" />
              </n-form-item>
              <n-divider>备份内容</n-divider>
              <n-form-item label="备份数据库" path="backupDatabase">
                <n-switch v-model:value="backupForm.backupDatabase" />
              </n-form-item>
              <n-form-item label="备份上传文件" path="backupUploads">
                <n-switch v-model:value="backupForm.backupUploads" />
              </n-form-item>
              <n-form-item label="备份配置文件" path="backupConfig">
                <n-switch v-model:value="backupForm.backupConfig" />
              </n-form-item>
              <n-divider>远程存储</n-divider>
              <n-form-item label="启用远程备份" path="enableRemoteBackup">
                <n-switch v-model:value="backupForm.enableRemoteBackup" />
              </n-form-item>
              <n-form-item label="远程存储类型" path="remoteStorageType">
                <n-select v-model:value="backupForm.remoteStorageType" :options="remoteStorageOptions" />
              </n-form-item>
              <n-form-item label="远程存储配置" path="remoteStorageConfig">
                <n-input
                  v-model:value="backupForm.remoteStorageConfig"
                  type="textarea"
                  placeholder="请输入远程存储配置(JSON格式)"
                  :rows="4"
                />
              </n-form-item>
              <n-form-item>
                <n-button type="primary" @click="saveBackupSettings" :loading="backupLoading">
                  保存备份设置
                </n-button>
                <n-button @click="manualBackup" style="margin-left: 12px;">
                  立即备份
                </n-button>
                <n-button @click="viewBackupHistory" style="margin-left: 12px;">
                  查看备份历史
                </n-button>
              </n-form-item>
            </n-form>
          </div>
        </n-tab-pane>

        <!-- 日志设置 -->
        <n-tab-pane name="logs" tab="日志设置">
          <div class="tab-content">
            <n-form
              ref="logsFormRef"
              :model="logsForm"
              label-placement="left"
              label-width="auto"
              require-mark-placement="right-hanging"
              class="settings-form"
            >
              <n-form-item label="启用系统日志" path="enableSystemLogs">
                <n-switch v-model:value="logsForm.enableSystemLogs" />
              </n-form-item>
              <n-form-item label="日志级别" path="logLevel">
                <n-select v-model:value="logsForm.logLevel" :options="logLevelOptions" />
              </n-form-item>
              <n-form-item label="日志保留天数" path="logRetentionDays">
                <n-input-number v-model:value="logsForm.logRetentionDays" :min="1" :max="365">
                  <template #suffix>天</template>
                </n-input-number>
              </n-form-item>
              <n-form-item label="日志文件最大大小" path="maxLogFileSize">
                <n-input-number v-model:value="logsForm.maxLogFileSize" :min="1" :max="1000">
                  <template #suffix>MB</template>
                </n-input-number>
              </n-form-item>
              <n-divider>日志类型</n-divider>
              <n-form-item label="记录用户操作日志" path="logUserActions">
                <n-switch v-model:value="logsForm.logUserActions" />
              </n-form-item>
              <n-form-item label="记录系统错误日志" path="logSystemErrors">
                <n-switch v-model:value="logsForm.logSystemErrors" />
              </n-form-item>
              <n-form-item label="记录数据库查询日志" path="logDatabaseQueries">
                <n-switch v-model:value="logsForm.logDatabaseQueries" />
              </n-form-item>
              <n-form-item label="记录API调用日志" path="logApiCalls">
                <n-switch v-model:value="logsForm.logApiCalls" />
              </n-form-item>
              <n-form-item label="记录支付相关日志" path="logPaymentActions">
                <n-switch v-model:value="logsForm.logPaymentActions" />
              </n-form-item>
              <n-divider>日志存储</n-divider>
              <n-form-item label="日志存储方式" path="logStorageType">
                <n-select v-model:value="logsForm.logStorageType" :options="logStorageOptions" />
              </n-form-item>
              <n-form-item label="外部日志服务" path="externalLogService">
                <n-select v-model:value="logsForm.externalLogService" :options="externalLogServiceOptions" />
              </n-form-item>
              <n-form-item label="外部日志配置" path="externalLogConfig">
                <n-input
                  v-model:value="logsForm.externalLogConfig"
                  type="textarea"
                  placeholder="请输入外部日志服务配置(JSON格式)"
                  :rows="4"
                />
              </n-form-item>
              <n-form-item>
                <n-button type="primary" @click="saveLogsSettings" :loading="logsLoading">
                  保存日志设置
                </n-button>
                <n-button @click="viewCurrentLogs" style="margin-left: 12px;">
                  查看当前日志
                </n-button>
                <n-button @click="exportLogs" style="margin-left: 12px;">
                  导出日志
                </n-button>
              </n-form-item>
            </n-form>
          </div>
        </n-tab-pane>
      </n-tabs>
    </n-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { useMessage } from 'naive-ui'
import type { FormInst } from 'naive-ui'

const message = useMessage()

// 表单引用
const basicFormRef = ref<FormInst | null>(null)
const aiFormRef = ref<FormInst | null>(null)
const emailFormRef = ref<FormInst | null>(null)
const siteFormRef = ref<FormInst | null>(null)
const securityFormRef = ref<FormInst | null>(null)
const paymentFormRef = ref<FormInst | null>(null)
const notificationFormRef = ref<FormInst | null>(null)
const seoFormRef = ref<FormInst | null>(null)
const backupFormRef = ref<FormInst | null>(null)
const logsFormRef = ref<FormInst | null>(null)

// 加载状态
const basicLoading = ref(false)
const aiLoading = ref(false)
const emailLoading = ref(false)
const siteLoading = ref(false)
const securityLoading = ref(false)
const paymentLoading = ref(false)
const notificationLoading = ref(false)
const seoLoading = ref(false)
const backupLoading = ref(false)
const logsLoading = ref(false)

// 基本设置表单
const basicForm = reactive({
  siteName: '',
  siteDescription: '',
  siteLogo: '',
  siteFavicon: '',
  keywords: '',
  author: '',
  language: 'zh-CN',
  timezone: 'Asia/Shanghai',
  dateFormat: 'YYYY-MM-DD',
  timeFormat: 'HH:mm:ss'
})

// AI设置表单
const aiForm = reactive({
  openaiApiKey: '',
  openaiApiUrl: 'https://api.openai.com/v1',
  aiModel: 'gpt-3.5-turbo',
  aiTemperature: 0.7,
  aiMaxTokens: 2048,
  aiTopP: 1.0,
  aiFrequencyPenalty: 0,
  aiPresencePenalty: 0
})

// 邮件设置表单
const emailForm = reactive({
  smtpHost: '',
  smtpPort: 587,
  smtpUsername: '',
  smtpPassword: '',
  smtpEncryption: 'tls',
  senderEmail: '',
  senderName: ''
})

// 站点设置表单
const siteForm = reactive({
  maintenanceMode: false,
  allowRegistration: true,
  emailVerification: true,
  commentModeration: true,
  uploadLimit: 10,
  allowedFileTypes: ['jpg', 'jpeg', 'png', 'gif', 'pdf', 'doc', 'docx'],
  enableCache: true,
  cacheTime: 60,
  enableSSL: false,
  enableCDN: false,
  cdnUrl: ''
})

// 安全设置表单
const securityForm = reactive({
  enableTwoFactor: false,
  maxLoginAttempts: 5,
  lockoutDuration: 15,
  sessionTimeout: 120,
  minPasswordLength: 8,
  passwordComplexity: ['uppercase', 'lowercase', 'number'],
  apiRateLimit: 100,
  ipWhitelist: [] as string[],
  ipBlacklist: [] as string[]
})

// 支付设置表单
const paymentForm = reactive({
  enableAlipay: false,
  alipayAppId: '',
  alipayPrivateKey: '',
  alipayPublicKey: '',
  enableWechatPay: false,
  wechatMchId: '',
  wechatApiKey: '',
  wechatAppId: '',
  paymentNotifyUrl: '',
  orderTimeout: 30,
  refundReview: true
})

// 通知设置表单
const notificationForm = reactive({
  enableEmailNotification: true,
  enableSmsNotification: false,
  enableInAppNotification: true,
  notifyNewUserRegistration: true,
  notifyLoginAbnormal: true,
  notifyPasswordChange: true,
  notifyNewOrder: true,
  notifyOrderPaid: true,
  notifyOrderShipped: true,
  notifyOrderCompleted: true,
  notifySystemError: true,
  notifyMaintenance: true,
  notificationEmail: ''
})

// SEO设置表单
const seoForm = reactive({
  enableSEO: true,
  siteTitleSuffix: '',
  siteKeywords: '',
  siteDescription: '',
  siteAuthor: '',
  siteCopyright: '',
  siteIcp: '',
  sitePoliceIcp: '',
  analyticsCode: '',
  robotsContent: '',
  autoGenerateSitemap: true,
  urlRewrite: true,
  custom404Content: ''
})

// 备份设置表单
const backupForm = reactive({
  enableAutoBackup: true,
  backupFrequency: 'daily',
  backupRetentionDays: 30,
  backupStorageLocation: 'local',
  backupDatabase: true,
  backupUploads: true,
  backupConfig: true,
  enableRemoteBackup: false,
  remoteStorageType: 's3',
  remoteStorageConfig: ''
})

// 日志设置表单
const logsForm = reactive({
  enableSystemLogs: true,
  logLevel: 'info',
  logRetentionDays: 30,
  maxLogFileSize: 100,
  logUserActions: true,
  logSystemErrors: true,
  logDatabaseQueries: false,
  logApiCalls: true,
  logPaymentActions: true,
  logStorageType: 'file',
  externalLogService: 'none',
  externalLogConfig: ''
})

// 选项数据
const languageOptions = [
  { label: '简体中文', value: 'zh-CN' },
  { label: '繁體中文', value: 'zh-TW' },
  { label: 'English', value: 'en-US' },
  { label: '日本語', value: 'ja-JP' }
]

const timezoneOptions = [
  { label: '北京时间 (UTC+8)', value: 'Asia/Shanghai' },
  { label: '东京时间 (UTC+9)', value: 'Asia/Tokyo' },
  { label: '纽约时间 (UTC-5)', value: 'America/New_York' },
  { label: '伦敦时间 (UTC+0)', value: 'Europe/London' }
]

const aiModelOptions = [
  { label: 'GPT-3.5 Turbo', value: 'gpt-3.5-turbo' },
  { label: 'GPT-4', value: 'gpt-4' },
  { label: 'GPT-4 Turbo', value: 'gpt-4-turbo' },
  { label: 'GPT-4o', value: 'gpt-4o' }
]

const encryptionOptions = [
  { label: 'TLS', value: 'tls' },
  { label: 'SSL', value: 'ssl' },
  { label: '无', value: 'none' }
]

const backupFrequencyOptions = [
  { label: '每小时', value: 'hourly' },
  { label: '每天', value: 'daily' },
  { label: '每周', value: 'weekly' },
  { label: '每月', value: 'monthly' }
]

const backupStorageOptions = [
  { label: '本地存储', value: 'local' },
  { label: '云存储', value: 'cloud' },
  { label: '混合存储', value: 'hybrid' }
]

const remoteStorageOptions = [
  { label: 'Amazon S3', value: 's3' },
  { label: '阿里云OSS', value: 'aliyun' },
  { label: '腾讯云COS', value: 'tencent' },
  { label: '七牛云', value: 'qiniu' }
]

const logLevelOptions = [
  { label: '调试', value: 'debug' },
  { label: '信息', value: 'info' },
  { label: '警告', value: 'warning' },
  { label: '错误', value: 'error' },
  { label: '致命', value: 'fatal' }
]

const logStorageOptions = [
  { label: '文件存储', value: 'file' },
  { label: '数据库存储', value: 'database' },
  { label: '混合存储', value: 'hybrid' }
]

const externalLogServiceOptions = [
  { label: '不使用', value: 'none' },
  { label: 'ELK Stack', value: 'elk' },
  { label: 'Graylog', value: 'graylog' },
  { label: 'Splunk', value: 'splunk' },
  { label: 'Datadog', value: 'datadog' }
]

// 保存基本设置
const saveBasicSettings = async () => {
  basicLoading.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('基本设置保存成功')
  } catch {
    message.error('基本设置保存失败')
  } finally {
    basicLoading.value = false
  }
}

// 保存AI设置
const saveAISettings = async () => {
  aiLoading.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('AI设置保存成功')
  } catch {
    message.error('AI设置保存失败')
  } finally {
    aiLoading.value = false
  }
}

// 测试AI连接
const testAIConnection = async () => {
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('AI连接测试成功')
  } catch {
    message.error('AI连接测试失败')
  }
}

// 保存邮件设置
const saveEmailSettings = async () => {
  emailLoading.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('邮件设置保存成功')
  } catch {
    message.error('邮件设置保存失败')
  } finally {
    emailLoading.value = false
  }
}

// 测试邮件连接
const testEmailConnection = async () => {
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('邮件连接测试成功')
  } catch {
    message.error('邮件连接测试失败')
  }
}

// 保存站点设置
const saveSiteSettings = async () => {
  siteLoading.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('站点设置保存成功')
  } catch {
    message.error('站点设置保存失败')
  } finally {
    siteLoading.value = false
  }
}

// 保存安全设置
const saveSecuritySettings = async () => {
  securityLoading.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('安全设置保存成功')
  } catch {
    message.error('安全设置保存失败')
  } finally {
    securityLoading.value = false
  }
}

// 保存支付设置
const savePaymentSettings = async () => {
  paymentLoading.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('支付设置保存成功')
  } catch {
    message.error('支付设置保存失败')
  } finally {
    paymentLoading.value = false
  }
}

// 测试支付连接
const testPaymentConnection = async () => {
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('支付连接测试成功')
  } catch {
    message.error('支付连接测试失败')
  }
}

// 保存通知设置
const saveNotificationSettings = async () => {
  notificationLoading.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('通知设置保存成功')
  } catch {
    message.error('通知设置保存失败')
  } finally {
    notificationLoading.value = false
  }
}

// 测试通知
const testNotification = async () => {
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('通知测试成功')
  } catch {
    message.error('通知测试失败')
  }
}

// 保存SEO设置
const saveSEOSettings = async () => {
  seoLoading.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('SEO设置保存成功')
  } catch {
    message.error('SEO设置保存失败')
  } finally {
    seoLoading.value = false
  }
}

// 生成Sitemap
const generateSitemap = async () => {
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('Sitemap生成成功')
  } catch {
    message.error('Sitemap生成失败')
  }
}

// 保存备份设置
const saveBackupSettings = async () => {
  backupLoading.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('备份设置保存成功')
  } catch {
    message.error('备份设置保存失败')
  } finally {
    backupLoading.value = false
  }
}

// 手动备份
const manualBackup = async () => {
  try {
    await new Promise(resolve => setTimeout(resolve, 2000))
    message.success('手动备份成功')
  } catch {
    message.error('手动备份失败')
  }
}

// 查看备份历史
const viewBackupHistory = async () => {
  message.info('备份历史功能开发中...')
}

// 保存日志设置
const saveLogsSettings = async () => {
  logsLoading.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('日志设置保存成功')
  } catch {
    message.error('日志设置保存失败')
  } finally {
    logsLoading.value = false
  }
}

// 查看当前日志
const viewCurrentLogs = async () => {
  message.info('当前日志功能开发中...')
}

// 导出日志
const exportLogs = async () => {
  try {
    // 模拟导出日志
    const logData = '系统日志数据...'
    const blob = new Blob([logData], { type: 'text/plain' })
    const url = URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.href = url
    link.download = `system-logs-${new Date().toISOString().split('T')[0]}.txt`
    link.click()
    URL.revokeObjectURL(url)
    message.success('日志导出成功')
  } catch {
    message.error('日志导出失败')
  }
}

// 加载设置数据
const loadSettings = async () => {
  try {
    await new Promise(resolve => setTimeout(resolve, 500))
  } catch {
    message.error('设置数据加载失败')
  }
}

onMounted(() => {
  loadSettings()
})
</script>

<style scoped>
.settings-container {
  padding: 24px;
  min-height: 100vh;
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
