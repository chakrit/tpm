package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"tpm/controllers/render"
	"tpm/models"
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
		fmt.Fprintln(resp, "model save error!")
	}

	http.Redirect(resp, r, "/?highlight="+strconv.Itoa(prob.ID), http.StatusFound)
}

func (p Problems) Edit(resp http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(resp, "Edit me please")
}
