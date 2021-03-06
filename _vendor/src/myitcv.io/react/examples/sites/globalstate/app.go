package main

import (
	r "myitcv.io/react"
	"myitcv.io/react/examples/sites/globalstate/state"
)

type AppDef struct {
	r.ComponentDef
}

type AppState struct {
	hideViewer bool
}

func App() *AppDef {
	res := new(AppDef)
	r.BlessElement(res, nil)
	return res
}

func (a *AppDef) Render() r.Element {
	var viewer *r.DivDef
	var showHide *r.ButtonDef

	if a.State().hideViewer {
		showHide = r.Button(
			&r.ButtonProps{OnClick: hideClick{a, false}},
			r.S("Show viewer"),
		)
	} else {
		viewer = r.Div(nil,
			r.H3(nil, r.S("Person Viewer")),
			PersonViewer(),
		)
		showHide = r.Button(
			&r.ButtonProps{OnClick: hideClick{a, true}},
			r.S("Hide viewer"),
		)
	}

	return r.Div(&r.DivProps{ClassName: "container"},
		r.H3(nil, r.S("Person Chooser")),
		PersonChooser(PersonChooserProps{
			PersonState: state.State.CurrentPerson(),
		}),
		r.HR(nil),
		showHide,
		viewer,
	)
}

type hideClick struct {
	*AppDef
	showHide bool
}

func (h hideClick) OnClick(e *r.SyntheticMouseEvent) {
	h.AppDef.SetState(AppState{
		hideViewer: h.showHide,
	})
}
