package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"tpm/controllers/render"
	"tpm/models"
	"tpm/views"

	"github.com/go-chi/chi/v5"
)

type Problems struct{}

func (p Problems) New(resp http.ResponseWriter, r *http.Request) {
	if err := render.View(resp, "problems_new", nil); err != nil {
		fmt.Fprintln(resp, "view error!!!!")
	}
}

func (p Problems) Create(resp http.ResponseWriter, r *http.Request) {
	prob, err := models.CreateProblem(&models.Problem{
		Summary:     r.PostFormValue("summary"),
		Description: r.PostFormValue("description"),
	})
	if err != nil {
		fmt.Fprintln(resp, "db error", err)
	} else {
		http.Redirect(resp, r, "/?highlight="+strconv.Itoa(prob.ID), http.StatusFound)
	}
}

func (p Problems) Edit(resp http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(resp, "Edit me please")
}

func (p Problems) View(resp http.ResponseWriter, r *http.Request) {
	problemId, err := strconv.Atoi(chi.URLParam(r, "problemId"))
	if err != nil {
		fmt.Fprintln(resp, "invalid problem id")
		return
	}

	prob, err := models.GetProblem(problemId)
	if err != nil {
		fmt.Fprintln(resp, "Database error", err)
		return
	}

	viewData := views.ProblemViewViewData{
		Problem: prob,
	}
	if err := render.View(resp, "problems_view", viewData); err != nil {
		fmt.Fprintln(resp, "view render error", err)
	}
}
