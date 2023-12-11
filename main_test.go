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
