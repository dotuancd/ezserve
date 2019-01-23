package controllers

import (
	"github.com/dotuancd/ezserve/app"
	"github.com/dotuancd/ezserve/app/http/errors"
	"github.com/dotuancd/ezserve/app/http/res"
	m "github.com/dotuancd/ezserve/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	App *app.App
}

var (
	ErrUserIsEmpty = errors.New("empty_username", "The email cannot be empty", http.StatusUnprocessableEntity)

	ErrPasswordIsEmpty = errors.New("password_empty", "The password cannot be empty", http.StatusUnprocessableEntity)

	ErrLoginFailed = errors.New("credential_is_mismatch", "The credential a is mismatch ", http.StatusUnauthorized)

)

type loginRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (ac *AuthController) Login () errors.HandlerFunc {
	return func(c *gin.Context) error {
		credential := loginRequest{}

		err := c.ShouldBindJSON(&credential)

		if err != nil {
			return err
		}

		user := m.User{}
		ac.App.DB.First(&user, m.User{Email: credential.Email})

		if user.ID == 0 {
			return ErrLoginFailed
		}

		res.NewSuccess(c).
			Extra("user", user).
			Send()

		return nil
	}
}