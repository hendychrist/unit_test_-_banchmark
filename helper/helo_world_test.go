package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

// Table Benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Hendy Panjang",
			request: "Hendy Panjang",
		},
		{
			name:    "Christian Panjang",
			request: "Christian Panjang",
		},
		{
			name:    "Reyhan Panjang",
			request: "Reyhan Panjang",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

// Sub Benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Hendy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Hendy")
		}
	})
	b.Run("Christian", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Christian")
		}
	})
}

// Benchmark
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}

// Benchmark
func BenchmarkHelloHendy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Hendy")
	}
}

// Unit Test Table
func TestTable(t *testing.T) {
	test := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Hendy",
			request:  "Hendy",
			expected: "Hello Hendy",
		},
		{
			name:     "Jeri",
			request:  "Jeri",
			expected: "Hello Jeri",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
	}

	// Looping function dengan for looping
	for _, test := range test {

		// .Run fungsinya menjalankan semua tesitng pada 1 function tersebut
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}

}

// Sub test
func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Hendy")
		require.Equal(t, "Hello Hendy", result, "TestSkip v1")
	})
	t.Run("Reyhan", func(t *testing.T) {
		result := HelloWorld("Reyhan")
		require.Equal(t, "Hello Reyhan", result, "TestSkip v2")
	})
	t.Run("Hendy", func(t *testing.T) {
		result := HelloWorld("Hendy")
		require.Equal(t, "Hello Hendy", result, "TestSkip v2")
	})
}

func TestMain(m *testing.M) {
	// Before
	fmt.Println("BEFORE UNIT TEST")

	m.Run() // -> No tests to run

	fmt.Println("AFTER UNIT TEST")
}

// t.Skip
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Hendy")
	require.Equal(t, "TestSKip", result, "TestSkip muncul kalau di window")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result != Hello Eko")
	fmt.Println("TestHelloWorldAssert after assert")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result != Hello Eko")
	fmt.Println("TestHelloWorldAssert after assert")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")

	if result != "Hello Eko" {
		//panic("Result is not 'Hello Eko'")
		//t.Fail() // Sedangkan fail akan run dibawahnya walaupun menemui fail
		t.Error("--> result != Hello Eko")
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldHendy(t *testing.T) {
	result := HelloWorld("Hendy")

	if result != "Hello Hendy" {
		//panic("Result is not 'Hello Hendy'")
		//t.FailNow() // Menemui FailNow dia akan berenti tidak akan run dibawahnya
		t.Fatal("--> result != Hello Hendy")
	}
	fmt.Println("TestHelloWorldHendy Done")
}

func TestHelloWorldJeri(t *testing.T) {
	result := HelloWorld("Jeri")

	if result != "Hello Jeri" {
		//panic("Result is not 'Hello Hendy'")
		//t.FailNow() // Menemui FailNow dia akan berenti tidak akan run dibawahnya
		t.Fatal("--> result != Hello Jeri")
	}
	fmt.Println("TestHelloWorldJeri Done")
}
