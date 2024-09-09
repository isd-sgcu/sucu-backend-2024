package constant

const (
	IMAGE         string = "IMAGE"
	DOCS          string = "DOCS"
	MAX_FILE_SIZE int64  = 32 * 1024 * 1024 // max 32 MB per file
)

var AllowedImageFileTypes = [...]string{
	".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".tiff", ".svg",
}

var AllowedDocsFileTypes = [...]string{
	".pdf", ".doc", ".docx", ".txt", ".rtf", ".odt",
	".xls", ".xlsx", ".csv",
	".ppt", ".pptx",
	".md", ".markdown",
}
