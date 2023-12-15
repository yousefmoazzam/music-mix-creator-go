package mixcreator

import (
	"errors"
	"fmt"
	"os"
	"path"
	"slices"
	"strings"
)

const FFMPEG_PATH = "/usr/bin/ffmpeg"
const FFPROBE_PATH = "/usr/bin/ffprobe"
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

func GenerateConvertedOutputFilepaths(songPaths *[]string, outDirPath *string) []string {
    var convertedSongPaths []string

    for i := range *songPaths {
        filename := path.Base((*songPaths)[i])
        convertedSongPaths = append(
            convertedSongPaths,
            path.Join(*outDirPath, filename),
        )
    }

    return convertedSongPaths
}

func GenerateInputFilesFlags(songPaths *[]string) []string {
    var args []string
    for i := range *songPaths {
        args = append(args, []string { "-i", (*songPaths)[i] }...)
    }
    args = append(args, SILENCE_INPUT...)
    return args
}

func GenerateConcatArgsTrims(noOfSongFiles int) string {
    var vals []string
    for i := 0; i < noOfSongFiles - 1; i++ {
        silenceTrim := fmt.Sprintf(
            "[%d]atrim=duration=1[g%d]", noOfSongFiles, i,
        )
        vals = append(vals, silenceTrim)
    }
    res := strings.Join(vals, ";") + ";"
    return res
}

func GenerateConcatArgsFileOrdering(noOfSongFiles int) string {
    var res []string
    for i := 0; i < noOfSongFiles - 1; i++ {
        songSilencePair := fmt.Sprintf("[%d][g%d]", i, i)
        res = append(res, songSilencePair)
    }
    finalSongNoSilence := fmt.Sprintf("[%d]", noOfSongFiles - 1)
    res = append(res, finalSongNoSilence)
    return strings.Join(res, "")
}

func GenerateConcatArgsFinalPart(noOfSongFiles int) string {
    noOfSilences := noOfSongFiles - 1
    noOfAudioPieces := noOfSongFiles + noOfSilences
    return fmt.Sprintf("concat=n=%d:v=0:a=1", noOfAudioPieces)
}

func GenerateConcatArgs(noOfSongFiles int) string {
    trimsPart := GenerateConcatArgsTrims(noOfSongFiles)
    orderingPart := GenerateConcatArgsFileOrdering(noOfSongFiles)
    concatPart := GenerateConcatArgsFinalPart(noOfSongFiles)
    return trimsPart + orderingPart + concatPart
}

func CheckIfConvertedAudioFilesExist(outDir string, inputFilePaths *[]string) bool {
    convertedAudioDir := path.Join(outDir, CONVERTED_OUT_DIR)
    contents, err := os.ReadDir(convertedAudioDir)

    if errors.Is(err, os.ErrNotExist) {
        return false
    }

    if len(contents) != len(*inputFilePaths) {
        return false
    }

    var inputFileNames []string
    for i := range *inputFilePaths {
        inputFileNames = append(
            inputFileNames,
            path.Base((*inputFilePaths)[i]),
        )
    }

    for i := range contents {
        if !slices.Contains(inputFileNames, contents[i].Name()) {
            return false
        }
    }

    return true
}

func GenerateffprobeCommand(mixFilePath string) (string, []string) {
    args := []string {
        "-show_entries",
        "format=duration",
        "-v",
        "quiet",
        "-of",
        "csv",
        mixFilePath,
    }
    return FFPROBE_PATH, args
}
