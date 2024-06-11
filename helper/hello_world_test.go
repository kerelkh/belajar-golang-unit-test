package helper

import (
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// benchmark mirip dengan testing (untuk menghitung performa function)
// function benchmark diawali dengan Benchmark sama seperti Test
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kerel")
	}
}

func BenchmarkHelloWorldKerelAfif(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kerel Khalif")
	}
}

// sub benchmark
func BenchmarkHelloWorldSub(b *testing.B) {
	benchmarks := []struct {
		name string
	}{
		{"Kerel"},
		{"Kerel Khalif"},
		{"Afif"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bm.name)
			}
		})
	}
}

// function TestMain untuk menjalankan sekali per package
func TestMain(m *testing.M) {
	//sebelum test dimulai
	fmt.Println("Before all tests on package helper")

	//menjalankan semua test dalam package yang sama
	code := m.Run()

	//setelah test dijalankan
	fmt.Println("After all tests on package helper")
	os.Exit(code)
}

// Awal function harus Test
func TestHelloWorld(t *testing.T) {
	if got, want := HelloWorld("Kerel Khalif Afif"), "Hello, Kerel Khalif Afif"; got != want {
		t.Fail()
	}

	fmt.Println("TestHelloWorld done")
}

func TestNamaLain(t *testing.T) {
	if got, want := HelloWorld("Afif"), "Hello, Afif"; got != want {
		t.FailNow()
	}

	//tidak keluar karena t.FailNow()
	fmt.Println("TestNamaLain dont")
}

func TestError(t *testing.T) {
	if got, want := HelloWorld("Testing error"), "Hello, Testing error"; got != want {
		t.Error("Result must be hello, testing error")
	}
}

func TestFatal(t *testing.T) {
	if got, want := HelloWorld("Testing fatal"), "Hello, Testing fatal"; got != want {
		t.Fatal("Result must be Hello, Testing fatal")
	}
}

//cara jalankan test
//untuk menjalankan test dari root folder => go test -v ./...
//go test (untuk menjalankan semua test)
//go test -v (untuk tampilan yang lebih menarik)
//go test -v -run=function (untuk testing function tertentu)

//untuk mengagalkan unit test bisa menggunakan Fail() dan FailNow()
//Fail() => jika gagal namun tetap melanjutkan eksekusi unit test
//FailNow() => jika gagal akan menhentikan unit test saat ini
//Error() => gagal akan log kemudian menjalankan Fail()
//Fatal() => gagal akan log kemudian menjalankan FailNow()

//untuk assert gunakan library seperti testify/assert
//assert sama seperti Fail() diujung gagal test

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Kereldd")
	assert.Equal(t, "Hello, Kerel", result, "Return must be Hello, Kerel")
	fmt.Println("Test Hello World Assertion gagal")

}

//sedangkan testify/require saat gagal akan menjalankan FailNow() diujung gagal test saat ini

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("required")
	require.Equal(t, "Hello, require", result, "Return must be Hello, require")
	fmt.Println("test Hello World Require gagal")
}

// skip test untuk tidak menjalankan test jika kondisi tertentu (misal os != mac)
func TestHelloWorldSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Skipping test on macOS")
	}

	result := HelloWorld("Hello, this")
	require.Equal(t, "Hello, require", result, "Return must be Hello, require")
}
