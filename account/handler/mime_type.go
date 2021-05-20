package handler

var validImageTypes = map[string]bool{
    "image/jpeg": true,
    "image/png":  true,
}

func isAllowedImageType(mimetype string) bool {
    _, exists := validImageTypes[mimetype]
    return exists
}
