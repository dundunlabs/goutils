package slices

func Map[S any, E any](s []S, fn func(v S) E) (r []E) {
	for i, v := range s {
		r[i] = fn(v)
	}
	return
}
