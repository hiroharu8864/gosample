package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	log.Print("program started.")

	// クロージャーを配列にする
	funcs := []func(){
		func() {
			log.Print("sleep1 started.")
			time.Sleep(1 * time.Second)
			log.Print("sleep1 finished.")
		},
		func() {
			log.Print("sleep2 started.")
			time.Sleep(2 * time.Second)
			log.Print("sleep2 finished.")
		},
		func() {
			log.Print("sleep3 started.")
			time.Sleep(3 * time.Second)
			log.Print("sleep3 finished.")
		},
		func() {
			log.Print("sleep4 started.")
			time.Sleep(4 * time.Second)
			log.Print("sleep4 finished.")
		},
	}

	var waitGroup sync.WaitGroup

	// 関数の数だけ並列化
	for _, sleep := range funcs {
		waitGroup.Add(1) // 待つ数をインクリメント

		go func(function func()) {
			defer waitGroup.Done() // 待つ数をデクリメント
			function()
		}(sleep)

	}

	waitGroup.Wait() // ゼロになるまで、処理をブロック

	log.Print("all finished.")
}
