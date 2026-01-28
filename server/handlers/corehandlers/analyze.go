package corehandlers

import (
	"net/http"
	"server/api/apirequests"
	"server/api/apiresponses"
	"server/handlers/generated/gencore"
	"server/sqlc/sqlgen"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h Handler) AnalyzeCollection(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	userID, err := apirequests.User(r)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	req, err := apirequests.Request[gencore.AnalyzeCollectionRequest](r)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	analysis, err := h.Core.AnalyzeCollection(
		r.Context(),
		*userID,
		id,
		sqlgen.AnalysisType(req.Type),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	apiresponses.Success(w, gencore.CollectionAnalysis{
		Id:     analysis.ID,
		Result: string(analysis.Result),
		Type:   gencore.CollectionAnalysisType(analysis.Type),
	})
}

func (handler Handler) GetCollectionAnalyses(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	userID, err := apirequests.User(r)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	collectionID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid collection id", http.StatusBadRequest)
		return
	}

	analyses, err := handler.Core.GetCollectionAnalyses(r.Context(), *userID, collectionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := gencore.CollectionAnalyses{}
	for _, i := range analyses {
		result = append(result, gencore.CollectionAnalysis{
			Id:     i.ID,
			Type:   gencore.CollectionAnalysisType(i.Type),
			Result: string(i.Result),
		})
	}

	apiresponses.Success(w, result)
}

func (handler Handler) GetAnalysis(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, analysisID openapi_types.UUID) {
	analysis, err := handler.Queries.GetAnalysis(r.Context(), analysisID)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	result := gencore.CollectionAnalysis{
		Id:     analysis.ID,
		Result: string(analysis.Result),
		Type:   gencore.CollectionAnalysisType(analysis.Type),
	}

	apiresponses.Success(w, result)
}
