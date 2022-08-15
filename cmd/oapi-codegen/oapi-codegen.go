package main

import (
	"flag"
	"fmt"
	"github.com/KosyanMedia/oapi-codegen/v2/pkg/codegen"
	"github.com/KosyanMedia/oapi-codegen/v2/pkg/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"runtime/debug"
	"strings"
)

func errExit(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

var (
	flagConfigFile   string
	flagPrintVersion bool
	flagGenerate     string
)

type configuration struct {
	codegen.Configuration `yaml:",inline"`

	// OutputFile is the filename to output.
	OutputFile string `yaml:"output,omitempty"`
}

func main() {
	flag.StringVar(&flagConfigFile, "config", "", "a YAML config file that controls oapi-codegen behavior")
	flag.BoolVar(&flagPrintVersion, "version", false, "when specified, print version and exit")
	flag.StringVar(&flagGenerate, "generate", "", "Comma-separated list of subsets to be generated")

	flag.Parse()

	if flagPrintVersion {
		bi, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Fprintln(os.Stderr, "error reading build info")
			os.Exit(1)
		}
		fmt.Println(bi.Main.Path + "/cmd/oapi-codegen")
		fmt.Println(bi.Main.Version)
		return
	}

	if flag.NArg() < 1 {
		fmt.Println("Please specify a path to a OpenAPI 3.0 spec file")
		os.Exit(1)
	}

	var opts configuration
	// We simply read the configuration from disk.
	if flagConfigFile != "" {
		buf, err := ioutil.ReadFile(flagConfigFile)
		if err != nil {
			errExit("error reading config file '%s': %v", flagConfigFile, err)
		}
		err = yaml.Unmarshal(buf, &opts)
		if err != nil {
			errExit("error parsing'%s' as YAML: %v", flagConfigFile, err)
		}
	}

	applyFlags(&opts)

	// Ensure default values are set if user hasn't specified some needed
	// fields.
	opts.Configuration = opts.UpdateDefaults()

	// Now, ensure that the config options are valid.
	if err := opts.Validate(); err != nil {
		errExit("configuration error: %v", err)
	}

	swagger, err := util.LoadSwagger(flag.Arg(0))
	if err != nil {
		errExit("error loading swagger spec in %s\n: %s", flag.Arg(0), err)
	}

	code, err := codegen.Generate(swagger, opts.Configuration)
	if err != nil {
		errExit("error generating code: %s\n", err)
	}

	if opts.OutputFile != "" {
		err = ioutil.WriteFile(opts.OutputFile, []byte(code), 0644)
		if err != nil {
			errExit("error writing generated code to file: %s", err)
		}
	} else {
		fmt.Print(code)
	}
}

func applyFlags(opts *configuration) {
	for _, subset := range strings.Split(flagGenerate, ",") {
		switch subset {
		case "chi-server":
			opts.Generate.ChiServer = true
		case "echo-server":
			opts.Generate.EchoServer = true
		case "gin-server":
			opts.Generate.GinServer = true
		case "client":
			opts.Generate.Client = true
		case "models":
			opts.Generate.Models = true
		case "embedded-spec":
			opts.Generate.EmbeddedSpec = true
		}
	}
}
