package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"login/mock_storage"
	"login/model"
	"login/storage"
	"testing"
	"time"
)

//Unit Test
func TestUserService_UpdatePassword(t *testing.T) {
	type fields struct {
		r        *storage.Storage
		hashCost int
	}
	cost := 6
	ctrl := gomock.NewController(t)
	dummyError := errors.New("dummy err")
	mockUser := mock_storage.NewMockIUserRepository(ctrl)
	mockUser.EXPECT().GetUserByEmail(context.Background(), "test").
		Return(&model.User{Email: "test", Password: []byte("123")}, nil)
	mockUser.EXPECT().GetUserByEmail(context.Background(), "not_found").Return(&model.User{}, dummyError)
	mockUser.EXPECT().Update(context.Background(), gomock.Any()).Return(nil)
	f := fields{r: &storage.Storage{DB: nil, UserRepo: mockUser,
		BookRepo: mock_storage.NewMockIBookRepository(ctrl), RecordRepo: mock_storage.NewMockIRecordRepository(ctrl)},
		hashCost: cost}

	type args struct {
		ctx context.Context
		req model.UserEmailPassReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			fields:  f,
			args:    args{ctx: context.Background(), req: model.UserEmailPassReq{Email: "test", Password: "123"}},
			wantErr: false,
		},
		{
			name:    "user not found",
			fields:  f,
			args:    args{ctx: context.Background(), req: model.UserEmailPassReq{Email: "not_found", Password: "123"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				r:        tt.fields.r,
				hashCost: tt.fields.hashCost,
			}
			gotResp, err := s.UpdatePassword(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdatePassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if compareDates(gotResp.UpdatedAt, time.Now()) {
				t.Errorf("Resp UpdatedAt error")
			}
		})
	}

}
func compareDates(resp time.Time, now time.Time) bool {
	year, month, day := resp.Date()
	y1, m1, d1 := now.Date()
	return year == y1 && month == m1 && day == d1
}
