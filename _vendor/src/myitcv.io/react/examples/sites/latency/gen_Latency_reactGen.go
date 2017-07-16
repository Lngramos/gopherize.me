// Code generated by reactGen. DO NOT EDIT.

package main

import "myitcv.io/react"

func (l *LatencyDef) ShouldComponentUpdateIntf(nextProps, prevState, nextState interface{}) bool {
	res := false

	v := prevState.(LatencyState)
	res = !v.EqualsIntf(nextState) || res
	return res
}

// SetState is an auto-generated proxy proxy to update the state for the
// Latency component.  SetState does not immediately mutate l.State()
// but creates a pending state transition.
func (l *LatencyDef) SetState(state LatencyState) {
	l.ComponentDef.SetState(state)
}

// State is an auto-generated proxy to return the current state in use for the
// render of the Latency component
func (l *LatencyDef) State() LatencyState {
	return l.ComponentDef.State().(LatencyState)
}

// IsState is an auto-generated definition so that LatencyState implements
// the myitcv.io/react.State interface.
func (l LatencyState) IsState() {}

var _ react.State = LatencyState{}

// GetInitialStateIntf is an auto-generated proxy to GetInitialState
func (l *LatencyDef) GetInitialStateIntf() react.State {
	return LatencyState{}
}

func (l LatencyState) EqualsIntf(val interface{}) bool {
	return l == val.(LatencyState)
}
