# traze-goclient

A cli client for traze game written in golang.

Reads game state from mqtt broker and displays board on a console buffer.

Configuration with brokers url is located in `./conf`. You may copy json conf file to `/etc/goclient`.

Build it **(go installation required)**

    make build

Run executable

    ./traze-goclient

Build and run it at once **(go installation required)**

    make run

## Todo

Add functionality to join a game and run a bike.
