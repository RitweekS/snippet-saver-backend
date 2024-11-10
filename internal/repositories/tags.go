package repositories

import (
	"fmt"
	"snippet-saver/internal/models"

	"gorm.io/gorm"
)

var TagsInstance Tags = &TagsImpl{}

type Tags interface {
	CreateTags(tx *gorm.DB, tag models.SnippetTag) (bool , error)
	UpdateTags(tx *gorm.DB, tag models.SnippetTag,snippetId int) (bool , error)

}

type TagsImpl struct{}



func (n *TagsImpl) CreateTags(tx *gorm.DB, tag models.SnippetTag) (bool, error) {
    result := tx.Table("tags").Create(&tag)

    if result.Error != nil{
        return false,result.Error
    }

    return true,nil
}
func (n *TagsImpl) UpdateTags(tx *gorm.DB, tag models.SnippetTag,snippetId int) (bool, error) {
    if tag.ID == 0 {
        return false, fmt.Errorf("tag ID is required for update")
    }

    result := tx.Table("tags").Where("id = ?", snippetId).Updates(map[string]interface{}{
        "name":       tag.Name,
        "updated_at": tag.UpdatedAt,
    })

    if result.Error != nil {
        return false, result.Error
    }

    return true, nil
}