module github.com/hedzr/go-test-app

go 1.12

// replace github.com/hedzr/errors v1.1.18

// replace gopkg.in/hedzr/errors.v2 => ../errors

// replace github.com/hedzr/logex => ../logex

// replace github.com/hedzr/cmdr => ../cmdr

// replace github.com/hedzr/cmdr-addons => ../cmdr-addons

// replace github.com/kardianos/service => ../../kardianos/service

require (
	github.com/Shopify/sarama v1.26.4
	github.com/hedzr/cmdr v1.6.47
	github.com/hedzr/cmdr-addons v1.0.11
	github.com/hedzr/logex v1.1.8
	github.com/sirupsen/logrus v1.6.0
	golang.org/x/net v0.0.0-20200520004742-59133d7f0dd7
	golang.org/x/text v0.3.2
	gopkg.in/hedzr/errors.v2 v2.0.12
)
