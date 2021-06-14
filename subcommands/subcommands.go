//* Lol wat?
package subcommands

//* Import all the things
import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli"

	"nebula/subcommands/installers"
)

func Install(c *cli.Context) {

	// TODO: Acctually check if the package exists
	var foundPackage bool = true

	//* Get the name of the package
	var packageName string = c.Args()[0]

	//* Check if the package exists
	if foundPackage {

		//* Run the installer
		installers.AppInstaller(c, packageName)

		//* if the package doesn't exist
	} else if !foundPackage {

		//* Tell the user what happened
		color.New(color.FgHiRed).Printf("Error:")
		fmt.Println(" Could not find any package called \"" + packageName + "\"")

		//* Exit with a non-zero exit code
		os.Exit(1)

	}

}

func Search(c *cli.Context) {

	//* Search all packages
	color.New(color.FgHiBlue).Printf("==>")
	fmt.Println(" Searching for \"" + c.Args()[0] + "\"")

}

func println(message string) { fmt.Println(message) }

func Licence() {
	println("BSD-2-Clause Licence")
	println("")
	println("Copyright (c) 2021-Present Ruben Flores. All rights reserved.")
	println("")
	println("Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:")
	println("")
	println(" 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.")
	println("")
	println(" 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.")
	println("")
	println("THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS \"AS IS\" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.")
}
