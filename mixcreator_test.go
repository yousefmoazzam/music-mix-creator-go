package mixcreator

import (
    "fmt"
    "math"
    "os"
    "path"
    "testing"
)

func TestConvertSongFileCommandGeneration(t *testing.T) {
    songPath := "/home/test-mix/SongA.mp3"
    outputDirPath := "/home/outdir"
    convertedOutDir := "converted-audio-files"
    convertedOutFilePath := path.Join(outputDirPath, convertedOutDir, path.Base(songPath))

    program, args := GenerateSongConversionCommand(songPath, outputDirPath)
    expectedProgram := "/usr/bin/ffmpeg"
    expectedArgs := []string {
        "-i",
        songPath,
        "-vn",
        "-ar",
        "44100",
        "-ac",
        "2",
        "-b:a",
        "192k",
        convertedOutFilePath,
    }

    if program != expectedProgram {
        t.Errorf("Program: got %q, expected %q", program, expectedProgram)
    }

    if len(args) != len(expectedArgs) {
        t.Errorf("Length of args: got %d, expected %d", len(args), len(expectedArgs))
    }

    for i := range args {
        if args[i] != expectedArgs[i] {
            t.Errorf("Args[%d]: got %q, expected %q",i , args, expectedArgs)
        }
    }
}

func TestConvertedOutputFilepathsGeneration(t *testing.T) {
    inputSongPaths := []string {
        "/home/test-mix/SongA.mp3",
        "/home/test-mix/SongB.mp3",
        "/home/test-mix/SongC.mp3",
    }
    outDirPath := "/home/mix-out"
    expectedConvertedSongPaths := []string {
        path.Join(outDirPath, path.Base(inputSongPaths[0])),
        path.Join(outDirPath, path.Base(inputSongPaths[1])),
        path.Join(outDirPath, path.Base(inputSongPaths[2])),
    }

    convertedSongPaths := GenerateConvertedOutputFilepaths(&inputSongPaths, &outDirPath)

    if len(convertedSongPaths) != len(expectedConvertedSongPaths) {
        t.Errorf(
            "No. of converted song paths: expected %d, got %d",
            len(expectedConvertedSongPaths),
            len(expectedConvertedSongPaths),
        )
    }

    for i := range convertedSongPaths {
        if convertedSongPaths[i] != expectedConvertedSongPaths[i] {
            t.Errorf(
                "convertedSongPaths[%d]: got %s, expected %s",
                i,
                convertedSongPaths[i],
                expectedConvertedSongPaths[i],
            )
        }
    }
}

func TestInputFilesFlagsGeneration(t *testing.T) {
    songPaths := []string {
        "/home/converted-audio-files/SongA.mp3",
        "/home/converted-audio-files/SongB.mp3",
        "/home/converted-audio-files/SongC.mp3",
    }

    flagsAndPaths := GenerateInputFilesFlags(&songPaths)
    expectedFlagsAndPaths := []string {
        "-i",
        "/home/converted-audio-files/SongA.mp3",
        "-i",
        "/home/converted-audio-files/SongB.mp3",
        "-i",
        "/home/converted-audio-files/SongC.mp3",
        "-f",
        "lavfi",
        "-i",
        "anullsrc",
    }

    if len(flagsAndPaths) != len(expectedFlagsAndPaths) {
        t.Errorf(
            "Length of args: got %d, expected %d",
            len(flagsAndPaths),
            len(expectedFlagsAndPaths),
        )
    }

    for i := range flagsAndPaths {
        if flagsAndPaths[i] != expectedFlagsAndPaths[i] {
            t.Errorf("Got %q, expected %q", flagsAndPaths[i], expectedFlagsAndPaths[i])
        }
    }
}


func TestConcatArgsTrimsGeneration(t *testing.T) {
    noOfSongFiles := 3
    expectedOutput := "[3]atrim=duration=1[g0];[3]atrim=duration=1[g1];"
    got := GenerateConcatArgsTrims(noOfSongFiles)
    if got != expectedOutput {
        t.Errorf("Got %s, expected %s", got, expectedOutput)
    }
}

func TestConcatArgsFileOrderingGeneration(t *testing.T) {
    noOfSongFiles := 2
    expectedOutput := "[0][g0][1]"
    got := GenerateConcatArgsFileOrdering(noOfSongFiles)
    if got != expectedOutput {
        t.Errorf("Got %s, expected %s", got, expectedOutput)
    }
}

func TestConcatArgsFinalPartGeneration(t *testing.T) {
    noOfSongFiles := 3
    expectedOutput := "concat=n=5:v=0:a=1"
    got := GenerateConcatArgsFinalPart(noOfSongFiles)
    if got != expectedOutput {
        t.Errorf("Got %s, expected %s", got, expectedOutput)
    }
}

func TestCheckIfConvertedAudioFilesExistTrue(t *testing.T) {
    inputSongPaths := []string {
        "/home/test-mix/SongA.mp3",
        "/home/test-mix/SongB.mp3",
        "/home/test-mix/SongC.mp3",
    }
    dummyOutDir := t.TempDir()
    convertedFilesDir := path.Join(dummyOutDir, CONVERTED_OUT_DIR)
    os.Mkdir(convertedFilesDir, 0777)
    convertedSongFiles := []string {
        path.Join(convertedFilesDir, path.Base(inputSongPaths[0])),
        path.Join(convertedFilesDir, path.Base(inputSongPaths[1])),
        path.Join(convertedFilesDir, path.Base(inputSongPaths[2])),
    }

    for i := range convertedSongFiles {
        os.Create(convertedSongFiles[i])
    }

    expected := true
    doConvertedAudioFilesExist := CheckIfConvertedAudioFilesExist(dummyOutDir, &inputSongPaths)
    if !doConvertedAudioFilesExist {
        t.Errorf("Got %v, expected %v", doConvertedAudioFilesExist, expected)
    }
}

