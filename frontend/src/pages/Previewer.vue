<template>
  <div class="h-screen flex flex-col bg-gray-50">
    <!-- 顶部工具栏 -->
    <header
      class="flex items-center justify-between px-4 py-2 bg-white shadow-sm border-b"
    >
      <router-link
        to="/"
        class="text-sm px-3 py-1 rounded bg-gray-100 hover:bg-gray-200 transition"
      >
        返回主页
      </router-link>
      <div class="text-gray-600 text-sm">
        {{ selectedImage?.name }}
      </div>
      <div class="w-20"></div>
      <!-- 占位保持对称 -->
    </header>

    <!-- 图片预览区域 -->
    <div
      class="relative flex-1 flex items-center justify-center bg-gray-100 overflow-hidden"
    >
      <!-- 左箭头 -->
      <button
        class="absolute left-4 z-10 bg-white bg-opacity-80 hover:bg-opacity-100 rounded-full p-2 shadow"
        @click="goToPrevious"
        :disabled="currentIndex <= 0"
        title="上一张"
      >
        <svg
          class="w-5 h-5 text-gray-700"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M15 19l-7-7 7-7"
          />
        </svg>
      </button>

      <!-- 图片 -->
      <img
        v-if="selectedImage"
        :src="'/files/image/' + selectedImage.path"
        :alt="selectedImage.name"
        class="max-w-full max-h-full object-contain"
        @error="onThumbnailError"
      />

      <!-- 右箭头 -->
      <button
        class="absolute right-4 z-10 bg-white bg-opacity-80 hover:bg-opacity-100 rounded-full p-2 shadow"
        @click="goToNext"
        :disabled="currentIndex >= imageFiles.length - 1"
        title="下一张"
      >
        <svg
          class="w-5 h-5 text-gray-700"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M9 5l7 7-7 7"
          />
        </svg>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, watch, ref } from "vue";
import { useRouter } from "vue-router";
import { useImageStore } from "@/stores/image";
import { storeToRefs } from "pinia";

// 引入状态
const imageStore = useImageStore();
const { imageFiles, selectedImage, thumbnailSize } = storeToRefs(imageStore);
const { selectImage, onThumbnailError } = imageStore;

// 计算当前图片在列表中的索引
const currentIndex = computed(() =>
  imageFiles.value.findIndex((img) => img.path === selectedImage.value?.path)
);

// 向前切换
const goToPrevious = () => {
  if (currentIndex.value > 0) {
    selectImage(imageFiles.value[currentIndex.value - 1]);
  }
};

// 向后切换
const goToNext = () => {
  if (currentIndex.value < imageFiles.value.length - 1) {
    selectImage(imageFiles.value[currentIndex.value + 1]);
  }
};
</script>
