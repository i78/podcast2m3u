# Podcast2m3u
Podcast2m3u is a simple tool to convert Podcast feeds to m3u playlists.

## What do I need it for?
This tool is primarily targeted to users who built a home stereo with mpd as control unit. As mpd does not support podcasts, Podcast2M3u in a cronjob might be your friend.

## Usage
Podcast to stdout:
```
podcast2m3u --podcast https://feeds.example.com/example.m3u
```

just pipe the output into your target file to save it for mpd.

