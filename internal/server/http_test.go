package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manaty226/proglog/testutil"
)

type want struct {
	status  int
	rspFile string
}

func TestHandleProduce(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		reqFile string
		want    want
	}{
		"ok": {
			reqFile: "testdata/http/produce/ok_req.json",
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/http/produce/ok_rsp.json",
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/",
				bytes.NewReader(testutil.LoadFile(t, tt.reqFile)),
			)
			httpsrv := newHTTPServer()
			httpsrv.handleProduce(w, r)
			resp := w.Result()
			testutil.AssertResponse(
				t, resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile),
			)
		})
	}
}

func TestHandleConsume(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		reqFile string
		want    want
	}{
		"ok": {
			reqFile: "testdata/http/consume/ok_req.json",
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/http/consume/ok_rsp.json",
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodGet,
				"/",
				bytes.NewReader(testutil.LoadFile(t, tt.reqFile)),
			)
			httpsrv := newHTTPServer()
			record := Record{Value: []byte("TGV0J3MgR28gIzMK"), Offset: 0}
			_, err := httpsrv.Log.Append(record)
			if err != nil {
				t.Fatal("test initialize error: %w", err)
			}
			httpsrv.handleProduce(w, r)
			resp := w.Result()
			testutil.AssertResponse(
				t, resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile),
			)
		})
	}
}
