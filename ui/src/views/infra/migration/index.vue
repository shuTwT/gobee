<script setup lang="ts">
import { ref, computed } from 'vue'
import {
  NSteps,
  NStep,
  NCard,
  NRadioGroup,
  NRadio,
  NSpace,
  NUpload,
  NInput,
  NButton,
  NAlert,
  NProgress,
  NStatistic,
  NIcon,
} from 'naive-ui'
import { CloudUploadOutline, Link, Folder, DocumentText, CheckmarkCircleOutline, AlertCircleOutline } from '@vicons/ionicons5'
import type { UploadFileInfo } from 'naive-ui'

// 步骤状态
const currentStep = ref(1)
const totalSteps = 3

// 来源类型 - 扩展了详细描述
const sourceType = ref<string>('')
const sourceTypes = [
  { 
    value: 'rss', 
    label: 'RSS Feed',
    description: '从RSS订阅源导入内容，支持文件上传或URL订阅',
    icon: Link
  },
  { 
    value: 'atom', 
    label: 'Atom Feed',
    description: '从Atom订阅源导入内容，支持文件上传或URL订阅',
    icon: Link
  },
  { 
    value: 'md', 
    label: '本地MD文件',
    description: '从本地Markdown文件导入内容，支持单文件或文件夹批量导入',
    icon: DocumentText
  },
]

// RSS/Atom 选项
const rssImportType = ref<string>('')
const atomImportType = ref<string>('')

// 本地MD文件选项
const mdImportType = ref<string>('')

// 文件列表
const rssFileList = ref<UploadFileInfo[]>([])
const atomFileList = ref<UploadFileInfo[]>([])
const mdFileList = ref<UploadFileInfo[]>([])

// 链接输入
const rssUrl = ref('')
const atomUrl = ref('')

// 导入状态
const importProgress = ref(0)
const isImporting = ref(false)
const hasConfirmedImport = ref(false)
const migrationResult = ref<{
  status: 'success' | 'failed' | null
  message: string
  stats: {
    total: number
    success: number
    failed: number
  }
  errors?: string[]
}>({
  status: null,
  message: '',
  stats: {
    total: 0,
    success: 0,
    failed: 0
  },
  errors: []
})

// 待迁移数据预览
const previewData = computed(() => {
  if (sourceType.value === 'rss') {
    if (rssImportType.value === 'file') {
      return {
        type: 'file',
        count: rssFileList.value.length,
        files: rssFileList.value.map(f => f.name)
      }
    } else {
      return {
        type: 'url',
        url: rssUrl.value
      }
    }
  } else if (sourceType.value === 'atom') {
    if (atomImportType.value === 'file') {
      return {
        type: 'file',
        count: atomFileList.value.length,
        files: atomFileList.value.map(f => f.name)
      }
    } else {
      return {
        type: 'url',
        url: atomUrl.value
      }
    }
  } else if (sourceType.value === 'md') {
    return {
      type: 'file',
      count: mdFileList.value.length,
      files: mdFileList.value.map(f => f.name)
    }
  }
  return null
})

// 验证步骤二是否可以继续
const isValidStep2 = computed(() => {
  if (sourceType.value === 'rss') {
    if (rssImportType.value === 'file') {
      return rssFileList.value.length > 0
    } else if (rssImportType.value === 'url') {
      return rssUrl.value.trim() !== ''
    }
    return false
  } else if (sourceType.value === 'atom') {
    if (atomImportType.value === 'file') {
      return atomFileList.value.length > 0
    } else if (atomImportType.value === 'url') {
      return atomUrl.value.trim() !== ''
    }
    return false
  } else if (sourceType.value === 'md') {
    if (mdImportType.value === 'folder' || mdImportType.value === 'file') {
      return mdFileList.value.length > 0
    }
    return false
  }
  return false
})

// 验证步骤三是否可以确认迁移
const isValidStep3 = computed(() => {
  return hasConfirmedImport.value === false && previewData.value !== null
})

// 计算步骤标题
const stepTitles = computed(() => {
  return [
    '选择来源',
    '配置导入方式',
    '预览与导入'
  ]
})

// 计算当前步骤的标题
const currentStepTitle = computed(() => {
  return stepTitles.value[currentStep.value - 1]
})

// 下一步
const nextStep = () => {
  if (currentStep.value < totalSteps) {
    currentStep.value++
  }
}

