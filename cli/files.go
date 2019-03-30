package cli

// StaticFiles is the static list of content for the test-files repo
// must be kept up to date with the repository
func StaticFiles() *Content {
	return &Content{
		Files: []ContentFile{
			ContentFile{
				Path: "audio/6-channel.avi",
			},
			ContentFile{
				Path: "audio/itsy-spider.mp3",
			},
			ContentFile{
				Path: "audio/mono.mov",
			},
			ContentFile{
				Path: "audio/multi-stream.mp4",
			},
			ContentFile{
				Path: "audio/old-mcdonald.m4a",
			},
			ContentFile{
				Path: "audio/stereo.mov",
			},
			ContentFile{
				Path:     "image/kids.jpg",
				Keywords: []string{"jpeg"},
			},
			ContentFile{
				Path: "image/warlock.png",
			},
			ContentFile{
				Path:     "text/brightcove-brightcove-es.vtt",
				Keywords: []string{"captions", "subtitles", "spanish"},
			},
			ContentFile{
				Path:     "text/brightcove-brightcove.vtt",
				Keywords: []string{"captions", "subtitles"},
			},
			ContentFile{
				Path:     "text/styling-demo.dfxp",
				Keywords: []string{"captions", "subtitles", "ttml"},
			},
			ContentFile{
				Path:     "text/styling-demo.scc",
				Keywords: []string{"captions", "subtitles"},
			},
			ContentFile{
				Path:     "text/vtt-inline-style.vtt",
				Keywords: []string{"captions", "subtitles"},
			},
			ContentFile{
				Path: "video/brightcove-brightcove.mp4",
			},
			ContentFile{
				Path:     "video/unicornlapse.mp4",
				Keywords: []string{"short"},
			},
			ContentFile{
				Path:     "video/vertical-video.mov",
				Keywords: []string{"portrait"},
			},
			ContentFile{
				Path:     "video/video-only.mp4",
				Keywords: []string{"noaudio"},
			},
		},
	}
}
