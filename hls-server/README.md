# A HLS music stream server

## Convert mp3 to m3u8

```sh
ffmpeg -i Carla\ Bruni\ -\ You\ Belong\ To\ Me.mp3 -c:a libmp3lame -b:a 128k -map 0:0 -f segment -segment_time 10 -segment_list outputlist.m3u8 -segment_format mpegts output%03d.ts
```

## Up Server

```sh
go run main.go
```

## URL

[http://localhost:8080/bruni/outputlist.m3u8](http://localhost:8080/bruni/outputlist.m3u8)

## GCP

[LINK](https://radiogaga-server.df.r.appspot.com/bruni/outputlist.m3u8)
