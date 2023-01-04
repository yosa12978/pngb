package dtos

import (
	"strings"

	"github.com/yosa12978/pngb/internal/pkg/helpers"
	"github.com/yosa12978/pngb/internal/pkg/models"
)

type PostCreateDTO struct {
	Text string   `json:"text"`
	Imgs []string `json:"imgs"`
}

func (dto *PostCreateDTO) Map() (models.Post, error) {
	if strings.Replace(dto.Text, " ", "", -1) == "" && (len(dto.Imgs) <= 0 || len(dto.Imgs) > 4) {
		return models.Post{}, helpers.ErrBadRequest
	}
	post := models.Post{
		Text: dto.Text,
		Imgs: dto.Imgs,
	}
	return post, nil
}
