package mixcreator

import (
    "fmt"
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
