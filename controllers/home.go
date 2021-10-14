package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"tpm/controllers/render"
	"tpm/models"
	"tpm/views"
)

type Home struct{}

func (c Home) Index(resp http.ResponseWriter, r *http.Request) {
	problems, err := models.ListAllProblems()
	if err != nil {
		fmt.Fprintln(resp, "db error:", err)
		return
	}

	highlight, err := strconv.Atoi(r.URL.Query().Get("highlight"))
	if err != nil {
		highlight = 0
	}

	viewData := &views.HomeIndexViewData{
		Highlight: highlight,
		Problems:  problems,
	}
	if err := render.View(resp, "home_index", viewData); err != nil {
		fmt.Fprintln(resp, "view error!!!!", err)
	}
}
