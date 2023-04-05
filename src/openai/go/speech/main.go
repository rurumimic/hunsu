package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/MarkKremer/microphone"
	"github.com/faiface/beep/wav"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing required argument: output file name")
		return
	}
	fmt.Println("Recording. Press Ctrl-C to stop.")

	err := microphone.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer microphone.Terminate()

	fmt.Println("1")

	stream, format, err := microphone.OpenDefaultStream(44100, 2)
	if err != nil {
		log.Fatal(err)
	}
	// Close the stream at the end if it hasn't already been
	// closed explicitly.
	defer stream.Close()

	fmt.Println("2")

	filename := os.Args[1]

	fmt.Println(filename)

	if !strings.HasSuffix(filename, ".wav") {
		filename += ".wav"
	}
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("3")

	// Stop the stream when the user tries to quit the program.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	go func() {
		<-sig
		stream.Stop()
		stream.Close()
	}()

	stream.Start()

	// Encode the stream. This is a blocking operation because
	// wav.Encode will try to drain the stream. However, this
	// doesn't happen until stream.Close() is called.
	err = wav.Encode(f, stream, format)
	if err != nil {
		log.Fatal(err)
	}
}
