package condition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCondition_Validate(t *testing.T) {
	type fields struct {
		currentEarnings    float64
		growthRate         float64
		terminalGrowthRate float64
		discountRate       float64
		years              int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Valid condition",
			fields: fields{
				currentEarnings:    1000,
				growthRate:         0.05,
				terminalGrowthRate: 0.03,
				discountRate:       0.1,
				years:              10,
			},
			wantErr: false,
		},
		{
			name: "Invalid current earnings",
			fields: fields{
				currentEarnings:    0,
				growthRate:         0.05,
				terminalGrowthRate: 0.03,
				discountRate:       0.1,
				years:              10,
			},
			wantErr: true,
		},
		{
			name: "Invalid growth rate",
			fields: fields{
				currentEarnings:    1000,
				growthRate:         1.5,
				terminalGrowthRate: 0.03,
				discountRate:       0.1,
				years:              10,
			},
			wantErr: true,
		},
		{
			name: "Invalid growth rate",
			fields: fields{
				currentEarnings:    1000,
				growthRate:         0.5,
				terminalGrowthRate: 1.5,
				discountRate:       0.1,
				years:              10,
			},
			wantErr: true,
		},
		{
			name: "Invalid discount rate",
			fields: fields{
				currentEarnings:    1000,
				growthRate:         0.5,
				terminalGrowthRate: 0.5,
				discountRate:       1.1,
				years:              10,
			},
			wantErr: true,
		},
		{
			name: "Invalid years",
			fields: fields{
				currentEarnings:    1000,
				growthRate:         0.5,
				terminalGrowthRate: 0.5,
				discountRate:       0.1,
				years:              0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Condition{
				currentEarnings:    tt.fields.currentEarnings,
				growthRate:         tt.fields.growthRate,
				terminalGrowthRate: tt.fields.terminalGrowthRate,
				discountRate:       tt.fields.discountRate,
				years:              tt.fields.years,
			}
			err := c.Validate()
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
