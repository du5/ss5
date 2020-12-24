package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/armon/go-socks5"
)

var flags struct {
	User string
	Pass string
	Port int
	Addr string
}

func main() {
	flag.StringVar(&flags.User, "u", "", "Auth user")
	flag.StringVar(&flags.Pass, "s", "", "Auth password")
	flag.IntVar(&flags.Port, "p", 8080, "Custom proxy port")
	flag.StringVar(&flags.Addr, "b", "0.0.0.0", "Custom bind IP")
	flag.Parse()

	if flags.User != "" && flags.Pass == "" {
		panic("Pass can not be empty.")
	}

	if flags.User == "" && flags.Pass != "" {
		panic("User can not be empty.")
	}

	if flags.Port < 1 || flags.Port > 65535 {
		panic("Port input error.")
	}

	authorization := "without password."
	conf := &socks5.Config{}

	if flags.User != "" && flags.Pass != "" {
		authorization = fmt.Sprintf("authorization by [%s:%s]", flags.User, flags.Pass)

		cred := socks5.StaticCredentials{flags.User: flags.Pass}
		conf.Credentials = cred
	}

	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	log.Printf("SOCKS5 started on %s:%d, %s", flags.Addr, flags.Port, authorization)

	if err := server.ListenAndServe("tcp", fmt.Sprintf("%s:%d", flags.Addr, flags.Port)); err != nil {
		panic(err)
	}
}
