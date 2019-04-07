package dtc

import "path/filepath"

func generateTranscodeVideoCommand(workspace string, j *Job) []string {
	command := []string{}
	command = append(command, "transcode-video", "--no-log")
	command = append(command, "--crop", j.Item.Crop)
	if j.Item.ForcedRate != "" {
		command = append(command, "--force-rate", j.Item.ForcedRate)
	}
	for _, f := range j.Item.Filters {
		command = append(command, "--filter", f)
	}
	command = append(
		command, "--output",
		filepath.Join(
			workspace, j.OutboxPath, j.Item.SubPath, j.Item.FileName,
		),
	)
	command = append(command, filepath.Join(workspace, j.InboxPath,
		j.Item.SubPath, j.Item.FileName))

	return command
}
