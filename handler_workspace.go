package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/davifrjose/My_Turn/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerWorkSpacesCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	Address      string     `json:"address"`
	UserID       uuid.UUID  `json:"user_id"`
	DisplayName  string     `json:"display_name"`
	OpeningTime  time.Time  `json:"opening_time"`
	ClosingTime  time.Time  `json:"closing_time"`
	Logo         string    `json:"logo"`
	Description  string    `json:"description"`

	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
	}

	user, err := cfg.DB.SelectUserById(r.Context(), params.UserID)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "This user does not exist")
	}

	description := sql.NullString{}
	if params.Description != "" {
		description.String = params.Description
		description.Valid = true
	}

	logo := sql.NullString{}
	if params.Logo != "" {
		logo.String = params.Logo
		logo.Valid = true
	}


	workspace, err := cfg.DB.CreateWorkspaces(r.Context(), database.CreateWorkspacesParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		Name: params.Name,
		Email: params.Email,
		Address: params.Address,
		UserID: user.ID,
		DisplayName: params.DisplayName,
		OpeningTime: params.OpeningTime,
		ClosingTime: params.ClosingTime,
		Logo: logo,
		Description: description,

	})



	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't create workspace")
		return
	}
	respondWithJSON(w, http.StatusCreated, databaseWorkspaceToWorkspace(workspace))
}