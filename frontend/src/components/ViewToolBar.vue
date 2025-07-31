<template>
  <div class="flex items-center space-x-4">
    <span class="text-sm text-gray-600">
      {{ imageFiles.length }} 个图片文件

      {{ thumbnailSize }}
    </span>
    <div class="flex items-center space-x-2">
      <span class="text-xs text-gray-500">缩略图大小:</span>
      <input
        v-model.number="thumbnailSize"
        type="range"
        min="80"
        max="300"
        step="30"
        class="w-16 h-1 bg-gray-300 rounded-lg appearance-none cursor-pointer"
      />
    </div>
  </div>

  <div class="flex items-center space-x-2">
    <button
      @click="viewMode = 'grid'"
      :class="[
        'p-2 rounded-lg',
        viewMode === 'grid'
          ? 'bg-blue-100 text-blue-600'
          : 'text-gray-500 hover:bg-gray-100',
      ]"
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
          d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z"
        ></path>
      </svg>
    </button>
    <button
      @click="viewMode = 'list'"
      :class="[
        'p-2 rounded-lg',
        viewMode === 'list'
          ? 'bg-blue-100 text-blue-600'
          : 'text-gray-500 hover:bg-gray-100',
      ]"
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
          d="M4 6h16M4 10h16M4 14h16M4 18h16"
        ></path>
      </svg>
    </button>
  </div>
</template>

<script setup lang="ts">
import { useImageStore } from "@/stores/image";
import { storeToRefs } from "pinia";

const imageStore = useImageStore();

const { imageFiles, thumbnailSize, viewMode } = storeToRefs(imageStore);
</script>
