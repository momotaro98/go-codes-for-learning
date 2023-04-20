//go:build linux && !appengine
// +build linux,!appengine

package main

import (
	"fmt"
	"log"
	"unsafe"

	"golang.org/x/sys/unix"
)

func main() {
	// inotify利用の初期化
	// fd(ファイルディスクリプタ)を取得している。
	fd, err := unix.InotifyInit1(unix.IN_CLOEXEC | unix.IN_NONBLOCK)
	if err != nil {
		log.Fatal(err)
	}
	defer unix.Close(fd)

	// 監視対象のファイルを設定
	// fdが渡されている。
	wd, err := unix.InotifyAddWatch(fd, "test1.log", unix.IN_ALL_EVENTS)
	if err != nil {
		log.Fatal(err)
	}
	defer unix.InotifyRmWatch(fd, uint32(wd)) // 処理が終われば監視対象から外すようにする

	fmt.Printf("WD is %d\n", wd)

	for {
		// Room for at least 128 events
		buffer := make([]byte, unix.SizeofInotifyEvent*128)
		bytesRead, err := unix.Read(fd, buffer) // 監視対象のfdをFor Loopで読み取る。
		if err != nil {
			log.Fatal(err)
		}

		if bytesRead < unix.SizeofInotifyEvent {
			// No point trying if we don't have at least one event
			continue
		}

		fmt.Printf("Size of InotifyEvent is %s\n", unix.SizeofInotifyEvent)
		fmt.Printf("Bytes read: %d\n", bytesRead)

		offset := 0
		for offset < bytesRead-unix.SizeofInotifyEvent {
			event := (*unix.InotifyEvent)(unsafe.Pointer(&buffer[offset]))
			fmt.Printf("%+v\n", event)

			if (event.Mask & unix.IN_ACCESS) > 0 {
				fmt.Printf("Saw IN_ACCESS for %+v\n", event)
			}

			// We need to account for the length of the name
			offset += unix.SizeofInotifyEvent + int(event.Len)
		}
	}

}
