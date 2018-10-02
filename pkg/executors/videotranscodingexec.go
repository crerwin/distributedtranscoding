package executors

const detectCropCommand = "detect-crop"
const queryLogCommand = "query-handbrake-log"
const transcodeVideoCommand = "transcode-video"
const convertVideoCommand = "convert-video"

type videoTranscodingExecutor struct {
	// these functions interface with Don Melton's video-transcoding gem
	// https://github.com/donmelton/video_transcoding
	executor Executor
}

func NewVideoTranscodingExecutor() *videoTranscodingExecutor {
	v := new(videoTranscodingExecutor)
	v.executor = NewCmdExecutor("detect-crop")
	return v
}

func (e *videoTranscodingExecutor) DetectCrop(args ...string) string {
	output, _ := e.executor.Execute()
	return output
}

func (e *videoTranscodingExecutor) TranscodeVideo(args ...string) string {
	output, _ := e.executor.Execute(args...)
	return output
}
