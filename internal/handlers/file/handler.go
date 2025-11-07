package file

import (
	"gobee/ent"
	"gobee/ent/file"
	"gobee/internal/database"
	"gobee/internal/services/storage"
	"gobee/pkg/domain/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ListFile(c *fiber.Ctx) error {
	client := database.DB
	files, err := client.File.Query().All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("文件列表获取成功", files))
}

func ListFilePage(c *fiber.Ctx) error {
	client := database.DB
	pageQuery := model.PageQuery{}
	if err := c.QueryParser(&pageQuery); err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	count, err := client.File.Query().Count(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	files, err := client.File.Query().
		Offset((pageQuery.Page - 1) * pageQuery.Size).
		Limit(pageQuery.Size).
		All(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	pageResult := model.PageResult[*ent.File]{

		Total:   int64(count),
		Records: files,
	}
	return c.JSON(model.NewSuccess("文件列表获取成功", pageResult))
}

func QueryFile(c *fiber.Ctx) error {
	client := database.DB
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	file, err := client.File.Query().Where(file.ID(id)).First(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("文件查询成功", file))
}

func DeleteFile(c *fiber.Ctx) error {
	client := database.DB
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusBadRequest, err.Error()))
	}
	err = client.File.DeleteOneID(id).Exec(c.Context())
	if err != nil {
		return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
	}
	return c.JSON(model.NewSuccess("文件删除成功", nil))
}

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

		client := database.DB
		// 保存到数据库
		newFile, err := client.File.Create().
			SetName(file.Filename).
			SetPath(strategy.BasePath).
			SetURL(fullUrl).
			SetType(file.Header.Get("Content-Type")).
			SetSize(strconv.FormatInt(file.Size, 10)).
			Save(c.Context())
		if err != nil {
			return c.JSON(model.NewError(fiber.StatusInternalServerError, err.Error()))
		}
		results = append(results, newFile)
	}

	return c.JSON(model.NewSuccess("文件上传成功", results))
}
