package git

import (
	"strings"

	"github.com/kataras/iris/v12"

	cfg "kumu-exam/config"
	git "kumu-exam/src/services/git"
)

func GetUsers(ctx iris.Context) {
	params := git.ListForm{}
	if err := ctx.ReadJSON(&params); err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title(cfg.InvalidPayload))
		return
	}

	// validate params elements
	validParams := []string{}
	for _, p := range params.Users {
		if tmp := strings.TrimSpace(p); tmp != "" {
			validParams = append(validParams, tmp)
		}
	}
	if len(validParams) <= 0 {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title(cfg.InvalidPayload))
		return
	}

	data, err := git.GetUsers(git.ListForm{Users: validParams})
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title(err.Error()))
		return
	}

	ctx.JSON(iris.Map{"status": iris.StatusOK, "data": data})
}
