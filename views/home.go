package views

import "tpm/models"

type HomeIndexViewData struct {
	Highlight int
	Problems  []*models.Problem
}
