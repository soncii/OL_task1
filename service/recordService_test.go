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

//UNIT TEST
func TestRecordService_Create(t *testing.T) {
	type fields struct {
		r *storage.Storage
	}
	type args struct {
		ctx context.Context
		req model.RecordCreateReq
	}
	ctrl := gomock.NewController(t)
	dummyError := errors.New("dummy err")
	mockUser := mock_storage.NewMockIUserRepository(ctrl)
	mockBook := mock_storage.NewMockIBookRepository(ctrl)
	mockRecord := mock_storage.NewMockIRecordRepository(ctrl)
	mockRecord.EXPECT().CreateRecord(context.Background(), gomock.Any()).Return(nil).AnyTimes()
	mockUser.EXPECT().GetUserByID(context.Background(), uint(1)).Return(nil, nil).AnyTimes()
	mockUser.EXPECT().GetUserByID(context.Background(), uint(0)).Return(nil, dummyError).AnyTimes()
	mockBook.EXPECT().GetBookByID(context.Background(), uint(1)).Return(nil, nil).AnyTimes()
	mockBook.EXPECT().GetBookByID(context.Background(), uint(0)).Return(nil, dummyError).AnyTimes()

	f := fields{r: &storage.Storage{DB: nil, UserRepo: mockUser,
		BookRepo: mockBook, RecordRepo: mockRecord}}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.RecordCreateResp
		wantErr bool
	}{
		{
			name:   "success",
			fields: f,
			args: args{ctx: context.Background(),
				req: model.RecordCreateReq{
					UID: "1",
					BID: "1",
				}},
			want:    model.RecordCreateResp{RID: 1, CreatedAt: time.Now()},
			wantErr: false,
		},
		{
			name:   "uid invalid",
			fields: f,
			args: args{ctx: context.Background(),
				req: model.RecordCreateReq{
					UID: "invalid",
					BID: "1",
				}},
			want:    model.RecordCreateResp{},
			wantErr: true,
		},
		{
			name:   "bid invalid",
			fields: f,
			args: args{ctx: context.Background(),
				req: model.RecordCreateReq{
					UID: "1",
					BID: "invalid",
				}},
			want:    model.RecordCreateResp{},
			wantErr: true,
		},
		{
			name:   "uid negative",
			fields: f,
			args: args{ctx: context.Background(),
				req: model.RecordCreateReq{
					UID: "-1",
					BID: "1",
				}},
			want:    model.RecordCreateResp{},
			wantErr: true,
		},
		{
			name:   "bid negative",
			fields: f,
			args: args{ctx: context.Background(),
				req: model.RecordCreateReq{
					UID: "1",
					BID: "-1",
				}},
			want:    model.RecordCreateResp{},
			wantErr: true,
		},
		{
			name:   "user doesn't exist",
			fields: f,
			args: args{ctx: context.Background(),
				req: model.RecordCreateReq{
					UID: "0",
					BID: "1",
				}},
			want:    model.RecordCreateResp{},
			wantErr: true,
		},
		{
			name:   "book doesn't exist",
			fields: f,
			args: args{ctx: context.Background(),
				req: model.RecordCreateReq{
					UID: "1",
					BID: "0",
				}},
			want:    model.RecordCreateResp{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RecordService{
				r: tt.fields.r,
			}
			got, err := s.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if compareDates(got.CreatedAt, time.Now()) {
				t.Errorf("Resp UpdatedAt error")
			}
		})
	}
}
