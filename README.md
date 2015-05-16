# What is `gw`?

`gw` is a `gradle`/`gradlew` wrapper written in `Go`. It's a clone of `gdub`
that works the same way (https://github.com/dougborg/gdub). The reason I
wanted to create this in Go is to be able to run it on OSX, Unix and Windows.

## Building

Since I do not have any binaries yet, you have to build it yourself.

    go build

Run it using the `gw` command:

    ./gw
