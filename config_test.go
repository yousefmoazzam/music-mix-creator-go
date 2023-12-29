package mixcreator

import (
    "fmt"
    "os"
    "path"
    "testing"
)

func TestAudioFilepathsArgValidatorNotEnoughFiles(t *testing.T) {
    songFilePathsArg := []string {}
    expectedRes := false
    expectedErrMessage := fmt.Sprintf(
        "Number of audio filepaths provided should be at least 2, the number " +
        "given is: %d",
        len(songFilePathsArg),
    )
    res, err := ValidateAudioFilepathsArg(&songFilePathsArg)
    ValidationAssertionsHelper(t, res, expectedRes, err, expectedErrMessage)
}

func TestAudioFilepathsArgValidatorNonExistentFile(t *testing.T) {
    tmpDir := t.TempDir()
    existingFile := "songA.mp3"
    nonExistentFile := "songB.mp3"
    songFilePathsArg := []string {
        path.Join(tmpDir, existingFile),
        path.Join(tmpDir, nonExistentFile),
    }
    os.Create(songFilePathsArg[0])
    expectedRes := false
    expectedErrMessage := fmt.Sprintf(
        "One of the input audio filepaths doesn't exist: %s",
        songFilePathsArg[1],
    )
    res, err := ValidateAudioFilepathsArg(&songFilePathsArg)
    ValidationAssertionsHelper(t, res, expectedRes, err, expectedErrMessage)
}

func TestImageFilepathArgValidatorFileExists(t *testing.T) {
    nonExistentImageFilePath := "/home/pictures/image.jpg"
    expectedRes := false
    expectedErrorMessage := fmt.Sprintf(
        "Image filepath doesn't exist: %s", nonExistentImageFilePath,
    )
    res, err := ValidateImageFilepathArg(&nonExistentImageFilePath)
    ValidationAssertionsHelper(t, res, expectedRes, err, expectedErrorMessage)
}

func TestOutputDirValidatorDirDoesntExistNorParent(t *testing.T) {
    outputDir := "/tmp/non-existent-parent/subdir"
    expectedRes := false
    parentDir := path.Dir(outputDir)
    expectedErrMessage := fmt.Sprintf(
        "Output dir can only be created if the parent dir exists, but the " +
        "parent dir %s doesn't exist",
        parentDir,
    )
    res, err := ValidateOutputDirArg(&outputDir)
    ValidationAssertionsHelper(t, res, expectedRes, err, expectedErrMessage)
}

func TestOutputDirValidatorDirParentExistsWriteError(t *testing.T) {
    tmpDir := t.TempDir()
    parentDirPath := path.Join(tmpDir, "parent")
    os.Mkdir(parentDirPath, 0000)
    outputDirPath := path.Join(parentDirPath, "output")
    expectedRes := false
    expectedErrMessage := fmt.Sprintf(
        "Unable to create output dir %s, please check permissions of parent dir %s",
        outputDirPath,
        parentDirPath,
    )
    res, err := ValidateOutputDirArg(&outputDirPath)
    ValidationAssertionsHelper(t, res, expectedRes, err, expectedErrMessage)
    // TODO: Check if this is necessary
    // Change permissions of parent dir so it can be cleaned up when test exits
    os.Chmod(parentDirPath, 0755)
}

func TestOutputDirValidatorArgExistsButIsFile(t *testing.T) {
    outputDir := t.TempDir()
    dummyFile := "some-file.txt"
    filePath := path.Join(outputDir, dummyFile)
    os.Create(filePath)
    expectedRes := false
    expectedErrMessage := fmt.Sprintf(
        "Given output dir arg exists, but is a file: %s",
        filePath,
    )
    res, err := ValidateOutputDirArg(&filePath)
    ValidationAssertionsHelper(t, res, expectedRes, err, expectedErrMessage)
}

func TestOutputDirValidatorDirExistsIsDir(t *testing.T) {
    outputDir := t.TempDir()
    expectedRes := true
    res, err := ValidateOutputDirArg(&outputDir)

    if res != expectedRes {
        t.Errorf(
            "Bool: expected %v, got %v", expectedRes, res,
        )
    }

    if err != nil {
        t.Errorf(
            "Expected nil, got error %s", err.Error(),
        )
    }
}

func ValidationAssertionsHelper(
    t *testing.T,
    res bool,
    expectedRes bool,
    err error,
    expectedErrMessage string,
) {
    if res != expectedRes {
        t.Errorf(
            "Bool: expected %v, got %v",
            expectedRes,
            res,
        )
    }

    if err == nil {
        t.Error("Expected error, got nil")
    }

    if err.Error() != expectedErrMessage {
        t.Errorf(
            "Error message: expected %s, got %s",
            expectedErrMessage,
            err.Error(),
        )
    }
}
