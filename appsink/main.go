package main

import (
	"fmt"
	"github.com/notedit/gst"
)

func main() {

	pipeline, err := gst.ParseLaunch("videotestsrc  num-buffers=10 ! appsink name=sink")

	if err != nil {
		panic("pipe error")
	}

	element := pipeline.GetByName("sink")

	pipeline.SetState(gst.StatePlaying)

	for {

		sample, err := element.PullSample()
		if err != nil {
			if element.IsEOS() == true {
				fmt.Println("eos")
				return
			} else {
				fmt.Println(err)
				continue
			}
		}
		fmt.Println("got sample", sample.Duration)
	}

	pipeline.SetState(gst.StateNull)

	pipeline = nil
	element = nil
}
