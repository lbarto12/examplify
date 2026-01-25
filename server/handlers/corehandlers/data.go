package corehandlers

import (
	"net/http"
	"server/api/apirequests"
	"server/api/apiresponses"
	"server/handlers/generated/gencore"
	"server/sqlc/sqlgen"

	"github.com/labstack/gommon/log"
)

// (GET /core/course/{courseID}/collections)
func (handler Handler) GetCourseCollections(w http.ResponseWriter, r *http.Request, courseID string) {
	userID, err := apirequests.User(r)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	qtx := sqlgen.New(handler.Services.Postgres)

	collections, err := qtx.GetCourseCollections(r.Context(), sqlgen.GetCourseCollectionsParams{
		UserID:   *userID,
		CourseID: courseID,
	})
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Internal Error", http.StatusInternalServerError)
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
func (handler Handler) GetCourses(w http.ResponseWriter, r *http.Request) {
	userID, err := apirequests.User(r)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	qtx := sqlgen.New(handler.Services.Postgres)

	courses, err := qtx.GetCourses(r.Context(), *userID)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	apiresponses.Success(w, courses)

}
