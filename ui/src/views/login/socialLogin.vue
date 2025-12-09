<script lang="ts" setup>
import LoginLayout from './components/loginLayout.vue';

const route = useRoute()
const message = useMessage()

const code = route.query.code as string

const tryLogin = async () => {
  if (!code) {
    message.error("缺少授权码")
    return
  }
}

onMounted(()=>{
  tryLogin()
})
</script>
<template>
  <LoginLayout>
    <div></div>
  </LoginLayout>
</template>
