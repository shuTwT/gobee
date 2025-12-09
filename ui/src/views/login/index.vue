<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import {
  NForm,
  NFormItem,
  NInput,
  NButton,
  NIcon,
  NDivider,
  NSpace,
  NText,
} from 'naive-ui'
import { LogoGithub, LogoGoogle } from '@vicons/ionicons5'
import { useUserStore } from '@/stores/modules/user'
import { initRouter } from '@/router/utils'
import LoginLayout from './components/loginLayout.vue'

const message = useMessage()
const router = useRouter()

const loginForm = reactive({
  email: '',
  password: '',
  remember: false,
})

const loading = ref(false)

const handleLogin = async () => {
  if (!loginForm.email || !loginForm.password) {
    alert('请输入用户名和密码')
    return
  }

  loading.value = true

  useUserStore()
    .loginByUsername({ ...loginForm })
    .then(() => {
      return initRouter()
        .then(() => {
          router.push('/').then(() => {
            message.success('登录成功')
          })
        })
        .catch((err: Error) => {
          console.error(err)
        })
    })
    .catch((err: Error) => {
      console.error(err)
    })
    .finally(() => {
      loading.value = false
    })
}

const handleSocialLogin = (provider: string) => {
  const redirect_uri = 'http://localhost:5379/console/social-login'
  const socialMap: Record<string, any> = {
    github: {
      client_id: '932203383335936000',
      login() {
        window.location.href = `https://github.com/login/oauth/authorize?client_id=${this.client_id}&redirect_uri=${redirect_uri}&response_type=code`
      },
    },
    gitee: {
      client_id: 'a588d3a3ce3d6c407c9721991423d237167ab0415b2f8d42761ccf527405bf3d',
      login() {
        window.location.href = `https://gitee.com/oauth/authorize?client_id=${this.client_id}&redirect_uri=${redirect_uri}&response_type=code`
      },
    },
  }
  if (!socialMap[provider]) {
    alert('不支持该社交登录')
    return
  }
  socialMap[provider].login()
}

const handleRegister = () => {
  router.push('/register')
}


onMounted(() => {
  // 初始化逻辑
})
</script>

<template>
  <LoginLayout>
    <div class="p-6">

      <!-- 用户名密码登录 -->
      <n-form :model="loginForm" class="space-y-4">
        <n-form-item label="邮箱" path="email">
          <n-input v-model:value="loginForm.email" placeholder="请输入邮箱" size="large" clearable>
            <template #prefix>
              <n-icon class="text-gray-400">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  fill="currentColor"
                  class="w-5 h-5"
                >
                  <path
                    d="M12 2.5a5.5 5.5 0 0 1 3.096 10.047 8.75 8.75 0 0 1-2.528 5.962.5.5 0 0 1-.668-.186 6.5 6.5 0 0 0-1.332-2.745.5.5 0 0 1 .186-.668A8.75 8.75 0 0 1 12 2.5Z"
                  />
                </svg>
              </n-icon>
            </template>
          </n-input>
        </n-form-item>

        <n-form-item label="密码" path="password">
          <n-input
            v-model:value="loginForm.password"
            type="password"
            placeholder="请输入密码"
            size="large"
            show-password-on="click"
            clearable
            @keyup.enter="handleLogin"
          >
            <template #prefix>
              <n-icon class="text-gray-400">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  fill="currentColor"
                  class="w-5 h-5"
                >
                  <path
                    fill-rule="evenodd"
                    d="M12 1.5a5.25 5.25 0 0 0-5.25 5.25v3a3 3 0 0 0-3 3v6.75a3 3 0 0 0 3 3h10.5a3 3 0 0 0 3-3v-6.75a3 3 0 0 0-3-3v-3A5.25 5.25 0 0 0 12 1.5Zm-3.75 5.25v3h7.5v-3a3.75 3.75 0 1 0-7.5 0Z"
                    clip-rule="evenodd"
                  />
                </svg>
              </n-icon>
            </template>
          </n-input>
        </n-form-item>

        <n-form-item>
          <n-button
            type="primary"
            size="large"
            :loading="loading"
            :disabled="loading"
            @click="handleLogin"
            class="!w-full"
          >
            登录
          </n-button>
        </n-form-item>
      </n-form>

      <!-- 分割线 -->
      <n-divider>
        <n-text type="info">或者</n-text>
      </n-divider>

      <!-- 社交登录 -->
      <div class="mb-6">
        <n-space justify="center" :wrap="false" :size="16">
          <n-button
            quaternary
            size="large"
            @click="handleSocialLogin('github')"
            class="border border-gray-300 hover:border-gray-400 flex-1"
          >
            <template #icon>
              <n-icon><LogoGithub /></n-icon>
            </template>
            GitHub
          </n-button>

          <n-button
            quaternary
            size="large"
            @click="handleSocialLogin('gitee')"
            class="border border-gray-300 hover:border-gray-400 flex-1"
          >
            <template #icon>
              <n-icon><LogoGoogle /></n-icon>
            </template>
            Google
          </n-button>
        </n-space>
      </div>

      <!-- 注册链接 -->
      <div class="text-center text-gray-600">
        <span>还没有账号？</span>
        <n-button text type="primary" @click="handleRegister" class="font-medium">
          立即注册
        </n-button>
      </div>
    </div>
  </LoginLayout>
</template>

