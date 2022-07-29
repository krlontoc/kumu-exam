package authenticator

import (
	"github.com/kataras/iris/v12"
)

func AuthToken(ctx iris.Context) {
	tokenString, err := GenerateToken()
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().Title(err.Error()))
		return
	}

	ctx.JSON(iris.Map{"status": iris.StatusOK, "data": tokenString})
}
