package main

import (
	"fmt"
	"github.com/notedit/gst"
)

func main() {


	pipelineStr := "videotestsrc is-live=true ! video/x-raw,format=I420,framerate=15/1 ! x264enc aud=false bframes=0 speed-preset=veryfast key-int-max=15 ! video/x-h264,stream-format=byte-stream,profile=baseline ! h264parse ! rtph264pay config-interval=-1  pt=%d ! appsink name=appsink"
	pipelineStr = fmt.Sprintf(pipelineStr, 100)

	err := gst.CheckPlugins([]string{"x264","rtp","videoparsersbad"})

	if err != nil {
		fmt.Println(err)
	}

	pipeline, err := gst.ParseLaunch(pipelineStr)

	if err != nil {
		panic(err)
	}

	element := pipeline.GetByName("appsink")

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

