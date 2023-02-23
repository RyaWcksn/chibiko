package usecases

import (
	"context"
	"testing"

	"github.com/RyaWcksn/chibiko/api/v1/repositories"
	"github.com/RyaWcksn/chibiko/forms"
	"github.com/RyaWcksn/chibiko/pkgs/encryptions"
	"github.com/golang/mock/gomock"
)

func TestUsecaseImpl_Encode(t *testing.T) {
	ctrl := gomock.NewController(t)

	clientMock := repositories.NewMockIDatabase(ctrl)

	type args struct {
		ctx     context.Context
		payload *forms.EncodeRequest
	}
	tests := []struct {
		name     string
		args     args
		wantMock func()
		wantResp string
		wantErr  bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: &forms.EncodeRequest{
					Url:         "https://google.com",
					IsTemporary: 0,
				},
			},
			wantMock: func() {
				clientMock.EXPECT().Save(gomock.Any(), gomock.Any()).AnyTimes().Return(int64(1), nil)
			},
			wantResp: encryptions.Encode(int64(1)),
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantMock()
			uc := &UsecaseImpl{
				dbPort: clientMock,
			}
			gotResp, err := uc.Encode(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("UsecaseImpl.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResp != tt.wantResp {
				t.Errorf("UsecaseImpl.Encode() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
