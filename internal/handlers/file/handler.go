package file

import (
	"gobee/internal/model"
	"gobee/internal/services/storage"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}

	files := form.File["file"]

	if len(files) == 0 {
		return c.JSON(model.NewError(fiber.StatusBadRequest, "请选择要上传的文件"))
	}

	var storageStrategyID int

	// 检查存储器策略
	if formValue, ok := form.Value["storage_strategy"]; ok {
		if len(formValue) > 0 {
			storageStrategyID, err = strconv.Atoi(formValue[0])
			if err != nil {
				return c.JSON(model.NewError(fiber.StatusBadRequest, "无效的存储策略ID"))
			}
		}
	}

	var results []map[string]interface{}

	for _, file := range files {
		fileInfo := map[string]interface{}{
			"filename": file.Filename,
			"size":     file.Size,
			"mimeType": file.Header.Get("Content-Type"),
			"status":   "success",
			"url":      "/api/file/" + file.Filename,
		}

		// 获取存储策略
		strategy, err := storage.GetStorageStrategyByID(storageStrategyID)
		if err != nil {
			return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
		}
		uploader, err := storage.GetUploader(strategy)
		if err != nil {
			return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
		}
		// 上传文件
		// 打开文件获取 io.Reader
		f, err := file.Open()
		if err != nil {
			return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
		}
		defer f.Close()

		url, err := uploader.Upload(file.Filename, f, file.Size, file.Header.Get("Content-Type"))
		if err != nil {
			return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
		}
		fileInfo["url"] = url

		// 保存文件
		results = append(results, fileInfo)
	}

	return c.JSON(model.NewSuccess("文件上传成功", results))
}
