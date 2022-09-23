package math

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func Routes( /* any dependency injection comes here*/ ) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/greatest-common-divisor", GreatestCommonDivisor)
	return r
}

// Return greater common divisor (GCD) of specified numbers.
// @Summary Return greater common divisor (GCD) of specified numbers.
// @Tags /math
// @ID math/greatest-common-divisor
// @Produce json
// @Success 200 {object} uint
// @Router /math/greatest-common-divisor [post]
// @Param model body UintArray true "All numbers must be greater than zero."
func GreatestCommonDivisor(w http.ResponseWriter, r *http.Request) {

	data := UintArray{}

	if err := render.Bind(r, &data); err != nil {
		var err = &ErrResponse{
			HTTPStatusCode: 400,
			StatusText:     err.Error(),
		}
		render.Render(w, r, err)
		return
	}

	//todo: https://stackoverflow.com/questions/16628088/euclidean-algorithm-gcd-with-multiple-numbers
	var resp uint = 0
	for _, el := range data {
		resp += el
	}

	render.JSON(w, r, &resp)
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

type UintArray []uint

func (a *UintArray) Bind(r *http.Request) error {
	if a == nil {
		return errors.New("Missing request body.")
	}
	return nil
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}
