// Command create_mod packages a module source directory into a Go module zip
// in the canonical proxy.golang.org layout.
//
// Usage: create_mod <module-path> <version> <source-dir> <output-zip>
package main

import (
	"fmt"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "usage: create_mod <module-path> <version> <source-dir> <output-zip>")
		os.Exit(1)
	}
	modPath, version, srcDir, outZip := os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	f, err := os.Create(outZip)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create %s: %v\n", outZip, err)
		os.Exit(1)
	}
	defer f.Close()

	m := module.Version{Path: modPath, Version: version}
	if err := zip.CreateFromDir(f, m, srcDir); err != nil {
		fmt.Fprintf(os.Stderr, "CreateFromDir: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("wrote %s for %s@%s\n", outZip, modPath, version)
}