// 上一步
const prevStep = () => {
  if (currentStep.value > 1) {
    currentStep.value--
  }
}

// 重置步骤
const resetSteps = () => {
  currentStep.value = 1
  sourceType.value = ''
  rssImportType.value = ''
  atomImportType.value = ''
  mdImportType.value = ''
  rssFileList.value = []
  atomFileList.value = []
  mdFileList.value = []
  rssUrl.value = ''
  atomUrl.value = ''
  importProgress.value = 0
  isImporting.value = false
  hasConfirmedImport.value = false
  migrationResult.value = {
    status: null,
    message: '',
    stats: {
      total: 0,
      success: 0,
      failed: 0
    },
    errors: []
  }
}

// 开始导入 - 步骤二到步骤三
const startImport = () => {
  nextStep()
}

// 确认迁移 - 步骤三执行迁移
const confirmImport = () => {
  hasConfirmedImport.value = true
  isImporting.value = true
  importProgress.value = 0
  
  // 模拟导入进度
  const interval = setInterval(() => {
    importProgress.value += 10
    if (importProgress.value >= 100) {
      clearInterval(interval)
      isImporting.value = false
      
      // 模拟迁移结果
      setTimeout(() => {
        migrationResult.value = {
          status: 'success',
          message: '数据迁移完成',
          stats: {
            total: 25,
            success: 23,
            failed: 2
          },
          errors: [
            'file1.md: 解析错误',
            'file2.md: 缺少必要字段'
          ]
        }
      }, 500)
    }
  }, 300)
}

// 文件上传变化
const handleFileChange = (fileList: UploadFileInfo[], type: string) => {
  switch (type) {
    case 'rss':
      rssFileList.value = fileList
      break
    case 'atom':
      atomFileList.value = fileList
      break
    case 'md':
      mdFileList.value = fileList
      break
  }
}

// 计算当前文件数量
const getFileCount = computed(() => {
  if (sourceType.value === 'rss') {
    return rssFileList.value.length
  } else if (sourceType.value === 'atom') {
    return atomFileList.value.length
  } else if (sourceType.value === 'md') {
    return mdFileList.value.length
  }
  return 0
})

// 计算当前导入方式文本
const getImportMethodText = computed(() => {
  if (sourceType.value === 'rss') {
    return rssImportType.value === 'file' ? '上传RSS文件' : '输入RSS链接'
  } else if (sourceType.value === 'atom') {
    return atomImportType.value === 'file' ? '上传Atom文件' : '输入Atom链接'
  } else if (sourceType.value === 'md') {
    return mdImportType.value === 'folder' ? '文件夹批量导入' : '单个MD文件导入'
  }
  return ''
})
</script>

