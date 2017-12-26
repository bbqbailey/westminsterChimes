/*
	Name: playWestminsterFile.go
	Author: Ben Bailey
	Date: Dec 18, 2017
	Purpose: Had some problems with programs, such as mpg123 that didn't want to exit properly
		if invoked via cron, so I ended up with mgp123 waiting process for every one that was
		run via cron; over a period of day, this was quite a few processes!  This was a problem,
		as typically I wanted to play the westminster chimes via cron.  So this is an attempt
		to fix that problem; testing has indicated that mpg123 will work properly from cron
		if I use cron to call an app, such as this one, then have it do the invoking of the
		mpg123 file.

		Couple of arguments:
		-h (default help provided by golang)
		-fileName - name of the file to play; default is westminster15Min.mp3
		-player   - name of an audio player; default is mpg123
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//player:="mpg123"
	defDirPtr := "/home/<CHANGE TO YOUR USERNAME!>/go/src/westminsterChimes/"

	fileNamePtr := flag.String("fileName", "westminster15Min.mp3", "Provide the filename to be played")
	playerPtr := flag.String("player", "mpg123", "Audio player to use.  Default is mpg123")
	//	optionsPtr := flag.String("option", " --play-and-exit ", " Defulat --play-and-exit is for clvc")
	flag.Parse()

	if _, err := os.Stat(defDirPtr + *fileNamePtr); os.IsNotExist(err) {
		// path/to/whatever does not exist
		fmt.Println("file ", defDirPtr+*fileNamePtr, " not found!  err: ", err)
	} else {
		args := []string{defDirPtr + *fileNamePtr} // + *optionsPtr}
		fmt.Printf("playWestminsterFile Command: %s %s\n", *playerPtr, args)
		err := exec.Command(*playerPtr, args...).Run()
		fmt.Println("err: ", err)
	}
}
