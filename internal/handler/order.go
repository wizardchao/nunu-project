package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"nunu-project/internal/service"
	"nunu-project/pkg/helper/resp"
)

type OrderHandler interface {
	GetOrderById(ctx *gin.Context)
	UpdateOrder(ctx *gin.Context)
}

type orderHandler struct {
	*Handler
	orderService service.OrderService
}

func NewOrderHandler(handler *Handler, orderService service.OrderService) OrderHandler {
	return &orderHandler{
		Handler:      handler,
		orderService: orderService,
	}
}

func (h *orderHandler) GetOrderById(ctx *gin.Context) {
	var params struct {
		Id int64 `form:"id" binding:"required"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	h.logger.Info(fmt.Sprintf("params:%d", params.Id))
	order, err := h.orderService.GetOrderById(params.Id)
	h.logger.Info("GetOrderByID", zap.Any("order", order))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, order)
}

func (h *orderHandler) UpdateOrder(ctx *gin.Context) {
	resp.HandleSuccess(ctx, nil)
}
