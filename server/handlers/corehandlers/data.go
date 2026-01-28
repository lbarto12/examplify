package corehandlers

import (
	"net/http"
	"server/api/apirequests"
	"server/api/apiresponses"
	"server/handlers/generated/gencore"
	"server/sqlc/sqlgen"
)

// (GET /core/course/{courseID}/collections)
// TODO: Add pagination - limit/offset parameters (e.g., ?limit=50&offset=0)
// For production, this endpoint should support pagination to handle large datasets
func (handler Handler) GetCourseCollections(w http.ResponseWriter, r *http.Request, courseID string) {
	userID, err := apirequests.User(r)
	if err != nil {
		apiresponses.BadRequest(w, "Invalid request", err)
		return
	}

	collections, err := handler.Queries.GetCourseCollections(r.Context(), sqlgen.GetCourseCollectionsParams{
		UserID:   *userID,
		CourseID: courseID,
	})
	if err != nil {
		apiresponses.InternalError(w, "Internal Error", err)
		return
	}

	names := gencore.Collections{}
	for _, i := range collections {
		names = append(names, gencore.Collection{
			ID:     i.ID,
			Course: i.Course,
			Title:  i.Title,
			Type:   i.Type,
		})
	}

	apiresponses.Success(w, names)
}

// (GET /core/courses)
// TODO: Add pagination - limit/offset parameters (e.g., ?limit=50&offset=0)
// For production, this endpoint should support pagination to handle large datasets
func (handler Handler) GetCourses(w http.ResponseWriter, r *http.Request) {
	userID, err := apirequests.User(r)
	if err != nil {
		apiresponses.BadRequest(w, "Invalid request", err)
		return
	}

	courses, err := handler.Queries.GetCourses(r.Context(), *userID)
	if err != nil {
		apiresponses.InternalError(w, "Internal Error", err)
		return
	}

	apiresponses.Success(w, courses)

}
