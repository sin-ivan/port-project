package parser_test

import (
	"log"
	"os"
	"path"
	"testing"

	models "github.com/sin-ivan/port-project/api-service/port"
	parser "github.com/sin-ivan/port-project/api-service/port/parser"
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

	gp := os.Getenv("GOPATH")
	ap := path.Join(gp, "src/github.com/port-project/api-service/port/parser")

	filePath := ap + "/test_files/ports_broken.json"
	log.Println("File path:", filePath)

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

	gp := os.Getenv("GOPATH")
	ap := path.Join(gp, "src/github.com/port-project/api-service/port/parser")

	filePath := ap + "/test_files/ports_empty.json"

	e := p.Parse(filePath)
	if e == nil {
		t.Error("broken file should not be parsed")
	}
}

func TestParseValidFile(t *testing.T) {

	parseHandler := func(port *models.Port) {
	}

	p := parser.NewPortParser(parseHandler)

	gp := os.Getenv("GOPATH")
	ap := path.Join(gp, "src/github.com/port-project/api-service/port/parser")
	filePath := ap + "/test_files/ports_valid.json"

	e := p.Parse(filePath)
	if e != nil {
		t.Error("valid file should be parsed")
	}
}
