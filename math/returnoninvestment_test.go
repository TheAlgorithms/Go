package math

import (
	"testing"
)

func TestReturnOnInvestment(t *testing.T) {
	tests := []struct {
		name    string
		gain    float64
		cost    float64
		want    float64
		wantErr bool
	}{
		{"positive ROI", 1000, 500, 100, false},
		{"zero ROI", 500, 500, 0, false},
		{"negative ROI", 200, 500, -60, false},
		{"total loss", 0, 500, -100, false},
		{"zero cost", 1000, 0, 0, true},
		{"negative cost", 1000, -100, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReturnOnInvestment(tt.gain, tt.cost)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReturnOnInvestment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("ReturnOnInvestment() = %v, want %v", got, tt.want)
			}
		})
	}
}
