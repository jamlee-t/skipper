package config

import (
	"flag"

	log "github.com/sirupsen/logrus"
)

type deprecatedFlag struct {
	defaults map[string]interface{}
}

func initDeprecated() *deprecatedFlag {
	return &deprecatedFlag{defaults: make(map[string]interface{})}
}

func (df *deprecatedFlag) BoolVar(p *bool, name string, value bool, usage string) {
	df.defaults[name] = value
	flag.BoolVar(p, name, value, usage)
}

func (df *deprecatedFlag) StringVar(p *string, name string, value string, usage string) {
	df.defaults[name] = value
	flag.StringVar(p, name, value, usage)
}

func (df *deprecatedFlag) warn() {
	for name, value := range df.defaults {
		f := flag.Lookup(name)
		getter, ok := f.Value.(flag.Getter)
		if ok && getter.Get() != value {
			log.Warnf("%s: %s", f.Name, f.Usage)
		}
	}
}
