<template>
  <div
    class="relative bg-gray-100 flex items-center justify-center"
    :style="{ height: height + 'px' }"
    ref="container"
  >
    <!-- 懒加载图片 -->
    <img
      :data-src="src"
      :src="shouldLoad ? src : placeholder"
      :alt="alt"
      class="max-w-full max-h-full object-cover transition-opacity duration-300"
      :class="{ 'opacity-0': !loaded }"
      @load="onLoad"
      @error="$emit('error')"
      loading="lazy"
      decoding="async"
    />

    <!-- 加载占位符 -->
    <div
      v-if="!loaded"
      class="absolute inset-0 flex items-center justify-center bg-gray-100"
    >
      <slot name="placeholder">
        <svg
          class="w-8 h-8 text-gray-400 animate-pulse"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="1"
            d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"
          />
        </svg>
      </slot>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from "vue";

interface Props {
  src: string;
  alt?: string;
  height?: number;
}

const props = withDefaults(defineProps<Props>(), {
  alt: "",
  height: 120,
});

const loaded = ref(false);
const shouldLoad = ref(false);
const container = ref<HTMLElement | null>(null);
let observer: IntersectionObserver | null = null;

const placeholder =
  "data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7";

const onLoad = () => {
  loaded.value = true;
};

onMounted(() => {
  if (!("IntersectionObserver" in window)) {
    shouldLoad.value = true;
    return;
  }

  observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          shouldLoad.value = true;
          observer?.disconnect();
        }
      });
    },
    {
      rootMargin: "100px",
      threshold: 0.01,
    }
  );

  if (container.value) {
    observer.observe(container.value);
  }
});

onUnmounted(() => {
  observer?.disconnect();
});
</script>
