package nmr

import (
	"encoding/binary"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

// exptPath is the location containing an individual Bruker
// experiment, i.e. the folder containing fid, acqu, etc.
func ReadBruker(expPath string) (*Expt, error) {

	// Derive experiment name and number
	// E.g. /a/b/c/d/1 => name = d, number = 1
	number, err := strconv.Atoi(path.Base(expPath))
	if err != nil {
		return nil, errors.New("Bad experiment number")
	}
	name := path.Base(path.Dir(expPath))

	acquVars := parseAcqu(path.Join(expPath, "acqu"))
	ns, _ := strconv.Atoi(acquVars["NS"])
	td, _ := strconv.Atoi(acquVars["TD"])
	sw, _ := strconv.ParseFloat(acquVars["NS"], 64)
	sfo1, _ := strconv.ParseFloat(acquVars["SFO1"], 64)
	o1, _ := strconv.ParseFloat(acquVars["O1"], 64)

	return &Expt{Name: name, Number: number,
		NS: ns, SW: sw,
		O1P: o1 / sfo1, TD: td}, nil
}

func parseAcqu(acquPath string) map[string]string {
	acquBytes, _ := ioutil.ReadFile(acquPath)
	acqu := string(acquBytes)
	declRegexp, _ := regexp.Compile(`^[\$ ]*(\S+?)= ?(.+)`)
	vars := make(map[string]string)
	// Declaration begin with ##
	decls := strings.Split(acqu, "##")
	// decls := declExpt.Split(acqu, 10)
	for _, decl := range decls {
		match := declRegexp.FindStringSubmatch(decl)
		if len(match) == 3 {
			key, value := strings.ToUpper(match[1]), match[2]
			vars[key] = value
		}
	}
	return vars
}

func readFID(name string, td int) ([]int32, error) {
	fidFile, err := os.Open(name)
	if err != nil {
		return nil, errors.New("Failed to open fid file " + name)
	}

	stat, _ := os.Stat(name)
	fileSize := int(stat.Size())
	if td != 0 && fileSize != td {
		return nil, errors.New("Given td and fid file size disagree")
	}
	td = fileSize

	fid := make([]int32, td)
	if binary.Read(fidFile, binary.LittleEndian, fid) != nil {
		return nil, errors.New("Decoding error")
	}
	return fid, nil
}
