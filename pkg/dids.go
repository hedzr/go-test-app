// +build linux,windows,posix,freebsd,darwin,aix,arm_linux,netbsd,openbsd,plan9,solaris

package pkg

// GOARCH GOOS
// amd64p32 nacl

func TestNacl() string {
  return "normal"
}
