package repositories

import (
	"snippet-saver/internal/models"

	"gorm.io/gorm"
)

var TagsInstance Tags = &TagsImpl{}

type Tags interface {
	CreateTags(tx *gorm.DB, tag models.SnippetTag) (bool , error)

}

type TagsImpl struct{}



func (n *TagsImpl) CreateTags(tx *gorm.DB, tag models.SnippetTag) (bool, error) {
    result := tx.Table("tags").Create(&tag)

    if result.Error != nil{
        return false,result.Error
    }

    return true,nil
}
