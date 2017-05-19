package models

type PackParams struct {
	ApkVersion      string   `json:"apkVersion"`
	CheckedChannels []string `json:"checkedChannels"`
}

func AddPackingJob(packParams PackParams) bool {
	//todo
	return false
}
