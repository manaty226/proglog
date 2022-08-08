package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manaty226/proglog/testutil"
)

func TestHandleProduce_OK(t *testing.T) {
	t.Parallel()

	type want struct {
		status  int
		rspFile string
	}
	tests := map[string]struct {
		reqFile string
		want    want
	}{
		"ok": {
			reqFile: "testdata/http/ok_req.json",
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/http/ok_rsp.json",
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
			if resp.StatusCode != tt.want.status {
				t.Errorf("want %d but got: %d", resp.StatusCode, tt.want.status)
			}
		})
	}
}
