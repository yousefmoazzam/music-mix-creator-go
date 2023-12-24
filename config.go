package mixcreator

import (
    "errors"
    "fmt"
    "io/fs"
    "os"
    "path"
    "path/filepath"
)

func ValidateInputDirArg(arg *string) (bool, error) {
    fileInfo, err := os.Stat(*arg)

    if errors.Is(err, fs.ErrNotExist) {
        message := fmt.Sprintf(
            "Input directory arg does not exist: %s",
            *arg,
        )
        return false, errors.New(message)
    }

    if !fileInfo.IsDir() {
        message := fmt.Sprintf(
            "Input directory arg is not a directory: %s",
            *arg,
        )
        return false, errors.New(message)
    }

    contents, _ := os.ReadDir(*arg)
    if len(contents) == 0 {
        message := fmt.Sprintf(
            "Input directory is empty: %s",
            *arg,
        )
        return false, errors.New(message)
    }

    globPatterns := []string {
        "*.mp3",
        "*.m4a",
        "*.aac",
        "*.wav",
        "*.flac",
    }
    var globResults []string
    for i := range globPatterns {
        filePathPattern := path.Join(*arg, globPatterns[i])
        res, _ := filepath.Glob(filePathPattern)
        globResults = append(globResults, res...)
    }

    if len(globResults) < 2 {
        message := fmt.Sprintf(
            "No. of audio files: need at least 2 audio files in " +
            "input dir, found %d",
            len(globResults),
        )
        return false, errors.New(message)
    }

    imageFileGlobPatterns := []string {
        "*.jpg",
        "*.jpeg",
        "*.png",
    }
    var imageFileGlobResults []string
    for i := range imageFileGlobPatterns {
        imageFilePathPattern := path.Join(*arg, imageFileGlobPatterns[i])
        res, _ := filepath.Glob(imageFilePathPattern)
        imageFileGlobResults = append(imageFileGlobResults, res...)
    }

    if len(imageFileGlobResults) == 0 {
        message := "Image file: detected no image file in input dir"
        return false, errors.New(message)
    }

    if len(imageFileGlobResults) > 1 {
        message := fmt.Sprintf(
            "Image file: detected %d image files in input dir, please provide only one",
            len(imageFileGlobResults),
        )
        return false, errors.New(message)
    }

    return true, nil
}
