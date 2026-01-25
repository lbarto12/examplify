package corehandlers

import (
	"net/http"
	"server/api/apirequests"
	"server/api/apiresponses"
	"server/core"
	"server/handlers/generated/gencore"
	"server/sqlc/sqlgen"

	"github.com/labstack/gommon/log"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// (POST /core/new-collection)
func (handler Handler) NewCollection(w http.ResponseWriter, r *http.Request) {
	userID, err := apirequests.User(r)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	request, err := apirequests.Request[gencore.NewCollectionRequest](r)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	collection, err := handler.Core.CreateCollection(r.Context(), *userID, core.Collection{
		Title:  request.Title,
		Type:   request.Type,
		Course: request.Course,
	})
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	apiresponses.Success(w, gencore.NewCollectionResponse{
		CollectionID: collection.ID,
	})
}

// (POST /core/upload-file)
func (handler Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	userID, err := apirequests.User(r)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	request, err := apirequests.Request[gencore.UploadFileRequest](r)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	uploadURL, err := handler.Core.CreateDocument(r.Context(), *userID, core.Document{
		CollectionID: request.CollectionID,
		MimeType:     request.MimeType,
	})
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	apiresponses.Success(w, gencore.UploadFileResponse{
		UploadURL: uploadURL.String(),
	})
}

// (GET /core/collection)
func (handler Handler) GetCollection(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	userID, err := apirequests.User(r)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	collection, err := handler.Core.GetCollection(r.Context(), *userID, id)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	apiresponses.Success(w, gencore.Collection{
		ID:     collection.ID,
		Title:  collection.Title,
		Course: collection.Course,
		Type:   collection.Type,
	})
}

// (GET /core/document)
func (handler Handler) GetDocument(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	userID, err := apirequests.User(r)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	document, err := handler.Core.GetDocument(r.Context(), *userID, id)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	downloadURL, err := handler.Core.PresignedGetDocument(r.Context(), *userID, document.ID)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	apiresponses.Success(w, gencore.Document{
		ID:           document.ID,
		CollectionID: document.CollectionID,
		MimeType:     document.MimeType,
		DownloadURL:  downloadURL.String(),
	})
}

func (handler Handler) FilterCollections(w http.ResponseWriter, r *http.Request, courseID string, pType string) {
	userID, err := apirequests.User(r)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	qtx := sqlgen.New(handler.Services.Postgres)

	filtered, err := qtx.FilterCollections(r.Context(), sqlgen.FilterCollectionsParams{
		UserID: *userID,
		Course: courseID,
		Type:   pType,
	})

	result := gencore.Collections{}
	for _, i := range filtered {
		result = append(result, gencore.Collection{
			ID:     i.ID,
			Course: i.Course,
			Title:  i.Title,
			Type:   i.Type,
		})
	}

	apiresponses.Success(w, result)
}