<template>
  <div class="container-fluid p-6">
    <n-card class="max-w-4xl mx-auto" title="数据迁移">
      <!-- 步骤导航 -->
      <n-steps
        :current="currentStep"
        :status="currentStep === 3 && hasConfirmedImport ? 'finish' : 'process'"
        class="mb-8"
      >
        <n-step
          v-for="(title, index) in stepTitles"
          :key="index"
          :title="title"
        />
      </n-steps>

      <!-- 步骤内容 -->
      <div class="step-content">
        <!-- 步骤一：选择来源 -->
        <div v-if="currentStep === 1" class="step-one">
          <n-alert title="选择数据来源" type="info" class="mb-6">
            请选择您要导入的数据来源类型
          </n-alert>
          
          <n-space vertical size="large" class="w-full">
            <div class="selection-section">
              <!-- 卡片式选择组件 -->
              <n-space wrap>
                <n-card
                  v-for="item in sourceTypes"
                  :key="item.value"
                  class="source-card"
                  :bordered="sourceType === item.value"
                  :style="{
                    cursor: 'pointer',
                    flex: '1 1 calc(33.333% - 16px)',
                    border: sourceType === item.value ? '1px solid var(--n-color-target)' : '1px solid var(--n-border-color)',
                    boxShadow: sourceType === item.value ? '0 4px 12px rgba(0, 0, 0, 0.15)' : 'none'
                  }"
                  @click="sourceType = item.value"
                >
                  <div class="source-card-header">
                    <n-icon size="24" class="source-icon" :component="item.icon" />
                    <h3 class="source-title">{{ item.label }}</h3>
                  </div>
                  <div class="source-description">{{ item.description }}</div>
                </n-card>
              </n-space>
            </div>
            
            <n-space justify="end">
              <n-button
                type="primary"
                @click="nextStep"
                :disabled="!sourceType"
                size="large"
              >
                下一步
              </n-button>
            </n-space>
          </n-space>
        </div>

        <!-- 步骤二：配置导入方式 -->
        <div v-if="currentStep === 2" class="step-two">
          <n-alert 
            :title="`${sourceTypes.find(t => t.value === sourceType)?.label} 导入配置`" 
            type="info" 
            class="mb-6"
          >
            请配置导入方式
          </n-alert>
          
          <!-- RSS Feed 配置 -->
          <div v-if="sourceType === 'rss'" class="source-config">
            <n-radio-group v-model:value="rssImportType" class="mb-6">
              <n-radio value="file">上传RSS文件</n-radio>
              <n-radio value="url">输入RSS订阅链接</n-radio>
            </n-radio-group>
            
            <!-- 上传RSS文件 -->
            <div v-if="rssImportType === 'file'" class="import-option">
              <n-upload
                v-model:file-list="rssFileList"
                multiple
                list-type="text"
                @change="(e) => handleFileChange(e.fileList, 'rss')"
              >
                <n-button size="large">
                  <template #icon>
                    <CloudUploadOutline />
                  </template>
                  选择RSS文件
                </n-button>
              </n-upload>
            </div>
            
            <!-- 输入RSS链接 -->
            <div v-else-if="rssImportType === 'url'" class="import-option">
              <n-input
                v-model:value="rssUrl"
                placeholder="请输入RSS订阅链接"
                size="large"
                class="w-full"
              >
            <template #prefix>
              <n-icon size="24" class="source-icon" :component="Link" />
            </template>
            </n-input>
            </div>
          </div>
          
          <!-- Atom Feed 配置 -->
          <div v-else-if="sourceType === 'atom'" class="source-config">
            <n-radio-group v-model:value="atomImportType" class="mb-6">
              <n-radio value="file">上传Atom文件</n-radio>
              <n-radio value="url">输入Atom订阅链接</n-radio>
            </n-radio-group>
            
            <!-- 上传Atom文件 -->
            <div v-if="atomImportType === 'file'" class="import-option">
              <n-upload
                v-model:file-list="atomFileList"
                multiple
                list-type="text"
                @change="(e) => handleFileChange(e.fileList, 'atom')"
              >
                <n-button size="large">
                  <template #icon>
                    <CloudUploadOutline />
                  </template>
                  选择Atom文件
                </n-button>
              </n-upload>
            </div>
            
            <!-- 输入Atom链接 -->
            <div v-else-if="atomImportType === 'url'" class="import-option">
              <n-input
                v-model:value="atomUrl"
                placeholder="请输入Atom订阅链接"
                size="large"
                class="w-full"
              >
            <template #prefix>
              <n-icon size="24" class="source-icon" :component="Link" />
            </template>
            </n-input>
            </div>
          </div>
          
          <!-- 本地MD文件配置 -->
          <div v-else-if="sourceType === 'md'" class="source-config">
            <n-radio-group v-model:value="mdImportType" class="mb-6">
              <n-radio value="folder">选择文件夹批量导入</n-radio>
              <n-radio value="file">选择单个MD文件导入</n-radio>
            </n-radio-group>
            
            <!-- 文件夹导入 -->
            <div v-if="mdImportType === 'folder'" class="import-option">
              <n-upload
                v-model:file-list="mdFileList"
                multiple
                list-type="text"
                directory
                @change="(e) => handleFileChange(e.fileList, 'md')"
              >
                <n-button size="large">
                  <template #icon>
                    <Folder />
                  </template>
                  选择文件夹
                </n-button>
              </n-upload>
            </div>
            
            <!-- 单文件导入 -->
            <div v-else-if="mdImportType === 'file'" class="import-option">
              <n-upload
                v-model:file-list="mdFileList"
                multiple
                list-type="text"
                :accept="'.md,.markdown'"
                @change="(e) => handleFileChange(e.fileList, 'md')"
              >
                <n-button size="large">
                  <template #icon>
                    <DocumentText />
                  </template>
                  选择MD文件
                </n-button>
              </n-upload>
            </div>
          </div>
          
          <n-space justify="space-between" class="mt-8">
            <n-button @click="prevStep" size="large">
              上一步
            </n-button>
            <n-button
              type="primary"
              @click="startImport"
              :disabled="!isValidStep2"
              size="large"
            >
              开始导入
            </n-button>
          </n-space>
        </div>

        <!-- 步骤三：预览与导入 -->
        <div v-if="currentStep === 3" class="step-three">
          <n-alert 
            v-if="!hasConfirmedImport" 
            title="数据预览" 
            type="info" 
            class="mb-6"
          >
            请确认您要导入的数据信息
          </n-alert>
          
          <!-- 数据预览区域 -->
          <div v-if="!hasConfirmedImport" class="preview-section">
            <n-card class="preview-card" title="待迁移数据信息" size="small">
              <div class="preview-info">
                <div class="preview-item">
                  <span class="preview-label">导入来源：</span>
                  <span class="preview-value">{{ sourceTypes.find(t => t.value === sourceType)?.label }}</span>
                </div>
                <div class="preview-item">
                  <span class="preview-label">导入方式：</span>
                  <span class="preview-value">{{ getImportMethodText }}</span>
                </div>
                <div class="preview-item">
                  <span class="preview-label">数据数量：</span>
                  <span class="preview-value">{{ getFileCount }} 项</span>
                </div>
                
                <!-- 文件列表预览 -->
                <div v-if="previewData?.type === 'file' && (previewData.count || 0) > 0" class="file-list-preview">
                  <h4 class="preview-subtitle">文件列表：</h4>
                  <div class="file-list">
                    <div v-for="(file, index) in previewData.files?.slice(0, 5)" :key="index" class="file-item">
                      <n-icon :component="DocumentText" size="16" />
                      <span class="file-name">{{ file }}</span>
                    </div>
                    <div v-if="(previewData.files?.length || 0) > 5" class="file-more">
                      ... 还有 {{ (previewData.files?.length || 0) - 5 }} 个文件
                    </div>
                  </div>
                </div>
                
                <!-- URL预览 -->
                <div v-else-if="previewData?.type === 'url'" class="url-preview">
                  <h4 class="preview-subtitle">订阅链接：</h4>
                  <div class="url-item">
                    <n-icon :component="Link" size="16" />
                    <a :href="previewData.url" target="_blank" class="url-link">{{ previewData.url }}</a>
                  </div>
                </div>
              </div>
            </n-card>
            
            <n-space justify="space-between" class="mt-8">
              <n-button @click="prevStep" size="large">
                上一步
              </n-button>
              <n-button
                type="primary"
                @click="confirmImport"
                :disabled="!isValidStep3"
                size="large"
              >
                确认迁移
              </n-button>
            </n-space>
          </div>
          
          <!-- 导入进度 -->
          <div v-if="isImporting" class="import-progress-section">
            <n-card title="导入进度" size="small" class="progress-card">
              <n-progress :percentage="importProgress" type="line" :show-indicator="true" />
              <div class="progress-text">正在导入数据，请稍候...</div>
            </n-card>
          </div>
          
          <!-- 迁移结果 -->
          <div v-if="hasConfirmedImport && !isImporting" class="result-section">
            <n-alert 
              :title="migrationResult.message" 
              :type="migrationResult.status === 'success' ? 'success' : 'error'" 
              class="mb-6"
            >
              <template v-if="migrationResult.status === 'success'">
                <n-icon :component="CheckmarkCircleOutline" />
                <span>数据迁移成功！</span>
              </template>
              <template v-else-if="migrationResult.status === 'failed'">
                <n-icon :component="AlertCircleOutline" />
                <span>数据迁移失败！</span>
              </template>
            </n-alert>
            
            <div class="stats-section">
              <n-grid :cols="3" :x-gap="16" class="stats-grid">
                <n-gi>
                  <n-card class="stat-card">
                    <n-statistic label="总数据量" :value="migrationResult.stats.total" />
                  </n-card>
                </n-gi>
                <n-gi>
                  <n-card class="stat-card success-card">
                    <n-statistic label="成功" :value="migrationResult.stats.success" />
                  </n-card>
                </n-gi>
                <n-gi>
                  <n-card class="stat-card error-card">
                    <n-statistic label="失败" :value="migrationResult.stats.failed" />
                  </n-card>
                </n-gi>
              </n-grid>
            </div>
            
            <!-- 错误信息 -->
            <div v-if="migrationResult.status === 'failed' && migrationResult.errors && migrationResult.errors.length > 0" class="errors-section">
              <n-card title="错误信息" size="small" type="error">
                <ul class="error-list">
                  <li v-for="(error, index) in migrationResult.errors" :key="index" class="error-item">
                    <n-icon :component="AlertCircleOutline" size="16" class="error-icon" />
                    <span>{{ error }}</span>
                  </li>
                </ul>
              </n-card>
            </div>
            
            <n-space justify="space-between" class="mt-8">
              <n-button @click="prevStep" size="large">
                上一步
              </n-button>
              <n-button
                type="primary"
                @click="resetSteps"
                size="large"
              >
                完成
              </n-button>
            </n-space>
          </div>
        </div>
      </div>
    </n-card>
  </div>
