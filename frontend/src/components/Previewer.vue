<template>
  <div class="p-4">
    <h3 class="text-sm font-medium text-gray-700 mb-3">图片预览</h3>

    <div v-if="selectedImage" class="p-4 rounded-lg">
      <div class="flex flex-col items-center">
        <!-- 缩略图 -->
        <div
          class="w-48 h-48 mb-4 rounded-lg overflow-hidden bg-gray-100 flex items-center justify-center"
        >
          <img
            :src="'/files/image/' + selectedImage.path"
            :alt="selectedImage.name"
            class="w-full h-full object-cover"
            @error="onThumbnailError"
          />
        </div>

        <!-- 文件信息 -->
        <div class="text-center w-full">
          <div
            class="text-sm font-medium text-gray-800 truncate px-1"
            :title="selectedImage.name"
          >
            {{ selectedImage.name }}
          </div>
          <div class="text-sm text-gray-500 mt-1">
            {{ formatFileSize(selectedImage.size) }}
          </div>
        </div>
      </div>
    </div>

    <div v-else class="text-center py-20 text-gray-400">
      <svg
        class="w-12 h-12 mx-auto mb-3 text-gray-300"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"
        />
      </svg>
      <p class="text-sm">请选择一张图片</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useImageStore } from "@/stores/image";
import { storeToRefs } from "pinia";

const imageStore = useImageStore();

const { selectedImage } = storeToRefs(imageStore);

const { formatFileSize, onThumbnailError } = useImageStore();
</script>
