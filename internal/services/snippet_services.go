package services

import (
	"snippet-saver/internal/database"
	"snippet-saver/internal/dto/request"
	"snippet-saver/internal/dto/response"
	"snippet-saver/internal/models"
	"snippet-saver/internal/repositories"
	"time"
)

var SnippetInstance Snippet = &SnippetImpl{}

type Snippet interface {
	CreateSnippet(userId int,requestBody request.CreateSnippetRequest) (bool , error)
    GetAllSnippet(userId int)([]response.GetSnippetResponse,error)
}

type SnippetImpl struct{}

func (n *SnippetImpl) CreateSnippet(userId int, requestBody request.CreateSnippetRequest) (bool, error) {
    // Start a transaction
    tx := database.DB.Begin()
    currentTime := time.Now()

    codeSnippetModel := models.CodeSnippet{
        UserID:      userId,
        Language:    requestBody.Language,
        Title:       requestBody.Title,
        CodeContent: requestBody.CodeSnippet,
        Note:        requestBody.Note,
        CreatedAt:   &currentTime,
    }

    createSnippet, createSnippetErr := repositories.CodeSnippetInstance.CreateSnippet(tx,codeSnippetModel)
    if createSnippetErr != nil {
        tx.Rollback() 
        return false, createSnippetErr
    }

    snippetID := createSnippet.ID

    // Loop through the tags and create each one
    for _, tagName := range requestBody.Tags {
        tag := models.SnippetTag{
            Name:          tagName,
            CodeSnippetID: snippetID, 
            CreatedAt:     &currentTime,
        }

        // Create the tag
        _,createTagError := repositories.TagsInstance.CreateTags(tx, tag)
        if createTagError != nil {
            tx.Rollback() // Rollback on error
            return false, createTagError
        }
    }

    // If everything is successful, commit the transaction
    if err := tx.Commit(); err != nil {
        return false, err.Error
    }

    return true, nil
}
type SnippetWithTags struct {
	Id          int
	Language    string
	Title       string
	CodeSnippet string
	Tags        string // To hold comma-separated tag names
}

func (n *SnippetImpl) GetAllSnippet(userId int) ([]response.GetSnippetResponse, error) {
	return  repositories.CodeSnippetInstance.GetAllSnippet(userId)
}