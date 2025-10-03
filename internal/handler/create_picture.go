package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreatePicture godoc
// @Summary      Create pictures for a trip
// @Description  Upload multiple pictures for a specific trip by trip ID
// @Tags         Pictures
// @Accept       multipart/form-data
// @Produce      json
// @Param        trip_id   path      string  true  "trip_id"
// @Param        images    formData  file    true  "Images"
// @Security     BearerAuth
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /picture/{trip_id} [post]
func (h *PictureHandler) CreatePicture(c *gin.Context) {
	ctx := c.Request.Context()

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid form data"})
		return
	}

	files := form.File["images"]
	filesPath := []string{}

	for _, file := range files {
		fileType := filepath.Ext(file.Filename)
		baseName := strings.TrimSuffix(filepath.Base(file.Filename), fileType)
		filename := strings.ReplaceAll(strings.ToLower(baseName), " ", "-") + "-" + fmt.Sprintf("%v", time.Now().Unix()) + fileType

		path := "http://localhost:8000/images/multiple/" + filename
		filesPath = append(filesPath, path)
		if err := os.MkdirAll("./public/multiple", os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot create directory"})
			return
		}
		out, err := os.Create("./public/multiple/" + filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		reader, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			out.Close()
			return
		}

		if _, err := io.Copy(out, reader); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			reader.Close()
			out.Close()
			return
		}

		reader.Close()
		out.Close()
	}

	tripIDstr := c.Param("trip_id")
	tripID, err := uuid.Parse(tripIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid trip_id"})
		return
	}

	PicturesPath, err := h.pictureService.CreatePicture(ctx, filesPath, tripID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	fmt.Println(PicturesPath)
	c.JSON(http.StatusCreated, gin.H{"message": "created", "data": PicturesPath})
}
