// Copyright (C) 2018 Jeffrey L. Overbey.  Use of this source code is governed
// by a BSD-style license posted at http://blog.jeff.over.bz/license/
package refactoring

import (
	"fmt"
	"strconv"
	"time"

	"github.com/godoctor/godoctor/analysis/names"
	"github.com/godoctor/godoctor/refactoring"
	"github.com/godoctor/godoctor/text"
)

var CurrentYear string = strconv.Itoa(time.Now().Year())

type AddCopyright struct {
	refactoring.RefactoringBase
}

func (r *AddCopyright) Description() *refactoring.Description {
	return &refactoring.Description{
		Name: "Add Copyright Header",
		//          ----+----1----+----2----+----3----+----4----+----5
		Synopsis:  "Add a copyright header to a file",
		Usage:     "addcopyright <text>",
		Multifile: false,
		Params: []refactoring.Parameter{{
			Label:        "Copyright Owner:",
			Prompt:       "Name to insert into the copyright text.",
			DefaultValue: ""}},
		Hidden: false,
	}
}

func (r *AddCopyright) Run(config *refactoring.Config) *refactoring.Result {
	r.Init(config, r.Description())
	r.Log.ChangeInitialErrorsToWarnings()
	if r.Log.ContainsErrors() {
		return &r.Result
	}

	extent := r.findInComments("Copyright")
	if extent != nil {
		file := r.Program.Fset.File(r.File.Package)
		startPos := file.Pos(extent.Offset)
		endPos := file.Pos(extent.OffsetPastEnd())

		r.Log.Error("An existing copyright was found.")
		r.Log.AssociatePos(startPos, endPos)
		return &r.Result
	}

	r.addCopyright(config.Args[0].(string))
	r.FormatFileInEditor()
	return &r.Result
}

func (r *AddCopyright) findInComments(text string) *text.Extent {
	occurrences := names.FindInComments(text, r.File, nil, r.Program.Fset)
	if len(occurrences) == 0 {
		return nil
	}
	return occurrences[0]
}

func (r *AddCopyright) addCopyright(name string) {
	extentToReplace := &text.Extent{0, 0}
	possibleSpace := " "
	if name == "" {
		possibleSpace = ""
	}
	text := fmt.Sprintf("// Copyright %s%s%s.  All rights reserved.\n",
		CurrentYear, possibleSpace, name)
	r.Edits[r.Filename].Add(extentToReplace, text)
}