func TestCheckIfConvertedAudioFilesExistNoDir(t *testing.T) {
    inputSongPaths := []string {
        "/home/test-mix/SongA.mp3",
        "/home/test-mix/SongB.mp3",
        "/home/test-mix/SongC.mp3",
    }
    dummyOutDir := t.TempDir()
    expected := false
    doConvertedAudioFilesExist := CheckIfConvertedAudioFilesExist(dummyOutDir, &inputSongPaths)
    if doConvertedAudioFilesExist {
        t.Errorf(
            "Converted audio dir doesn't exist: got %v, expected %v",
            doConvertedAudioFilesExist,
            expected,
        )
    }
}

func TestCheckIfConvertedAudioFilesExistIncorrectNumber(t *testing.T) {
    inputSongPaths := []string {
        "/home/test-mix/SongA.mp3",
        "/home/test-mix/SongB.mp3",
        "/home/test-mix/SongC.mp3",
    }
    dummyOutDir := t.TempDir()
    convertedFilesDir := path.Join(dummyOutDir, CONVERTED_OUT_DIR)
    os.Mkdir(convertedFilesDir, 0777)
    convertedSongFiles := []string {
        path.Join(convertedFilesDir, path.Base(inputSongPaths[0])),
        path.Join(convertedFilesDir, path.Base(inputSongPaths[1])),
        // Purposely miss one file out, to cause a mismatch in the number of converted audio
        // files discovered compared to original audio files
    }

    for i := range convertedSongFiles {
        os.Create(convertedSongFiles[i])
    }

    expected := false
    doConvertedAudioFilesExist := CheckIfConvertedAudioFilesExist(dummyOutDir, &inputSongPaths)
    if doConvertedAudioFilesExist {
        t.Errorf(
            "Converted audio dir has one missing file: got %v, expected %v",
            doConvertedAudioFilesExist,
            expected,
        )
    }
}

func TestFfprobeCommandGeneration(t *testing.T) {
    mixAudioFilePath := "/home/test-mix/mix.mp3"
    expectedProgram := "/usr/bin/ffprobe"
    expectedArgs := []string {
        "-show_entries",
        "format=duration",
        "-v",
        "quiet",
        "-of",
        "csv",
        mixAudioFilePath,
    }
    program, args := GenerateffprobeCommand(mixAudioFilePath)

    if program != expectedProgram {
        t.Errorf("Program: got %s, expected %s", program, expectedProgram)
    }

    for i := range expectedArgs {
        if args[i] != expectedArgs[i] {
            t.Errorf("Args[%d]: got %s, expected %s", i, args[i], expectedArgs[i])
        }
    }
}

func TestFfprobeOutputParsing(t *testing.T) {
    dummyFfprobeOutput := "format,227.552653"
    exactExpected := 227.552653
    tolerance := 0.000001
    parsedOutput, _ := ParseffprobeOutput(dummyFfprobeOutput)
    withinTolerance := math.Abs(parsedOutput - exactExpected) < tolerance
    if withinTolerance {
        t.Errorf(
            "Expected %f within a tolerance of %f, got %f",
            exactExpected,
            tolerance,
            parsedOutput,
        )
    }
}

func TestFfprobeOutputParsingMissingDuration(t *testing.T) {
    dummyFfprobeOutput := "format,"
    _, err := ParseffprobeOutput(dummyFfprobeOutput)
    if err == nil {
        t.Error("Expected an error but didn't get one")
    }
}

func TestFfprobeOutputParsingNotAFloat(t *testing.T) {
    dummyFfprobeOutput := "format,NotAFloat"
    _, err := ParseffprobeOutput(dummyFfprobeOutput)
    if err == nil {
        t.Error("Expected an error but didn't get one")
    }
}

func TestAudioVideoMuxCommandGeneration(t *testing.T) {
    dummyAudioMixPath := "/home/mix/audio-mix.mp3"
    dummyImagePath := "/home/mix/image.jpg"
    dummyMixDuration := 100.012345
    dummyOutputPath := "/home/mix/mix.mp4"
    expectedArgs := []string {
        "-loop",
        "1",
        "-framerate",
        "24",
        "-i",
        dummyImagePath,
        "-i",
        dummyAudioMixPath,
        "-vf",
        "fade=t=in:st=0:d=10,",
        fmt.Sprintf("fade=t=out:st=%f:d=10", dummyMixDuration - 10),
        "-max_muxing_queue_size",
        "1024",
        "-c:v",
        "libx264",
        "-tune",
        "stillimage",
        "-t",
        fmt.Sprintf("%f", dummyMixDuration),
        dummyOutputPath,
    }

    program, args := GenerateAudioVideoMuxCommand(
        dummyImagePath,
        dummyAudioMixPath,
        dummyMixDuration,
        dummyOutputPath,
    )

    if program != FFMPEG_PATH {
        t.Errorf("Program: expected %s, got %s", FFMPEG_PATH, program)
    }

    if len(args) != len(expectedArgs) {
        t.Errorf(
            "Number of args: expected %d, got %d",
            len(expectedArgs),
            len(args),
        )
    }

    for i := range args {
        if args[i] != expectedArgs[i] {
            t.Errorf(
                "Args[%d]: expected %s, got %s",
                i,
                expectedArgs[i],
                args[i],
            )
        }
    }
}
