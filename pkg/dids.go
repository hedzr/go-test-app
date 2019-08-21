// +build darwin dragonfly freebsd linux netbsd openbsd windows aix arm_linux plan9 solaris
// +build !nacl,!amd64p32

package pkg

// GOARCH GOOS
// amd64p32 nacl

func TestNacl() string {
  return "normal"
}
