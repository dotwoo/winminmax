package winminmax

import (
	"testing"
)

func TestWinMinMax_Get(t *testing.T) {
	type fields struct {
		s [3]minMaxSample
	}
	tests := []struct {
		name   string
		fields fields
		want   uint32
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WinMinMax{
				s: tt.fields.s,
			}
			if got := w.Get(); got != tt.want {
				t.Errorf("WinMinMax.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWinMinMax_Reset(t *testing.T) {
	type fields struct {
		s [3]minMaxSample
	}
	type args struct {
		t    uint32
		meas uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint32
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WinMinMax{
				s: tt.fields.s,
			}
			if got := w.Reset(tt.args.t, tt.args.meas); got != tt.want {
				t.Errorf("WinMinMax.Reset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWinMinMax_RunningMax(t *testing.T) {
	type fields struct {
		s [3]minMaxSample
	}
	type args struct {
		win  uint32
		t    uint32
		meas uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint32
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WinMinMax{
				s: tt.fields.s,
			}
			if got := w.RunningMax(tt.args.win, tt.args.t, tt.args.meas); got != tt.want {
				t.Errorf("WinMinMax.RunningMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWinMinMax_RunningMin(t *testing.T) {
	type fields struct {
		s [3]minMaxSample
	}
	type args struct {
		win  uint32
		t    uint32
		meas uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint32
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WinMinMax{
				s: tt.fields.s,
			}
			if got := w.RunningMin(tt.args.win, tt.args.t, tt.args.meas); got != tt.want {
				t.Errorf("WinMinMax.RunningMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
