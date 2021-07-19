package main

//go:generate go run . -o ./pkg/

import (
	"flag"
	"log"
	"os"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/cmd/gir_generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir_generate/genutil"
	"github.com/diamondburned/gotk4/gir/girgen"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
)

var (
	output  string
	verbose bool
	listPkg bool
)

func init() {
	flag.StringVar(&output, "o", "", "output directory to mkdir in")
	flag.BoolVar(&verbose, "v", verbose, "log verbosely (debug mode)")
	flag.BoolVar(&listPkg, "l", listPkg, "only list packages and exit")
	flag.Parse()

	if !listPkg && output == "" {
		log.Fatalln("Missing -o output directory.")
	}

	if !verbose {
		verbose = os.Getenv("GIR_VERBOSE") == "1"
	}
}

func main() {
	var repos gir.Repositories

	// Load all of gotk4's packages first.
	genutil.MustAddPackages(&repos, gendata.Packages)
	// Get a map of exteral imports for packages that gotk4 already generates.
	overrides := genutil.LoadExternOverrides(gotk4Module, repos)

	// Add our own packages in.
	genutil.MustAddPackages(&repos, packages)
	// Dump the added packages down.
	genutil.PrintAddedPkgs(repos)

	if listPkg {
		return
	}

	gen := girgen.NewGenerator(repos, genutil.ModulePath(adwaitaModule, overrides))
	gen.Logger = log.New(os.Stderr, "girgen: ", log.Lmsgprefix)
	gen.AddFilters(gendata.Filters)
	gen.AddFilters(filters)
	gen.ApplyPreprocessors(gendata.Preprocessors)

	if verbose {
		gen.LogLevel = logger.Debug
	}

	if err := genutil.CleanDirectory(output, pkgExceptions); err != nil {
		log.Fatalln("failed to clean output directory:", err)
	}

	if errors := genutil.GeneratePackages(gen, output, packages); len(errors) > 0 {
		for _, err := range errors {
			log.Println("generation error:", err)
		}
		os.Exit(1)
	}

	if err := genutil.EnsureDirectory(output, pkgExceptions, pkgGenerated); err != nil {
		log.Fatalln("error verifying generation:", err)
	}
}
