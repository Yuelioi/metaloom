<template>
  <div class="p-4 h-full">
    <div class="flex flex-col h-full">
      <h3 class="text-sm font-medium text-gray-700 mb-3">文件资源管理器</h3>

      <div class="flex-1 overflow-y-auto overflow-x-hidden">
        <div class="py-2">
          <ExplorerTreeNode
            v-for="item in explorerItems"
            :key="item.path"
            :item="item"
            :level="0"
            :selected-path="selectedPath"
            @item-click="handleItemClick"
            @toggle-folder="toggleFolder"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref, watch } from "vue";
import ExplorerTreeNode from "@/components/ExplorerTreeNode.vue";

import { App, NavigationItem } from "@app";

import { useImageStore } from "@/stores/image";
import { storeToRefs } from "pinia";

const imageStore = useImageStore();

const { selectedImage, imageFiles, currentFolder } = storeToRefs(imageStore);

const explorerItems = ref<NavigationItem[]>([]);
const selectedPath = ref<string>("");

// 切换文件夹展开/收起
const toggleFolder = async (item: NavigationItem) => {
  item.expanded = !item.expanded;

  // 如果是展开操作且没有子项目，则请求加载
  if (item.expanded && !item.children) {
    App.ShowExplorerItems(item.path).then((data) => {
      item.children = data;
    });
  }
};

const findItemByPath = (
  items: NavigationItem[],
  targetPath: string
): NavigationItem | null => {
  for (const item of items) {
    if (item.path === targetPath) {
      return item;
    }
    if (item.children) {
      const found = findItemByPath(item.children, targetPath);
      if (found) return found;
    }
  }
  return null;
};

// 处理项目点击
const handleItemClick = (item: NavigationItem) => {
  selectedPath.value = item.path;

  // 如果是文件，则打开预览
  if (item.type === "file") {
    App.ShowImage(item.path).then((data) => {
      selectedImage.value = data;
    });
  } else {
    currentFolder.value = item.path;
  }
};

watch(currentFolder, (newVal) => {
  App.ShowFolders(currentFolder.value).then((data) => {
    imageFiles.value = data.images;
  });
});

onMounted(async () => {
  App.GetExplorer().then((data) => {
    explorerItems.value = data;
  });
});
</script>

<style scoped>
/* 自定义滚动条样式 */
</style>
