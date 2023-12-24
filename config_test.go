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

func TestInputDirArgValidatorIsNonEmpty(t *testing.T) {
    inputDir := t.TempDir()
    expectedRes := false
    expectedErrMessage := fmt.Sprintf(
        "Input directory is empty: %s",
        inputDir,
    )
    res, err := ValidateInputDirArg(&inputDir)

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

func TestInputDirArgValidatorHasEnoughAudioFiles(t *testing.T) {
    inputDir := t.TempDir()
    audioFile := "songA.mp3"
    audioFilePath := path.Join(inputDir, audioFile)
    os.Create(audioFilePath)
    nonAudioFiles := []string {
        "fileA.txt",
        "fileB.mp4",
        "fileC.md",
        "image.jpg",
    }
    var nonAudioFilePaths []string
    for i := range nonAudioFiles {
        nonAudioFilePaths = append(
            nonAudioFilePaths,
            path.Join(inputDir, nonAudioFiles[i]),
        )
        os.Create(nonAudioFilePaths[i])
    }
    expectedRes := false
    expectedErrMessage := "No. of audio files: need at least 2 audio files in " +
        "input dir, found 1"
    res, err := ValidateInputDirArg(&inputDir)

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

func TestInputDirArgValidatorAtLeastOneImage(t *testing.T) {
    inputDir := t.TempDir()
    songFiles := []string {
        "songA.mp3",
        "songB.mp3",
    }
    for i := range songFiles {
        songFilePath := path.Join(inputDir, songFiles[i])
        os.Create(songFilePath)
    }
    expectedRes := false
    expectedErrMessage := "Image file: detected no image file in input dir"
    res, err := ValidateInputDirArg(&inputDir)

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
