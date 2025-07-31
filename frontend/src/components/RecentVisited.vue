<template>
  <div class="p-4 border-b border-gray-100">
    <h3 class="text-sm font-medium text-gray-700 mb-3">最近访问</h3>
    <div class="space-y-2">
      <div
        v-for="folder in recentFolders"
        :key="folder"
        @click="selectFolder(folder)"
        class="p-2 rounded-lg hover:bg-gray-50 cursor-pointer text-sm text-gray-600 truncate"
      >
        {{ folder }}
      </div>
      <div v-if="recentFolders.length === 0" class="text-xs text-gray-400">
        暂无最近访问
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useImageStore } from "@/stores/image";
import { storeToRefs } from "pinia";

const imageStore = useImageStore();

const { recentFolders } = storeToRefs(imageStore);
const { selectFolder } = useImageStore();
</script>
