package mixcreator

import (
    "errors"
    "fmt"
    "io/fs"
    "os"
    "path"
)

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

func ValidateImageFilepathArg(arg *string) (bool, error) {
    _, err := os.Stat(*arg)
    if err != nil && errors.Is(err, fs.ErrNotExist) {
        message := fmt.Sprintf(
            "Image filepath doesn't exist: %s", *arg,
        )
        return false, errors.New(message)
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
