package emails

import (
	"context"
	"reflect"
	"testing"
)

func TestEmailService_LookupStatus(t *testing.T) {
	type fields struct {
		emailRepository EmailRepository
		senders         EmailSenders
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *LookupEmail
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EmailService{
				emailRepository: tt.fields.emailRepository,
				senders:         tt.fields.senders,
			}
			got, err := e.LookupStatus(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LookupStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LookupStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmailService_Send(t *testing.T) {
	type fields struct {
		emailRepository EmailRepository
		senders         EmailSenders
	}
	type args struct {
		ctx     context.Context
		email   SendEmail
		backend string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EmailService{
				emailRepository: tt.fields.emailRepository,
				senders:         tt.fields.senders,
			}
			got, err := e.Send(tt.args.ctx, tt.args.email, tt.args.backend)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Send() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEmailService(t *testing.T) {
	type args struct {
		sender     EmailSender
		repository EmailRepository
	}
	tests := []struct {
		name string
		args args
		want *EmailService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmailService(tt.args.sender, tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmailService() = %v, want %v", got, tt.want)
			}
		})
	}
}
