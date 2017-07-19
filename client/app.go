// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

//go:generate reactGen

package main

import (
	r "myitcv.io/react"

	"fmt"
)

type appDef struct {
	r.ComponentDef
}

func app() *appDef {
	res := &appDef{}
	r.BlessElement(res, nil)
	return res
}

func (a appDef) Render() r.Element {
	return r.Div(
		&r.DivProps{ClassName: "container mt-1"},
		r.Div(
			&r.DivProps{ClassName: "row"},
			r.Div(
				&r.DivProps{ClassName: "col-xs-8"},
				r.Div(
					&r.DivProps{ClassName: "preview"},
					r.Img(
						&r.ImgProps{
							ClassName: "img-responsive",
							Src:       "https://storage.googleapis.com/gopherizeme.appspot.com/artwork/010-Body/blue_gopher.png",
						},
					),
				),
			),
			r.Div(
				&r.DivProps{ClassName: "col-xs-4"},
				r.Button(
					&r.ButtonProps{ClassName: "btn btn-default", OnClick: buttonHandler{}},
					r.S("Shuffle"),
				),
				r.Button(
					&r.ButtonProps{ClassName: "btn btn-default", OnClick: buttonHandler{}},
					r.S("Reset"),
				),
			),
		),
	)
}

type buttonHandler struct{}

func (h buttonHandler) OnClick(e *r.SyntheticMouseEvent) {
	e.PreventDefault()
	fmt.Println("button clicked", e)
}
