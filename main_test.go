package main

import (
	"log"
	"testing"

	auth "kumu-exam/src/authenticator"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
)

func TestInitApp(t *testing.T) {
	// generate token tobe use for test
	token, err := auth.GenerateToken()
	if err != nil {
		log.Panic("Unable to generate token for testing")
	}

	a := initApp()
	e := httptest.New(t, a)

	resp := map[string]interface{}{
		"status":  iris.StatusOK,
		"message": "Hi! This is KUMU Coding Challenge",
	}
	e.GET("/").Expect().Status(iris.StatusOK).JSON().Equal(resp)

	e.POST("/auth/token").Expect().Status(iris.StatusOK).JSON().Object().Keys().Contains("data", "status")

	body := map[string]interface{}{
		"users": []interface{}{
			"some-user-1",
			"krlontoc",
			"some-user-2",
			"",
			"krlontoc",
		},
	}
	e.GET("/api/v1/git-users").WithHeaders(map[string]string{"Authorization": "Bearer " + token}).
		WithJSON(body).Expect().Status(iris.StatusOK).JSON().Object().Keys().Contains("data", "status")

}
