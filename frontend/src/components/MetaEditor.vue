<template>
  <div class="p-4 border-b border-gray-100">
    <div class="flex items-center justify-between">
      <h3 class="text-lg font-medium text-gray-800">元数据</h3>
      <button
        v-if="selectedImage && !selectedImage.hasMetadata"
        @click="createMetadata"
        class="text-sm bg-green-500 hover:bg-green-600 text-white px-3 py-1 rounded-lg transition-colors"
      >
        创建
      </button>
    </div>
  </div>

  <div v-if="selectedImage" class="flex-1 p-4 overflow-y-auto">
    <!-- 基本信息 -->
    <div class="mb-6">
      <h4 class="text-sm font-medium text-gray-700 mb-3">基本信息</h4>
      <div class="space-y-3">
        <div>
          <label class="block text-xs font-medium text-gray-600 mb-1"
            >标题</label
          >
          <input
            v-model="metadata.title"
            type="text"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="输入图片标题"
          />
        </div>
        <div>
          <label class="block text-xs font-medium text-gray-600 mb-1"
            >描述</label
          >
          <textarea
            v-model="metadata.description"
            rows="3"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500 resize-none"
            placeholder="输入图片描述"
          ></textarea>
        </div>
      </div>
    </div>

    <!-- 标签 -->
    <div class="mb-6">
      <h4 class="text-sm font-medium text-gray-700 mb-3">标签</h4>
      <div class="flex flex-wrap gap-2 mb-3">
        <span
          v-for="(tag, index) in metadata.tags"
          :key="index"
          class="inline-flex items-center px-2 py-1 bg-blue-100 text-blue-800 text-xs rounded-lg"
        >
          {{ tag }}
          <button @click="removeTag(index)" class="ml-1 hover:text-blue-600">
            <svg
              class="w-3 h-3"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M6 18L18 6M6 6l12 12"
              ></path>
            </svg>
          </button>
        </span>
      </div>
      <div class="flex">
        <input
          v-model="newTag"
          @keyup.enter="addTag"
          type="text"
          class="flex-1 px-3 py-2 border border-gray-300 rounded-l-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          placeholder="添加标签"
        />
        <button
          @click="addTag"
          class="px-3 py-2 bg-blue-500 hover:bg-blue-600 text-white rounded-r-lg text-sm transition-colors"
        >
          添加
        </button>
      </div>
    </div>

    <!-- 自定义字段 -->
    <div class="mb-6">
      <h4 class="text-sm font-medium text-gray-700 mb-3">自定义字段</h4>
      <div class="space-y-3">
        <div
          v-for="(field, index) in metadata.fields"
          :key="index"
          class="flex space-x-2"
        >
          <input
            v-model="field.key"
            type="text"
            class="flex-1 px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="字段名"
          />
          <input
            v-model="field.value"
            type="text"
            class="flex-1 px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="值"
          />
          <button
            @click="removeCustomField(index)"
            class="p-2 text-red-500 hover:text-red-700 hover:bg-red-50 rounded-lg"
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
                d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
              ></path>
            </svg>
          </button>
        </div>
        <button
          @click="addCustomField"
          class="w-full px-3 py-2 border border-dashed border-gray-300 rounded-lg text-sm text-gray-500 hover:border-gray-400 hover:text-gray-600 transition-colors"
        >
          + 添加自定义字段
        </button>
      </div>
    </div>

    <!-- 保存按钮 -->
    <div class="flex space-x-3">
      <button
        @click="saveMetadata"
        class="flex-1 bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded-lg text-sm font-medium transition-colors"
      >
        保存
      </button>
      <button
        @click="resetMetadata"
        class="px-4 py-2 border border-gray-300 text-gray-700 rounded-lg text-sm font-medium hover:bg-gray-50 transition-colors"
      >
        重置
      </button>
    </div>
  </div>

  <div v-else class="flex-1 flex items-center justify-center text-gray-400">
    <div class="text-center">
      <svg
        class="w-16 h-16 mx-auto mb-3 text-gray-300"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
        ></path>
      </svg>
      <p class="text-sm">选择图片查看元数据</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useImageStore } from "@/stores/image";
import { storeToRefs } from "pinia";

const imageStore = useImageStore();

const { metadata, selectedImage } = storeToRefs(imageStore);
const {
  resetMetadata,
  saveMetadata,
  addCustomField,
  removeCustomField,
  addTag,
  removeTag,
  createMetadata,
  newTag,
} = useImageStore();
</script>
