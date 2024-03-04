package main

import (
	"time"

	"github.com/davifrjose/My_Turn/internal/database"
	"github.com/google/uuid"
)


type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		Email: user.Email,
		Name: user.Name,
		Password: user.Password,
	}
}
type Workspace struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	UserID      uuid.UUID `json:"user_id"`
	DisplayName string    `json:"display_name"`
	OpeningTime time.Time `json:"opening_time"`
	ClosingTime time.Time `json:"closing_time"`
	Logo *string `json:"logo"`
	Description *string `json:"description"`
}

func databaseWorkspaceToWorkspace(workspace database.Workspace) Workspace {
	var description *string
	if workspace.Description.Valid {
		description = &workspace.Description.String
	}

	var logo *string
	if workspace.Logo.Valid {
		logo = &workspace.Logo.String
	}
	return Workspace{
		ID:        workspace.ID,
		CreatedAt: workspace.CreatedAt,
		Email: workspace.Email,
		Name: workspace.Name,
		Address: workspace.Address,
		UserID: workspace.UserID,
		DisplayName: workspace.DisplayName,
		OpeningTime: workspace.OpeningTime,
		ClosingTime: workspace.ClosingTime,
		Logo: logo,
		Description: description,
	}
}