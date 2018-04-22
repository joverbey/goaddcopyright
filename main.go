// Copyright (C) 2018 Jeffrey L. Overbey.  Use of this source code is governed
// by a BSD-style license posted at http://blog.jeff.over.bz/license/
package main

import (
	"os"

	"github.com/godoctor/godoctor/engine"
	"github.com/godoctor/godoctor/engine/cli"
	"github.com/joverbey/goaddcopyright/refactoring"
)

func main() {
	engine.AddRefactoring("addcopyright", new(refactoring.AddCopyright))
	os.Exit(cli.Run("Add Copyright Header", os.Stdin, os.Stdout, os.Stderr, os.Args))
}
