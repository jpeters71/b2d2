# b2d2
## Description
This project represents code used to run my robotic bartender.  The drinks, ingredients, and recipes are all stored in a SQLite DB.  There are services designed to pull out relevant information and a web server designed to serve up a web UI.  This is a work in process, but eventually there will be an additional service that will be used to actually turn on the pumps, etc.with a GPIO interface.  This interface will be implemented for Raspberry Pi but will need to be stubbed out for all other platforms.

## Building
I've setup a Makefile to build everything, though I haven't tested it extensively.  Before you start, make sure you have Go and an implementation of GCC setup and in your path (on Windows, I'm using MinGW implementation of GCC).  Then, to get the required Go dependencies run this once: `make setup`

After that, doing a simple `make build` or just a `make` should work.

## Running
I've included a "seed" DB that has a number of drinks/ingredients in it so you should be able to run this out of the box.  Once built, start the server from the /b2d2/bar2d2 directory like so:
`ui/ui` (or, under Windows `ui/ui.exe`).  If you want to start it somewhere else, you can specify on the command line where the root bar2d2 directory is using the `--root` argument.  Something like:
`./ui --root=/home/bozo/bar2d2`

One of these days, I'm going to break out using command line arguments, the path to the DB, web files, etc. to make it a little more portable.

## Work in Process
This is a work in process so if you want to play with it...Good Luck!

