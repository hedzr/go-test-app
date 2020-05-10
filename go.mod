module github.com/hedzr/go-test-app

go 1.12

// replace github.com/hedzr/errors v1.1.18

// replace gopkg.in/hedzr/errors.v2 => ../errors

// replace github.com/hedzr/logex => ../logex

replace github.com/hedzr/cmdr => ../cmdr

// replace github.com/hedzr/cmdr-addons => ../cmdr-addons

// replace github.com/kardianos/service => ../../kardianos/service

require (
	github.com/Shopify/sarama v1.24.1
	github.com/hedzr/cmdr v1.6.35
	github.com/hedzr/logex v1.1.8
	github.com/klauspost/cpuid v1.2.2 // indirect
	github.com/sirupsen/logrus v1.4.2
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3
	golang.org/x/text v0.3.0
	gopkg.in/hedzr/errors.v2 v2.0.12
	gopkg.in/yaml.v2 v2.2.8
)
