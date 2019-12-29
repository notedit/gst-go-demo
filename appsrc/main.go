package main

import (
	"fmt"
	"time"

	"github.com/notedit/gst"
)

func main() {

	pipeline, err := gst.ParseLaunch("appsrc name=mysource format=time is-live=true do-timestamp=true ! videoconvert ! autovideosink")

	if err != nil {
		panic("pipeline error")
	}

	videoCap := gst.CapsFromString("video/x-raw,format=RGB,width=320,height=240,bpp=24,depth=24")

	element := pipeline.GetByName("mysource")

	element.SetObject("caps", videoCap)

	pipeline.SetState(gst.StatePlaying)

	i := 0
	for {

		if i > 100 {
			break
		}

		data := make([]byte, 320*240*3)

		err := element.PushBuffer(data)

		if err != nil {
			fmt.Println("push buffer error")
			break
		}

		fmt.Println("push one")
		i++
		time.Sleep(50000000)
	}

	pipeline.SetState(gst.StateNull)

	pipeline = nil
	element = nil
	videoCap = nil
}
