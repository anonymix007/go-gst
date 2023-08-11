package gst

// #include "gst.go.h"
import "C"

import (
	"errors"
	"fmt"
	"strings"
	"unsafe"

	"github.com/go-gst/go-glib/glib"
)

// Pipeline is a go implementation of a GstPipeline.
type Pipeline struct {
	*Bin
	bus *Bus
}

// FromGstPipelineUnsafeFull wraps the given pipeline pointer.
func FromGstPipelineUnsafeFull(pipeline unsafe.Pointer) *Pipeline {
	return &Pipeline{Bin: &Bin{&Element{wrapObject(glib.TransferFull(pipeline))}}}
}

// FromGstPipelineUnsafeNone wraps the given pipeline pointer.
func FromGstPipelineUnsafeNone(pipeline unsafe.Pointer) *Pipeline {
	return &Pipeline{Bin: &Bin{&Element{wrapObject(glib.TransferNone(pipeline))}}}
}

// NewPipeline allocates and returns a new empty pipeline. If name is empty, one
// is generated by gstreamer.
func NewPipeline(name string) (*Pipeline, error) {
	var cChar *C.char
	if name != "" {
		cChar = C.CString(name)
		defer C.free(unsafe.Pointer(cChar))
	}
	pipeline := C.gst_pipeline_new((*C.gchar)(cChar))
	if pipeline == nil {
		return nil, errors.New("Could not create new pipeline")
	}
	return FromGstPipelineUnsafeNone(unsafe.Pointer(pipeline)), nil
}

// NewPipelineFromString creates a new gstreamer pipeline from the given launch string.
func NewPipelineFromString(launchv string) (*Pipeline, error) {
	if len(strings.Split(launchv, "!")) < 2 {
		return nil, fmt.Errorf("Given string is too short for a pipeline: %s", launchv)
	}
	cLaunchv := C.CString(launchv)
	defer C.free(unsafe.Pointer(cLaunchv))
	var gerr *C.GError
	pipeline := C.gst_parse_launch((*C.gchar)(cLaunchv), (**C.GError)(&gerr))
	if gerr != nil {
		defer C.g_error_free((*C.GError)(gerr))
		errMsg := C.GoString(gerr.message)
		return nil, errors.New(errMsg)
	}
	return FromGstPipelineUnsafeNone(unsafe.Pointer(pipeline)), nil
}

// Instance returns the native GstPipeline instance.
func (p *Pipeline) Instance() *C.GstPipeline { return C.toGstPipeline(p.Unsafe()) }

// GetPipelineBus returns the message bus for this pipeline.
func (p *Pipeline) GetPipelineBus() *Bus {
	if p.bus == nil {
		cBus := C.gst_pipeline_get_bus((*C.GstPipeline)(p.Instance()))
		p.bus = FromGstBusUnsafeFull(unsafe.Pointer(cBus))
	}
	return p.bus
}

// GetPipelineClock returns the global clock for this pipeline.
func (p *Pipeline) GetPipelineClock() *Clock {
	cClock := C.gst_pipeline_get_pipeline_clock((*C.GstPipeline)(p.Instance()))
	return FromGstClockUnsafeFull(unsafe.Pointer(cClock))
}

/*
SetAutoFlushBus can be used to disable automatically flushing the message bus
when a pipeline goes to StateNull.

Usually, when a pipeline goes from READY to NULL state, it automatically flushes
all pending messages on the bus, which is done for refcounting purposes, to break
circular references.

This means that applications that update state using (async) bus messages (e.g. do
certain things when a pipeline goes from PAUSED to READY) might not get to see
messages when the pipeline is shut down, because they might be flushed before they
can be dispatched in the main thread. This behaviour can be disabled using this function.

It is important that all messages on the bus are handled when the automatic flushing
is disabled else memory leaks will be introduced.
*/
func (p *Pipeline) SetAutoFlushBus(b bool) {
	C.gst_pipeline_set_auto_flush_bus(p.Instance(), gboolean(b))
}

// Start is the equivalent to calling SetState(StatePlaying) on the underlying GstElement.
func (p *Pipeline) Start() error {
	return p.SetState(StatePlaying)
}
