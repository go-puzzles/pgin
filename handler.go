// File:		handler.go
// Created by:	Hoven
// Created on:	2024-09-19
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package pgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func parseRequestParams(c *gin.Context, obj any) (err error) {
	if err = c.Bind(obj); err != nil {
		return
	}

	if len(c.Params) > 0 {
		if err = c.BindUri(obj); err != nil {
			return
		}
	}

	if len(c.Request.Header) > 0 {
		if err = c.BindHeader(obj); err != nil {
			return
		}
	}

	if len(c.Request.URL.Query()) > 0 {
		if err = c.BindQuery(obj); err != nil {
			return
		}
	}

	return
}

type requestHandler[Q any] func(c *gin.Context, req *Q)

func RequestHandler[Q any](fn requestHandler[Q]) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestPtr := new(Q)

		if err := parseRequestParams(c, requestPtr); err != nil {
			c.JSON(http.StatusBadRequest, ErrorRet(http.StatusBadRequest, err.Error()))
			return
		}

		fn(c, requestPtr)
	}
}

type requestResponseHandler[Q any, P any] func(c *gin.Context, req *Q) (resp *P, err error)

func RequestResponseHandler[Q any, P any](fn requestResponseHandler[Q, P]) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestPtr := new(Q)

		if err := parseRequestParams(c, requestPtr); err != nil {
			c.JSON(http.StatusBadRequest, ErrorRet(http.StatusBadRequest, err.Error()))
			return
		}

		resp, err := fn(c, requestPtr)
		currentStatus := c.Writer.Status()

		if err != nil {
			if currentStatus <= http.StatusBadRequest {
				currentStatus = http.StatusBadRequest
			}

			c.JSON(currentStatus, ErrorRet(currentStatus, err.Error()))
			return
		}

		c.JSON(http.StatusOK, SuccessRet(resp))
	}
}
