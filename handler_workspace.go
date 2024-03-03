package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/davifrjose/My_Turn/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerWorkSpacesCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string
		Email string
		Address string
		UserID      uuid.UUID
		DisplayName string
		OpeningTime time.Time
		ClosingTime time.Time

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

	})



	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't create workspace")
		return
	}
	respondWithJSON(w, http.StatusCreated, databaseWorkspaceToWorkspace(workspace))
}