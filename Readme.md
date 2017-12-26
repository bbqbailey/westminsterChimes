# Project Title
playWestminsterFile

Small golang program to play Westminster chimes, invoked via crontab (on your choices):

 - Hour
 - 15 min
 - 30 min
 - 45 min

Tested on Ubuntu 17.10

Assumes the use of Linux, but should work on other platforms as well, as long as the platform can run golang (go)

## Copyright (c) 2017 bbqbailey AKA Ben Bailey

Refer to LICENSE.md for full insight.

## Getting Started
 - All sounds played are .mp3 files of Westminster Chimes.
 - golang program "playWestminsterFile" is invoked via crontab entery (or other of your choice, e.g., command line)
 - "playWestminsterFile" invokes the music player (my suggestion is mpg123 for Linux; not sure for other computer types), passing it the name of the audio .mp3 file to play (could be any mp3 file, not just the westminster mp3 files)
 - Your placement in crontab determines when the "playWestminsterFile" is invoked, the specified command line argument determines which file is played.

These instructions will get you a copy of the project up and running on your local machine.

### Prerequisites
These instructions have only been tested on Ubuntu 17.10  If your machine is different then you may need to alter different elements of this installation and method of invokation.

Software:
 - mpg123 (note: a lightweight audio player)
 - playWestminsterFile built as executable and marked as executable in file system. (note: built from this repository source code)
 - mp3 files to be played are on file system. (note: you can use those in this repository, or create your own, and reflect in your crontab entry)
 - entries in crontab made to support invoking at desired entervals (note: entervals such as hourly, 1/4 hour, etc)

Hardware:

 - Speakers are available and have been tested to play sound from this computer.

This program assumes the use of mpg123 as the audio player, although you can edit the program source go file to place any audio player to be invoked.  If you do edit, then you will need to rebuild the application.  Therefore, download and install mpg123 onto your Linux machine if that is your desired player.  I tried various players, and this one was light-weight and easy to invoke, and also shutdown after playing.  Some others wanted X11 to popup to play, while others didn't want to exit after playing.  Note: I originally tried playing the sound files directly from crontab, but ran into problems with the invoked program (vlc, mpg123, etc) having issues when called in this manner.

### Installing
If you do not have golang installed on your Ubuntu 17.10 system:
 - Install golan onto your Ubuntu 17.10 system in order to build the executable "playWestminsterFile" from the golang source code.  Please consult www.golang.org for information associated with this.
 - Create the directory $GOPATH/src/github.com/bbqbailey/westminsterChimes

Issue the following commands:
 - cd $GOPATH/src/github.com/bbqbailey/westminsterChimes  (note: this changes your working directory to that needed for this program)
 - go get github.com/bbqbailey/westminsterChimes  (note: this places the playWestminsterFile and the associated chimes mp3 files into this directory)
 - go build playWestminsterFile.go (note: this produces your executable playWestminsterFile program)
 - chmod 744 playWestminsterFile program) (note: this makes your program executable by you; refer to 'man chmod' if you desire different access)
 - ./playWestminsterFileprogram -fileName=westminsterHour.mp3  (note: this will test your installation and build
 - crontab -e (this opens your personal crontab for you to edit.  Refer to 'man crontab' if you desire to run as root or other entity.  


Installing mpg123 on Ubuntu 17.10:
 - sudo apt install mpg123

This will install mpg123 as an audio player onto a Ubuntu 17.10 machine.

## How To Invoke

The program is invoked in the following form from the command line: 

   ./playWestminsterChimes -fileName=NameOfAudioFileToPlay.mp3

## Crontab Setup Examples

You can invoke the program via multiple entry in crontab.  I have set my Ubuntu 17.10 computer  up to play Hour, 15 minute, 30 minute, and 45 minute.

Crontab expects data in the following format, where: minute (m), hour (h), day-of-month (dom), month (mon), day-of-week (dow) or '*' for any in a field.  Note that I'm also directing any output to /tmp/cronlog for debug purposes.  You can omit this, or use it to test then delete later

 - I have set my machine up to play Hour, 15 minute, 30 minute, and 45 minute. The following is my entry for those periods.  I am not playing any bells for-the-hour; e.g., at 4 I don't have it play the hour and also ring bells 4 times, I just play the hour.  Note that the hour (h) is based on 24 hours; e.g., 2 PM would be 2 + 12 = 14.

  (note: pay close attention as I can't control how the following lines appear, but should be one line per entry):

  m    h    dom    mon    dow    command 

  00   *    *      *      *      /home/<yourhomename>/go/src/github.com/bbqbailey/westminsterchimes/playWestminsterFile -fileName=westminsterHour.mp3 >> /tmp/cronlog # 00 minutes is on-hour

  15   *    *      *      *      /home/<yourhomename>/go/src/github.com/bbqbailey/westminsterchimes/playWestminsterFile -fileName=westminster15Min.mp3 >> /tmp/cronlog # 15 minutes is 1/4 hour

  30   *    *      *      *      /home/<yourhomename>/go/src/github.com/bbqbailey/westminsterchimes/playWestminsterFile -fileName=westminster30Min.mp3 >> /tmp/cronlog # 30 minutes is 1/2 hour

  45   *    *      *      *      /home/<yourhomename>/go/src/github.com/bbqbailey/westminsterchimes/playWestminsterFile -fileName=westminster45Min.mp3 >> /tmp/cronlog # 45 minutes is 3/4 hour


- If you desire to have the hourly bells to play, then modify the '00' entry above to the following 
  
  (note: pay close attention as I can't control how the following lines appear, but should be one line per entry):

  m    h    dom    mon    dow    command

  00   12    *      *      *      /home/<yourhomename>/go/src/github.com/bbqbailey/westminsterchimes/playWestminsterFile -fileName=westminster12-Oclock.mp3 >> /tmp/cronlog # noon

  00   13    *      *      *      /home/<yourhomename>/go/src/github.com/bbqbailey/westminsterchimes/playWestminsterFile -fileName=westminster1-Oclock.mp3 >> /tmp/cronlog # 1 PM

  00   14    *      *      *      /home/<yourhomename>/go/src/github.com/bbqbailey/westminsterchimes/playWestminsterFile -fileName=westminster2-Oclock.mp3 >> /tmp/cronlog # 2 PM

  00   15    *      *      *      /home/<yourhomename>/go/src/github.com/bbqbailey/westminsterchimes/playWestminsterFile -fileName=westminster3-Oclock.mp3 >> /tmp/cronlog # 3 PM

  ...

  00   21    *      *      *      /home/<yourhomename>/go/src/github.com/bbqbailey/westminsterchimes/playWestminsterFile -fileName=westminster3-Oclock.mp3 >> /tmp/cronlog # 9 PM


## Running the tests

From the command line, issue:
 - ./playWestminsterFileprogram -fileName=westminsterHour.mp3  (note: this will test your installation and build

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

* **BBQBailey** - *Initial work* - (https://github.com/bbqbailey/westminsterChimes

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip: https://freesound.org/people/acclivity/sounds/14416/  I took the 12 hour strike that was provided from this site and edited it via audacity to get the 15 - hour chimes as well as the *-Oclock rings.
* Inspiration: I have a history of sitting all day at my desk working without getting up much, and as a result (conjecture) have experienced pulmonary embolism that have twice place me in the hospital.  I needed to be reminded to get up from my desk to walk every half hour, and decided to use Westminster chimes as my reminder.  I editied (with audacity) these files, and added my wife's pretty voice to the end of hour and half hour so after it chimes, she says "need to get up and walk Baby!".
* etc


