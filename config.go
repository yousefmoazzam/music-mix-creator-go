package mixcreator

import (
    "errors"
    "fmt"
    "io/fs"
    "os"
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

    return true, nil
}
