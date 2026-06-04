package main

import (
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"

	"github.com/msoap/raphanus"
)

var version = "dev"

const (
	// defaultPort - default port for http-server
	defaultPort = 8080

	// shBasicAuthVar - name of env var for basic auth credentials
	shBasicAuthVar = "SH_BASIC_AUTH"

	// defaultShellPOSIX - shell executable by default in POSIX systems
	defaultShellPOSIX = "sh"

	// defaultShellWindows - shell executable by default in Windows
	defaultShellWindows = "cmd"

	// defaultShellPlan9 - shell executable by default in Plan9
	defaultShellPlan9 = "rc"

	maxHTTPCode            = 1000
	maxMemoryForUploadFile = 65536
)

// indexTmpl - template for index page
const indexTmpl = `<!DOCTYPE html>
<!-- Served by shell2http/%s -->
<html>
<head>
    <title>❯ shell2http</title>
    <style>
    body {
        font-family: sans-serif;
    }
    li {
        list-style-type: none;
    }
    li:before {
        content: "❯";
        padding-right: 5px;
    }
    </style>
</head>
<body>
	<h1>shell2http</h1>
	<ul>
		%s
	</ul>
	Get from: <a href="https://github.com/msoap/shell2http">github.com/msoap/shell2http</a>
</body>
</html>
`

// command - one command
type command struct {
	path       string
	cmd        string
	httpMethod string
	handler    http.HandlerFunc
}

// parsePathAndCommands - get all commands with pathes
func parsePathAndCommands(args []string) ([]command, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getShellAndParams - get default shell and command
func getShellAndParams(cmd string, appConfig Config) (shell string, params []string, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// sh -c "cmd"

// custom shell

// getShellHandler - get handler function for one shell command
func getShellHandler(appConfig Config, shell string, params []string, cacheTTL raphanus.DB) func(http.ResponseWriter, *http.Request) {
	_ = "STUB: not implemented"
	return nil
}

// execShellCommand - execute shell command, returns bytes out and error
func execShellCommand(appConfig Config, shell string, params []string, req *http.Request, cacheTTL raphanus.DB) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// cache hit
// TODO: save exit code in cache

// #nosec

// get request body data data to stdin of script (if not parse form vars above)

// setupHandlers - setup http handlers
func setupHandlers(cmdHandlers []command, appConfig Config, cacheTTL raphanus.DB) ([]command, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// map[path][http-method]handler

// --------------

// --------------

// responseWrite - write text to response
func responseWrite(rw io.Writer, text string) { _ = "STUB: not implemented"; return }

// setCGIEnv - set some CGI variables
func setCGIEnv(cmd *exec.Cmd, req *http.Request, appConfig Config) {
	_ = "STUB: not implemented"
	// set HTTP_* variables
	return
}

/*
	parse headers from script output:

Header-name1: value1\n
Header-name2: value2\n
\n
text
*/
func parseCGIHeaders(shellOut string) (string, map[string]string) {
	_ = "STUB: not implemented"
	return "", nil
}

// headers is not valid, return all text

// headers don't found, return all text

// getForm - parse form into environment vars, also handle uploaded files
func getForm(cmd *exec.Cmd, req *http.Request, checkFormRe *regexp.Regexp) (func(), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// handle uploaded files, save all to temporary files and set variables filename_XXX, filepath_XXX

// isMultipartFormData - check header for multipart/form-data
func isMultipartFormData(headers http.Header) bool { _ = "STUB: not implemented"; return false }

// proxySystemEnv - proxy some system vars
func proxySystemEnv(cmd *exec.Cmd, appConfig Config) { _ = "STUB: not implemented"; return }

// errChain - handle errors on few functions
func errChain(chainFuncs ...func() error) error { _ = "STUB: not implemented"; return nil }

// errChainAll - handle errors on few functions, exec all func and returns the first error
func errChainAll(chainFuncs ...func() error) error { _ = "STUB: not implemented"; return nil }

func main() {
	appConfig, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	cmdHandlers, err := parsePathAndCommands(flag.Args())
	if err != nil {
		log.Fatalf("failed to parse arguments: %s", err)
	}

	var cacheTTL raphanus.DB
	if appConfig.cache > 0 {
		cacheTTL = raphanus.New()
	}

	cmdHandlers, err = setupHandlers(cmdHandlers, *appConfig, cacheTTL)
	if err != nil {
		log.Fatal(err)
	}
	for _, handler := range cmdHandlers {
		handlerFunc := handler.handler
		if len(appConfig.auth.users) > 0 {
			handlerFunc = mwBasicAuth(handlerFunc, appConfig.auth)
		}
		if appConfig.oneThread {
			handlerFunc = mwOneThread(handlerFunc)
		}
		handlerFunc = mwLogging(mwCommonHeaders(handlerFunc))

		http.HandleFunc(handler.path, handlerFunc)
		log.Printf("register: %s (%s)\n", handler.path, handler.cmd)
	}

	listener, err := net.Listen("tcp", net.JoinHostPort(appConfig.host, strconv.Itoa(appConfig.port)))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("listen %s\n", appConfig.readableURL(listener.Addr()))

	if len(appConfig.cert) > 0 && len(appConfig.key) > 0 {
		log.Fatal(http.ServeTLS(listener, nil, appConfig.cert, appConfig.key))
	} else {
		log.Fatal(http.Serve(listener, nil))
	}
}
