// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=sbus -id dma_descriptor -d VecType=dma_descriptor_vec -d Type=dma_descriptor github.com/platinasystems/go/elib/hw/dma_mem.tmpl]

package sbus

import (
	"github.com/platinasystems/go/elib"
	"github.com/platinasystems/go/elib/hw"

	"reflect"
	"unsafe"
)

type dma_descriptor_vec []dma_descriptor

func fromByteSlice_dma_descriptor(b []byte, l, c uint) (x dma_descriptor_vec) {
	s := uint(unsafe.Sizeof(x[0]))
	if l == 0 {
		l = uint(len(b)) / s
		c = uint(cap(b))
	}
	return *(*dma_descriptor_vec)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&b[0])),
		Len:  int(l),
		Cap:  int(c / s),
	}))
}

func (x dma_descriptor_vec) toByteSlice() []byte {
	l := uint(len(x))
	l *= uint(unsafe.Sizeof(x[0]))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&x[0])),
		Len:  int(l),
		Cap:  int(l)}))
}

func dma_descriptorAllocAligned(n, a uint) (x dma_descriptor_vec, id elib.Index) {
	var b []byte
	var c uint
	b, id, _, c = hw.DmaAllocAligned(n*uint(unsafe.Sizeof(x[0])), a)
	x = fromByteSlice_dma_descriptor(b, n, c)
	return
}

func dma_descriptorAlloc(n uint) (x dma_descriptor_vec, id elib.Index) {
	return dma_descriptorAllocAligned(n, 0)
}

func dma_descriptorNew() (x dma_descriptor_vec, id elib.Index) { return dma_descriptorAlloc(1) }

func (x *dma_descriptor_vec) Free(id elib.Index) {
	hw.DmaFree(id)
	*x = nil
}

func (x *dma_descriptor_vec) Get(id elib.Index) {
	*x = fromByteSlice_dma_descriptor(hw.DmaGetData(id), 0, 0)
}

func (x *dma_descriptor) PhysAddress() uintptr {
	return hw.DmaPhysAddress(uintptr(unsafe.Pointer(x)))
}
