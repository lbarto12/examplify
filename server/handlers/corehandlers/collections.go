package corehandlers

import (
	"net/http"
	"server/api/apirequests"
	"server/api/apiresponses"
	"server/api/validation"
	"server/business/core"
	"server/handlers/generated/gencore"
	"server/sqlc/sqlgen"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// (POST /core/new-collection)
func (handler Handler) NewCollection(w http.ResponseWriter, r *http.Request) {
	userID, err := apirequests.User(r)
	if err != nil {
		apiresponses.BadRequest(w, "Invalid Request", err)
		return
	}

	request, err := apirequests.Request[gencore.NewCollectionRequest](r)
	if err != nil {
		apiresponses.BadRequest(w, "Invalid Request", err)
		return
	}

	// Validate collection title is non-empty
	if err := validation.ValidateNonEmpty("title", request.Title); err != nil {
		apiresponses.BadRequest(w, err.Error(), err)
		return
	}

	// Auto-create course if it doesn't exist
	exists, err := handler.Queries.CourseExists(r.Context(), sqlgen.CourseExistsParams{
		Name:   request.Course,
		UserID: *userID,
	})
	if err != nil {
		apiresponses.InternalError(w, "Internal Error", err)
		return
	}

	if !exists {
		err = handler.Queries.CreateCourse(r.Context(), sqlgen.CreateCourseParams{
			Name:      request.Course,
			CreatorID: *userID,
		})
		if err != nil {
			apiresponses.InternalError(w, "Failed to create course", err)
			return
		}
	}

	collection, err := handler.Core.CreateCollection(r.Context(), *userID, core.Collection{
		Title:  request.Title,
		Type:   request.Type,
		Course: request.Course,
	})
	if err != nil {
		apiresponses.InternalError(w, "Internal Error", err)
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
		apiresponses.BadRequest(w, "Invalid Request", err)
		return
	}

	request, err := apirequests.Request[gencore.UploadFileRequest](r)
	if err != nil {
		apiresponses.BadRequest(w, "Invalid Request", err)
		return
	}

	// Validate MIME type is allowed
	if err := validation.ValidateMimeType(request.MimeType); err != nil {
		apiresponses.BadRequest(w, err.Error(), err)
		return
	}

	uploadURL, err := handler.Core.CreateDocument(r.Context(), *userID, core.Document{
		CollectionID: request.CollectionID,
		MimeType:     request.MimeType,
	})
	if err != nil {
		apiresponses.InternalError(w, "Internal Error", err)
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
		apiresponses.BadRequest(w, "Invalid Request", err)
		return
	}

	collection, err := handler.Core.GetCollection(r.Context(), *userID, id)
	if err != nil {
		apiresponses.InternalError(w, "Internal Error", err)
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
		apiresponses.BadRequest(w, "Invalid Request", err)
		return
	}

	document, err := handler.Core.GetDocument(r.Context(), *userID, id)
	if err != nil {
		apiresponses.InternalError(w, "Internal Error", err)
		return
	}

	downloadURL, err := handler.Core.PresignedGetDocument(r.Context(), *userID, document.ID)
	if err != nil {
		apiresponses.InternalError(w, "Internal Error", err)
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
		apiresponses.BadRequest(w, "Invalid Request", err)
		return
	}

	filtered, err := handler.Queries.FilterCollections(r.Context(), sqlgen.FilterCollectionsParams{
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

// TODO: Add pagination - limit/offset parameters (e.g., ?limit=50&offset=0)
// For production, this endpoint should support pagination to handle large datasets
func (handler Handler) GetCollectionDocuments(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	userID, err := apirequests.User(r)
	if err != nil {
		apiresponses.BadRequest(w, "Invalid Request", err)
		return
	}

	docs, err := handler.Core.GetCollectionDocuments(r.Context(), *userID, id)
	if err != nil {
		apiresponses.InternalError(w, "Internal Error", err)
		return
	}

	result := gencore.Documents{}
	for _, i := range docs {
		download, err := handler.Core.PresignedGetDocument(r.Context(), *userID, i.ID)
		if err != nil {
			apiresponses.InternalError(w, "Internal Error", err)
			return
		}

		result = append(result, gencore.Document{
			ID:           i.ID,
			CollectionID: i.CollectionID,
			MimeType:     i.MimeType,
			DownloadURL:  download.String(),
		})
	}

	apiresponses.Success(w, result)
}
