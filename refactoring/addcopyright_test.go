// Copyright (C) 2018 Jeffrey L. Overbey.  Use of this source code is governed
// by a BSD-style license posted at http://blog.jeff.over.bz/license/
package refactoring_test

import (
	"testing"

	"github.com/godoctor/godoctor/engine"
	"github.com/godoctor/godoctor/refactoring/testutil"
	"github.com/joverbey/goaddcopyright/refactoring"
)

func TestRefactorings(t *testing.T) {
	engine.AddRefactoring("addcopyright", new(refactoring.AddCopyright))

	refactoring.CurrentYear = "YYYY"

	const directory = "testdata/"
	testutil.TestRefactorings(directory, t)
}
