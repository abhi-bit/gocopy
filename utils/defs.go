package gocopy

// FileStat struct
type FileStat struct {
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	Checksum string `json:"checksum"`
}
