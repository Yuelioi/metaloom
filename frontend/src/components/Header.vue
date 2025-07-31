<template>
  <!-- 顶部工具栏 -->
  <div
    class="bg-white border-b border-gray-200 px-4 py-3 flex items-center justify-between"
  >
    <div class="flex items-center space-x-4">
      <h1 class="text-xl font-semibold text-gray-800">图片管理器</h1>
      <button
        @click="openFolder"
        class="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-lg flex items-center space-x-2 transition-colors"
      >
        <svg
          class="w-4 h-4"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-5l-2-2H5a2 2 0 00-2 2z"
          ></path>
        </svg>
        <span>打开文件夹</span>
      </button>
      <!-- <button @click="router.push('/previewer')">tiaojian</button> -->
    </div>
    <div class="text-sm text-gray-500">
      {{ currentFolder || "未选择文件夹" }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { useImageStore } from "@/stores/image";
import { useRouter } from "vue-router";
import { storeToRefs } from "pinia";

import { App } from "@app";

const imageStore = useImageStore();

const router = useRouter();

const { currentFolder } = storeToRefs(imageStore);

const openFolder = () => {
  App.OpenFolderDialog().then((folder) => {
    if (folder) {
      currentFolder.value = folder;
    }
  });
};
</script>
