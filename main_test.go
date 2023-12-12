package main

import (
    "testing"
    "path"
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
