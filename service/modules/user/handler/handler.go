package handler

import (
	"fmt"
	"net/http"

	"github.com/DeniesKresna/gohelper/utstring"
	"github.com/DeniesKresna/sined/service/extensions/helper"
	"github.com/DeniesKresna/sined/service/extensions/terror"
	"github.com/DeniesKresna/sined/service/modules/user/usecase"
	"github.com/DeniesKresna/sined/types/constants"
	"github.com/DeniesKresna/sined/types/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.IUsecase
}

func UserCreateHandler(userUsecase usecase.IUsecase) UserHandler {
	return UserHandler{
		userUsecase: userUsecase,
	}
}

func ResponseJson(ctx *gin.Context, data interface{}) {
	httpStatusCode := http.StatusOK
	resp, ok := data.(*terror.ErrorModel)
	if ok {
		responseData := models.Response{
			ResponseDesc: resp.Message,
			ResponseCode: resp.Code,
		}

		appEnv := utstring.GetEnv(constants.ENV_APP_ENV, "local")
		if appEnv != "production" {
			responseData.ResponseTrace = resp.Trace
		}

		ctx.JSON(httpStatusCode, responseData)
		return
	}

	responseData := models.Response{
		ResponseDesc: "Success",
		ResponseCode: "00",
	}

	if helper.IsStruct(data) || helper.IsMap(data) || helper.IsSlice(data) {
		responseData.ResponseData = data
	} else {
		responseData.ResponseDesc = fmt.Sprintf("%v", data)
	}

	ctx.JSON(httpStatusCode, responseData)
}
