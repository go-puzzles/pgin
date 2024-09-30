// File:		engine.go
// Created by:	Hoven
// Created on:	2024-09-24
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package pgin

import "github.com/gin-gonic/gin"

func Default(opts ...gin.OptionFunc) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()
	engine.Use(LoggerMiddleware(), gin.Recovery())
	return engine.With(opts...)
}
