package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"nunu-project/internal/service"
	"nunu-project/pkg/helper/resp"
)

type BlogHandler interface {
	GetBlogById(ctx *gin.Context)
	UpdateBlog(ctx *gin.Context)
	GetBlogList(ctx *gin.Context)
}

type blogHandler struct {
	*Handler
	blogService service.BlogService
}

func (h *blogHandler) GetBlogList(ctx *gin.Context) {
	blogList, err := h.blogService.GetBlogList()
	if err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}
	data := make(map[string]interface{})
	data["list"] = blogList
	resp.HandleSuccess(ctx, data)
}

func NewBlogHandler(handler *Handler, blogService service.BlogService) BlogHandler {
	return &blogHandler{
		Handler:     handler,
		blogService: blogService,
	}
}

func (h *blogHandler) GetBlogById(ctx *gin.Context) {
	var params struct {
		Id int64 `form:"id" binding:"required"`
	}

	h.logger.Info("fhfghfhj")
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	blog, err := h.blogService.GetBlogById(params.Id)
	h.logger.Info("GetBlogByID", zap.Any("blog", blog))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, blog)
}

func (h *blogHandler) UpdateBlog(ctx *gin.Context) {
	resp.HandleSuccess(ctx, nil)
}
