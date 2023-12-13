package integration_tests

import (
	"mixcreator"
	"strings"
	"testing"
)

func TestConcatArgsGeneration(t *testing.T) {
    noOfSongFiles := 4
    expectedTrims := []string {
        "[4]atrim=duration=1[g0]",
        "[4]atrim=duration=1[g1]",
        "[4]atrim=duration=1[g2]",
    }
    expectedTrimsPart := strings.Join(expectedTrims, ";") + ";"
    expectedOrderingPart := "[0][g0][1][g1][2][g2][3]"
    expectedConcatPart := "concat=n=7:v=0:a=1"
    expectedConcatArgs := expectedTrimsPart + expectedOrderingPart + expectedConcatPart
    concatArgs := mixcreator.GenerateConcatArgs(noOfSongFiles)

    if concatArgs != expectedConcatArgs {
        t.Errorf("Got %s, expected %s", concatArgs, expectedConcatArgs)
    }
}
