// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loop

import (
	"github.com/platinasystems/go/elib"
	"github.com/platinasystems/go/elib/elog"

	"fmt"
)

type poller_elog_event_type byte

const (
	poller_wake poller_elog_event_type = iota + 1
	poller_wait
	poller_suspend
	poller_resume
	poller_resumed
	poller_activate
)

var poller_elog_event_type_names = [...]string{
	poller_wake:     "wake",
	poller_wait:     "wait",
	poller_suspend:  "suspend",
	poller_resume:   "resume",
	poller_resumed:  "resumed",
	poller_activate: "activate",
}

func (t poller_elog_event_type) String() string {
	return elib.Stringer(poller_elog_event_type_names[:], int(t))
}

type pollerElogEvent struct {
	event_type   poller_elog_event_type
	flags        byte
	poller_index uint32
	node_name    elog.StringRef
}

func (e *pollerElogEvent) Elog(l *elog.Log) {
	pi := ""
	if e.poller_index != ^uint32(0) {
		pi = fmt.Sprintf("%d", e.poller_index)
	}
	l.Logf("loop%s %s %s", pi, e.event_type, l.GetString(e.node_name))
	if e.flags != 0 {
		l.Logf("new flags: %s", node_flags(e.flags))
	}
}

func (n *Node) pollerElog(t poller_elog_event_type, f node_flags) {
	if elog.Enabled() {
		c := elog.GetCaller(elog.PointerToFirstArg(&n))
		le := pollerElogEvent{
			event_type:   t,
			poller_index: uint32(n.activePollerIndex),
			flags:        byte(f),
			node_name:    n.elogNodeName,
		}
		elog.Addc(&le, c)
	}
}

type callEvent struct {
	active_index uint32
	n_vectors    uint32
	is_input     bool
	node_name    elog.StringRef
}

func (e *callEvent) Elog(l *elog.Log) {
	n := l.GetString(e.node_name)
	nv := e.n_vectors
	if e.is_input {
		l.Logf("loop%d %s in %d", e.active_index, n, nv)
	} else {
		l.Logf("loop%d %s(%d)", e.active_index, n, nv)
	}
}
