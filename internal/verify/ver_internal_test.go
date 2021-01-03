package verify_test

import (
	"github.com/GhvstCode/twillight/internal/app"
	"github.com/GhvstCode/twillight/internal/utils"
	"github.com/GhvstCode/twillight/internal/verify"
	"reflect"
	"testing"
)

func TestInternalNewVerificationService(t *testing.T) {
	type args struct {
		APIClient    app.Client
		friendlyName string
		opts         utils.VerOpts
	}
	tests := []struct {
		name    string
		args    args
		want    *verify.ResponseVerifyService
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := verify.InternalNewVerificationService(tt.args.APIClient, tt.args.friendlyName, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("InternalNewVerificationService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InternalNewVerificationService() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInternalStartPsd2Verification(t *testing.T) {
	type args struct {
		APIClient  app.Client
		serviceSid string
		to         string
		channel    string
		amount     string
		payee      string
	}
	tests := []struct {
		name    string
		args    args
		want    *verify.ResponseSendToken
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := verify.InternalStartPsd2Verification(tt.args.APIClient, tt.args.serviceSid, tt.args.to, tt.args.channel, tt.args.amount, tt.args.payee)
			if (err != nil) != tt.wantErr {
				t.Errorf("InternalStartPsd2Verification() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InternalStartPsd2Verification() got = %v, want %v", got, tt.want)
			}
		})
	}
}
