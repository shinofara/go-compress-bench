package compressor

import (
	"os"
	"testing"
)

func BenchmarkCompress(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		compress()
	}
	b.StopTimer()

	os.Remove("sample.zip")
}

func BenchmarkCompressNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		compressNew()
	}

	os.Remove("sample.zip")
}
