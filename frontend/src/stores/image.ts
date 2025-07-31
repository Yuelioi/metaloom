import { defineStore } from "pinia";
import { ref, reactive } from "vue";
import { Metadata, Image } from "@app";

export const useImageStore = defineStore("image", () => {
  const selectedImage = ref<Image | null>(null);
  const imageFiles = ref<Image[]>([]);
  const viewMode = ref<"grid" | "list">("grid");
  const thumbnailSize = ref<number>(120);
  const newTag = ref<string>("");

  const metadata = reactive<Metadata>({
    title: "",
    description: "",
    tags: [],
    fields: [],
  });

  const recentFolders = ref<string[]>([]);
  const currentFolder = ref("");

  const selectFolder = (folder: string) => {
    currentFolder.value = folder;
  };

  const selectImage = (image: Image) => {
    selectedImage.value = image;
    loadMetadata(image);
  };

  const toggleImageSelection = (image: Image) => {
    console.log("切换选中状态:", image.name);
  };

  const openImagePreview = (image: Image) => {
    console.log("打开图片预览:", image.name);
  };

  const loadMetadata = (image: Image) => {
    if (image.hasMetadata) {
      metadata.title = "Sample Title";
      metadata.description = "This is a sample description";
      metadata.tags = ["sample", "test", "photo"];
      metadata.fields = [];
    } else {
      resetMetadata();
    }
  };

  const saveMetadata = () => {
    console.log("保存元数据:", metadata);
  };

  const createMetadata = () => {
    if (selectedImage.value) {
      selectedImage.value.hasMetadata = true;
      resetMetadata();
    }
  };

  const resetMetadata = () => {
    metadata.title = "";
    metadata.description = "";
    metadata.tags = [];
    metadata.fields = [];
  };

  const addTag = () => {
    const tag = newTag.value.trim();
    if (tag && !metadata.tags.includes(tag)) {
      metadata.tags.push(tag);
      newTag.value = "";
    }
  };

  const removeTag = (index: number) => {
    metadata.tags.splice(index, 1);
  };

  const addCustomField = () => {
    // metadata.customFields.push({ key: "", value: "" });
  };

  const removeCustomField = (index: number) => {
    metadata.fields.splice(index, 1);
  };

  const formatFileSize = (bytes: number): string => {
    if (bytes === 0) return "0 Bytes";
    const k = 1024;
    const sizes = ["Bytes", "KB", "MB", "GB"];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + " " + sizes[i];
  };

  const onThumbnailError = (event: Event) => {
    const img = event.target as HTMLImageElement;
    img.src =
      "data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMTIwIiBoZWlnaHQ9IjEyMCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48cmVjdCB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIiBmaWxsPSIjZjNmNGY2Ii8+PHRleHQgeD0iNTAlIiB5PSI1MCUiIGZvbnQtZmFtaWx5PSJBcmlhbCIgZm9udC1zaXplPSIxMiIgZmlsbD0iIzk5OSIgdGV4dC1hbmNob3I9Im1pZGRsZSIgZHk9Ii4zZW0iPkVSUk9SPC90ZXh0Pjwvc3ZnPg==";
  };

  return {
    selectedImage,
    imageFiles,
    viewMode,
    thumbnailSize,
    newTag,
    metadata,
    currentFolder,
    recentFolders,
    selectFolder,
    selectImage,
    toggleImageSelection,
    openImagePreview,
    loadMetadata,
    saveMetadata,
    createMetadata,
    resetMetadata,
    addTag,
    removeTag,
    addCustomField,
    removeCustomField,
    formatFileSize,
    onThumbnailError,
  };
});
