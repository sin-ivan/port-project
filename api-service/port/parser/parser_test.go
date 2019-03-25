package parser_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	models "github.com/port-project/api-service/port"
	parser "github.com/port-project/api-service/port/parser"
)

func TestParseEmptyFileString(t *testing.T) {

	parseHandler := func(port *models.Port) {
		t.Error("no item should be parsed from empty file")
	}

	p := parser.NewPortParser(parseHandler)
	err := p.Parse("")
	if err == nil {
		t.Error("broken file should not be parsed")
	}
}

func TestParseBrokenFile(t *testing.T) {

	parseHandler := func(port *models.Port) {
		t.Error("no item should be parsed from empty file")
	}

	p := parser.NewPortParser(parseHandler)

	filePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		t.Error("failed reading test file for parsing")
	}
	filePath += "/test_files/ports_broken.json"

	e := p.Parse(filePath)
	if e == nil {
		t.Error("broken file should not be parsed")
	}
}

func TestParseEmptyFile(t *testing.T) {

	parseHandler := func(port *models.Port) {
		t.Error("no item should be parsed from empty file")
	}

	p := parser.NewPortParser(parseHandler)

	filePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		t.Error("failed reading test file for parsing")
	}
	filePath += "/test_files/ports_empty.json"

	e := p.Parse(filePath)
	if e == nil {
		t.Error("broken file should not be parsed")
	}
}

func TestParseValidFile(t *testing.T) {

	parseHandler := func(port *models.Port) {
	}

	p := parser.NewPortParser(parseHandler)

	fPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	log.Println(fPath)
	if err != nil {
		t.Error("failed reading test file for parsing")
	}
	fPath += "/test_files/ports_valid.json"

	e := p.Parse(fPath)
	if e != nil {
		t.Error("valid file should be parsed")
	}
}
