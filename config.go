package main

import (
	"fmt"
	"regexp"
)

type authUsers struct {
	users map[string]string
}

func (au *authUsers) String() string { _ = "STUB: not implemented"; return "" }

func (au *authUsers) Set(in string) error { _ = "STUB: not implemented"; return nil }

func (au *authUsers) add(user, pass string) { _ = "STUB: not implemented"; return }

func (au authUsers) isAllow(user, pass string) bool { _ = "STUB: not implemented"; return false }

// Config - config struct
type Config struct {
	port          int            // server port
	cache         int            // caching command out (in seconds)
	timeout       int            // timeout for shell command (in seconds)
	host          string         // server host
	exportVars    string         // list of environment vars for export to script
	shell         string         // custom shell
	defaultShell  string         // shell by default
	defaultShOpt  string         // shell option for one-liner (-c or /C)
	cert          string         // SSL certificate
	key           string         // SSL private key path
	auth          authUsers      // basic authentication
	exportAllVars bool           // export all current environment vars
	setCGI        bool           // set CGI variables
	setForm       bool           // parse form from URL
	noIndex       bool           // don't generate index page
	addExit       bool           // add /exit command
	oneThread     bool           // run each shell commands in one thread
	showErrors    bool           // returns the standard output even if the command exits with a non-zero exit code
	includeStderr bool           // also returns output written to stderr (default is stdout only)
	intServerErr  bool           // return 500 error if shell status code != 0
	formCheckRe   *regexp.Regexp // regexp for check form fields
}

// getConfig - parse arguments
func getConfig() (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

// setup log file

// readableURL - get readable URL for logging
func (cfg Config) readableURL(addr fmt.Stringer) string { _ = "STUB: not implemented"; return "" }
