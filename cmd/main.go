package main

import (
	"log"
	"time"

	internal "github.com/rogaliiik/playlist/internal/playlist"
)

func main() {
	p := internal.NewPlaylist()
	p.Wg.Add(1)
	defer p.Wg.Wait()

	p.AddSong(internal.NewSong("Song 1", 1))
	p.AddSong(internal.NewSong("Song 2", 1))
	p.AddSong(internal.NewSong("Song 3", 1))

	go p.Broadcast()
	p.Play()
	time.Sleep(time.Second * 3)
	p.AddSong(internal.NewSong("Song 4", 1))
	err := p.Prev()
	if err != nil {
		log.Println(err)
	}
	err = p.Next()
	if err != nil {
		log.Println(err)
	}
	p.AddSong(internal.NewSong("Song 5", 1))
	p.Pause()
	p.Shutdown()

}
