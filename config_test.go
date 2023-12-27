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
    ValidationAssertionsHelper(t, res, expectedRes, err, expectedErrMessage)
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
    ValidationAssertionsHelper(t, res, expectedRes, err, expectedErr)
}

func TestInputDirArgValidatorIsNonEmpty(t *testing.T) {
    inputDir := t.TempDir()
    expectedRes := false
    expectedErrMessage := fmt.Sprintf(
        "Input directory is empty: %s",
        inputDir,
    )
    res, err := ValidateInputDirArg(&inputDir)
    ValidationAssertionsHelper(t, res, expectedRes, err, expectedErrMessage)
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
    ValidationAssertionsHelper(t, res, expectedRes, err, expectedErrMessage)
}

func TestInputDirArgValidatorImages(t *testing.T) {
    tests := map[string]struct {
        imageFiles []string
        expectedErrMessage string
    } {
        "no images": {
            imageFiles: []string{},
            expectedErrMessage:  "Image file: detected no image file in input dir",
        },
        "too many images": {
            imageFiles: []string {
                "image1.jpg",
                "image2.png",
            },
            expectedErrMessage: "Image file: detected 2 image files in input dir, " +
                "please provide only one",
        },
    }

    songFiles := []string {
        "songA.mp3",
        "songB.mp3",
    }
    expectedRes := false

    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            inputDir := t.TempDir()
            for i := range songFiles {
                songFilePath := path.Join(inputDir, songFiles[i])
                os.Create(songFilePath)
            }

            for i := range test.imageFiles {
                imageFilePath := path.Join(inputDir, test.imageFiles[i])
                os.Create(imageFilePath)
            }

            res, err := ValidateInputDirArg(&inputDir)

            ValidationAssertionsHelper(t, res, expectedRes, err, test.expectedErrMessage)
        })
    }
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
