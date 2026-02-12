package _39

import "strings"

//https://100go.co/#under-optimized-strings-concatenation-39

func concat1Bad(values []string) string {
	s := ""
	for _, value := range values {
		//Why is this bad? Hint: Immutability of strings
		s += value
	}
	return s
}

func concat2Better(values []string) string {
	sb := strings.Builder{}
	for _, value := range values {
		//Why is this better?
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}

func concat3Best(values []string) string {
	total := 0
	for i := 0; i < len(values); i++ {
		total += len(values[i])
	}

	sb := strings.Builder{}
	//Grow grows b's capacity, if necessary, to guarantee space for another n bytes
	sb.Grow(total)
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}

//BenchmarkConcatV1-4             16      72291485 ns/op
//BenchmarkConcatV2-4           1188        878962 ns/op
//BenchmarkConcatV3-4           5922        190340 ns/op
