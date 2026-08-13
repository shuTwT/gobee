<script setup lang="ts">
import { computed } from 'vue'
import { useUserStore } from '@/stores/modules/user'

const router = useRouter()
const userStore = useUserStore()
const dialog = useDialog()
const message = useMessage()

const gotoUserCenter = () => {
  router.push({ name: 'UserCenter' })
}

// 头像显示：有头像URL则用URL，否则用用户名首字母
const avatarText = computed(() => {
  const name = userStore.nickname || userStore.username || 'U'
  return name.charAt(0).toUpperCase()
})

const hasAvatar = computed(() => !!userStore.avatar)

// 角色显示文本
const roleLabel = computed(() => {
  const roles = userStore.roles || []
  if (roles.includes('admin')) return '管理员'
  if (roles.includes('common')) return '普通用户'
  return roles[0] || '访客'
})

function logout() {
  dialog.create({
    type: 'info',
    title: '提示',
    content: '确定要登出吗',
    positiveText: '确定',
    negativeText: '不确定',
    onPositiveClick: () => {
      userStore.logOut().then(() => {
        message.success('登出成功')
      })
    },
    onNegativeClick: () => {},
  })
}
</script>

<template>
  <div class="flex items-center justify-center">
    <n-popover
      trigger="click"
      placement="right-end"
      :show-arrow="false"
      style="padding: 0; background: transparent; border: none"
    >
      <template #trigger>
        <n-avatar
          class="cursor-pointer"
          :src="hasAvatar ? userStore.avatar : undefined"
          :style="!hasAvatar ? { color: '#fff', backgroundColor: '#18a058' } : {}"
        >
          {{ avatarText }}
        </n-avatar>
      </template>

      <!-- Discord 风格用户卡片 -->
      <div class="discord-user-card">
        <!-- 顶部横幅区 -->
        <div class="card-banner"></div>

        <!-- 主体信息区 -->
        <div class="card-body">
          <!-- 头像（底部超出横幅） -->
          <div class="card-avatar-wrapper">
            <n-avatar
              class="card-avatar"
              :src="hasAvatar ? userStore.avatar : undefined"
              :style="!hasAvatar ? { color: '#fff', backgroundColor: '#5865f2' } : {}"
            >
              {{ avatarText }}
            </n-avatar>
            <div class="card-status-dot online"></div>
          </div>

          <!-- 用户信息 -->
          <div class="card-info">
            <div class="card-name" @click="gotoUserCenter">
              {{ userStore.nickname || userStore.username }}
            </div>
            <div class="card-username">{{ userStore.username }}</div>
          </div>

          <!-- 分隔线 -->
          <div class="card-divider"></div>

          <!-- 成员信息 -->
          <div class="card-section">
            <div class="card-section-title">成员信息</div>
            <div class="card-section-row">
              <span class="card-label">角色</span>
              <span class="card-role-badge">{{ roleLabel }}</span>
            </div>
          </div>

          <!-- 操作按钮 -->
          <div class="card-actions">
            <button class="card-action-btn" @click="gotoUserCenter">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
                ></path>
              </svg>
              <span>个人中心</span>
            </button>
            <button class="card-action-btn danger" @click="logout">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
                ></path>
              </svg>
              <span>退出登录</span>
            </button>
          </div>
        </div>
      </div>
    </n-popover>
  </div>
</template>

<style scoped lang="scss">
.discord-user-card {
  width: 300px;
  border-radius: 8px;
  overflow: hidden;
  background: #111214;
  color: #dbdee1;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.24);
}

/* 顶部横幅 */
.card-banner {
  height: 60px;
  background: linear-gradient(135deg, #5865f2 0%, #7289da 100%);
}

/* 主体 */
.card-body {
  position: relative;
  padding: 0 16px 16px;
}

/* 头像 */
.card-avatar-wrapper {
  position: relative;
  margin-top: -34px;
  margin-bottom: 12px;
  width: fit-content;
}

.card-avatar {
  width: 68px;
  height: 68px;
  border: 6px solid #111214;
  border-radius: 50%;
  background: #5865f2;
}

.card-status-dot {
  position: absolute;
  bottom: 6px;
  right: 6px;
  width: 16px;
  height: 16px;
  border-radius: 50%;
  border: 4px solid #111214;
  box-sizing: content-box;

  &.online {
    background: #23a559;
  }

  &.offline {
    background: #80848e;
  }

  &.busy {
    background: #f23f42;
  }
}

/* 用户信息 */
.card-info {
  margin-bottom: 12px;
}

.card-name {
  font-size: 16px;
  font-weight: 600;
  color: #f2f3f5;
  cursor: pointer;
  transition: text-decoration 0.15s;

  &:hover {
    text-decoration: underline;
  }
}

.card-username {
  font-size: 14px;
  color: #949ba4;
  margin-top: 2px;
}

/* 分隔线 */
.card-divider {
  height: 1px;
  background: #232428;
  margin: 12px 0;
}

/* 信息区块 */
.card-section {
  margin-bottom: 12px;
}

.card-section-title {
  font-size: 12px;
  font-weight: 700;
  color: #b5bac1;
  text-transform: uppercase;
  margin-bottom: 8px;
  letter-spacing: 0.02em;
}

.card-section-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 14px;
}

.card-label {
  color: #b5bac1;
}

.card-role-badge {
  display: inline-flex;
  align-items: center;
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 12px;
  font-weight: 500;
  background: rgba(88, 101, 242, 0.15);
  color: #5865f2;
}

/* 操作按钮区 */
.card-actions {
  display: flex;
  flex-direction: column;
  gap: 6px;
  margin-top: 12px;
}

.card-action-btn {
  display: flex;
  align-items: center;
  gap: 10px;
  width: 100%;
  padding: 8px 12px;
  border: none;
  border-radius: 4px;
  background: #2b2d31;
  color: #dbdee1;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.15s, color 0.15s;
  text-align: left;

  &:hover {
    background: #35373c;
  }

  &.danger {
    color: #f23f42;

    &:hover {
      background: rgba(242, 63, 66, 0.1);
    }
  }
}

/* 暗色模式适配（跟随系统，但卡片本身就是暗色风格） */
:global(.dark) {
  /* 已是暗色风格，无需调整 */
}
</style>
