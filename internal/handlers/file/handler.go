package file

import (
	"gobee/ent"
	"gobee/internal/services/infra/file"
	"gobee/internal/services/storage"
	"gobee/pkg/domain/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type FileHandler interface {
	ListFile(c *fiber.Ctx) error
	ListFilePage(c *fiber.Ctx) error
	QueryFile(c *fiber.Ctx) error
	DeleteFile(c *fiber.Ctx) error
	Upload(c *fiber.Ctx) error
}

type FileHandlerImpl struct {
	fileService file.FileService
}

func NewFileHandlerImpl(fileService file.FileService) *FileHandlerImpl {
	return &FileHandlerImpl{fileService: fileService}
}

func (h *FileHandlerImpl) ListFile(c *fiber.Ctx) error {
	files, err := h.fileService.ListFile(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("文件列表获取成功", files))
}

func (h *FileHandlerImpl) ListFilePage(c *fiber.Ctx) error {
	pageQuery := model.PageQuery{}
	if err := c.QueryParser(&pageQuery); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}

	count, files, err := h.fileService.ListFilePage(c.Context(), pageQuery.Page, pageQuery.Size)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	pageResult := model.PageResult[*ent.File]{
		Total:   int64(count),
		Records: files,
	}
	return c.JSON(model.NewSuccess("文件列表获取成功", pageResult))
}

func (h *FileHandlerImpl) QueryFile(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}

	file, err := h.fileService.QueryFile(c.Context(), id)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("文件查询成功", file))
}

func (h *FileHandlerImpl) DeleteFile(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}

	err = h.fileService.DeleteFile(c.Context(), id)
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("文件删除成功", nil))
}

func (h *FileHandlerImpl) Upload(c *fiber.Ctx) error {
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

	var results []*ent.File

	for _, file := range files {

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

		// 保存文件
		var fullUrl string

		if strategy.Domain != "" {
			fullUrl = strategy.Domain + "/" + url
		} else {
			fullUrl = strategy.Endpoint + "/" + url
		}

		// 保存到数据库
		newFile, err := h.fileService.CreateFile(c.Context(), file.Filename, strategy.BasePath, fullUrl, file.Header.Get("Content-Type"), strconv.FormatInt(file.Size, 10))
		if err != nil {
			return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
		}
		results = append(results, newFile)
	}

	return c.JSON(model.NewSuccess("文件上传成功", results))
}
