package hello

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	//before
	fmt.Println("Before unit test")

	m.Run()

	//after
	fmt.Println("After unit test")

}
func TestSayHello(t *testing.T) {
	result := SayHello("Najib")
	assert.Equal(t, "Hello Najib", result, "Result buan Hello Najib") // assert call Error when error through to next code
	fmt.Println("DOneeeeeeeeeeeeeeeeee Test Say Hello ")
}

func TestSayHi(t *testing.T) {
	result := SayHi("Najib")
	require.Equal(t, "Hi Najib", result, "result bukan Hi Najib") //require call Fatal when error it can be stopped
	fmt.Println("DOneeeeeeeeeeeeeeeeee Test Say Hi ")

}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Gabisa running di mac")
	}
	result := SayHello("ajib")
	require.Equal(t, "Hello Najib", result, "Bukan Hello Najib") //require(variable test, expected result, actual result, message)
}

func TestSubTest(t *testing.T) {
	t.Run(
		"Hello", func(t *testing.T) {
			result := SayHello("Najib")
			require.Equal(t, "Hello Najib", result, "Bukan Hello Najib")
		})

	t.Run(
		"Hi", func(t *testing.T) {
			result := SayHi("Najib")
			require.Equal(t, "Hi Najib", result, "Bukan Hi Najib")
		})

}

func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "SayHello(Najib)",
			request:  "Najib",
			expected: "Hello Najib",
		},
		{
			name:     "SayHello(Alyasyfi)",
			request:  "Alyasyfi",
			expected: "Hello Alyasyfi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SayHello(test.request)
			require.Equal(t, test.expected, result, "Bukan Hello"+test.request)
		})
	}
}

//Benchmark
func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Najib")
	}
}

func BenchmarkSayHi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHi("Najib")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("SayHello", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Najib")
		}
	})

	b.Run("SayHi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHi("Najib")
		}
	})

}

func BenchmarkTable(b *testing.B) {
	benchmarkVar := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:    "SayHello(Najib)",
			request: "Najib",
		},
		{
			name:    "SayHello(Alyasyfi)",
			request: "Alyasyfi",
		},
	}

	for _, benchmarks := range benchmarkVar {
		b.Run(benchmarks.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SayHello(benchmarks.request)
			}
		})
	}

}
