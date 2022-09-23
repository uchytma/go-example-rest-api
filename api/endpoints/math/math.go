package math

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func Routes( /* any dependency injection comes here*/ ) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", GetMathExample)
	return r
}

// @Summary get example object for test purposes.
// @ID math/test_get
// @Produce json
// @Success 200 {object} GetMathExampleResponse
// @Router /math [get]
// @Tags /math
// Render response with example JSON object
func GetMathExample(w http.ResponseWriter, r *http.Request) {
	resp := GetMathExampleResponse{
		ID:   51121,
		Name: "Test"}
	render.JSON(w, r, &resp)
}

// Test structure for example response
type GetMathExampleResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
