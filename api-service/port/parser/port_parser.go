package parser

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	models "github.com/port-project/api-service/port"
)

type portParser struct {
	handler func(*models.Port)
}

// NewPortParser is used to create repository instance
func NewPortParser(handler func(*models.Port)) PortParser {
	// map is used for testing purposes
	return &portParser{handler}
}

func (p *portParser) Parse(filePath string) {
	file, err := os.Open(filePath)
	//handleError(err)
	if err != nil {
		log.Fatal("Failed reading resource file")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	dec := json.NewDecoder(reader)

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	// parse values from the reader
	for dec.More() {
		identifier, err := dec.Token()
		ID := fmt.Sprintf("%v", identifier)
		if err != nil {
			log.Fatal(err)
		}
		// decode an array value (Port)
		var result *models.Port
		err = dec.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		port := &models.Port{
			ID:          ID,
			Name:        result.Name,
			City:        result.City,
			Country:     result.Country,
			Coordinates: result.Coordinates,
			Alias:       result.Alias,
			Regions:     result.Regions,
			Province:    result.Province,
			Timezone:    result.Timezone,
			Unlocs:      result.Unlocs,
			Code:        result.Code,
		}

		// submit parsed value to the handler
		p.handler(port)
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%T: %v\n", t, t)
}
