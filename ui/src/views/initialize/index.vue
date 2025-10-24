<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-2xl w-full space-y-8">
      <div>
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
          系统初始化
        </h2>
        <p class="mt-2 text-center text-sm text-gray-600">
          请完成以下步骤来初始化您的系统
        </p>
      </div>

      <div class="bg-white p-6 rounded-lg shadow-md">
        <div class="mb-6">
          <div class="flex items-center">
            <div class="flex-1 border-t-2 border-gray-300"></div>
            <div class="flex items-center justify-center mx-4">
              <span :class="[
                'flex items-center justify-center w-8 h-8 rounded-full font-bold',
                currentStep > 1 ? 'bg-green-600 text-white' : currentStep === 1 ? 'bg-blue-600 text-white' : 'bg-gray-300 text-gray-600'
              ]">1</span>
              <span class="ml-2 font-medium" :class="currentStep >= 1 ? 'text-gray-900' : 'text-gray-500'">选择数据库</span>
            </div>
            <div class="flex-1 border-t-2 border-gray-300"></div>
            <div class="flex items-center justify-center mx-4">
              <span :class="[
                'flex items-center justify-center w-8 h-8 rounded-full font-bold',
                currentStep > 2 ? 'bg-green-600 text-white' : currentStep === 2 ? 'bg-blue-600 text-white' : 'bg-gray-300 text-gray-600'
              ]">2</span>
              <span class="ml-2 font-medium" :class="currentStep >= 2 ? 'text-gray-900' : 'text-gray-500'">连接设置</span>
            </div>
            <div class="flex-1 border-t-2 border-gray-300"></div>
            <div class="flex items-center justify-center mx-4">
              <span :class="[
                'flex items-center justify-center w-8 h-8 rounded-full font-bold',
                currentStep > 3 ? 'bg-green-600 text-white' : currentStep === 3 ? 'bg-blue-600 text-white' : 'bg-gray-300 text-gray-600'
              ]">3</span>
              <span class="ml-2 font-medium" :class="currentStep >= 3 ? 'text-gray-900' : 'text-gray-500'">管理员账户</span>
            </div>
            <div class="flex-1 border-t-2 border-gray-300"></div>
          </div>
        </div>

        <form @submit.prevent="submitForm">
          <!-- 步骤1: 选择数据库类型 -->
          <div v-show="currentStep === 1">
            <div class="mb-6">
              <h3 class="text-lg font-medium text-gray-900 mb-4">选择数据库类型</h3>
              <div class="grid grid-cols-1 gap-4 sm:grid-cols-3">
                <div class="relative">
                  <input type="radio" id="sqlite" v-model="formData.dbType" value="sqlite" class="hidden peer" checked>
                  <label for="sqlite"
                    class="flex flex-col items-center justify-between p-4 text-gray-500 bg-white border border-gray-200 rounded-lg cursor-pointer peer-checked:border-blue-600 peer-checked:text-blue-600 hover:bg-gray-50">
                    <div class="w-full text-center">
                      <svg class="w-10 h-10 mx-auto mb-2" fill="currentColor" viewBox="0 0 24 24">
                        <path
                          d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-1 17.93c-3.95-.49-7-3.85-7-7.93 0-.62.08-1.21.21-1.79L9 15v1c0 1.1.9 2 2 2v1.93zm6.9-2.54c-.26-.81-1-1.39-1.9-1.39h-1v-3c0-.55-.45-1-1-1H8v-2h2c.55 0 1-.45 1-1V7h2c1.1 0 2-.9 2-2v-.41c2.93 1.19 5 4.06 5 7.41 0 2.08-.8 3.97-2.1 5.39z" />
                      </svg>
                      <h4 class="text-base font-semibold">SQLite</h4>
                      <p class="text-sm">轻量级文件数据库，适合小型应用</p>
                    </div>
                  </label>
                </div>

                <div class="relative">
                  <input type="radio" id="mysql" v-model="formData.dbType" value="mysql" class="hidden peer">
                  <label for="mysql"
                    class="flex flex-col items-center justify-between p-4 text-gray-500 bg-white border border-gray-200 rounded-lg cursor-pointer peer-checked:border-blue-600 peer-checked:text-blue-600 hover:bg-gray-50">
                    <div class="w-full text-center">
                      <svg class="w-10 h-10 mx-auto mb-2" fill="currentColor" viewBox="0 0 24 24">
                        <path
                          d="M12 3c-4.97 0-9 4.03-9 9s4.03 9 9 9 9-4.03 9-9-4.03-9-9-9zm0 16c-3.87 0-7-3.13-7-7s3.13-7 7-7 7 3.13 7 7-3.13 7-7 7zm-1.46-5.47L8.41 11.4l-1.06 1.06 3.18 3.18 6-6-1.06-1.06-4.93 4.95z" />
                      </svg>
                      <h4 class="text-base font-semibold">MySQL</h4>
                      <p class="text-sm">流行的开源关系型数据库</p>
                    </div>
                  </label>
                </div>

                <div class="relative">
                  <input type="radio" id="postgresql" v-model="formData.dbType" value="postgresql" class="hidden peer">
                  <label for="postgresql"
                    class="flex flex-col items-center justify-between p-4 text-gray-500 bg-white border border-gray-200 rounded-lg cursor-pointer peer-checked:border-blue-600 peer-checked:text-blue-600 hover:bg-gray-50">
                    <div class="w-full text-center">
                      <svg class="w-10 h-10 mx-auto mb-2" fill="currentColor" viewBox="0 0 24 24">
                        <path
                          d="M19 5v14H5V5h14m0-2H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2z" />
                        <path d="M14 17H7v-2h7v2zm3-4H7v-2h10v2zm0-4H7V7h10v2z" />
                      </svg>
                      <h4 class="text-base font-semibold">PostgreSQL</h4>
                      <p class="text-sm">高级开源关系型数据库</p>
                    </div>
                  </label>
                </div>
              </div>
            </div>

            <div class="flex justify-end">
              <button type="button" @click="nextStep"
                class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">
                下一步
              </button>
            </div>
          </div>

          <!-- 步骤2: 数据库连接参数 -->
          <div v-show="currentStep === 2">
            <div class="mb-6">
              <h3 class="text-lg font-medium text-gray-900 mb-4">数据库连接参数</h3>

              <!-- SQLite 参数 -->
              <div v-show="formData.dbType === 'sqlite'">
                <div class="mb-4">
                  <label for="sqlite-file" class="block text-sm font-medium text-gray-700 mb-1">数据库文件路径</label>
                  <input type="text" id="sqlite-file" v-model="formData.sqlite.file"
                    class="appearance-none rounded-md relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm"
                    placeholder="./data.db">
                  <p class="mt-1 text-sm text-gray-500">默认将在应用根目录创建数据库文件</p>
                </div>
              </div>

              <!-- MySQL 参数 -->
              <div v-show="formData.dbType === 'mysql'">
                <div class="grid grid-cols-2 gap-4">
                  <div class="mb-4">
                    <label for="mysql-host" class="block text-sm font-medium text-gray-700 mb-1">主机地址</label>
                    <input type="text" id="mysql-host" v-model="formData.mysql.host"
                      class="appearance-none rounded-md relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm"
                      placeholder="localhost">
                  </div>
                  <div class="mb-4">
                    <label for="mysql-port" class="block text-sm font-medium text-gray-700 mb-1">端口</label>
                    <input type="text" id="mysql-port" v-model="formData.mysql.port"
                      class="appearance-none rounded-md relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm"
                      placeholder="3306">
                  </div>
                </div>
                <div class="grid grid-cols-2 gap-4">
                  <div class="mb-4">
                    <label for="mysql-user" class="block text-sm font-medium text-gray-700 mb-1">用户名</label>
                    <input type="text" id="mysql-user" v-model="formData.mysql.user"
                      class="appearance-none rounded-md relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm"
                      placeholder="root">
                  </div>
                  <div class="mb-4">
                    <label for="mysql-password" class="block text-sm font-medium text-gray-700 mb-1">密码</label>
                    <input type="password" id="mysql-password" v-model="formData.mysql.password"
                      class="appearance-none rounded-md relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm">
                  </div>
                </div>
                <div class="mb-4">
                  <label for="mysql-database" class="block text-sm font-medium text-gray-700 mb-1">数据库名</label>
                  <input type="text" id="mysql-database" v-model="formData.mysql.database"
                    class="appearance-none rounded-md relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm"
                    placeholder="gobee">
                </div>
              </div>

              <!-- PostgreSQL 参数 -->
              <div v-show="formData.dbType === 'postgresql'">
                <div class="grid grid-cols-2 gap-4">
                  <div class="mb-4">
                    <label for="pg-host" class="block text-sm font-medium text-gray-700 mb-1">主机地址</label>
                    <input type="text" id="pg-host" v-model="formData.postgresql.host"
                      class="appearance-none rounded-md relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm"
                      placeholder="localhost">
                  </div>
                  <div class="mb-4">
                    <label for="pg-port" class="block text-sm font-medium text-gray-700 mb-1">端口</label>
                    <input type="text" id="pg-port" v-model="formData.postgresql.port"
                      class="appearance-none rounded-md relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm"
                      placeholder="5432">
                  </div>
                </div>
                <div class="grid grid-cols-2 gap-4">
                  <div class="mb-4">
                    <label for="pg-user" class="block text-sm font-medium text-gray-700 mb-1">用户名</label>
                    <input type="text" id="pg-user" v-model="formData.postgresql.user"
                      class="appearance-none rounded-md relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm"
                      placeholder="postgres">
                  </div>
                  <div class="mb-4">
                    <label for="pg-password" class="block text-sm font-medium text-gray-700 mb-1">密码</label>
                    <input type="password" id="pg-password" v-model="formData.postgresql.password"
                      class="appearance-none rounded-md relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm">
                  </div>
                </div>
                <div class="mb-4">
                  <label for="pg-database" class="block text-sm font-medium text-gray-700 mb-1">数据库名</label>
                  <input type="text" id="pg-database" v-model="formData.postgresql.database"
                    class="appearance-none rounded-md relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm"
                    placeholder="gobee">
                </div>
              </div>
            </div>

            <div class="flex justify-between">
              <button type="button" @click="prevStep"
                class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md shadow-sm text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">
                上一步
              </button>
              <button type="button" @click="nextStep"
                class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">
                下一步
              </button>
            </div>
          </div>

          <!-- 步骤3: 超级管理员账户 -->
          <div v-show="currentStep === 3">
            <div class="mb-6">
              <h3 class="text-lg font-medium text-gray-900 mb-4">设置超级管理员账户</h3>
              <div class="mb-4">
                <label for="admin-username" class="block text-sm font-medium text-gray-700 mb-1">用户名</label>
                <input type="text" id="admin-username" v-model="formData.admin.username" required
                  class="appearance-none rounded-md relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm"
                  placeholder="admin">
              </div>
              <div class="mb-4">
                <label for="admin-email" class="block text-sm font-medium text-gray-700 mb-1">电子邮箱</label>
                <input type="email" id="admin-email" v-model="formData.admin.email" required
                  class="appearance-none rounded-md relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm"
                  placeholder="admin@example.com">
              </div>
              <div class="mb-4">
                <label for="admin-password" class="block text-sm font-medium text-gray-700 mb-1">密码</label>
                <input type="password" id="admin-password" v-model="formData.admin.password" required
                  class="appearance-none rounded-md relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm"
                  placeholder="请输入至少8位密码">
              </div>
              <div class="mb-4">
                <label for="admin-password-confirm" class="block text-sm font-medium text-gray-700 mb-1">确认密码</label>
                <input type="password" id="admin-password-confirm" v-model="formData.admin.passwordConfirm" required
                  class="appearance-none rounded-md relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm"
                  placeholder="请再次输入密码">
                <p v-if="passwordError" class="mt-1 text-sm text-red-600">{{ passwordError }}</p>
              </div>
            </div>

            <div class="flex justify-between">
              <button type="button" @click="prevStep"
                class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md shadow-sm text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">
                上一步
              </button>
              <button type="submit"
                class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500">
                完成初始化
              </button>
            </div>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';
