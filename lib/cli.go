package lib

import (
	"flag"
	"fmt"
)

// Parse the flags and args
func ParseCLIArgs() (map[string]string, error) {
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
		if !IsValidDir(p) {
			return nil, fmt.Errorf("cli: %s is not a valid directory", p)
		}

		if IsDirectoryUnique(p) {
			Directories[p] = NewDirectory(p)
		}
	}
	return f, nil
}
