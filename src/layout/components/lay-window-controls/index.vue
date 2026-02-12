<script setup lang="ts">
import { ref, onMounted } from "vue";
import Minus from "~icons/ri/subtract-line";
import FullScreen from "~icons/ri/fullscreen-line";
import ExitFullScreen from "~icons/ri/fullscreen-exit-line";
import Close from "~icons/ri/close-line";

import {
  WindowMinimize,
  WindowMaximize,
  WindowClose,
  WindowIsMaximised
} from "../../../../wailsjs/go/main/App";

const isMaximized = ref(false);

// 检查窗口状态
const checkWindowState = async () => {
  try {
    isMaximized.value = await WindowIsMaximised();
  } catch (e) {
    console.error("Failed to get window state:", e);
  }
};

// 最小化
const handleMinimize = async () => {
  try {
    await WindowMinimize();
  } catch (e) {
    console.error("Failed to minimize window:", e);
  }
};

// 最大化/恢复
const handleMaximize = async () => {
  try {
    await WindowMaximize();
    // 延迟检查状态，因为最大化动画需要时间
    setTimeout(checkWindowState, 300);
  } catch (e) {
    console.error("Failed to maximize window:", e);
  }
};

// 关闭
const handleClose = async () => {
  try {
    await WindowClose();
  } catch (e) {
    console.error("Failed to close window:", e);
  }
};

onMounted(() => {
  checkWindowState();
  // 监听窗口大小变化
  window.addEventListener("resize", checkWindowState);
});
</script>

<template>
  <div class="window-controls">
    <div
      class="control-btn minimize"
      :title="'最小化'"
      @click="handleMinimize"
    >
      <IconifyIconOffline :icon="Minus" />
    </div>
    <div
      class="control-btn maximize"
      :title="isMaximized ? '还原' : '最大化'"
      @click="handleMaximize"
    >
      <IconifyIconOffline
        :icon="isMaximized ? ExitFullScreen : FullScreen"
      />
    </div>
    <div class="control-btn close" :title="'关闭'" @click="handleClose">
      <IconifyIconOffline :icon="Close" />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.window-controls {
  display: flex;
  align-items: center;
  height: 48px;
  margin-left: 8px;
  -webkit-app-region: no-drag; // 允许点击

  .control-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 40px;
    height: 100%;
    font-size: 16px;
    color: #666;
    cursor: pointer;
    transition: all 0.2s;

    &:hover {
      background-color: rgb(0 0 0 / 5%);
    }

    &.minimize:hover {
      color: #409eff;
    }

    &.maximize:hover {
      color: #67c23a;
    }

    &.close:hover {
      color: #fff;
      background-color: #f56c6c;
    }
  }
}

.dark {
  .window-controls {
    .control-btn {
      color: #ccc;

      &:hover {
        background-color: rgb(255 255 255 / 10%);
      }

      &.close:hover {
        background-color: #f56c6c;
      }
    }
  }
}
</style>
