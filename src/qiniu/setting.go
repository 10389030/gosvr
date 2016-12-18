package qiniu

const (
	QN_AccessKey = "2GUPGX1epE1AXzrnfJTmA1kkLNdto9Eodf6aoSPy"
	QN_SecretKey = "K2LEU-lrpVYtlx70rZGL8j5ioiULwrc9atVJ2C2J"
)

var DefaultUploadPolicy = UploadPolicy{
	InsertOnly: 1,
	Scope:      "assets",
	SaveKey:    "$(etag)",
	FsizeLimit: 100 << 20,
	ReturnBody: `{"key": $(key), "hash": $(etag), "w": $(imageInfo.width), "h": $(imageInfo.height), "fsize": $(fsize), "mimeType": $(mimeType), "fname": $(fname)}`,
}
