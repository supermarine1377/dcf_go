package dcf

import "testing"

func TestDCF(t *testing.T) {
	type args struct {
		currentCashFlow float64
		growthRate      float64
		disountRate     float64
		evebit          float64
		years           int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1st",
			args: args{
				currentCashFlow: 60,
				growthRate:      0.15,
				disountRate:     0.082,
				evebit:          15.0,
				years:           10,
			},
			want: 6703.72,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DCF(tt.args.currentCashFlow, tt.args.growthRate, tt.args.disountRate, tt.args.evebit, tt.args.years); got != tt.want {
				t.Errorf("DCF() = %v, want %v", got, tt.want)
			}
		})
	}
}
