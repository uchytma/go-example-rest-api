package math

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/uchytma/go-example-rest-api/internal/math"
)

func Routes( /* any dependency injection comes here*/ ) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/greatest-common-divisor", GreatestCommonDivisor)
	r.Post("/least-common-multiple", LeastCommonMultiple)
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
		RenderErrorResponse(w, r, &err)
		return
	}

	resp, err := math.Gcd{}.Calculate(data...)

	if err != nil {
		RenderErrorResponse(w, r, &err)
		return
	}

	render.JSON(w, r, &resp)
}

// Return least common multiple (LCM) of specified numbers.
// @Summary Return least common multiple (LCM) of specified numbers.
// @Tags /math
// @ID math/least-common-multiple
// @Produce json
// @Success 200 {object} uint
// @Router /math/least-common-multiple [post]
// @Param model body UintArray true "All numbers must be greater than zero."
func LeastCommonMultiple(w http.ResponseWriter, r *http.Request) {

	data := UintArray{}

	if err := render.Bind(r, &data); err != nil {
		RenderErrorResponse(w, r, &err)
		return
	}

	resp, err := math.Lcm{}.Calculate(data...)

	if err != nil {
		RenderErrorResponse(w, r, &err)
		return
	}

	render.JSON(w, r, &resp)
}

func RenderErrorResponse(w http.ResponseWriter, r *http.Request, errorData *error) {
	var err = &ErrResponse{
		HTTPStatusCode: 400,
		StatusText:     (*errorData).Error(),
	}
	render.Render(w, r, err)
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
	HTTPStatusCode int    `json:"-"`
	StatusText     string `json:"status"`
}
