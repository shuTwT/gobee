<script setup lang="ts">
import { isFunction } from 'lodash-es'
import { closeDialog, dialogStore } from '.'
import type { DialogOptions, EventType, ModalButtonProps } from './types'
import { ExpandSharp } from '@vicons/ionicons5'

defineOptions({
  name: 'DialogManager',
})

const sureBtnMap = ref<Record<string, any>>({})

const fullscreenStyle = (options:DialogOptions)=>options.fullscreen?{
  maxHeight:'100vh',
  height:'100vh',
  width:'100%'
}:{
  maxHeight:'80vh',
  width:options.width
}

const footerButtons = computed(() => {
  return (options: DialogOptions) => {
    return options?.footerButtons?.length! > 0
      ? options.footerButtons
      : ([
          {
            label: '取消',
            quaternary: true,
            btnClick: ({ dialog: { options, index } }) => {
              const done = () => closeDialog(options!, index!, { command: 'cancel' })
              if (options?.beforeCancel && isFunction(options?.beforeCancel)) {
                options.beforeCancel(done, { options, index: Number(index) })
              }
              done()
            },
          },
          {
            label: '确定',
            type: 'primary',
            quaternary: true,
            btnClick: ({ dialog: { options, index } }) => {
              if (options?.sureBtnLoading) {
                sureBtnMap.value[index!] = Object.assign({}, sureBtnMap.value[index!], {
                  loading: true,
                })
              }
              const closeLoading = () => {
                if (options?.sureBtnLoading) {
                  sureBtnMap.value[index!].loading = false
                }
              }
              const done = () => {
                closeLoading()
                closeDialog(options!, index!, { command: 'sure' })
              }
              if (options?.beforeSure && isFunction(options?.beforeSure)) {
                options.beforeSure(done, { options, index: Number(index), closeLoading })
              } else {
                done()
              }
            },
          },
        ] as Array<ModalButtonProps>)
  }
})

function eventsCallback(
  event: EventType,
  options: DialogOptions,
  index: number,
  isClickFullScreen = false,
) {
  if (!isClickFullScreen) options.fullscreen = false
  if (options?.[event] && isFunction(options?.[event])) {
    return options?.[event]({ options, index })
  }
}

function handleClose(options: DialogOptions, index: number, args = { command: 'close' }) {
  closeDialog(options, index, args)
}
</script>
<template>
  <n-modal
    v-for="(options, index) in dialogStore"
    preset="card"
    :show="options.visible"
    :key="index"
    :closable="true"
    v-bind="options"
    :style="fullscreenStyle(options)"
    @after-enter="eventsCallback('open', options, index)"
    @close="handleClose(options, index)"
  >
    <!-- header -->
    <template #header>
      <span>{{ options.title }}</span>
    </template>
    <template #header-extra>
      <n-button tabindex="-1" :quaternary="true" style="width: 22px;height: 22px;padding: 0;" @click="()=>{
        options.fullscreen = !options.fullscreen;
        eventsCallback(
          'fullscreenCallBack',
          { ...options, fullscreen:options.fullscreen },
          index,
          true
        );
      }">
        <template #icon>
          <n-icon size="16">
            <expand-sharp />
          </n-icon>
        </template>
      </n-button>
    </template>
    <!-- content -->
     <n-scrollbar style="height: calc(80vh - 128px);">
      <div class="px-3">
        <component v-bind="options.props" :is="options.contentRenderer({ options, index })" />
      </div>
     </n-scrollbar>

    <!-- footer -->
    <template #footer>
      <template v-if="options.footer">
        <component :is="options.footer" />
      </template>
      <span v-else class="flex justify-end items-center">
        <template v-for="(btn, key) in footerButtons(options)" :key="key">
          <n-popconfirm v-if="btn.popconfirm" v-bind="btn.popconfirm">
            <template #trigger>
              <n-button v-bind="btn">{{ btn.label }}</n-button>
            </template>
          </n-popconfirm>
          <n-button
            v-else
            v-bind="btn"
            :loading="key === 1 && sureBtnMap[index]?.loading"
            @click="
              btn.btnClick!({
                dialog: { options, index },
                button: { btn, index: key },
              })
            "
            >{{ btn.label }}</n-button
          >
        </template>
      </span>
    </template>
  </n-modal>
</template>
