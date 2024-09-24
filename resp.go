// File:		resp.go
// Created by:	Hoven
// Created on:	2024-09-19
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package pgin

type Ret struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func SuccessRet(data any) *Ret {
	return &Ret{Code: 0, Data: data, Message: "success"}
}

func ErrorRet(code int, message string) *Ret {
	return &Ret{Code: code, Data: nil, Message: message}
}
