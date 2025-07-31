package services

type NavigationItem struct {
	Name     string           `json:"name"`
	Path     string           `json:"path"`
	Type     string           `json:"type"` // "drive", "folder"
	Icon     string           `json:"icon"`
	Children []NavigationItem `json:"children,omitempty"`
	Expanded bool             `json:"expanded"`
}

type Image struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Size        int64  `json:"size"`        // 字节数
	Thumbnail   string `json:"thumbnail"`   // 缩略图
	HasMetadata bool   `json:"hasMetadata"` // 是否有元数据
	Dimensions  string `json:"dimensions"`  // 图片尺寸
}

type Folder struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type DirectoryContent struct {
	Folders []Folder `json:"folders"`
	Images  []Image  `json:"images"`
}

type Field struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Metadata struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Fields      []Field  `json:"fields"`
}
