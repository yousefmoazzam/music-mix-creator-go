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

func ValidateAudioFilepathsArg(paths *[]string) (bool, error) {
    if len(*paths) < 2 {
        message := fmt.Sprintf(
            "Number of audio filepaths provided should be at least 2, the number " +
            "given is: %d",
            len(*paths),
        )
        return false, errors.New(message)
    }

    var err error

    for _, filePath := range *paths {
        _, err = os.Stat(filePath)
        if err != nil && errors.Is(err, fs.ErrNotExist) {
            message := fmt.Sprintf(
                "One of the input audio filepaths doesn't exist: %s",
                filePath,
            )
            return false, errors.New(message)
        }
    }

    return true, nil
}

func ValidateOutputDirArg(arg *string) (bool, error) {
    argFileInfo, argFileErr := os.Stat(*arg)
    if argFileErr == nil && !argFileInfo.IsDir() {
        message := fmt.Sprintf(
            "Given output dir arg exists, but is a file: %s",
            *arg,
        )
        return false, errors.New(message)
    }

    parentDir := path.Dir(*arg)
    _, err := os.Stat(parentDir)

    if errors.Is(err, fs.ErrNotExist) {
        message := fmt.Sprintf(
            "Output dir can only be created if the parent dir exists, but the " +
            "parent dir %s doesn't exist",
            parentDir,
        )
        return false, errors.New(message)
    }

    dirCreationErr := os.Mkdir(*arg, 0755)
    if dirCreationErr != nil && errors.Is(dirCreationErr, fs.ErrPermission) {
        message := fmt.Sprintf(
            "Unable to create output dir %s, please check permissions of parent dir %s",
            *arg,
            parentDir,
        )
        return false, errors.New(message)
    }

    return true, nil
}
