package handlers

import (
	"net/http"
	"testing"
)

func TestHandlerImpl_Encode(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		h       *HandlerImpl
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HandlerImpl{}
			if err := h.Encode(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("HandlerImpl.Encode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
