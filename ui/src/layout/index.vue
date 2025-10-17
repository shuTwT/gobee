<script setup lang="ts">
import type { MenuOption } from 'naive-ui'
import { ProLayout, useLayoutMenu } from 'pro-naive-ui'
import { RouterLink } from 'vue-router'
import { useAppStore } from '@/stores/modules/app'
import { storeToRefs } from 'pinia'
import { usePermissionStore } from '@/stores/modules/permission'

const appStore = useAppStore()
const permissionStore = usePermissionStore()
const {wholeMenus} = storeToRefs(permissionStore)

console.log(wholeMenus.value)

const menuOptions: MenuOption[] = [
  {
    label: () => {
      return h(RouterLink,
        {
          to:{
            name:'Home'
          },
        },
        {
          default:()=>'仪表盘'
        }
      )
    },
    key: 'dashboard',
  },
  {
    label: () => {
      return h(RouterLink,
        {
          to:{
            name:'ModelSchemaManagement'
          },
        },
        {
          default:()=>'内容模型管理'
        }
      )
    },
    key: 'content-model-management',
  },
  {
    label: () => {
      return h(RouterLink,
        {
          to:{
            name:'ModelContentManagement'
          },
        },
        {
          default:()=>'内容管理'
        }
      )
    },
    key: 'content-management',
  },
  {
    label: () => {
      return h(RouterLink,
        {
          to:{
            name:'PayChannelManagement'
          },
        },
        {
          default:()=>'支付渠道'
        }
      )
    },
    key: 'payment-channel',
  },
  {
    label: () => {
      return h(RouterLink,
        {
          to:{
            name:'PayOrderManagement'
          },
        },
        {
          default:()=>'支付订单'
        }
      )
    },
    key: 'payment-order',
  },
  {
    label: () => {
      return h(RouterLink,
        {
          to:{
            name:'UserManagement'
          },
        },
        {
          default:()=>'用户权限管理'
        }
      )
    },
    key: 'user-management',
  },
  {
    label: () => {
      return h(RouterLink,
        {
          to:{
            name:'WebhookManagement'
          },
        },
        {
          default:()=>'webhook'
        }
      )
    },
    key: 'webhook-management',
  },
  {
    label: () => {
      return h(RouterLink,
        {
          to:{
            name:'FileManagement',
            params:{}
          },
        },
        {
          default:()=>'文件管理'
        }
      )
    },
    key: 'file-management',
  },
    {
    label: () => {
      return h(RouterLink,
        {
          to:{
            name:'ApiManagement'
          },
        },
        {
          default:()=>'API管理'
        }
      )
    },
    key: 'api-management',
  },
  {
    label: () => {
      return h(RouterLink,
        {
          to:{
            name:'SystemSettings'
          },
        },
        {
          default:()=>'系统设置'
        }
      )
    },
    key: 'system-settings',
  },

]

const {isMobile,showFooter,showTabbar,showLogo,showSidebar,showNav,collapsed,navFixed,footerFixed,layoutMode}  = storeToRefs(appStore)
const navHeight = ref(50)
const sidebarWidth = ref(224)
const tabbarHeight = ref(38)
const footerHeight = ref(50)
const sidebarCollapsedWidth = ref(58)

const { layout, verticalLayout } = useLayoutMenu({
  mode:layoutMode,
  menus: menuOptions,
})

const hasHorizontalMenu = computed(() =>
  ['horizontal', 'mixed-two-column', 'mixed-sidebar'].includes(layoutMode.value),
)

// function updateMode(v: ProLayoutMode) {
//   mode.value = v
//   isMobile.value = v === 'mobile'
// }

function logout() {
  if (confirm('确定要登出吗？')) {
    // 发送登出请求
    fetch('/logout', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
    })
      .then((response) => {
        if (response.ok) {
          window.location.href = '/login'
        } else {
          alert('登出失败，请重试')
        }
      })
      .catch((error) => {
        console.error('登出错误:', error)
        alert('登出失败，请重试')
      })
  }
}
</script>
<template>
  <div class="h-dvh w-dvw">
    <pro-layout
      v-model:collapsed="collapsed"
      :mode="layoutMode"
      :show-nav="showNav"
      :show-logo="showLogo"
      :is-mobile="isMobile"
      :nav-fixed="navFixed"
      :nav-height="navHeight"
      :show-footer="showFooter"
      :show-tabbar="showTabbar"
      :show-sidebar="showSidebar"
      :footer-fixed="footerFixed"
      :footer-height="footerHeight"
      :sidebar-width="sidebarWidth"
      :tabbar-height="tabbarHeight"
      :sidebar-collapsed-width="sidebarCollapsedWidth"
      logo-class="flex justify-center"
    >
      <template #logo>logo</template>
      <template #nav-left>
        <span>left</span>
        <n-popover v-if="isMobile" trigger="click" style="padding: 0">
          <template #trigger>
            <n-button type="primary" size="small"> 菜单 </n-button>
          </template>
          <n-scrollbar class="flex-[1_0_0]">
            <n-menu v-bind="verticalLayout.verticalMenuProps" :collapsed="false" />
          </n-scrollbar>
        </n-popover>
      </template>
      <template #nav-center>
        <n-menu v-if="hasHorizontalMenu" v-bind="layout.horizontalMenuProps" />
      </template>
      <template #footer>footer</template>
      <template #sidebar>
        <n-scrollbar class="flex-[1_0_0]">
          <n-menu v-bind="layout.verticalMenuProps" :collapsed-width="sidebarCollapsedWidth" />
        </n-scrollbar>
        <n-divider />

        <!-- 用户信息卡片 -->
        <div class="p-2">
          <div class="bg-gray-100 dark:bg-gray-800 rounded-lg py-4 px-2 shadow-sm">
            <div class="flex items-center space-x-3">
              <n-avatar
                :style="{
                  color: 'yellow',
                  backgroundColor: 'red',
                }"
              >
                M
              </n-avatar>
              <div class="flex-1 min-w-0">
                <p class="text-sm font-medium text-gray-900 dark:text-white truncate">
                  {CurrentUser.Name}
                </p>
                <p class="text-xs text-gray-500 dark:text-gray-400 truncate">{CurrentUser.Email}</p>
              </div>
              <button
                @click="logout()"
                class="p-2 text-gray-400 hover:text-gray-600 dark:hover:text-gray-300 rounded-lg hover:bg-gray-200 dark:hover:bg-gray-700 transition-colors"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
                  ></path>
                </svg>
              </button>
            </div>
          </div>
        </div>
      </template>
      <slot v-if="$slots.default" />
      <router-view v-else />
    </pro-layout>
  </div>
</template>
