// 接口定义
export interface ImageFile {
  path: string;
  name: string;
  size: number;
  thumbnail: string;
  hasMetadata: boolean;
  dimensions: string;
}

export interface Metadata {
  title: string;
  description: string;
  tags: string[];
  customFields: Array<{ key: string; value: string }>;
}

export interface RecentFolder {
  path: string;
  name: string;
}
