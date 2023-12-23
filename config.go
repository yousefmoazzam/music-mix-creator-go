package mixcreator

import (
    "errors"
    "fmt"
    "io/fs"
    "os"
)

func ValidateInputDirArg(arg *string) (bool, error) {
    _, err := os.Stat(*arg)

    if errors.Is(err, fs.ErrNotExist) {
        message := fmt.Sprintf(
            "Input directory arg does not exist: %s",
            *arg,
        )
        return false, errors.New(message)
    }

    return true, nil
}
