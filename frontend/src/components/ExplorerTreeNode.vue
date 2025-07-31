<template>
  <div class="select-none">
    <div
      class="cursor-pointer transition-colors duration-150 relative text-sm flex items-center min-h-8 px-3 py-1 hover:bg-gray-200"
      :class="[
        selectedPath === item.path
          ? 'bg-blue-600 text-white hover:bg-blue-700'
          : '',
        { 'pl-3': level === 0, [`pl-${3 + level * 5}`]: level > 0 },
      ]"
      @click="$emit('item-click', item)"
    >
      <span
        v-if="item.type === 'folder' || item.type === 'drive'"
        class="w-4 h-4 inline-flex items-center justify-center mr-1 text-xs text-gray-600 transition-transform duration-150 hover:bg-black hover:bg-opacity-10 rounded-sm"
        :class="[
          selectedPath === item.path
            ? 'text-gray-200 hover:bg-white hover:bg-opacity-20'
            : '',
          item.expanded ? 'transform rotate-0' : '',
        ]"
        @click.stop="$emit('toggle-folder', item)"
      >
        {{ item.expanded ? "▼" : "▶" }}
      </span>
      <span class="mr-2 text-sm w-5 text-center">{{ item.icon }}</span>
      <span class="flex-1 whitespace-nowrap overflow-hidden text-ellipsis">
        {{ item.name }}
      </span>
    </div>

    <!-- 递归渲染子节点 -->
    <div
      v-if="item.expanded && item.children"
      class="border-l border-gray-300 ml-5"
    >
      <ExplorerTreeNode
        v-for="child in item.children"
        :key="child.path"
        :item="child"
        :level="level + 1"
        :selected-path="selectedPath"
        @item-click="$emit('item-click', $event)"
        @toggle-folder="$emit('toggle-folder', $event)"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { NavigationItem } from "@app";

defineProps<{
  item: NavigationItem;
  level: number;
  selectedPath: string;
}>();

defineEmits<{
  "item-click": [item: NavigationItem];
  "toggle-folder": [item: NavigationItem];
}>();
</script>
