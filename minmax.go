package winminmax

type minMaxSample struct {
	t uint32
	v uint32
}

// WinMinMax 结构体
type WinMinMax [3]minMaxSample

// Get 返回最大或最小值
func (w *WinMinMax) Get() uint32 {
	return w[0].v
}

// Reset 重置结构
func (w *WinMinMax) Reset(t, meas uint32) uint32 {
	w[0].t, w[1].t, w[2].t = t, t, t
	w[0].v, w[1].v, w[2].v = meas, meas, meas
	return w[0].v
}

// RunningMax 更新最大值
func (w *WinMinMax) RunningMax(win, t, meas uint32) uint32 {
	val := minMaxSample{t, meas}

	if meas >= w[0].v || /* found new max? */
		t-w[2].t > win { /* nothing left in window? */
		return w.Reset(t, meas) /* forget earlier samples */
	}

	if meas >= w[1].v {
		w[2], w[1] = val, val

	} else if meas >= w[2].v {
		w[2] = val
	}

	return w.subWinUpdate(win, &val)
}

// RunningMin 更新最小值
func (w *WinMinMax) RunningMin(win, t, meas uint32) uint32 {
	val := minMaxSample{t, meas}

	if meas <= w[0].v || /* found new min? */
		t-w[2].t > win { /* nothing left in window? */
		return w.Reset(t, meas) /* forget earlier samples */
	}

	if meas <= w[1].v {
		w[2], w[1] = val, val
	} else if meas <= w[2].v {
		w[2] = val
	}

	return w.subWinUpdate(win, &val)
}

func (w *WinMinMax) subWinUpdate(win uint32, val *minMaxSample) uint32 {
	dt := val.t - w[0].t

	if dt > win {
		/*
		 * Passed entire window without a new val so make 2nd
		 * choice the new val & 3rd choice the new 2nd choice.
		 * we may have to iterate this since our 2nd choice
		 * may also be outside the window (we checked on entry
		 * that the third choice was in the window).
		 */
		w[0] = w[1]
		w[1] = w[2]
		w[2] = *val
		if val.t-w[0].t > win {
			w[0] = w[1]
			w[1] = w[2]
			w[2] = *val
		}
	} else if w[1].t == w[0].t && dt > win/4 {
		/*
		 * We've passed a quarter of the window without a new val
		 * so take a 2nd choice from the 2nd quarter of the window.
		 */
		w[2], w[1] = *val, *val
	} else if w[2].t == w[1].t && dt > win/2 {
		/*
		 * We've passed half the window without finding a new val
		 * so take a 3rd choice from the last half of the window
		 */
		w[2] = *val
	}
	return w[0].v
}
