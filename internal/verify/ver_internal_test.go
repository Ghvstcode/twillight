package verify

import (
	"github.com/GhvstCode/twillight/internal/app"
	"github.com/GhvstCode/twillight/internal/utils"
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
		want    *ResponseVerifyService
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InternalNewVerificationService(tt.args.APIClient, tt.args.friendlyName, tt.args.opts)
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
		want    *ResponseSendToken
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InternalStartPsd2Verification(tt.args.APIClient, tt.args.serviceSid, tt.args.to, tt.args.channel, tt.args.amount, tt.args.payee)
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
