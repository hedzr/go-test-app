
## about go playground



golang playground source codes: https://go.googlesource.com/playground





https://stackoverflow.com/questions/36409302/which-packages-may-be-imported-in-the-go-playground



### 1

The About button on the Playground gives some hint:

> The playground can use most of the standard library, with some exceptions.

By the *standard library* the packages of the *standard library* are meant, which are listed on the [**Packages**](https://golang.org/pkg/) page, under the [**Standard library**](https://golang.org/pkg/#stdlib) section. Packages listed under the [**Other**](https://golang.org/pkg/#other) section do not qualify (which is what you have tried - package [`golang.org/x/exp/ebnf`](https://godoc.org/golang.org/x/exp/ebnf) falls under the experimental and deprecated packages listed under the **Other** category).

A must-read link if you want to know more about the Playground implementation:

[**The Go Blog: Inside the Go Playground**](https://blog.golang.org/playground)

Here is an exhaustive playground test to import all the standard library packages to show they at least can be imported, but that doesn't mean everything (*or even anything*) can be reasonably used from them. The only package from the standard library that gives a compile error is `runtime/cgo`; "packages" without a buildable Go source file are not included for obvious reasons (because a folder is [not a package](https://golang.org/ref/spec#Packages) if it does not contain at least one buildable Go source file).

Here's the [Playground Link](http://play.golang.org/p/YcdVP6e76s) to try it yourself.

```golang
package main

import (
    _ "archive/tar"
    _ "archive/zip"

    _ "bufio"
    _ "bytes"

    _ "compress/bzip2"
    _ "compress/flate"
    _ "compress/gzip"
    _ "compress/lzw"
    _ "compress/zlib"

    _ "container/heap"
    _ "container/list"
    _ "container/ring"

    _ "crypto"
    _ "crypto/aes"
    _ "crypto/cipher"
    _ "crypto/des"
    _ "crypto/dsa"
    _ "crypto/ecdsa"
    _ "crypto/elliptic"
    _ "crypto/hmac"
    _ "crypto/md5"
    _ "crypto/rand"
    _ "crypto/rc4"
    _ "crypto/rsa"
    _ "crypto/sha1"
    _ "crypto/sha256"
    _ "crypto/sha512"
    _ "crypto/subtle"
    _ "crypto/tls"
    _ "crypto/x509"
    _ "crypto/x509/pkix"

    _ "database/sql"
    _ "database/sql/driver"

    _ "debug/dwarf"
    _ "debug/elf"
    _ "debug/gosym"
    _ "debug/macho"
    _ "debug/pe"
    _ "debug/plan9obj"

    _ "encoding"
    _ "encoding/ascii85"
    _ "encoding/asn1"
    _ "encoding/base32"
    _ "encoding/base64"
    _ "encoding/binary"
    _ "encoding/csv"
    _ "encoding/gob"
    _ "encoding/hex"
    _ "encoding/json"
    _ "encoding/pem"
    _ "encoding/xml"

    _ "errors"
    _ "expvar"
    _ "flag"
    _ "fmt"

    _ "go/ast"
    _ "go/build"
    _ "go/constant"
    _ "go/doc"
    _ "go/format"
    _ "go/importer"
    _ "go/parser"
    _ "go/printer"
    _ "go/scanner"
    _ "go/token"
    _ "go/types"

    _ "hash"
    _ "hash/adler32"
    _ "hash/crc32"
    _ "hash/crc64"
    _ "hash/fnv"

    _ "html"
    _ "html/template"

    _ "image"
    _ "image/color"
    _ "image/color/palette"
    _ "image/draw"
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"

    _ "index/suffixarray"

    _ "io"
    _ "io/ioutil"

    _ "log"
    _ "log/syslog"

    _ "math"
    _ "math/big"
    _ "math/cmplx"
    _ "math/rand"

    _ "mime"
    _ "mime/multipart"
    _ "mime/quotedprintable"

    _ "net"
    _ "net/http"
    _ "net/http/cgi"
    _ "net/http/cookiejar"
    _ "net/http/fcgi"
    _ "net/http/httptest"
    _ "net/http/httputil"
    _ "net/http/pprof"
    _ "net/mail"
    _ "net/rpc"
    _ "net/rpc/jsonrpc"
    _ "net/smtp"
    _ "net/textproto"
    _ "net/url"

    _ "os"
    _ "os/exec"
    _ "os/signal"
    _ "os/user"

    _ "path"
    _ "path/filepath"

    _ "reflect"
    _ "regexp"
    _ "regexp/syntax"

    _ "runtime"
    // _ "runtime/cgo"  // ERROR: missing Go type information
                        // for global symbol: .dynsym size 60
    _ "runtime/debug"
    _ "runtime/pprof"
    _ "runtime/race"
    _ "runtime/trace"

    _ "sort"
    _ "strconv"
    _ "strings"
    _ "sync"
    _ "sync/atomic"
    _ "syscall"

    _ "testing"
    _ "testing/iotest"
    _ "testing/quick"

    _ "text/scanner"
    _ "text/tabwriter"
    _ "text/template"
    _ "text/template/parse"

    _ "time"
    _ "unicode"
    _ "unicode/utf16"
    _ "unicode/utf8"
    _ "unsafe"
)

func main() {
    println("ok")
}
```





### 2

> I had trouble finding a list of what packages may be imported in the go playground at http://play.golang.org/.

Finding an exact list will be all the more difficult than you now (May. 16th 2019)

- can [import any package referenced by `GOPROXY`](https://stackoverflow.com/a/27813778/6309)
- can import your *own* package (defined in the "play.ground") namespace.
  See the [announcement from Brad Fitzpatrick](https://twitter.com/bradfitz/status/1128747022503165952)

> And now the #golang playground has support for multiple files: [Example](https://play.golang.org/p/KLZR7NlVZNX)

[Dmitri Shuralyov adds](https://twitter.com/dmitshur/status/1128817738678448128):

> Which means you can also have multiple modules! And `go.mod` files get formatted now too: [Example](https://play.golang.org/p/w3XRk47wo5u)

```golang
package main

import (
    "fmt"

    "gopher.example/bar"
    "play.ground/foo"
)

func main() {
    fmt.Println(foo.Bar)
    fmt.Println()
    fmt.Println(bar.Baz)
    fmt.Println()
    fmt.Println("And go.mod files get formatted too. Try pressing the Format button!")
}
-- go.mod --

  module      "play.ground"

   replace (

       "gopher.example/bar"    => ./bar
 )

-- foo/foo.go --
package foo

const Bar = "The Go playground now has support for multiple files!"
-- bar/go.mod --
module gopher.example/bar
-- bar/bar.go --
package bar

const Baz = "Which means you can have multiple modules too!"
```





### 3

It appears to depend on the environment in which the playground is launched. The relevant code in the source seems to be the [`compileAndRun`](https://github.com/golang/playground/blob/master/sandbox/sandbox.go#L67-L89) func, especially line 88:

```golang
cmd.Env = []string{"GOOS=nacl", "GOARCH=amd64p32", "GOPATH=" + os.Getenv("GOPATH")}
```

Which gets the GOPATH from the environment.

Other than that, the playground source does not have any explicit whitelist (or blacklist) of importable packages.

It's probably key to note that the `compileAndRun` func has no code to `go get` packages, so whatever is already in the GOPATH is what is available.

A perusal of the Makefile and Dockerfile does not reveal the specific deployment steps taken in the canonical [http://play.golang.org](http://play.golang.org/) site, so we simply have to rely on the documentation that Markus W Mahlberg pointed out; i.e. "a subset of the stdlib".

Also, you can deploy your own version of the go playground, and give it whatever GOPATH environment you choose.





## 3rd-party tools



- https://github.com/xiam/go-playground
- 