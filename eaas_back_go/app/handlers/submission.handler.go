package handlers

import (
	"eaas_back/app/models"
	submissionSrv "eaas_back/app/services/submission"
	"eaas_back/utility"
	"encoding/json"
	"net/http"
	"time"
)

// SubmissionHandler - handles subsmission request
type SubmissionHandler struct {
	us submissionSrv.SubmissionServiceInterface
}

func NewSubmissionAPI(submissionService submissionSrv.SubmissionServiceInterface) *SubmissionHandler {
	return &SubmissionHandler{
		us: submissionService,
	}
}

// Get godoc
// @Summary Get Profile
// @Description Get submission profile info
// @Tags submissions
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Submission
// @Success 200 {object} errorRes
// @Security ApiKeyAuth
// @Router /submissions [get]
func (h *SubmissionHandler) Get(w http.ResponseWriter, r *http.Request) {
	user, err := h.us.Get(r.Context())

	if err != nil {
		utility.Response(w, utility.NewHTTPError(utility.InternalError, 500))
	} else {
		utility.Response(w, utility.SuccessPayload(user, ""))
	}
}

// Update godoc
// @Summary Update submission
// @Description Update submission info
// @Tags submissions
// @Accept  json
// @Produce  json
// @Param   payload     body    models.SubmissionUpdate     true        "Submission Data"
// @Success 200 {object} basicResponse
// @Success 200 {object} errorRes
// @Security ApiKeyAuth
// @Router /submissions [put]
func (h *SubmissionHandler) Update(w http.ResponseWriter, r *http.Request) {
	updateSubmission := new(models.SubmissionUpdate)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&updateSubmission)
	result := make(map[string]interface{})
	err := h.us.Update(r.Context(), utility.GetLoggedInUserID(r), updateSubmission)
	if err != nil {
		result = utility.NewHTTPCustomError(utility.BadRequest, err.Error(), http.StatusBadRequest)
		utility.Response(w, result)
		return
	}

	result = utility.SuccessPayload(nil, "Successfully updated")
	utility.Response(w, result)
}

func (h *SubmissionHandler) Create(w http.ResponseWriter, r *http.Request) {
	payload := new(submissionReq)
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&payload)
	requestSub := &models.Submission{
		Name:      payload.Name,
		Email:     payload.Email,
		Status:    payload.Status,
		CreatedAt: time.Now(),
	}
	result := make(map[string]interface{})

	err := h.us.Create(r.Context(), requestSub)
	if err != nil {
		result = utility.NewHTTPError(utility.EntityCreationError, http.StatusBadRequest)
	} else {
		result = utility.SuccessPayload(nil, "Successfully registered", 201)
	}
	utility.Response(w, result)
}
