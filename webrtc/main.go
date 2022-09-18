package main

import (
	"fmt"

	"github.com/notedit/gst"
)

const pipelineStr = `
webrtcbin bundle-policy=max-bundle name=pusher 
videotestsrc is-live=true pattern=ball ! videoconvert ! queue ! vp8enc deadline=1 ! rtpvp8pay ! application/x-rtp,media=video,encoding-name=VP8,payload=96 ! pusher.  
audiotestsrc is-live=true wave=red-noise ! audioconvert ! audioresample ! queue ! opusenc ! rtpopuspay ! queue ! application/x-rtp,media=audio,encoding-name=OPUS,payload=97 ! pusher.
`

func main() {

	pipeline, err := gst.ParseLaunch(pipelineStr)

	if err != nil {
		panic(err)
	}

	pipeline.SetState(gst.StatePlaying)

	bus := pipeline.GetBus()

	for {
		message := bus.Pull(gst.MessageAny)
		fmt.Println("message:", message.GetName())
		if message.GetType() == gst.MessageEos {
			break
		}
	}

	pipeline.SetState(gst.StateNull)

	pipeline = nil
}
