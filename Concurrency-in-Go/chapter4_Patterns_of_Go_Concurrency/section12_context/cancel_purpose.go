package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// $ go run XXX.go
// connot print greeting: context deadline exceeded
// cannot print farewell: context canceled

// 【ポイント】
// i.   mainはGreetingが失敗したら一番親のContextをキャンセルにする。
//
// ii.  Greetingは親のContextをcontext.WithTimeoutで囲む。
//      これによって1秒後には自分の子供をすべてキャンセルにする。この場合はlocaleをキャンセルにする。
//
// iii. 呼び出し元から渡されたContextがキャンセルされた場合(Done()からチャネルが返った場合)、
//      Contextがキャンセル理由と共にエラーを返す。このエラーはmainにまで伝搬され、これが i. での
//      キャンセルを発生させ、Farewell側のlocaleでもキャンセルになる。
//
// iv.  ここでデッドラインが設定されているかを確認する。もし設定されていて、システムの時計でそのデッドラインを
//      過ぎていたら、単純にcontextパッケージに定義されている特別なエラーであるDeadlineExceededを返します。

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Greeting
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreeting(ctx); err != nil {
			fmt.Printf("connot print greeting: %v\n", err)
			cancel() // i.
		}
	}()

	// Farewell
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(ctx); err != nil {
			fmt.Printf("cannot print farewell: %v\n", err)
		}
	}()

	wg.Wait()
}

func printGreeting(ctx context.Context) error {
	greeting, err := genGreeting(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func printFarewell(ctx context.Context) error {
	farewell, err := genFarewell(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func genGreeting(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second) // ii.
	defer cancel()

	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func genFarewell(ctx context.Context) (string, error) {
	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func locale(ctx context.Context) (string, error) {
	if deadline, ok := ctx.Deadline(); ok { // iv.
		if deadline.Sub(time.Now().Add(1*time.Minute)) <= 0 {
			return "", context.DeadlineExceeded
		}
	}

	select {
	case <-ctx.Done(): // iii.
		return "", ctx.Err()
	case <-time.After(1 * time.Minute):
	}
	return "EN/US", nil
}
