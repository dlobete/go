// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=elib -id Byte -d VecType=ByteVec -d Type=byte vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elib

type ByteVec []byte

func (p *ByteVec) Resize(n uint) {
	c := Index(cap(*p))
	l := Index(len(*p)) + Index(n)
	if l > c {
		c = NextResizeCap(l)
		q := make([]byte, l, c)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:l]
}

func (p *ByteVec) validate(new_len uint, zero byte) *byte {
	c := Index(cap(*p))
	lʹ := Index(len(*p))
	l := Index(new_len)
	if l <= c {
		// Need to reslice to larger length?
		if l > lʹ {
			*p = (*p)[:l]
			for i := lʹ; i < l; i++ {
				(*p)[i] = zero
			}
		}
		return &(*p)[l-1]
	}
	return p.validateSlowPath(zero, c, l, lʹ)
}

func (p *ByteVec) validateSlowPath(zero byte, c, l, lʹ Index) *byte {
	if l > c {
		cNext := NextResizeCap(l)
		q := make([]byte, cNext, cNext)
		copy(q, *p)
		for i := c; i < cNext; i++ {
			q[i] = zero
		}
		*p = q[:l]
	}
	if l > lʹ {
		*p = (*p)[:l]
	}
	return &(*p)[l-1]
}

func (p *ByteVec) Validate(i uint) *byte {
	var zero byte
	return p.validate(i+1, zero)
}

func (p *ByteVec) ValidateInit(i uint, zero byte) *byte {
	return p.validate(i+1, zero)
}

func (p *ByteVec) ValidateLen(l uint) (v *byte) {
	if l > 0 {
		var zero byte
		v = p.validate(l, zero)
	}
	return
}

func (p *ByteVec) ValidateLenInit(l uint, zero byte) (v *byte) {
	if l > 0 {
		v = p.validate(l, zero)
	}
	return
}

func (p ByteVec) Len() uint { return uint(len(p)) }
