package router

import (
	"embed"
	"io"
	"io/fs"
	"log"
	"mime"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func InitFrontendRes(app *fiber.App, frontendRes embed.FS) {

	// 载入静态资源
	distDir, err := fs.Sub(frontendRes, "ui/dist")
	if err != nil {
		log.Fatalln("静态资源载入失败")
	}

	app.Get("/console/*", func(c *fiber.Ctx) error {
		resp := c.Response()
		filePath := c.Params("*")
		var contentType string
		if file, err := distDir.Open(filePath); err == nil {
			defer file.Close()
			stat, err := file.Stat()
			if err == nil && !stat.IsDir() {
				// 不是文件夹
				ext := filepath.Ext(filePath)
				contentType = mime.TypeByExtension(ext)
				if contentType == "" {
					contentType = "application/octet-stream"
				}
				if ext == ".html" {
					// 是 html 文件

				}
				resp.Header.Set("Content-Type", contentType)
				_, err = io.Copy(c.Response().BodyWriter(), file)
				return err
			} else if err == nil && stat.IsDir() {
				//是文件夹，向下查找 index.html
				fillePathTmp := filepath.Join(filePath, "/index.html")
				if fileSub, err := distDir.Open(fillePathTmp); err == nil {
					defer fileSub.Close()
					statSub, err := fileSub.Stat()
					if err == nil && !statSub.IsDir() {
						if filepath.Ext(filePath) == ".html" {
							// 是 html 文件
							resp.Header.Set("Content-Type", "text/html")
							_, err = io.Copy(c.Response().BodyWriter(), file)
							return err
						}
					}
				}
			}

		}
		// If no file or directory/index.html found, try to serve root index.html
		if file, err := distDir.Open("index.html"); err == nil {
			defer file.Close()
			resp.Header.Set("Content-Type", "text/html")
			_, err = io.Copy(c.Response().BodyWriter(), file)
			return err
		}
		return c.SendStatus(fiber.StatusNotFound)
	})
}
