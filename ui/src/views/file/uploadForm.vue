<script lang="ts" setup>
import Uppy from '@uppy/core';
import Dashboard from '@uppy/dashboard';
import Zhcn from '@uppy/locales/lib/zh_CN';

import '@uppy/core/css/style.min.css';
import '@uppy/dashboard/css/style.min.css';

let uppy:Uppy|null = null;

onMounted(()=>{
  uppy = new Uppy({
    locale:Zhcn,
  }).use(Dashboard,{inline:true,target:"#uppy-dashboard"})
})

onBeforeUnmount(()=>{
  if(uppy){
    uppy.destroy()
    uppy = null
  }
})

</script>
<template>
  <div>
    <div id="uppy-dashboard"></div>
  </div>
</template>
