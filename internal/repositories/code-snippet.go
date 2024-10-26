package repositories

import (
	"snippet-saver/internal/database"
	"snippet-saver/internal/dto/response"
	"snippet-saver/internal/models"
	"strings"

	"gorm.io/gorm"
)




var CodeSnippetInstance CodeSnippet = &CodeSnippetImpl{}

type CodeSnippet interface {
	CreateSnippet(tx *gorm.DB, codeSnippetData models.CodeSnippet) (models.CodeSnippet , error)
    GetAllSnippet(userId int) ([]response.GetSnippetResponse, error)

}

type CodeSnippetImpl struct{}



func (n *CodeSnippetImpl) CreateSnippet(tx *gorm.DB, codeSnippetData models.CodeSnippet) (models.CodeSnippet, error) {
    result := tx.Table("code_snippets").Create(&codeSnippetData)

    if result.Error != nil{
        return models.CodeSnippet{},result.Error
    }

    return codeSnippetData,nil
}

type SnippetWithTags struct {
	Id          int
	Language    string
	Title       string
	CodeSnippet string
    Note        string
	Tags        string 
}



func (n *CodeSnippetImpl) GetAllSnippet(userId int) ([]response.GetSnippetResponse, error) {
   
    var snippets []SnippetWithTags
    err := database.DB.Table("code_snippets").
        Select("code_snippets.id, code_snippets.language, code_snippets.title, code_snippets.code_snippet,code_snippets.note, STRING_AGG(tags.name, ',') AS tags").
        Joins("LEFT JOIN tags ON tags.code_snippets=  code_snippets.id").
        Where("code_snippets.user_id = ?", userId).
        Group("code_snippets.id").
        Find(&snippets).Error

    if err != nil {
        return []response.GetSnippetResponse{}, err
    }


    var result []response.GetSnippetResponse
    for _, s := range snippets {
        tags := strings.Split(s.Tags, ",") // Convert comma-separated tags into a slice
        result = append(result, response.GetSnippetResponse{
            Id:          s.Id,
            Tags:        tags,
            Title:       s.Title,
            Note:        s.Note,
            CodeSnippet: s.CodeSnippet,
            Language:    s.Language,
        })
    }

    return result,nil

}
