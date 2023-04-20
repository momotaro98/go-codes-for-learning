
# Overview

`inotify` is an API of Linux.

manual page is here. https://man7.org/linux/man-pages/man7/inotify.7.html

There's a golang library named [fsnotify](https://github.com/fsnotify/fsnotify) which covers each Operating System including BSD, Windows. (BSD uses similar API named kqueue.)

# Code in this directory

Original code is from [here](https://github.com/tomnomnom/go-learning/blob/master/inotify.go).

In MacOS, [the code](./main.go) cannot be compiled because the `unix.InotifyXXX` function calls are only for Linux OS. (`// +build linux,!appengine`)

So, you need to build it with [Dockerfile](./Dockerfile) and compose file.

```
docker-compose build
```