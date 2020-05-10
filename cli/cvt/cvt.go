// Copyright © 2020 Hedzr Yeh.

package cvt

import (
	"bytes"
	"fmt"
	"github.com/hedzr/cmdr"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"gopkg.in/hedzr/errors.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Cvt(cmd *cmdr.Command, args []string) (err error) {
	fmt.Printf("- CVT: %v\n", args)

	ec := errors.NewContainer("Loops")
	defer func() { err = ec.Error() }()

	for ix, s := range args {
		fmt.Printf("%5d. %s => %s\n", ix, s, s)
		ec.Attach(dirWalk(s))
	}

	return
}

func dirWalk(dir string) (err error) {
	err = os.Chdir(dir)
	if err != nil {
		return
	}

	err = filepath.Walk(dir, func(pathX string, infoX os.FileInfo, errX error) error {
		// first thing to do, check error. and decide what to do about it
		if errX != nil {
			fmt.Printf("error 「%v」 at a path 「%q」\n", errX, pathX)
			return errX
		}

		runes := []rune(pathX)
		fmt.Printf("pathX: %v | runes = % X\n", pathX, runes)
		if r1, e1 := Utf8ToGbk([]byte(pathX)); e1 == nil {
			fmt.Printf("  err: %v | r1 = %v\n", e1, string(r1))
			if r2, e2 := Utf8ToGbk([]byte(r1)); e2 == nil {
				fmt.Printf("  err: %v | r2 = %v\n", e2, string(r2))
			}
		}

		r, err := charset.NewReader(strings.NewReader(pathX), "latin-1")
		if err != nil {
			logrus.Fatal(err)
		}
		result, err := ioutil.ReadAll(r)
		if err != nil {
			logrus.Fatal(err)
		}
		fmt.Printf("%s\n", result)

		fmt.Printf("pathX: %v\n", pathX)

		// find out if it's a dir or file, if file, print info
		if infoX.IsDir() {
			// fmt.Printf("is dir.\n")
		} else {
			// fmt.Printf("  dir: 「%v」\n", filepath.Dir(pathX))
			// fmt.Printf("  file name 「%v」\n", infoX.Name())
			// fmt.Printf("  extenion: 「%v」\n", filepath.Ext(pathX))
		}
		return nil
	})
	return
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
