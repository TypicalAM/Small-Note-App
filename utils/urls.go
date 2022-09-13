// Copyright 2022 Adam Piaseczny. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GenerateUrls	takes the router and generates appropriate urls for the found markdown files
// if err!=nil it returns the encountered error. The markdown files are assumed to be safe code.
func GenerateUrls(router *gin.Engine, rootDir string, files []string) error {
	router.LoadHTMLGlob("templates/index.tmpl")
	router.Static("/assets", "./assets")
	for _, file := range files {
		path := rootDir + "/" + file
		err, md := ProcessFile(path)
		if err != nil {
			return err
		}
		router.GET(file, func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
				"files":    files,
				"markdown": template.HTML(md),
			})
		})
	}
	return nil
}