</template>

<style scoped>
.step-content {
  padding: 20px 0;
}

/* 步骤一：来源选择卡片样式 */
.selection-section {
  width: 100%;
}

.source-card {
  border-radius: 8px;
  transition: all 0.3s ease;
  cursor: pointer;
  min-height: 120px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding: 20px;
}

.source-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.source-card-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 8px;
}

.source-icon {
  color: var(--primary-color);
}

.source-title {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
}

.source-description {
  font-size: 14px;
  color: #6b7280;
  line-height: 1.5;
}

/* 步骤二：配置导入方式 */
.source-config {
  background-color: #fafafa;
  padding: 24px;
  border-radius: 8px;
}

.import-option {
  background-color: white;
  padding: 20px;
  border-radius: 8px;
  border: 1px solid #e5e7eb;
}

/* 步骤三：预览与导入 */
.preview-section {
  margin-top: 20px;
}

.preview-card {
  margin-bottom: 20px;
}

.preview-info {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.preview-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 0;
  border-bottom: 1px solid #e5e7eb;
}

.preview-item:last-child {
  border-bottom: none;
}

.preview-label {
  font-weight: 600;
  color: #374151;
}

.preview-value {
  color: #6b7280;
  font-weight: 500;
}

.preview-subtitle {
  margin: 16px 0 8px 0;
  font-size: 16px;
  font-weight: 600;
  color: #374151;
}

