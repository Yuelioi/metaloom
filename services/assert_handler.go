package services

import (
	"fmt"
	"image"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/vp8"
	_ "golang.org/x/image/vp8l"
	_ "golang.org/x/image/webp"

	"github.com/disintegration/imaging"
)

// ImageHandler 处理图片资产请求
type ImageHandler struct {
	http.Handler
	cacheDir string // 缩略图缓存目录
}

type Config struct {
	RootPath string
}

// NewImageHandler 创建新的图片处理器
func NewImageHandler(cacheDir string) *ImageHandler {
	// 确保缓存目录存在
	os.MkdirAll(cacheDir, 0755)

	return &ImageHandler{
		cacheDir: cacheDir,
	}
}

// ServeHTTP 处理HTTP请求
func (h *ImageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// 只允许GET请求
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析请求路径
	requestPath := r.URL.Path

	// 处理不同类型的请求
	if strings.HasPrefix(requestPath, "/image") {
		h.handleImageAPI(w, r, requestPath)
	} else if strings.HasPrefix(requestPath, "/thumbnail") {
		h.handleThumbnailAPI(w, r, requestPath)
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func (h *ImageHandler) handleImageAPI(w http.ResponseWriter, r *http.Request, requestPath string) {
	imagePath := strings.TrimPrefix(requestPath, "/image/")
	if err := h.serveImage(w, imagePath, 0, 0); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// handleThumbnailAPI 处理缩略图请求
func (h *ImageHandler) handleThumbnailAPI(w http.ResponseWriter, r *http.Request, requestPath string) {
	parts := strings.SplitN(strings.TrimPrefix(requestPath, "/thumbnail/"), "/", 2)
	if len(parts) != 2 {
		http.Error(w, "Invalid thumbnail request format", http.StatusBadRequest)
		return
	}
	sizeStr, imagePath := parts[0], parts[1]

	width, height, err := h.parseSize(sizeStr)
	if err != nil {
		http.Error(w, "Invalid size parameter", http.StatusBadRequest)
		return
	}

	if err := h.serveImage(w, imagePath, width, height); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// serveImage 提供原图或缩略图（width/height > 0 时生成缩略图）
func (h *ImageHandler) serveImage(w http.ResponseWriter, imagePath string, width, height int) error {
	// URL 解码
	decodedPath, err := url.QueryUnescape(imagePath)
	if err != nil {
		return fmt.Errorf("Invalid path encoding")
	}

	// 验证和清理路径
	cleanPath, err := h.validateAndCleanPath(decodedPath)
	if err != nil {
		return err
	}

	// 检查文件存在性与类型
	if !h.fileExists(cleanPath) {
		return fmt.Errorf("Image not found")
	}
	if !h.isImageFile(cleanPath) {
		return fmt.Errorf("Not an image file")
	}

	// 获取真实文件路径（原图或缩略图）
	targetPath := cleanPath
	if width > 0 && height > 0 {
		targetPath, err = h.getThumbnail(cleanPath, width, height)
		if err != nil {
			return fmt.Errorf("Failed to generate thumbnail")
		}
	}

	// 设置响应头
	h.setImageHeaders(w, targetPath)

	// 打开文件并复制到响应中
	file, err := os.Open(targetPath)
	if err != nil {
		return fmt.Errorf("Failed to open file")
	}
	defer file.Close()

	_, err = io.Copy(w, file)
	if err != nil {
		return fmt.Errorf("Failed to send file")
	}

	return nil
}

// validateAndCleanPath 验证并清理文件路径
func (h *ImageHandler) validateAndCleanPath(path string) (string, error) {

	// 清理路径，防止目录遍历攻击
	cleanPath := filepath.Clean(path)

	// 转换为绝对路径
	absPath, err := filepath.Abs(cleanPath)
	if err != nil {
		fmt.Printf("Failed to get absolute path: %v\n", err)
		return "", fmt.Errorf("invalid path")
	}

	return absPath, nil
}

// fileExists 检查文件是否存在
func (h *ImageHandler) fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// isImageFile 检查是否为图片文件
func (h *ImageHandler) isImageFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	imageExts := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".tiff", ".tif"}

	for _, imageExt := range imageExts {
		if ext == imageExt {
			return true
		}
	}
	return false
}

// setImageHeaders 设置图片响应头
func (h *ImageHandler) setImageHeaders(w http.ResponseWriter, path string) {
	// 设置MIME类型
	ext := filepath.Ext(path)
	mimeType := mime.TypeByExtension(ext)
	if mimeType != "" {
		w.Header().Set("Content-Type", mimeType)
	}

	// 设置缓存头
	w.Header().Set("Cache-Control", "public, max-age=3600")
	w.Header().Set("Last-Modified", time.Now().Format(http.TimeFormat))
}

// parseSize 解析尺寸参数
func (h *ImageHandler) parseSize(sizeStr string) (int, int, error) {
	// 支持格式: "200x200", "200", "200x" (等比缩放)
	if strings.Contains(sizeStr, "x") {
		parts := strings.Split(sizeStr, "x")
		if len(parts) != 2 {
			return 0, 0, fmt.Errorf("invalid size format")
		}

		var width, height int
		var err error

		if parts[0] != "" {
			width, err = strconv.Atoi(parts[0])
			if err != nil {
				return 0, 0, err
			}
		}

		if parts[1] != "" {
			height, err = strconv.Atoi(parts[1])
			if err != nil {
				return 0, 0, err
			}
		}

		return width, height, nil
	} else {
		// 只指定一个尺寸，生成正方形缩略图
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			return 0, 0, err
		}
		return size, size, nil
	}
}

// getThumbnail 获取或生成缩略图
func (h *ImageHandler) getThumbnail(imagePath string, width, height int) (string, error) {
	// 生成缩略图文件名
	hash := h.generateHash(imagePath, width, height)
	thumbnailPath := filepath.Join(h.cacheDir, fmt.Sprintf("%s.jpg", hash))

	// 检查缓存是否存在且是最新的
	if h.isThumbnailValid(thumbnailPath, imagePath) {
		return thumbnailPath, nil
	}

	// 生成新的缩略图
	err := h.generateThumbnail(imagePath, thumbnailPath, width, height)
	if err != nil {
		return "", err
	}

	return thumbnailPath, nil
}

// generateHash 生成缩略图哈希名称
func (h *ImageHandler) generateHash(imagePath string, width, height int) string {
	// 简单的哈希算法，可以使用更复杂的如MD5
	return fmt.Sprintf("%x_%dx%d",
		strings.Replace(imagePath, string(os.PathSeparator), "_", -1),
		width, height)
}

// isThumbnailValid 检查缩略图是否有效
func (h *ImageHandler) isThumbnailValid(thumbnailPath, originalPath string) bool {
	thumbStat, err := os.Stat(thumbnailPath)
	if err != nil {
		return false
	}

	originalStat, err := os.Stat(originalPath)
	if err != nil {
		return false
	}

	// 如果原图更新时间比缩略图新，则缩略图无效
	return originalStat.ModTime().Before(thumbStat.ModTime())
}

// generateThumbnail 生成缩略图
func (h *ImageHandler) generateThumbnail(imagePath, thumbnailPath string, width, height int) error {
	// 打开原图
	src, err := imaging.Open(imagePath)
	if err != nil {
		return err
	}

	// 生成缩略图
	var thumbnail image.Image
	if width > 0 && height > 0 {
		// 指定了宽高，进行缩放并裁剪
		thumbnail = imaging.Fill(src, width, height, imaging.Center, imaging.Lanczos)
	} else if width > 0 {
		// 只指定宽度，等比缩放
		thumbnail = imaging.Resize(src, width, 0, imaging.Lanczos)
	} else if height > 0 {
		// 只指定高度，等比缩放
		thumbnail = imaging.Resize(src, 0, height, imaging.Lanczos)
	} else {
		return fmt.Errorf("invalid dimensions")
	}

	// 保存缩略图
	return imaging.Save(thumbnail, thumbnailPath)
}
