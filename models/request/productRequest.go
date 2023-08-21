package request

import "mime/multipart"

// pakai form
type ProductRequest struct {
	Name        string                `form:"name" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Stock       int                   `form:"stock" binding:"required"`
	Image       *multipart.FileHeader `form:"file"`
}
