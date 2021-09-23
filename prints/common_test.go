package prints

import "testing"

func BenchmarkCommonPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CommonPrint(".")
	}
}
