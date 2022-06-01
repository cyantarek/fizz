package marketing

import (
	"context"
	"reflect"
	"testing"
)

func TestMarketingService_GetCompleteStats(t *testing.T) {
	type fields struct {
		emailStats EmailStats
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Stats
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := marketing{
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
		emailStats EmailStats
	}
	tests := []struct {
		name string
		args args
		want *marketing
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
