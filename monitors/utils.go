package monitors

func trimLeft(in []float32, size int) []float32 {
	if len(in) > size {
		return in[len(in)-size:]
	}
	return in
}

func avg(in []float32) (t float32) {
	for _, v := range in {
		t += v
	}
	t = t / float32(len(in))
	return
}
