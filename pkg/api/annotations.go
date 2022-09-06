package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/pyroscope-io/pyroscope/pkg/model"
	"github.com/pyroscope-io/pyroscope/pkg/server/httputils"
	"github.com/pyroscope-io/pyroscope/pkg/service"
	"github.com/pyroscope-io/pyroscope/pkg/util/attime"
	"github.com/sirupsen/logrus"
)

type AnnotationsService interface {
	CreateAnnotation(ctx context.Context, params service.CreateAnnotationParams) (*model.Annotation, error)
}
type AnnotationsCtrl struct {
	log       *logrus.Logger
	svc       AnnotationsService
	httpUtils httputils.Utils
}

func NewAnnotationsCtrl(log *logrus.Logger, svc AnnotationsService, httpUtils httputils.Utils) *AnnotationsCtrl {
	return &AnnotationsCtrl{
		log:       log,
		svc:       svc,
		httpUtils: httpUtils,
	}
}

func (ctrl *AnnotationsCtrl) CreateHandler(w http.ResponseWriter, r *http.Request) {
	type createParams struct {
		AppName   string `json:"appName"`
		Timestamp int64  `json:"timestamp"`
		Content   string `json:"content"`
	}
	var params createParams
	//var params service.CreateAnnotationParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		ctrl.httpUtils.WriteInternalServerError(r, w, err, "failed to unmarshal JSON")
		return
	}

	if params.Timestamp == 0 {
		params.Timestamp = time.Now().Unix()
	}

	// TODO: parse parameters
	//	params := service.CreateAnnotationParams{
	//		AppName:   mux.Vars(r)["appName"],
	//		Timestamp: attime.Parse(v.Get("from")),
	//	}
	//
	annotation, err := ctrl.svc.CreateAnnotation(r.Context(), service.CreateAnnotationParams{
		AppName:   params.AppName,
		Timestamp: attime.Parse(strconv.FormatInt(params.Timestamp, 10)),
		Content:   params.Content,
	})
	if err != nil {
		// TODO: check parameter error
		ctrl.httpUtils.WriteInternalServerError(r, w, err, "failed to create annotation")
	}

	// TODO(eh-am): unify this with render.go
	type annotationsResponse struct {
		AppName   string `json:"appName"`
		Content   string `json:"content"`
		Timestamp int64  `json:"timestamp"`
	}
	annotationsResp := annotationsResponse{
		Content:   annotation.Content,
		Timestamp: annotation.From.Unix(),
	}
	ctrl.httpUtils.WriteResponseJSON(r, w, annotationsResp)
}