<script setup lang="ts">

const loginForm = ref({
  email:'',
  password:''
})
async function handleLogin() {


    try {
        const response = await fetch('/auth/login/password', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                email: loginForm.value.email,
                password: loginForm.value.password
            })
        });

        const result = await response.json();

        if (response.ok) {
            // 保存令牌到 localStorage
            localStorage.setItem('token', result.token);
            localStorage.setItem('user', JSON.stringify(result.user));

            // 跳转到控制台首页
            window.location.href = '/console';
        } else {
            throw new Error(result.error || '登录失败');
        }
    } catch (error:unknown) {
        if (error instanceof Error) {
            alert(error.message);
        } else {
            alert('登录失败');
        }
    }
}

onMounted(()=>{
  const token = localStorage.getItem('token');
    if (token) {
        // 如果已登录，直接跳转到控制台首页
        window.location.href = '/console';
    }
})
</script>
<template>
  <div class="container">
    <div class="row justify-content-center align-items-center min-vh-100">
      <div class="col-md-6 col-lg-4">
        <div class="card shadow">
          <div class="card-body p-5">
            <h3 class="text-center mb-4">登录</h3>
            <n-form id="loginForm">
              <n-form-item label="邮箱">
                <n-input v-model:value="loginForm.email" class="form-control" id="email" required />
              </n-form-item>
              <n-form-item label="密码">
                <n-input v-model:value="loginForm.password" type="password" class="form-control" id="password" required />
              </n-form-item>
              <n-form-item>
                <button type="submit" class="btn btn-primary" @click="handleLogin">登录</button>
              </n-form-item>
            </n-form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