import { useRouter } from 'vue-router';
import axios from 'axios';

const router = useRouter();
const currentStep = ref(1);
const passwordError = ref('');

const formData = reactive({
  dbType: 'sqlite',
  sqlite: {
    file: './data.db'
  },
  mysql: {
    host: 'localhost',
    port: '3306',
    user: '',
    password: '',
    database: 'gobee'
  },
  postgresql: {
    host: 'localhost',
    port: '5432',
    user: 'postgres',
    password: '',
    database: 'gobee'
  },
  admin: {
    username: '',
    email: '',
    password: '',
    passwordConfirm: ''
  }
});

const nextStep = () => {
  if (currentStep.value < 3) {
    currentStep.value++;
  }
};

const prevStep = () => {
  if (currentStep.value > 1) {
    currentStep.value--;
  }
};

const validateForm = () => {
  // 验证密码
  if (formData.admin.password !== formData.admin.passwordConfirm) {
    passwordError.value = '两次输入的密码不一致，请重新输入';
    return false;
  }
  
  if (formData.admin.password.length < 8) {
    passwordError.value = '密码长度至少为8位';
    return false;
  }
  
  passwordError.value = '';
  return true;
};

const submitForm = async () => {
  if (!validateForm()) {
    return;
  }
  
  try {
    // 准备提交的数据
    const submitData: any = {
      db_type: formData.dbType,
      admin_username: formData.admin.username,
      admin_email: formData.admin.email,
      admin_password: formData.admin.password,
      admin_password_confirm: formData.admin.passwordConfirm
    };
    
    // 根据数据库类型添加相应的参数
    if (formData.dbType === 'sqlite') {
      submitData.sqlite_file = formData.sqlite.file;
    } else if (formData.dbType === 'mysql') {
      submitData.mysql_host = formData.mysql.host;
      submitData.mysql_port = formData.mysql.port;
      submitData.mysql_user = formData.mysql.user;
      submitData.mysql_password = formData.mysql.password;
      submitData.mysql_database = formData.mysql.database;
    } else if (formData.dbType === 'postgresql') {
      submitData.pg_host = formData.postgresql.host;
      submitData.pg_port = formData.postgresql.port;
      submitData.pg_user = formData.postgresql.user;
      submitData.pg_password = formData.postgresql.password;
      submitData.pg_database = formData.postgresql.database;
    }
    
    // 发送请求
    const response = await axios.post('/initialize', submitData);
    
    // 处理响应
    if (response.data.success) {
      // 初始化成功，跳转到登录页
      router.push('/login');
    } else {
      // 显示错误信息
      alert(response.data.message || '初始化失败，请重试');
    }
  } catch (error) {
    console.error('初始化失败:', error);
    alert('初始化失败，请检查网络连接或联系管理员');
  }
};
</script>

<style scoped>
.step-content {
  transition: all 0.3s ease;
}
</style>