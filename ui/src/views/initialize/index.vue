<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
    <n-card class="max-w-2xl w-full" :bordered="false">
      <template #header>
        <div class="text-center">
          <h2 class="text-3xl font-bold text-gray-900">系统初始化</h2>
          <p class="mt-2 text-sm text-gray-600">请完成以下步骤来初始化您的系统</p>
        </div>
      </template>

      <n-steps :current="currentStep" class="mb-8">
        <n-step title="确认数据库" description="查看当前数据库类型" />
        <n-step title="管理员账户" description="注册超级管理员" />
      </n-steps>

      <n-spin :show="loading">
        <n-form ref="formRef" :model="formData" :rules="rules" label-placement="top" size="large">
          <n-form-item v-if="currentStep === 1" label="当前数据库类型">
            <n-tag :type="dbTypeTagType" size="large" round>
              {{ dbTypeText }}
            </n-tag>
          </n-form-item>

          <n-alert v-if="currentStep === 1" type="info" class="mb-6">
            <template #header>
              <span class="font-medium">数据库配置提示</span>
            </template>
            当前 {{ dbTypeText }} 数据库，如果不是您想要的数据库，请修改 config.toml 后重启服务
          </n-alert>

          <template v-if="currentStep === 2">
            <n-form-item label="用户名" path="username">
              <n-input v-model:value="formData.username" placeholder="请输入用户名" />
            </n-form-item>

            <n-form-item label="电子邮箱" path="email">
              <n-input v-model:value="formData.email" placeholder="请输入电子邮箱" />
            </n-form-item>

            <n-form-item label="密码" path="password">
              <n-input v-model:value="formData.password" type="password" show-password-on="click" placeholder="请输入至少8位密码" />
            </n-form-item>

            <n-form-item label="确认密码" path="confirmPassword">
              <n-input v-model:value="formData.confirmPassword" type="password" show-password-on="click" placeholder="请再次输入密码" />
            </n-form-item>
          </template>
        </n-form>
      </n-spin>

      <template #footer>
        <div class="flex justify-between">
          <n-button v-if="currentStep > 1" @click="prevStep" quaternary>
            上一步
          </n-button>
          <div class="flex-1"></div>
          <n-button v-if="currentStep === 1" type="primary" @click="nextStep" :loading="loading">
            下一步
          </n-button>
          <n-button v-else type="primary" @click="submitForm" :loading="submitting" :disabled="!isFormValid">
            完成初始化
          </n-button>
        </div>
      </template>
    </n-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { preInit, initialize } from '@/api/initialize';
import type { FormInst, FormRules } from 'naive-ui';

const message = useMessage();
const router = useRouter();
const formRef = ref<FormInst | null>(null);
const currentStep = ref(1);
const loading = ref(false);
const submitting = ref(false);
const dbType = ref('');

const formData = reactive({
  username: '',
  email: '',
  password: '',
  confirmPassword: ''
});

const rules: FormRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, message: '用户名至少3个字符', trigger: 'blur' }
  ],
  email: [
    { required: true, message: '请输入电子邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入有效的电子邮箱', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 8, message: '密码至少8位', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: '请确认密码', trigger: 'blur' },
    {
      validator: (rule, value) => {
        return value === formData.password;
      },
      message: '两次输入的密码不一致',
      trigger: 'blur'
    }
  ]
};

const dbTypeText = computed(() => {
  const typeMap: Record<string, string> = {
    'sqlite': 'SQLite',
    'mysql': 'MySQL',
    'postgresql': 'PostgreSQL'
  };
  return typeMap[dbType.value] || dbType.value.toUpperCase();
});

const dbTypeTagType = computed(() => {
  const typeMap: Record<string, 'success' | 'info' | 'warning'> = {
    'sqlite': 'success',
    'mysql': 'info',
    'postgresql': 'warning'
  };
  return typeMap[dbType.value] || 'info';
});

const isFormValid = computed(() => {
  return formData.username && 
         formData.email && 
         formData.password && 
         formData.confirmPassword &&
         formData.password === formData.confirmPassword &&
         formData.password.length >= 8;
});

const fetchDbType = async () => {
  loading.value = true;
  try {
    const response = await preInit();
    if (response.code === 200 && response.data) {
      dbType.value = response.data.dbType || 'sqlite';
    } else {
      message.error('获取数据库类型失败');
    }
  } catch (error) {
    console.error('获取数据库类型失败:', error);
    message.error('获取数据库类型失败');
  } finally {
    loading.value = false;
  }
};

const nextStep = () => {
  currentStep.value++;
};

const prevStep = () => {
  currentStep.value--;
};

const submitForm = async () => {
  if (!formRef.value) return;

  try {
    await formRef.value.validate();
  } catch (error) {
    return;
  }

  submitting.value = true;
  try {
    const response = await initialize({
      admin_username: formData.username,
      admin_email: formData.email,
      admin_password: formData.password,
      confirm_password: formData.confirmPassword
    });

    if (response.code === 200) {
      message.success('初始化成功');
      router.push('/login');
    } else {
      message.error(response.msg || '初始化失败，请重试');
    }
  } catch (error) {
    console.error('初始化失败:', error);
    message.error('初始化失败，请检查网络连接或联系管理员');
  } finally {
    submitting.value = false;
  }
};

onMounted(() => {
  fetchDbType();
});
</script>

<style scoped>
:deep(.n-card) {
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
}

:deep(.n-button--primary-type) {
  background-color: #18a058;
  border-color: #18a058;
}

:deep(.n-button--primary-type:hover) {
  background-color: #36ad6a;
  border-color: #36ad6a;
}

:deep(.n-button--primary-type:active) {
  background-color: #0c7a43;
  border-color: #0c7a43;
}

:deep(.n-steps .n-step-header__indicator) {
  background-color: #18a058;
}

:deep(.n-steps .n-step-header__title) {
  color: #18a058;
}
</style>
