package r3dOperation

import (
	"fmt"
	"testing"
	"os"
)

func TestBuildDate(t * testing.T) {
	fmt.Println(BuildDate())
}

	// Clip ClipInit(const char *);
	// //int ClipStatus();
	// int ClipVideoTrackCount(Clip);
	// int ClipWidth(Clip);
	// int ClipHeight(Clip);
	// float ClipVideoAudioFramerate(Clip);
	// float ClipTimecodeFramerate(Clip);
	// int ClipVideoFrameCount(Clip);
	// const char * ClipStartAbsoluteTimecode(Clip);
	// const char * ClipEndAbsoluteTimecode(Clip);
	// const char * ClipStartEdgeTimecode(Clip);
	// const char * ClipEndEdgeTimecode(Clip);
	// void ClipFree(Clip);
func TestClip(t * testing.T) {
	clip := ClipInit("/home/developer/Videos/B011_C010_060551.RDC/B011_C010_060551_002.R3D")
	fmt.Println(ClipVideoTrackCount(clip))
	fmt.Println(ClipWidth(clip))
	fmt.Println(ClipHeight(clip))
	fmt.Println(ClipVideoAudioFramerate(clip))
	fmt.Println(ClipTimecodeFramerate(clip))
	fmt.Println(ClipVideoFrameCount(clip))
	fmt.Println(ClipStartAbsoluteTimecode(clip))
	//fmt.Println(ClipEndAbsoluteTimecode(clip))
	fmt.Println(ClipStartEdgeTimecode(clip))
	fmt.Println(ClipEndEdgeTimecode(clip))
	fmt.Println(ClipMetaData(clip))
	file, _ := os.Create("test.png")
	defer file.Close()
	ClipDecodeFrame(clip, 1000, file)

	//ClipDecodeFrameFree(buffer)
	ClipFree(clip)
}