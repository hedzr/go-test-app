// +build !nacl,!amd64p32

package pkg

// GOARCH GOOS
// amd64p32 nacl

func TestNacl() string {
  return "normal"
}
