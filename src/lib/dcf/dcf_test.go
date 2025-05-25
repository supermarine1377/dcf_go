package dcf

import (
	"testing"

	"github.com/supermarine1377/dcf_go/src/lib/dcf/condition"
)

func TestDCF(t *testing.T) {
	type args struct {
		cr  float64
		gr  float64
		fgr float64
		dr  float64
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
				fgr: 0.05,
				dr:  0.09,
				y:   100,
			},
			want: 47109,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := condition.New(tt.args.cr, tt.args.gr, tt.args.fgr, tt.args.dr, tt.args.y)
			if got := DCF(c); got != tt.want {
				t.Errorf("DCF() = %v, want %v", got, tt.want)
			}
		})
	}
}
