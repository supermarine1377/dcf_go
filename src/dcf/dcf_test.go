package dcf

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/supermarine1377/dcf_go/src/dcf/condition"
)

func TestDCF(t *testing.T) {
	type args struct {
		cr  float64
		gr  float64
		tgr float64
		dr  float64
		gy  int
		y   int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "A case written in the book \"The Warren Buffett Way\" in page number 2766. This calcurates intrinic value of Coca-Cola in 1988",
			args: args{
				cr:  828,
				gr:  0.15,
				tgr: 0.05,
				dr:  0.09,
				gy:  10,
				y:   100,
			},
			want: 47109,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := condition.New(
				condition.WithCurrentEarnings(tt.args.cr),
				condition.WithGrowthRate(tt.args.gr),
				condition.WithTerminalGrowthRate(tt.args.tgr),
				condition.WithDiscountRate(tt.args.dr),
				condition.WithGrowthYears(tt.args.gy),
				condition.WithYears(tt.args.y),
			)
			assert.NoError(t, err)
			if got := DCF(c); got != tt.want {
				t.Errorf("DCF() = %v, want %v", got, tt.want)
			}
		})
	}
}
