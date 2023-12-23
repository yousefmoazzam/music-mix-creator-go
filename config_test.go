package mixcreator

import (
    "fmt"
    "os"
    "path"
    "testing"
)

func TestInputDirArgValidatorDoexExist(t *testing.T) {
    nonexistentDir := "/tmp/nonexistent-dir"
    expectedRes := false
    expectedErrMessage := fmt.Sprintf(
        "Input directory arg does not exist: %s",
        nonexistentDir,
    )
    res, err := ValidateInputDirArg(&nonexistentDir)

    if res != expectedRes {
        t.Errorf(
            "Bool: expected %v, got %v",
            expectedRes,
            res,
        )
    }

    if err == nil {
        t.Errorf("Expected error, got nil")
    }

    if err.Error() != expectedErrMessage {
        t.Errorf(
            "Error message: expected %s, got %s",
            expectedErrMessage,
            err.Error(),
        )
    }
}

func TestInputDirArgValidatorIsDir(t *testing.T) {
    inputDir := t.TempDir()
    fileNotDir := "test-file"
    fileNotDirPath := path.Join(inputDir, fileNotDir)
    os.Create(fileNotDirPath)
    expectedRes := false
    expectedErr := fmt.Sprintf(
        "Input directory arg is not a directory: %s",
        fileNotDirPath,
    )
    res, err := ValidateInputDirArg(&fileNotDirPath)

    if res != expectedRes {
        t.Errorf(
            "Bool: expected %v, got %v",
            expectedRes,
            res,
        )
    }

    if err == nil {
        t.Errorf("Expected error, got nil")
    }

    if err.Error() != expectedErr {
        t.Errorf(
            "Error message: expected %s, got %s",
            expectedErr,
            err.Error(),
        )
    }
}
