package main

import (
	"log"
	"os"
	"time"
)

/*
$ go run main.go or.go
14:07:02 ward: Hello, I'm irresponsible!
14:07:06 steward: ward unhealthy; restarting
14:07:06 ward: Hello, I'm irresponsible!
14:07:06 ward: I am halting.
14:07:10 steward: ward unhealthy; restarting
14:07:10 ward: Hello, I'm irresponsible!
14:07:10 ward: I am halting.
14:07:11 main: halting steward and ward.
14:07:11 Done
*/

type startGoroutineFn func(
	done <-chan interface{},
	pulseInterval time.Duration,
) (heartbeat <-chan interface{})

func newSteward( // Steward はここでは管理人の意味
	timeout time.Duration,
	startGoroutine startGoroutineFn,
) startGoroutineFn { // > 興味深いことに管理人自身は startGoroutineFn を返していて、これは管理人自身も監視可能であることを示しています。
	return func(
		done <-chan interface{},
		pulseInterval time.Duration,
	) <-chan interface{} {
		heartbeat := make(chan interface{})
		go func() {
			defer close(heartbeat)

			var wardDone chan interface{} // ward は中庭という意味
			// ここでは管理人に管理される対象のゴルーチンのこと
			var wardHeartbeat <-chan interface{}
			startWard := func() {
				wardDone = make(chan interface{})
				wardHeartbeat = startGoroutine(or(wardDone, done), timeout/2)
				// > ↑監視対象のゴルーチンを起動します。管理人が停止するか、管理人が中庭のゴルーチンを停止させたい
				// > 場合に対象のゴルーチンには停止してもらいたいので、両方のdoneチャンネルをorの中に内包します。
			}
			startWard()
			pulse := time.Tick(pulseInterval)

		monitorLoop:
			for {
				timeoutSignal := time.After(timeout)

				for {
					select {
					case <-pulse:
						select {
						case heartbeat <- struct{}{}: // 内側のforループで管理人が自身の鼓動を確実に外へと送信できるようにしています。
						default:
						}
					case <-wardHeartbeat: // 中庭の鼓動を受信したら、監視のループを継続する、という実装になっているのがわかります。
						continue monitorLoop
					case <-timeoutSignal: // タイムアウト期間内に中庭からの鼓動が受信できなければ、中庭に停止するようリクエストし、
						// 新しい中庭のゴルーチンを起動しはじめることを示している行です。その後、監視を続けます。
						log.Println("steward: ward unhealthy; restarting")
						close(wardDone)
						startWard()
						continue monitorLoop
					case <-done:
						return
					}
				}
			}
		}()

		return heartbeat
	}

}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	// 管理人動作確認のため、わざと失敗する仕事の関数を作って newSteward に渡す。
	doWork := func(done <-chan interface{}, _ time.Duration) <-chan interface{} {
		log.Println("ward: Hello, I'm irresponsible!")
		go func() {
			<-done
			log.Println("ward: I am halting.")
		}()
		return nil
	}
	doWorkWithSteward := newSteward( /*timeout*/ 4*time.Second /*startGoroutine*/, doWork)

	done := make(chan interface{})
	time.AfterFunc(9*time.Second, func() {
		log.Println("main: halting steward and ward.")
		close(done)
	})

	for range doWorkWithSteward(done, 4*time.Second) {
	}
	log.Println("Done")
}
