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

//replace github.com/hedzr/go-socketlib => ../go-socketlib

replace github.com/hedzr/go-coaplib => ../go-coaplib

replace github.com/plgd-dev/go-coap/v2 => ../../../../../../Downloads/tmp/00.go-need-study/go-coap

require (
	github.com/Shopify/sarama v1.26.4
	github.com/hedzr/cmdr v1.7.25
	github.com/hedzr/cmdr-addons v1.7.25
	github.com/hedzr/go-coaplib v0.0.0-00010101000000-000000000000
	github.com/hedzr/go-socketlib v0.2.5
	github.com/hedzr/log v0.2.2
	github.com/hedzr/logex v1.2.15
	github.com/plgd-dev/go-coap/v2 v2.0.4
	github.com/sirupsen/logrus v1.6.0
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	golang.org/x/text v0.3.3
	gopkg.in/hedzr/errors.v2 v2.1.1
)
