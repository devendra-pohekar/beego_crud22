package controllers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/astaxie/beego"
)

type FileUploadController struct {
	beego.Controller
}

func (c *FileUploadController) Upload() {
	file, header, err := c.GetFile("file")
	if err != nil {
		c.Ctx.WriteString(fmt.Sprintf("Get file error: %v", err))
		return
	}
	defer file.Close()
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(header.Filename))

	uploadDir := "./uploads/images/"
	if err := os.MkdirAll(uploadDir, 0777); err != nil {
		c.Ctx.WriteString(fmt.Sprintf("Failed to create upload directory: %v", err))
		return
	}
	filePath := filepath.Join(uploadDir, filename)
	out, err := os.Create(filePath)
	if err != nil {
		c.Ctx.WriteString(fmt.Sprintf("Failed to create the destination file: %v", err))
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		c.Ctx.WriteString(fmt.Sprintf("Failed to copy the file: %v", err))
		return
	}

	c.Ctx.WriteString(fmt.Sprintf("File uploaded successfully. File path: %s", filePath))
}
