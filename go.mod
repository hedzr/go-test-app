module github.com/hedzr/go-test-app

go 1.12

// replace github.com/hedzr/errors v1.1.18

// replace gopkg.in/hedzr/errors.v2 => ../errors

// replace github.com/hedzr/cmdr => ../cmdr

// replace github.com/hedzr/log => ../log

// replace github.com/hedzr/logex => ../logex

// replace github.com/hedzr/cmdr-addons => ../cmdr-addons

// replace github.com/kardianos/service => ../../kardianos/service

// replace github.com/hedzr/go-ringbuf => ../go-ringbuf

replace github.com/hedzr/go-socketlib => ../go-socketlib

require (
	github.com/Shopify/sarama v1.26.4
	github.com/hedzr/cmdr v1.7.3
	github.com/hedzr/cmdr-addons v1.1.3
	github.com/hedzr/go-socketlib v0.0.0-00010101000000-000000000000
	github.com/hedzr/log v0.1.15
	github.com/hedzr/logex v1.2.7
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.6.0
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	golang.org/x/text v0.3.3
	golang.org/x/tools v0.0.0-20200131211209-ecb101ed6550 // indirect
	gopkg.in/hedzr/errors.v2 v2.0.12
)
