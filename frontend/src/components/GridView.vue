<template>
  <div
    v-for="image in imageFiles"
    :key="image.path"
    @click="selectImage(image)"
    @dblclick="openImagePreview(image)"
    :class="[
      'group cursor-pointer rounded-lg overflow-hidden bg-white shadow-sm hover:shadow-md transition-all duration-200',
      selectedImage?.path === image.path
        ? 'ring-2 ring-blue-500 shadow-lg'
        : 'hover:ring-1 hover:ring-gray-300',
    ]"
  >
    <!-- 缩略图 -->
    <div
      class="relative bg-gray-100 flex items-center justify-center"
      :style="{ height: thumbnailSize + 'px' }"
      :ref="(el) => observeImage(el as HTMLElement, image.path)"
    >
      <!-- 懒加载图片 -->

      <Image
        :src="'/files/thumbnail/250x250/' + image.path"
        :alt="image.name"
        :height="thumbnailSize"
        @error="onThumbnailError"
      />

      <!-- 选中状态覆盖层 -->
      <div
        v-if="selectedImage?.path === image.path"
        class="absolute inset-0 bg-blue-500 bg-opacity-20 flex items-center justify-center"
      >
        <div
          class="w-8 h-8 bg-blue-500 rounded-full flex items-center justify-center"
        >
          <svg
            class="w-5 h-5 text-white"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M5 13l4 4L19 7"
            ></path>
          </svg>
        </div>
      </div>

      <!-- 悬停状态操作按钮 -->
      <div
        class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity duration-200"
      >
        <button
          @click.stop="toggleImageSelection(image)"
          class="w-6 h-6 bg-white bg-opacity-90 rounded-full flex items-center justify-center hover:bg-opacity-100 backdrop-blur-sm"
          :title="isImageFavorited(image) ? '取消收藏' : '添加收藏'"
        >
          <svg
            class="w-3 h-3"
            :class="isImageFavorited(image) ? 'text-red-500' : 'text-gray-600'"
            fill="currentColor"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"
            ></path>
          </svg>
        </button>
      </div>

      <!-- 元数据状态指示器 -->
      <div class="absolute bottom-2 left-2">
        <div
          v-if="image.hasMetadata"
          class="w-3 h-3 bg-green-400 rounded-full border-2 border-white shadow-sm"
          title="有元数据"
        ></div>
        <div
          v-else
          class="w-3 h-3 bg-gray-400 rounded-full border-2 border-white shadow-sm"
          title="无元数据"
        ></div>
      </div>

      <!-- 图片类型标识 -->
      <div
        v-if="getImageType(image.name)"
        class="absolute top-2 left-2 px-1.5 py-0.5 bg-black bg-opacity-60 text-white text-xs rounded backdrop-blur-sm"
      >
        {{ getImageType(image.name) }}
      </div>
    </div>

    <!-- 图片信息 -->
    <div class="p-3">
      <div
        class="text-sm font-medium text-gray-800 truncate mb-1"
        :title="image.name"
      >
        {{ image.name }}
      </div>
      <div class="flex items-center justify-between text-xs text-gray-500">
        <span>{{ image.dimensions }}</span>
        <span>{{ formatFileSize(image.size) }}</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, onMounted, onUnmounted } from "vue";
import { useImageStore } from "@/stores/image";
import { storeToRefs } from "pinia";

import Image from "@/components/Image.vue";

const imageStore = useImageStore();

const { imageFiles, selectedImage, thumbnailSize } = storeToRefs(imageStore);
const {
  formatFileSize,
  selectImage,
  toggleImageSelection,
  onThumbnailError,
  openImagePreview,
} = useImageStore();

// 图片加载状态
const visibleImages = reactive<Set<string>>(new Set());

// Intersection Observer 实例
let observer: IntersectionObserver | null = null;

// 初始化 Intersection Observer
const initObserver = () => {
  if (typeof IntersectionObserver === "undefined") {
    // 不支持 IntersectionObserver，回退到全部加载
    imageFiles.value.forEach((image) => {
      visibleImages.add(image.path);
    });
    return;
  }

  observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        const imagePath = entry.target.getAttribute("data-image-path");
        if (!imagePath) return;

        if (entry.isIntersecting) {
          visibleImages.add(imagePath);
          // 一旦加载就不再观察
          observer?.unobserve(entry.target);
        }
      });
    },
    {
      // 提前加载视口外 100px 的图片
      rootMargin: "100px",
      threshold: 0.01,
    }
  );
};

// 观察图片元素
const observeImage = (el: HTMLElement, imagePath: string) => {
  if (observer && el) {
    el.setAttribute("data-image-path", imagePath);
    observer.observe(el);
  }
};

// 检查图片是否被收藏（示例功能）
const isImageFavorited = (image: any) => {
  // 这里可以实现收藏逻辑
  return false;
};

// 获取图片类型
const getImageType = (filename: string) => {
  const ext = filename.split(".").pop()?.toLowerCase();
  const typeMap: Record<string, string> = {
    jpg: "JPG",
    jpeg: "JPG",
    png: "PNG",
    gif: "GIF",
    webp: "WEBP",
    svg: "SVG",
    bmp: "BMP",
    tiff: "TIFF",
    tif: "TIFF",
  };
  return ext ? typeMap[ext] : null;
};

onMounted(() => {
  initObserver();
});

onUnmounted(() => {
  if (observer) {
    observer.disconnect();
  }
});
</script>
