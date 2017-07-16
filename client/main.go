// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package main

import (
	r "myitcv.io/react"

	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document()

func main() {
	domTarget := document.GetElementByID("gopherize.me")

	d := r.Div(nil, r.S("Hello world"))

	r.Render(d, domTarget)
}
