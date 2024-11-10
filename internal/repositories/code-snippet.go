package repositories

import (
	"fmt"
	"snippet-saver/internal/database"
	"snippet-saver/internal/dto/response"
	"snippet-saver/internal/models"
	"strings"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

var CodeSnippetInstance CodeSnippet = &CodeSnippetImpl{}

type CodeSnippet interface {
	CreateSnippet(tx *gorm.DB, codeSnippetData models.CodeSnippet) (models.CodeSnippet, error)
	GetAllSnippet(userId int) ([]response.GetSnippetResponse, error)
	GetSnippetById(userId int, snippetId int) (response.GetSnippetResponse, error)
	DeleteSnippetById(userId int, snippetId int) (bool, error)
	UpdateSnippetById(tx *gorm.DB, codeSnippetData models.CodeSnippet, userId int, snippetId int) (bool, error)
}

type CodeSnippetImpl struct{}

func (n *CodeSnippetImpl) CreateSnippet(tx *gorm.DB, codeSnippetData models.CodeSnippet) (models.CodeSnippet, error) {
	result := tx.Table("code_snippets").Create(&codeSnippetData)

	if result.Error != nil {
		return models.CodeSnippet{}, result.Error
	}

	return codeSnippetData, nil
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

	return result, nil

}

func (n *CodeSnippetImpl) GetSnippetById(userId int, snippetId int) (response.GetSnippetResponse, error) {
	type Result struct {
		Id          int            `json:"id"`
		Language    string         `json:"language"`
		Title       string         `json:"title"`
		Note        string         `json:"note"`
		CodeSnippet string         `json:"code_snippet"`
		Tags        pq.StringArray `gorm:"type:text[]"`
	}
	var result Result
	queryErr := database.DB.Table("code_snippets").
		Select("code_snippets.id, code_snippets.language, code_snippets.title,code_snippets.note, code_snippets.code_snippet, array_agg(tags.name) AS tags").
		Joins("LEFT JOIN tags ON tags.code_snippets = code_snippets.id").
		Where("code_snippets.user_id = ? AND code_snippets.id = ?", userId, snippetId).
		Group("code_snippets.id").
		Scan(&result).Error

	if queryErr != nil {
		return response.GetSnippetResponse{}, queryErr
	}

	response := response.GetSnippetResponse{
		Id:          result.Id,
		Tags:        []string(result.Tags),
		Title:       result.Title,
		CodeSnippet: result.CodeSnippet,
		Note:        result.Note,
		Language:    result.Language,
	}

	if response.Tags == nil {
		response.Tags = []string{}
	}

	return response, nil
}

func (m *CodeSnippetImpl) DeleteSnippetById(userId int, snippetId int) (bool, error) {
	queryErr := database.DB.Table("code_snippets").
		Where("id = ? AND user_id = ?", snippetId, userId).
		Delete(&struct{}{}).Error

	if queryErr != nil {
		return false, queryErr
	}

	return true, nil
}

func (n *CodeSnippetImpl) UpdateSnippetById(tx *gorm.DB, codeSnippetData models.CodeSnippet, userId int, snippetId int) (bool, error) {

	updatedTime := time.Now()
	result := tx.Table("code_snippets").Where("id = ? AND user_id=?", snippetId, userId).Updates(map[string]interface{}{
		"language":     codeSnippetData.Language,
		"title":        codeSnippetData.Title,
		"code_snippet": codeSnippetData.CodeContent,
		"note":         codeSnippetData.Note,
		"updated_at":   &updatedTime,
	})

	fmt.Println("result", codeSnippetData.CodeContent, codeSnippetData.ID, codeSnippetData.Language, codeSnippetData.Note)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
