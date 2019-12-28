package main

import (
	"fmt"
	"strings"

	"github.com/notedit/gst"
)

func main() {

	pipeline, err := gst.PipelineNew("test-pipeline")

	if err != nil {
		panic(err)
	}

	source, _ := gst.ElementFactoryMake("uridecodebin", "source")
	convert, _ := gst.ElementFactoryMake("audioconvert", "convert")
	sink, _ := gst.ElementFactoryMake("autoaudiosink", "sink")

	pipeline.Add(source)
	pipeline.Add(convert)
	pipeline.Add(sink)

	convert.Link(sink)

	source.SetObject("uri", "https://www.freedesktop.org/software/gstreamer-sdk/data/media/sintel_trailer-480p.webm")

	source.SetPadAddedCallback(func(element *gst.Element, pad *gst.Pad) {
		capstr := pad.GetCurrentCaps().ToString()

		if strings.HasPrefix(capstr, "audio") {
			sinkpad := convert.GetStaticPad("sink")
			pad.Link(sinkpad)
		}

	})

	pipeline.SetState(gst.StatePlaying)

	bus := pipeline.GetBus()

	for {
		message := bus.Pull(gst.MessageError | gst.MessageEos)
		fmt.Println("message:", message.GetName())
		if message.GetType() == gst.MessageEos {
			break
		}
	}
}
