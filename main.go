package main

import "path"

const FFMPEG_PATH = "/usr/bin/ffmpeg"
const CONVERTED_OUT_DIR = "converted-audio-files"
const INPUT_FLAG = "-i"
var SONG_CONVERSION_FLAGS = []string {
    "-vn",
    "-ar",
    "44100",
    "-ac",
    "2",
    "-b:a",
    "192k",
}
var SILENCE_INPUT = []string { "-f", "lavfi", "-i", "anullsrc" }

func GenerateSongConversionCommand(songPath string, outPath string) (string, []string) {
    fileName := path.Base(songPath)
    outFilePath := path.Join(outPath, CONVERTED_OUT_DIR, fileName)
    inputFlagSlice := []string{ INPUT_FLAG, songPath }
    args := append(
        inputFlagSlice,
        SONG_CONVERSION_FLAGS...,
    )
    args = append(args, outFilePath)
    return FFMPEG_PATH, args
}

func GenerateInputFilesFlags(songPaths *[]string) []string {
    var args []string
    for i := range *songPaths {
        args = append(args, []string { "-i", (*songPaths)[i] }...)
    }
    args = append(args, SILENCE_INPUT...)
    return args
}
