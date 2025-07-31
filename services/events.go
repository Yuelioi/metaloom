package services

import (
	"context"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v3/pkg/application"
)

var imageExtensions = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true,
	".svg": true, ".ico": true, ".bmp": true, ".tiff": true, ".tif": true, ".avif": true,
	".apng": true, ".flif": true,
	".cr2": true, ".nef": true, ".arw": true, ".orf": true, ".raf": true,
	".rw2": true, ".dng": true, ".sr2": true, ".pef": true, ".raw": true,
	".hdr": true, ".exr": true,
}

func (a *App) OpenFolderDialog(ctx context.Context) string {
	paths, err := application.OpenFileDialog().
		CanCreateDirectories(false).
		CanChooseFiles(false).
		CanChooseDirectories(true).
		AllowsOtherFileTypes(true).
		ShowHiddenFiles(false).
		SetTitle("选择文件夹").PromptForMultipleSelection()
	if err != nil {
		return "There was an error!"
	}

	return paths[0]
}

func (a *App) GetExplorer(ctx context.Context) []NavigationItem {
	var drives []NavigationItem

	for _, folder := range []string{"Documents", "Downloads", "Desktop", "Pictures"} {
		folderPath := os.Getenv("USERPROFILE") + "\\" + folder
		if _, err := os.Stat(folderPath); err == nil {
			drives = append(drives, NavigationItem{
				Name: folder,
				Path: folderPath,
				Type: "folder",
				Icon: "📁",
			})
		}
	}

	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		drivePath := string(drive) + ":\\"
		if _, err := os.Stat(drivePath); err == nil {
			driveName := string(drive) + ": 盘"
			drives = append(drives, NavigationItem{
				Name: driveName,
				Path: drivePath,
				Type: "drive",
				Icon: "💾",
			})
		}
	}

	return drives
}

func (a *App) ShowFolders(path string) DirectoryContent {
	var content DirectoryContent

	files, err := os.ReadDir(path)
	if err != nil {
		log.Println("读取文件夹失败:", err)
		return content
	}

	for _, file := range files {
		fullPath := filepath.Join(path, file.Name())

		if file.IsDir() {
			content.Folders = append(content.Folders, Folder{
				Name: file.Name(),
				Path: fullPath,
			})
		} else {
			// 检查是否是图片文件
			ext := strings.ToLower(filepath.Ext(file.Name()))
			if imageExtensions[ext] {
				// 获取文件信息
				info, err := file.Info()
				if err != nil {
					continue
				}

				// 检查是否存在同名 .xmp 文件
				baseName := strings.TrimSuffix(file.Name(), ext)
				xmpPath := filepath.Join(path, baseName+".xmp")
				hasMetadata := false
				if _, err := os.Stat(xmpPath); err == nil {
					hasMetadata = true
				}

				// 获取图片尺寸
				dimensions := ""
				if dim, err := getImageDimensions(fullPath); err == nil {
					dimensions = dim
				}

				// 构建图片信息
				image := Image{
					Name:        file.Name(),
					Path:        fullPath,
					Size:        info.Size(),
					Thumbnail:   "",
					HasMetadata: hasMetadata,
					Dimensions:  dimensions,
				}

				content.Images = append(content.Images, image)
			}
		}
	}

	return content
}

func (a *App) ShowImage(path string) Image {
	var image Image

	// 判断文件是否存在
	info, err := os.Stat(path)
	if err != nil || info.IsDir() {
		return image
	}

	// 判断扩展名是否是图片
	ext := strings.ToLower(filepath.Ext(path))
	if !imageExtensions[ext] {
		return image
	}

	// 检查是否存在同名 .xmp 文件
	baseName := strings.TrimSuffix(filepath.Base(path), ext)
	dir := filepath.Dir(path)
	xmpPath := filepath.Join(dir, baseName+".xmp")

	hasMetadata := false
	if _, err := os.Stat(xmpPath); err == nil {
		hasMetadata = true
	}

	// 获取图片尺寸
	dimensions := ""
	if dim, err := getImageDimensions(path); err == nil {
		dimensions = dim
	}

	// 构建返回值
	image = Image{
		Name:        filepath.Base(path),
		Path:        path,
		Size:        info.Size(),
		Thumbnail:   "",
		HasMetadata: hasMetadata,
		Dimensions:  dimensions,
	}

	return image
}

func getImageDimensions(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	img, _, err := image.DecodeConfig(f)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%dx%d", img.Width, img.Height), nil
}

// 浏览器切换子文件夹
func (a *App) ShowExplorerItems(path string) []NavigationItem {
	var items []NavigationItem

	files, err := os.ReadDir(path)
	if err != nil {
		log.Println("读取文件夹失败:", err)
		return items
	}

	for _, file := range files {
		fullPath := filepath.Join(path, file.Name())

		if file.IsDir() {
			items = append(items, NavigationItem{
				Name: file.Name(),
				Path: fullPath,
				Type: "folder",
				Icon: "📁",
			})
		} else {
			ext := strings.ToLower(filepath.Ext(file.Name()))
			if imageExtensions[ext] {
				items = append(items, NavigationItem{
					Name: file.Name(),
					Path: fullPath,
					Type: "file",
					Icon: "🖼️",
				})
			}
		}
	}

	return items
}
