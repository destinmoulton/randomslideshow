package lib

import (
	"flag"
	"fmt"
	"os"
)

type CLIArgs struct {
	Flags map[string]string
	Paths []string
}

func GetCLIArgs() (*CLIArgs, error) {
	var ip string
	var port string
	flag.StringVar(&ip, "ip", "localhost", "server ip address")
	flag.StringVar(&port, "port", "3050", "server port")

	flag.Parse()

	var f = make(map[string]string)

	f["ip"] = ip
	f["port"] = port

	paths := flag.Args()
	if len(paths) == 0 {
		return nil, fmt.Errorf("cli: you must include a directory")
	}
	for _, p := range paths {
		if !isValidDir(p) {
			return nil, fmt.Errorf("cli: %s is not a valid directory", p)
		}
	}
	return &CLIArgs{f, paths}, nil
}

func isValidDir(dir string) bool {

	dirinfo, err := os.Stat(dir)

	return !os.IsNotExist(err) && dirinfo.IsDir()
}
