package otlp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/google/uuid"

	pushv1 "github.com/grafana/pyroscope/api/gen/proto/go/push/v1"
	typesv1 "github.com/grafana/pyroscope/api/gen/proto/go/types/v1"
	"github.com/grafana/pyroscope/pkg/tenant"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/pprofile/pprofileotlp"
)

type ingestHandler struct {
	pprofileotlp.UnimplementedGRPCServer
	svc PushService
	log log.Logger
}

// TODO(@petethepig): split http and grpc
type Handler interface {
	http.Handler
	pprofileotlp.GRPCServer
}

type PushService interface {
	Push(ctx context.Context, req *connect.Request[pushv1.PushRequest]) (*connect.Response[pushv1.PushResponse], error)
}

func NewOTLPIngestHandler(svc PushService, l log.Logger) Handler {
	return &ingestHandler{
		svc: svc,
		log: level.Error(l),
	}
}

// TODO(@petethepig): implement
// TODO(@petethepig): split http and grpc
func (h *ingestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

	req := &pushv1.PushRequest{}
	_, err := h.svc.Push(r.Context(), connect.NewRequest(req))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error: " + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// TODO(@petethepig): split http and grpc
func (h *ingestHandler) Export(ctx context.Context, er pprofileotlp.ExportRequest) (pprofileotlp.ExportResponse, error) {
	// TODO(@petethepig): make it tenant-aware
	ctx = tenant.InjectTenantID(ctx, tenant.DefaultTenantID)

	h.log.Log("msg", "Export called")

	rps := er.Profiles().ResourceProfiles()
	for i := 0; i < rps.Len(); i++ {
		rp := rps.At(i)

		labelsDst := []*typesv1.LabelPair{}
		// TODO(@petethepig): make labels work
		labelsDst = append(labelsDst, &typesv1.LabelPair{
			Name:  "__name__",
			Value: "process_cpu",
		})
		labelsDst = append(labelsDst, &typesv1.LabelPair{
			Name:  "service_name",
			Value: "otlp_test_app",
		})
		labelsDst = append(labelsDst, &typesv1.LabelPair{
			Name:  "__delta__",
			Value: "false",
		})
		labelsDst = append(labelsDst, &typesv1.LabelPair{
			Name:  "pyroscope_spy",
			Value: "unknown",
		})

		sps := rp.ScopeProfiles()
		for j := 0; j < sps.Len(); j++ {
			sp := sps.At(j)
			for k := 0; k < sp.Profiles().Len(); k++ {
				p := sp.Profiles().At(k)

				pprofBytes := pprofile.OprofToPprof(p)

				req := &pushv1.PushRequest{
					Series: []*pushv1.RawProfileSeries{
						{
							Labels: labelsDst,
							Samples: []*pushv1.RawSample{{
								RawProfile: pprofBytes,
								ID:         uuid.New().String(),
							}},
						},
					},
				}
				_, err := h.svc.Push(ctx, connect.NewRequest(req))
				if err != nil {
					h.log.Log("msg", "failed to push profile", "err", err)
					return pprofileotlp.NewExportResponse(), fmt.Errorf("failed to make a GRPC request: %w", err)
				}
			}
		}

	}

	return pprofileotlp.NewExportResponse(), nil
}