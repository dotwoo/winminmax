package winminmax

import (
	"testing"
)

func TestWinMinMax_Reset(t *testing.T) {
	type args struct {
		t    uint32
		meas uint32
	}
	tests := []struct {
		name string
		w    *WinMinMax
		args args
		want uint32
	}{
		// TODO: Add test cases.
		{
			name: "reset1",
			w:    &WinMinMax{},
			args: args{1, 1},
			want: 1,
		},
		{
			name: "reset2",
			w:    &WinMinMax{},
			args: args{1, 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.Reset(tt.args.t, tt.args.meas); got != tt.want {
				t.Errorf("WinMinMax.Reset() = %v, want %v", got, tt.want)
			}
			if tt.w[0].t != tt.args.t || tt.w[1].t != tt.args.t || tt.w[2].t != tt.args.t {
				t.Errorf("WinMinMax.Reset()  the t %v, want %v", tt.w[0].t, tt.args.t)
			}
			if tt.w[0].v != tt.args.meas || tt.w[1].v != tt.args.meas || tt.w[2].v != tt.args.meas {
				t.Errorf("WinMinMax.Reset()  the v %v, want %v", tt.w[0].v, tt.args.meas)
			}
		})
	}
}

func TestWinMinMax_Get(t *testing.T) {
	tests := []struct {
		name string
		w    *WinMinMax
		want uint32
	}{
		// TODO: Add test cases.
		{
			name: "get1",
			w:    &WinMinMax{{1, 1}, {0, 0}, {2, 2}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.Get(); got != tt.want {
				t.Errorf("WinMinMax.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWinMinMax_RunningMax(t *testing.T) {
	type args struct {
		win  uint32
		t    uint32
		meas uint32
		time uint32
	}
	tests := []struct {
		name string
		w    *WinMinMax
		args args
		want uint32
	}{
		// TODO: Add test cases.
		{
			name: "runningmax10",
			w:    &WinMinMax{{3, 3}, {2, 2}, {1, 1}},
			args: args{
				win:  10,
				t:    4,
				meas: 10,
			},
			want: 10,
		},
		{
			name: "runningmax11",
			w:    &WinMinMax{{3, 11}, {2, 2}, {1, 1}},
			args: args{
				win:  10,
				t:    4,
				meas: 10,
			},
			want: 11,
		},
		{
			name: "runningmax11_10times",
			w:    &WinMinMax{{3, 11}, {2, 2}, {1, 1}},
			args: args{
				win:  10,
				t:    4,
				meas: 10,
				time: 10,
			},
			want: 10,
		},
		{
			name: "runningmax11_20times",
			w:    &WinMinMax{{3, 40}, {2, 2}, {1, 1}},
			args: args{
				win:  10,
				t:    4,
				meas: 10,
				time: 20,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got uint32
			for i := uint32(0); i <= tt.args.time; i++ {
				got = tt.w.RunningMax(tt.args.win, tt.args.t+i, tt.args.meas+i)
			}
			if got != tt.want+tt.args.time {
				t.Errorf("WinMinMax.RunningMax() = %v, want %v", got, tt.want+tt.args.time)
			}
		})
	}
}

func TestWinMinMax_RunningMin(t *testing.T) {
	type args struct {
		win  uint32
		t    uint32
		meas uint32
		time uint32
	}
	tests := []struct {
		name string
		w    *WinMinMax
		args args
		want uint32
	}{
		// TODO: Add test cases.
		{
			name: "runningmin1",
			w:    &WinMinMax{{3, 1}, {2, 2}, {1, 1}},
			args: args{
				win:  10,
				t:    4,
				meas: 10,
			},
			want: 1,
		},
		{
			name: "runningmin1",
			w:    &WinMinMax{{3, 2}, {2, 3}, {1, 4}},
			args: args{
				win:  10,
				t:    4,
				meas: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got uint32
			for i := uint32(0); i <= tt.args.time; i++ {
				got = tt.w.RunningMin(tt.args.win, tt.args.t+i, tt.args.meas+i)
			}
			if got != tt.want {
				t.Errorf("WinMinMax.RunningMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
