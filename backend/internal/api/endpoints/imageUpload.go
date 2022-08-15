package endpoints

import (
	"errors"
	"github.com/ericlp/tasteit/backend/internal/common"
	"github.com/ericlp/tasteit/backend/internal/process"
	"github.com/ericlp/tasteit/backend/internal/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ImageUpload(c *gin.Context) {
	formFile, formHeader, err := c.Request.FormFile("file")
	if err != nil {
		log.Printf("Failed to retrieve file from form: %v", err)
		c.JSON(http.StatusBadRequest, common.Error(common.ResponseMissingFile))
		return
	}

	image, err := validation.ValidateFile(&formFile, formHeader)
	if err != nil {
		log.Printf("Failed to validate image: %v", err)

		if errors.Is(err, validation.ErrFiletypeNotSupported) {
			c.JSON(http.StatusBadRequest, common.Error(common.ResponseFileTypeNotSupported))
			return
		}

		c.JSON(http.StatusBadRequest, common.Error(common.ResponseBadImage))
		return
	}

	imageJson, err := process.UploadImage(image)
	if err != nil {
		log.Printf("Failed to handle image upload: %v", err)
		c.JSON(http.StatusInternalServerError, common.Error(common.ResponseFailedToSaveImage))
		return
	}

	c.JSON(http.StatusOK, common.Success(imageJson))
}
