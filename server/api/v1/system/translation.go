package system

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"hz-admin-base/global"
	"hz-admin-base/model/common/response"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type TranslationApi struct{}

// 翻译内容结构
type TranslationContent struct {
	Language string                            `json:"language"`
	Files    map[string]map[string]interface{} `json:"files"`
}

// 文件树节点
type FileNode struct {
	Title    string     `json:"title"`    // 显示名称
	Key      string     `json:"key"`      // 唯一标识
	Children []FileNode `json:"children"` // 子节点
	IsLeaf   bool       `json:"isLeaf"`   // 是否为叶子节点
	Path     string     `json:"path"`     // 文件路径
}

// 更新翻译请求
type UpdateTranslationRequest struct {
	Language string `json:"language" binding:"required"`
	File     string `json:"file" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

// 复制语言包请求
type CopyLanguageRequest struct {
	FromLang string `json:"fromLang" binding:"required"`
	ToLang   string `json:"toLang" binding:"required"`
}

// 对比翻译请求
type CompareTranslationRequest struct {
	SourceLang string `json:"sourceLang" binding:"required"`
	TargetLang string `json:"targetLang" binding:"required"`
	File       string `json:"file" binding:"required"`
}

// 对比结果
type CompareResult struct {
	MissingKeys   []string               `json:"missingKeys"`   // 缺失的键
	DifferentKeys []string               `json:"differentKeys"` // 值不同的键
	SourceContent map[string]interface{} `json:"sourceContent"` // 源语言内容
	TargetContent map[string]interface{} `json:"targetContent"` // 目标语言内容
}

// 批量更新请求
type BatchUpdateRequest struct {
	Updates []struct {
		Language string `json:"language" binding:"required"`
		File     string `json:"file" binding:"required"`
		Content  string `json:"content" binding:"required"`
	} `json:"updates" binding:"required"`
}

// 完整性检查结果
type IntegrityCheckResult struct {
	Language    string   `json:"language"`
	File        string   `json:"file"`
	MissingKeys []string `json:"missingKeys"`
	ExtraKeys   []string `json:"extraKeys"`
	EmptyValues []string `json:"emptyValues"`
	Coverage    float64  `json:"coverage"`
}

// GetFileTree 获取翻译文件树结构
func (t *TranslationApi) GetFileTree(ctx *gin.Context) {
	dir := global.GVA_CONFIG.System.TranslationDir

	tree := make([]FileNode, 0)
	languages, err := os.ReadDir(dir)
	if err != nil {
		response.FailWithMessage("Failed to read translation directory", ctx)
		return
	}

	// 递归构建文件树
	var buildFileTree func(path string, parentPath string) ([]FileNode, error)
	buildFileTree = func(path string, parentPath string) ([]FileNode, error) {
		entries, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}

		nodes := make([]FileNode, 0)
		for _, entry := range entries {
			relativePath := filepath.Join(parentPath, entry.Name())
			fullPath := filepath.Join(path, entry.Name())

			if entry.IsDir() {
				// 处理子目录
				children, err := buildFileTree(fullPath, relativePath)
				if err != nil {
					continue
				}
				nodes = append(nodes, FileNode{
					Title:    entry.Name(),
					Key:      relativePath,
					Children: children,
					IsLeaf:   false,
					Path:     relativePath,
				})
			} else if filepath.Ext(entry.Name()) == ".json" {
				// 处理 JSON 文件
				nodes = append(nodes, FileNode{
					Title:  entry.Name(),
					Key:    relativePath,
					IsLeaf: true,
					Path:   relativePath,
				})
			}
		}
		return nodes, nil
	}

	// 遍历语言目录
	for _, lang := range languages {
		if !lang.IsDir() {
			continue
		}

		// 构建语言节点的子树
		children, err := buildFileTree(
			filepath.Join(dir, lang.Name()),
			lang.Name(),
		)
		if err != nil {
			continue
		}

		langNode := FileNode{
			Title:    lang.Name(),
			Key:      lang.Name(),
			Children: children,
			IsLeaf:   false,
			Path:     lang.Name(),
		}

		tree = append(tree, langNode)
	}

	response.OkWithData(tree, ctx)
}

// GetFileContent 获取单个文件内容
func (t *TranslationApi) GetFileContent(ctx *gin.Context) {
	language := ctx.Query("language")
	file := ctx.Query("file")

	if language == "" || file == "" {
		response.FailWithMessage("Language and file parameters are required", ctx)
		return
	}

	dir := global.GVA_CONFIG.System.TranslationDir
	filePath := filepath.Join(dir, language, file)

	// 读取文件内容
	content, err := os.ReadFile(filePath)
	if err != nil {
		response.FailWithMessage("Failed to read file", ctx)
		return
	}

	// 验证 JSON 格式
	var jsonContent interface{}
	if err := json.Unmarshal(content, &jsonContent); err != nil {
		response.FailWithMessage("Invalid JSON format", ctx)
		return
	}

	response.OkWithData(string(content), ctx)
}

// UpdateTranslation 更新翻译内容
func (t *TranslationApi) UpdateTranslation(ctx *gin.Context) {
	var req UpdateTranslationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Invalid request parameters", ctx)
		return
	}

	dir := global.GVA_CONFIG.System.TranslationDir
	filePath := filepath.Join(dir, req.Language, req.File)

	// 确保目录存在
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		response.FailWithMessage(fmt.Sprintf("Failed to create directory: %v", err), ctx)
		return
	}

	// 验证 JSON 格式
	var jsonContent interface{}
	if err := json.Unmarshal([]byte(req.Content), &jsonContent); err != nil {
		response.FailWithMessage("Invalid JSON format", ctx)
		return
	}

	// 格式化 JSON 内容
	formattedContent, err := json.MarshalIndent(jsonContent, "", "    ")
	if err != nil {
		response.FailWithMessage("Failed to format JSON", ctx)
		return
	}

	// 写入文件
	if err := os.WriteFile(filePath, formattedContent, 0644); err != nil {
		response.FailWithMessage(fmt.Sprintf("Failed to write translation file: %v", err), ctx)
		return
	}

	response.Ok(ctx)
}

// CopyLanguage 复制语言包
func (t *TranslationApi) CopyLanguage(ctx *gin.Context) {
	var req CopyLanguageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Invalid request parameters", ctx)
		return
	}

	dir := global.GVA_CONFIG.System.TranslationDir
	srcPath := filepath.Join(dir, req.FromLang)
	destPath := filepath.Join(dir, req.ToLang)

	// 检查源语言是否存在
	if _, err := os.Stat(srcPath); os.IsNotExist(err) {
		response.FailWithMessage("Source language does not exist", ctx)
		return
	}

	// 检查目标语言是否已存在
	if _, err := os.Stat(destPath); !os.IsNotExist(err) {
		response.FailWithMessage("Target language already exists", ctx)
		return
	}

	// 复制整个目录
	err := filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 计算目标路径
		relPath, err := filepath.Rel(srcPath, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(destPath, relPath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		}

		// 复制文件
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(targetPath, data, info.Mode())
	})

	if err != nil {
		response.FailWithMessage("Failed to copy language pack", ctx)
		return
	}

	response.Ok(ctx)
}

// CompareTranslation 对比翻译内容
func (t *TranslationApi) CompareTranslation(ctx *gin.Context) {
	var req CompareTranslationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Invalid request parameters", ctx)
		return
	}

	dir := global.GVA_CONFIG.System.TranslationDir
	sourceFile := filepath.Join(dir, req.SourceLang, req.File)
	targetFile := filepath.Join(dir, req.TargetLang, req.File)

	// 读取源文件
	sourceContent := make(map[string]interface{})
	if content, err := os.ReadFile(sourceFile); err == nil {
		json.Unmarshal(content, &sourceContent)
	}

	// 读取目标文件
	targetContent := make(map[string]interface{})
	if content, err := os.ReadFile(targetFile); err == nil {
		json.Unmarshal(content, &targetContent)
	}

	// 对比结果
	result := CompareResult{
		MissingKeys:   make([]string, 0),
		DifferentKeys: make([]string, 0),
		SourceContent: sourceContent,
		TargetContent: targetContent,
	}

	// 检查缺失和不同的键
	for key := range sourceContent {
		if _, exists := targetContent[key]; !exists {
			result.MissingKeys = append(result.MissingKeys, key)
		} else if fmt.Sprintf("%v", sourceContent[key]) != fmt.Sprintf("%v", targetContent[key]) {
			result.DifferentKeys = append(result.DifferentKeys, key)
		}
	}

	response.OkWithData(result, ctx)
}

// ExportTranslation 导出翻译文件
func (t *TranslationApi) ExportTranslation(ctx *gin.Context) {
	language := ctx.Query("language")
	if language == "" {
		response.FailWithMessage("Language parameter is required", ctx)
		return
	}

	dir := global.GVA_CONFIG.System.TranslationDir
	langDir := filepath.Join(dir, language)

	// 检查语言目录是否存在
	if _, err := os.Stat(langDir); os.IsNotExist(err) {
		response.FailWithMessage("Language does not exist", ctx)
		return
	}

	// 创建临时目录
	tempDir := filepath.Join(os.TempDir(), "translation_export")
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)

	// 复制文件到临时目录
	err := filepath.Walk(langDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(langDir, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(tempDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(targetPath, data, info.Mode())
	})

	if err != nil {
		response.FailWithMessage("Failed to prepare export files", ctx)
		return
	}

	// 设置响应头
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s.json", language))

	// 读取并返回文件内容
	content, err := os.ReadFile(filepath.Join(tempDir, "translation.json"))
	if err != nil {
		response.FailWithMessage("Failed to read export file", ctx)
		return
	}

	ctx.Data(200, "application/json", content)
}

// ImportTranslation 导入翻译文件
func (t *TranslationApi) ImportTranslation(ctx *gin.Context) {
	language := ctx.PostForm("language")
	if language == "" {
		response.FailWithMessage("Language parameter is required", ctx)
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		response.FailWithMessage("Failed to get upload file", ctx)
		return
	}

	// 读取上传的文件
	uploadedFile, err := file.Open()
	if err != nil {
		response.FailWithMessage("Failed to open upload file", ctx)
		return
	}
	defer uploadedFile.Close()

	content, err := ioutil.ReadAll(uploadedFile)
	if err != nil {
		response.FailWithMessage("Failed to read upload file", ctx)
		return
	}

	// 验证JSON格式
	var translationData map[string]interface{}
	if err := json.Unmarshal(content, &translationData); err != nil {
		response.FailWithMessage("Invalid JSON format", ctx)
		return
	}

	dir := global.GVA_CONFIG.System.TranslationDir
	targetDir := filepath.Join(dir, language)

	// 确保目标目录存在
	os.MkdirAll(targetDir, 0755)

	// 写入文件
	targetFile := filepath.Join(targetDir, file.Filename)
	if err := os.WriteFile(targetFile, content, 0644); err != nil {
		response.FailWithMessage("Failed to write translation file", ctx)
		return
	}

	response.Ok(ctx)
}

// CheckIntegrity 检查翻译完整性
func (t *TranslationApi) CheckIntegrity(ctx *gin.Context) {
	dir := global.GVA_CONFIG.System.TranslationDir
	baseLang := "zh-CN" // 使用中文作为基准语言

	// 获取所有语言
	languages, err := os.ReadDir(dir)
	if err != nil {
		response.FailWithMessage("Failed to read translation directory", ctx)
		return
	}

	results := make([]IntegrityCheckResult, 0)

	// 读取基准语言的所有文件
	baseFiles, err := os.ReadDir(filepath.Join(dir, baseLang))
	if err != nil {
		response.FailWithMessage("Failed to read base language directory", ctx)
		return
	}

	// 遍历所有语言
	for _, lang := range languages {
		if !lang.IsDir() || lang.Name() == baseLang {
			continue
		}

		// 遍历基准语言的所有文件
		for _, baseFile := range baseFiles {
			if !strings.HasSuffix(baseFile.Name(), ".json") {
				continue
			}

			// 读取基准文件
			baseContent := make(map[string]interface{})
			content, err := os.ReadFile(filepath.Join(dir, baseLang, baseFile.Name()))
			if err != nil {
				continue
			}
			json.Unmarshal(content, &baseContent)

			// 读取目标语言文件
			targetContent := make(map[string]interface{})
			content, err = os.ReadFile(filepath.Join(dir, lang.Name(), baseFile.Name()))
			if err == nil {
				json.Unmarshal(content, &targetContent)
			}

			// 检查结果
			result := IntegrityCheckResult{
				Language:    lang.Name(),
				File:        baseFile.Name(),
				MissingKeys: make([]string, 0),
				ExtraKeys:   make([]string, 0),
				EmptyValues: make([]string, 0),
			}

			// 检查缺失的键和空值
			totalKeys := len(baseContent)
			translatedKeys := 0
			for key := range baseContent {
				if val, exists := targetContent[key]; !exists {
					result.MissingKeys = append(result.MissingKeys, key)
				} else {
					if fmt.Sprintf("%v", val) == "" {
						result.EmptyValues = append(result.EmptyValues, key)
					} else {
						translatedKeys++
					}
				}
			}

			// 检查多余的键
			for key := range targetContent {
				if _, exists := baseContent[key]; !exists {
					result.ExtraKeys = append(result.ExtraKeys, key)
				}
			}

			// 计算覆盖率
			if totalKeys > 0 {
				result.Coverage = float64(translatedKeys) / float64(totalKeys) * 100
			}

			results = append(results, result)
		}
	}

	response.OkWithData(results, ctx)
}

// BatchUpdate 批量更新翻译
func (t *TranslationApi) BatchUpdate(ctx *gin.Context) {
	var req BatchUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Invalid request parameters", ctx)
		return
	}

	dir := global.GVA_CONFIG.System.TranslationDir

	for _, update := range req.Updates {
		filePath := filepath.Join(dir, update.Language, update.File)

		// 确保目录存在
		os.MkdirAll(filepath.Dir(filePath), 0755)

		// 验证 JSON 格式
		var jsonContent interface{}
		if err := json.Unmarshal([]byte(update.Content), &jsonContent); err != nil {
			continue
		}

		// 直接写入文件
		os.WriteFile(filePath, []byte(update.Content), 0644)
	}

	response.Ok(ctx)
}

// ExportLanguage 导出语言包
func (t *TranslationApi) ExportLanguage(ctx *gin.Context) {
	var req struct {
		Language string `json:"language" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Invalid request parameters", ctx)
		return
	}

	dir := global.GVA_CONFIG.System.TranslationDir
	langDir := filepath.Join(dir, req.Language)

	// 检查语言目录是否存在
	if _, err := os.Stat(langDir); os.IsNotExist(err) {
		response.FailWithMessage("Language does not exist", ctx)
		return
	}

	// 创建临时目录
	tempDir := filepath.Join(os.TempDir(), "translation_export")
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)

	// 创建 zip 文件
	zipFile := filepath.Join(tempDir, req.Language+".zip")
	archive, err := os.Create(zipFile)
	if err != nil {
		response.FailWithMessage("Failed to create zip file", ctx)
		return
	}
	defer archive.Close()

	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	// 遍历语言目录
	err = filepath.Walk(langDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 获取相对路径
		relPath, err := filepath.Rel(langDir, path)
		if err != nil {
			return err
		}

		// 如果是根目录，跳过
		if relPath == "." {
			return nil
		}

		// 创建 zip 文件头
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// 设置文件名为相对路径
		header.Name = relPath

		// 如果是目录，确保以斜杠结尾
		if info.IsDir() {
			header.Name += "/"
		} else {
			// 如果是 JSON 文件，同时创建一个 JS 版本
			if filepath.Ext(path) == ".json" {
				// 读取 JSON 内容
				content, err := os.ReadFile(path)
				if err != nil {
					return err
				}

				// 创建 JS 文件
				jsHeader := &zip.FileHeader{
					Name:   strings.TrimSuffix(relPath, ".json") + ".js",
					Method: zip.Deflate,
				}
				jsWriter, err := zipWriter.CreateHeader(jsHeader)
				if err != nil {
					return err
				}

				// 写入 JS 内容
				jsContent := fmt.Sprintf("export default %s;", string(content))
				if _, err := jsWriter.Write([]byte(jsContent)); err != nil {
					return err
				}
			}
		}

		// 创建文件写入器
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// 如果是文件，写入内容
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		response.FailWithMessage("Failed to create zip archive", ctx)
		return
	}

	// 关闭 zip 写入器
	if err := zipWriter.Close(); err != nil {
		response.FailWithMessage("Failed to close zip writer", ctx)
		return
	}

	// 读取 zip 文件内容
	zipContent, err := os.ReadFile(zipFile)
	if err != nil {
		response.FailWithMessage("Failed to read zip file", ctx)
		return
	}

	// 设置响应头
	ctx.Header("Content-Type", "application/zip")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s.zip", req.Language))

	// 返回 zip 文件内容
	ctx.Data(200, "application/zip", zipContent)
}

