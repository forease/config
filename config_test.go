package config

import (
	//"errors"
	"flag"
	//"strings"
	"fmt"
	"testing"
)

func init() {
	flag.Parse()
}

var file = flag.String("f", "", "")

func TestParse(t *testing.T) {
	var f string
	if *file == "" {
		f = "toml.conf"
	} else {
		f = *file
	}
	c, err := NewConfig(f, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n\n======\n\n", c.Item)
		fmt.Println("String:")
		fmt.Println(c.String("servers.alpha.ip", "default"))
		fmt.Println("Int64:")
		fmt.Println(c.Int64("database.connection_max", 0))
		fmt.Println("Int:")
		fmt.Println(c.Int("database.connection_max2", 10))
		fmt.Println("Bool:")
		fmt.Println(c.Bool("database.enabled", false))
		fmt.Println(c.Bool("database.enabled2", false))
		fmt.Println("Float64:")
		fmt.Println(c.Float64("owner.float", 0.12))

		fmt.Println(c.Array("clients.hosts"))
	}
}
