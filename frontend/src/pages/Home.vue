<template>
  <div class="h-screen flex flex-col bg-gray-50">
    <!-- 头部 -->
    <Header></Header>

    <!-- 主内容区域 -->
    <div class="flex-1 flex overflow-hidden">
      <!-- 左侧-->
      <div class="w-80 bg-white border-r border-gray-200 flex flex-col">
        <!-- 资源管理器 -->
        <div class="flex-1 overflow-hidden"><Explorer></Explorer></div>

        <div class="border border-gray-200 h-[1px]"></div>

        <!-- 选择的图片预览 -->
        <div class="h-80"><Previewer></Previewer></div>
      </div>

      <!-- 中间 主要显示区域 -->
      <div class="flex-1 flex flex-col bg-gray-50">
        <!-- 工具栏 -->
        <div
          class="bg-white border-b border-gray-200 px-4 py-3 flex items-center justify-between"
        >
          <ViewToolBar></ViewToolBar>
        </div>

        <!-- 缩略图网格 -->
        <div class="flex-1 p-4 overflow-y-auto">
          <!-- 网格视图 -->
          <div v-if="viewMode === 'grid'" class="grid gap-4" :style="gridStyle">
            <GridView></GridView>
          </div>

          <!-- 列表视图 -->
          <div v-else class="space-y-2">
            <ListView></ListView>
          </div>

          <!-- 空状态 -->
          <div
            v-if="imageFiles.length === 0"
            class="flex-1 flex items-center justify-center"
          >
            <EmptyView> </EmptyView>
          </div>
        </div>
      </div>

      <!-- 右侧：元数据编辑区 -->
      <div class="w-96 bg-white border-l border-gray-200 flex flex-col">
        <MetaEditor></MetaEditor>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useImageStore } from "@/stores/image";
import { storeToRefs } from "pinia";
import { computed } from "vue";

import Header from "@/components/Header.vue";

import RecentVisited from "@/components/RecentVisited.vue";
import Explorer from "@/components/Explorer.vue";
import Previewer from "@/components/Previewer.vue";

import ViewToolBar from "@/components/ViewToolBar.vue";
import GridView from "@/components/GridView.vue";
import ListView from "@/components/ListView.vue";
import EmptyView from "@/components/EmptyView.vue";

import MetaEditor from "@/components/MetaEditor.vue";

const imageStore = useImageStore();

const { imageFiles, viewMode, thumbnailSize } = storeToRefs(imageStore);

const gridStyle = computed(() => ({
  gridTemplateColumns: `repeat(auto-fill, minmax(${
    thumbnailSize.value + 80
  }px, 1fr))`,
}));
</script>
