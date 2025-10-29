<script setup lang="ts">
import type { MenuOption } from 'naive-ui'
import { ProLayout, useLayoutMenu,type ProLayoutProps } from 'pro-naive-ui'
import { RouterLink, useRoute, useRouter } from 'vue-router'
import { useAppStore } from '@/stores/modules/app'
import { storeToRefs } from 'pinia'
import { usePermissionStore } from '@/stores/modules/permission'
import { useUserStore } from '@/stores/modules/user'
import { remainingPaths } from '@/router'

const route = useRoute()
const router = useRouter()
const appStore = useAppStore()
const permissionStore = usePermissionStore()

const userStore = useUserStore()
const message = useMessage()

window.$message = useMessage()
window.$dialog = useDialog()
window.$notification = useNotification()

type LayoutThemeOverrides = NonNullable<ProLayoutProps['builtinThemeOverrides']>

const layoutThemeOverrides :LayoutThemeOverrides= {
  color:'#f5f7fa'
}

const menuData = computed(() => permissionStore.wholeMenus)

const menuOptions = computed<MenuOption[]>(() =>
  menuData.value.map((item: any) => {
    if (item.children && item.children.length > 1) {
      return {
        label: item.meta?.title || item.name,
        key: item.name,
        children: item.children.map((child: any) => ({
          label: () => {
            return h(
              RouterLink,
              {
                to: {
                  name: child.name,
                },
              },
              {
                default: () => child.meta?.title || child.name,
              },
            )
          },
          key: child.name,
        })),
      }
    }
    return {
      label: () => {
        return h(
          RouterLink,
          {
            to: {
              name: item.name,
            },
          },
          {
            default: () => item.meta?.title || item.name,
          },
        )
      },
      key: item.name,
    }
  }),
)

const {
  isMobile,
  showFooter,
  showTabbar,
  showLogo,
  showSidebar,
  showNav,
  collapsed,
  navFixed,
  footerFixed,
  layoutMode,
} = storeToRefs(appStore)
const navHeight = ref(50)
const sidebarWidth = ref(224)
const tabbarHeight = ref(38)
const footerHeight = ref(50)
const sidebarCollapsedWidth = ref(58)
const loading = computed(() => menuData.value.length === 0)

const { layout, verticalLayout } = useLayoutMenu({
  mode: layoutMode,
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
    userStore.logOut().then(() => {
      router.push('/login').then(() => {
        message.success('登出成功')
      })
    })
  }
}

/** 判断路径是否参与菜单 */
function isRemaining(path: string) {
  return remainingPaths.includes(path)
}

function menuSelect(path: string) {
  if (permissionStore.wholeMenus.length === 0 || isRemaining(path)) return
}

watch(
  () => [route.path, usePermissionStore().wholeMenus],
  () => {
    if (route.path.includes('/redirect')) return
    menuSelect(route.path)
  },
)
</script>
<template>
  <div class="app-wrapper h-dvh w-dvw">
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
      :builtin-theme-overrides="layoutThemeOverrides"
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
            <n-spin v-if="loading" />
            <n-menu v-else v-bind="verticalLayout.verticalMenuProps" :collapsed="false" />
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
                  {{ userStore.username }}
                </p>
                <p class="text-xs text-gray-500 dark:text-gray-400 truncate">
                  {{ userStore.username }}
                </p>
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
      <div class="app-main app-main--vertical">
        <slot v-if="$slots.default" />
        <router-view v-else />
      </div>
    </pro-layout>
  </div>
</template>
<style scoped>
.app-main {
  position: relative;
  width: 100%;
  overflow-x: hidden;
  /* background-color: #f5f7fa; */
}
.app-main--vertical{
  display: flex;
  flex-direction: column;
}
</style>