.file-list-preview,
.url-preview {
  margin-top: 16px;
}

.file-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
  max-height: 150px;
  overflow-y: auto;
}

.file-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 0;
  font-size: 14px;
  color: #6b7280;
}

.file-more {
  font-size: 14px;
  color: #9ca3af;
  padding: 8px 0;
  font-style: italic;
}

.url-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 0;
  font-size: 14px;
}

.url-link {
  color: var(--primary-color);
  text-decoration: underline;
}

/* 导入进度 */
.import-progress-section {
  margin-top: 20px;
}

.progress-card {
  padding: 24px;
}

.progress-text {
  text-align: center;
  margin-top: 16px;
  color: #6b7280;
}

/* 迁移结果 */
.result-section {
  margin-top: 20px;
}

.stats-section {
  margin: 20px 0;
}

.stats-grid {
  gap: 16px;
}

.stat-card {
  text-align: center;
  padding: 24px;
  border-radius: 8px;
  transition: all 0.3s ease;
}

.stat-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.success-card {
  border: 1px solid #d1fae5;
  background-color: #f0fdf4;
}

.error-card {
  border: 1px solid #fee2e2;
  background-color: #fef2f2;
}

/* 错误信息 */
.errors-section {
  margin-top: 20px;
}

.error-list {
  margin: 0;
  padding: 0;
  list-style: none;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.error-item {
  display: flex;
  align-items: flex-start;
  gap: 8px;
  padding: 8px 0;
  color: #dc2626;
  font-size: 14px;
}

.error-icon {
  color: #dc2626;
  margin-top: 2px;
  flex-shrink: 0;
}

/* 通用样式 */
.step-one,
.step-two,
.step-three {
  max-width: 600px;
  margin: 0 auto;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .source-card {
    flex: 1 1 calc(50% - 16px) !important;
  }
  
  .stats-grid {
    grid-template-columns: 1fr !important;
  }
  
  .preview-item {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }
}

@media (max-width: 480px) {
  .source-card {
    flex: 1 1 100% !important;
  }
  
  .source-card-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }
}
</style>
