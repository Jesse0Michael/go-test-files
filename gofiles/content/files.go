package content

// StaticFiles is the static list of content for the test-files repo
// must be kept up to date with the repository
func StaticFiles() *Content {
	return &Content{
		Files: []File{
			File{
				Path: "audio/6-channel.avi",
			},
			File{
				Path: "audio/itsy-spider.mp3",
			},
			File{
				Path: "audio/mono.mov",
			},
			File{
				Path: "audio/multi-stream.mp4",
			},
			File{
				Path: "audio/old-mcdonald.m4a",
			},
			File{
				Path: "audio/stereo.mov",
			},
			File{
				Path:     "image/kids.jpg",
				Keywords: []string{"jpeg"},
			},
			File{
				Path: "image/warlock.png",
			},
			File{
				Path:     "text/brightcove-brightcove-es.vtt",
				Keywords: []string{"captions", "subtitles", "spanish"},
			},
			File{
				Path:     "text/brightcove-brightcove.vtt",
				Keywords: []string{"captions", "subtitles"},
			},
			File{
				Path:     "text/styling-demo.dfxp",
				Keywords: []string{"captions", "subtitles", "ttml"},
			},
			File{
				Path:     "text/styling-demo.scc",
				Keywords: []string{"captions", "subtitles"},
			},
			File{
				Path:     "text/vtt-inline-style.vtt",
				Keywords: []string{"captions", "subtitles"},
			},
			File{
				Path: "video/brightcove-brightcove.mp4",
			},
			File{
				Path:     "video/unicornlapse.mp4",
				Keywords: []string{"short"},
			},
			File{
				Path:     "video/vertical-video.mov",
				Keywords: []string{"portrait"},
			},
			File{
				Path:     "video/video-only.mp4",
				Keywords: []string{"noaudio"},
			},
		},
	}
}