// DeleteLanguage 删除语言包
func (t *TranslationApi) DeleteLanguage(ctx *gin.Context) {
	var req struct {
		Language string `json:"language" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Invalid request parameters", ctx)
		return
	}

	// 不允许删除中文包
	if req.Language == "zh-CN" {
		response.FailWithMessage("Cannot delete Chinese language pack", ctx)
		return
	}

	dir := global.GVA_CONFIG.System.TranslationDir
	langDir := filepath.Join(dir, req.Language)

	// 检查语言目录是否存在
	if _, err := os.Stat(langDir); os.IsNotExist(err) {
		response.FailWithMessage("Language does not exist", ctx)
		return
	}

	// 删除语言目录
	if err := os.RemoveAll(langDir); err != nil {
		response.FailWithMessage("Failed to delete language pack", ctx)
		return
	}

	response.Ok(ctx)
}

// DownloadLanguage 下载语言包
func (t *TranslationApi) DownloadLanguage(ctx *gin.Context) {
	language := ctx.Query("language")
	if language == "" {
		response.FailWithMessage("Language parameter is required", ctx)
		return
	}

	dir := global.GVA_CONFIG.System.TranslationDir
	langDir := filepath.Join(dir, language)

	// 检查语言目录是否存在
	if _, err := os.Stat(langDir); os.IsNotExist(err) {
		response.FailWithMessage("Language does not exist", ctx)
		return
	}

	// 获取所有文件列表
	var files []struct {
		Path    string `json:"path"`
		Content string `json:"content"`
	}

	err := filepath.Walk(langDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 获取相对路径
		relPath, err := filepath.Rel(langDir, path)
		if err != nil {
			return err
		}

		// 如果是根目录，跳过
		if relPath == "." {
			return nil
		}

		// 如果是文件，读取内容
		if !info.IsDir() {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			files = append(files, struct {
				Path    string `json:"path"`
				Content string `json:"content"`
			}{
				Path:    relPath,
				Content: string(content),
			})
		}

		return nil
	})

	if err != nil {
		response.FailWithMessage("Failed to read files", ctx)
		return
	}

	response.OkWithData(files, ctx)
}

// UploadLanguage 上传语言包
func (t *TranslationApi) UploadLanguage(ctx *gin.Context) {
	language := ctx.PostForm("language")
	if language == "" {
		response.FailWithMessage("Language parameter is required", ctx)
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		response.FailWithMessage("Failed to get upload file", ctx)
		return
	}

	// 读取上传的文件
	uploadedFile, err := file.Open()
	if err != nil {
		response.FailWithMessage("Failed to open upload file", ctx)
		return
	}
	defer uploadedFile.Close()

	// 创建临时目录
	tempDir := filepath.Join(os.TempDir(), "translation_upload")
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)

	// 保存上传的文件
	tempZip := filepath.Join(tempDir, file.Filename)
	dst, err := os.Create(tempZip)
	if err != nil {
		response.FailWithMessage("Failed to create temporary file", ctx)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, uploadedFile); err != nil {
		response.FailWithMessage("Failed to save uploaded file", ctx)
		return
	}

	// 解压文件
	reader, err := zip.OpenReader(tempZip)
	if err != nil {
		response.FailWithMessage("Failed to open zip file", ctx)
		return
	}
	defer reader.Close()

	// 获取目标目录
	dir := global.GVA_CONFIG.System.TranslationDir
	targetDir := filepath.Join(dir, language)

	// 解压文件
	for _, f := range reader.File {
		rc, err := f.Open()
		if err != nil {
			continue
		}

		path := filepath.Join(targetDir, f.Name)
		if f.FileInfo().IsDir() {
			continue // 跳过目录
		}

		// 检查目标文件是否存在
		if _, err := os.Stat(path); os.IsNotExist(err) {
			rc.Close()
			continue // 跳过不存在的文件
		}

		// 创建目标文件的目录
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			rc.Close()
			continue
		}

		// 打开目标文件
		dstFile, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			rc.Close()
			continue
		}

		// 复制文件内容
		_, err = io.Copy(dstFile, rc)
		dstFile.Close()
		rc.Close()
	}

	response.Ok(ctx)
}

// GetMessages 获取所有翻译内容
func (t *TranslationApi) GetMessages(ctx *gin.Context) {
	dir := global.GVA_CONFIG.System.TranslationDir
	messages := make(map[string]map[string]interface{})

	// 读取所有语言目录
	languages, err := os.ReadDir(dir)
	if err != nil {
		response.FailWithMessage("Failed to read translation directory", ctx)
		return
	}

	// 遍历每个语言目录
	for _, lang := range languages {
		if !lang.IsDir() {
			continue
		}

		langDir := filepath.Join(dir, lang.Name())
		langMessages := make(map[string]interface{})
		businessDir := make(map[string]interface{})
		// 递归读取所有JSON文件
		err := filepath.Walk(langDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && filepath.Ext(path) == ".json" {
				// 读取文件内容
				content, err := os.ReadFile(path)
				if err != nil {
					return err
				}

				// 解析JSON
				var jsonContent interface{}
				if err := json.Unmarshal(content, &jsonContent); err != nil {
					return err
				}

				// 获取相对路径作为键
				relPath, err := filepath.Rel(langDir, path)
				if err != nil {
					return err
				}

				// 移除.json后缀
				key := strings.TrimSuffix(relPath, ".json")
				if strings.Contains(key, "business/") {
					key = strings.Replace(key, "business/", "", 1)
					businessDir[key] = jsonContent
					return nil
				}
				langMessages[key] = jsonContent
			}

			return nil
		})

		if err != nil {
			continue
		}

		messages[lang.Name()] = langMessages
		messages[lang.Name()]["business"] = businessDir
	}
	timeZone := global.GVA_CONFIG.System.TimeZone
	//nowTime := time.Now().Unix()
	response.OkWithData(map[string]interface{}{
		"messages": messages,
		"timeZone": timeZone,
	}, ctx)
}
