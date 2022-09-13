// Copyright 2022 Adam Piaseczny. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/TypicalAM/Static-Site-Generator/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	rootDir := "markdown"
	router := gin.Default()
	err, files := utils.DiscoverFiles(rootDir)
	if err != nil {
		panic("File discovery failed")
	}
	err = utils.GenerateUrls(router, rootDir, files)
	if err != nil {
		panic("Url generation failed")
	}
	router.Run()
}
