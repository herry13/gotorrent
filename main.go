package main

import (
  //"fmt"
  "log"
  "time"

  "github.com/anacrolix/torrent"
  "github.com/anacrolix/tagflag"
)

var flags = struct {
  TorrentFile   string  `help:Torrent file to download`
}{}

func main() {
  log.Print("hello")

  tagflag.Parse(&flags)

  c, _ := torrent.NewClient(nil)
  defer c.Close()
  t, _ := c.AddTorrentFromFile(flags.TorrentFile)
  t.SetMaxEstablishedConns(3)
  <-t.GotInfo()

  t.DownloadAll()
  for {
    log.Print("PeerID:", c.PeerID())
    stop := true
    for _, tt := range c.Torrents() {
      log.Print("Name:", tt.Name(), " Length:", tt.Length(), " Completed:",
          tt.BytesCompleted(), " Missing:", tt.BytesMissing(), " Stats:",
          tt.Stats(), " Peers:", len(tt.KnownSwarm()))
      stop = stop && (t.BytesMissing() == 0)
    }
    if stop {
      break
    }
    time.Sleep(2 * time.Second)
  }

  log.Print("torrent dowloaded")
}
