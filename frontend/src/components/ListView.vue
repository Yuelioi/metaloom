<template>
  <div
    v-for="image in imageFiles"
    :key="image.path"
    :style="{ height: thumbnailSize + 'px' }"
    @click="selectImage(image)"
    :class="[
      'flex items-center p-3 rounded-lg cursor-pointer transition-colors',
      selectedImage?.path === image.path
        ? 'bg-blue-50 border border-blue-200'
        : 'bg-white hover:bg-gray-50',
    ]"
  >
    <div
      class="h-full rounded-lg overflow-hidden bg-gray-100 flex-shrink-0 mr-3"
    >
      <Image
        :src="'/files/thumbnail/250x250/' + image.path"
        :alt="image.name"
        :height="thumbnailSize"
        @error="onThumbnailError"
      />
    </div>
    <div class="flex-1 min-w-0">
      <div class="text-sm font-medium text-gray-800 truncate">
        {{ image.name }}
      </div>
      <div class="text-xs text-gray-500">
        {{ image.dimensions }} • {{ formatFileSize(image.size) }}
      </div>
    </div>
    <div class="flex items-center space-x-3">
      <div
        v-if="image.hasMetadata"
        class="w-2 h-2 bg-green-400 rounded-full"
        title="有元数据"
      ></div>
      <div
        v-else
        class="w-2 h-2 bg-gray-400 rounded-full"
        title="无元数据"
      ></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useImageStore } from "@/stores/image";
import { storeToRefs } from "pinia";
import Image from "@/components/Image.vue";

const imageStore = useImageStore();

const { imageFiles, selectedImage, thumbnailSize } = storeToRefs(imageStore);
const { formatFileSize, selectImage, onThumbnailError } = useImageStore();
</script>
