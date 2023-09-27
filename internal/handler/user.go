package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"nunu-project/internal/repository"
	"nunu-project/internal/service"
	"nunu-project/pkg/helper/resp"
	"strconv"
)

type UserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
	GetList(ctx *gin.Context)
}

type userHandler struct {
	*Handler
	userService service.UserService
}

const (
	page     = "1"
	pageSize = "10"
)

func NewUserHandler(handler *Handler, userService service.UserService) UserHandler {
	return &userHandler{
		Handler:     handler,
		userService: userService,
	}
}

func (h *userHandler) Register(ctx *gin.Context) {
	req := new(service.RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	if err := h.userService.Register(ctx, req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	resp.HandleSuccess(ctx, nil)
}

func (h *userHandler) Login(ctx *gin.Context) {
	var req service.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	token, err := h.userService.Login(ctx, &req)
	if err != nil {
		resp.HandleError(ctx, http.StatusUnauthorized, 1, err.Error(), nil)
		return
	}

	resp.HandleSuccess(ctx, gin.H{
		"accessToken": token,
	})
}

func (h *userHandler) GetProfile(ctx *gin.Context) {
	userId, ok := ctx.GetQuery("userId")
	if !ok {
		resp.HandleError(ctx, http.StatusUnauthorized, 1, "userId is not found", nil)
		return
	}
	if userId == "" {
		resp.HandleError(ctx, http.StatusUnauthorized, 1, "unauthorized", nil)
		return
	}

	user, err := h.userService.GetProfile(ctx, userId)
	if err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	resp.HandleSuccess(ctx, user)
}

func (h *userHandler) UpdateProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)

	var req service.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	if err := h.userService.UpdateProfile(ctx, userId, &req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	resp.HandleSuccess(ctx, nil)
}

func (h *userHandler) GetList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", page))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", pageSize))
	userOption := make([]repository.UserOptions, 0, 5)

	username, ok := ctx.GetQuery("username")
	if ok && username != "" {
		userOption = append(userOption, repository.WithUsername(username))
	}

	userId, ok := ctx.GetQuery("userId")
	if ok && userId != "" {
		userOption = append(userOption, repository.WithUserId(userId))
	}

	nickname, ok := ctx.GetQuery("nickname")
	if ok && nickname != "" {
		userOption = append(userOption, repository.WithUserNickname(nickname))
	}

	email, ok := ctx.GetQuery("email")
	if ok && email != "" {
		userOption = append(userOption, repository.WithEmail(email))
	}

	sortKey, sortKeyOk := ctx.GetQuery("sortType")
	sortValue, sortValueOk := ctx.GetQuery("sortValue")
	if sortKeyOk && sortValueOk && sortValue != "" && sortKey != "" {
		userOption = append(userOption, repository.WithOrder(sortValue+" "+sortKey))
	}

	userList, err := h.userService.GetList(ctx, page, pageSize, userOption...)
	if err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	//data := make([]interface{}, 0, page)
	//for _, el := range userList {
	//	el.CreatedTime, _ = el.CreatedAt.Value()
	//	data = append(data, el)
	//}
	resp.HandleSuccess(ctx, userList)
}
