package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/algoGuy/EasyMIDI/smfio"
)

func main() {
	file, _ := os.Open("./short.mid")
	defer file.Close()

	midi, err := smfio.Read(bufio.NewReader(file))

	if err != nil {
		fmt.Println(err)
	}

	track := midi.GetTrack(1)
	fmt.Println(midi.GetTracksNum())

	iter := track.GetIterator()
	var length int

	for iter.MoveNext() {
		e := iter.GetValue()
		var status int
		var pitch uint8
		if e.GetStatus() == 0x80 {
			status = 0
			pitch = uint8(e.GetData()[0])
			length += int(e.GetDTime())
		} else if e.GetStatus() == 0x90 {
			status = 1
			pitch = uint8(e.GetData()[0])
			length += int(e.GetDTime())
		} else if e.GetStatus() == 0x90 {
		} else {
			continue
		}
		fmt.Printf("Status: %v\tTime: %v\tPitch: %v\n", status, e.GetDTime(), pitch)
	}
	fmt.Printf("Total time: %v\n", length)
}
