package application

import (
	"context"
	"fizz/internal/core/application/dto"
	"fizz/internal/core/port/outgoing"
	"reflect"
	"testing"
)

func TestMarketingService_GetCompleteStats(t *testing.T) {
	type fields struct {
		emailStats outgoing.EmailStats
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []dto.Stats
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MarketingService{
				emailStats: tt.fields.emailStats,
			}
			got, err := m.GetCompleteStats(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCompleteStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCompleteStats() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMarketingService(t *testing.T) {
	type args struct {
		emailStats outgoing.EmailStats
	}
	tests := []struct {
		name string
		args args
		want *MarketingService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMarketingService(tt.args.emailStats); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMarketingService() = %v, want %v", got, tt.want)
			}
		})
	}
}
