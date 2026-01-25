package corehandlers

import (
	"net/http"
	"server/api/apirequests"
	"server/api/apiresponses"
	"server/core"
	"server/handlers/generated/gencore"

	"github.com/labstack/gommon/log"
)

//TODO: !!!!!! MAKE CREATE COLLECTION ENDPOINT

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
