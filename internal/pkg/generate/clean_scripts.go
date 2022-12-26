package generate

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/NethermindEth/sedge/configs"
	"gopkg.in/yaml.v2"
)

var protectedFlags = map[string]bool{
	"bootnodes":             true,
	"bootstrap-node":        true,
	"fallback-web3provider": true,
}

/*
cleanFlags
This function get the raw flags data from the generated docker compose and
remove the existing duplicates. In case of errors it returns the original
flags.
params :-
a. rawFlags any
Raw flags data from
returns :-
a. any
Flags after being processed
*/
func cleanFlags(rawFlags any) any {
	// Prepare raw flags
	isString := false
	flagsElems, ok := rawFlags.([]any) // Check if flags are in a list form
	if !ok {
		flagsString, ok := rawFlags.(string) // Check if flags are in a string form
		if ok {
			// Convert from string form to list form
			isString = true
			parts := strings.Split(flagsString, "\n")
			for _, part := range parts {
				if part != "" {
					flagsElems = append(flagsElems, part)
				}
			}
		} else {
			return rawFlags // No known format. Return original
		}
	}

	// Find existing duplicates
	ReFlag := regexp.MustCompile(`^ *--(?P<VAR>[a-zA-Z0-9_\-\.]+) *[= ]{1} *(?P<VAL>.+) *$`) // Flags regex
	existingFlags := make(map[string]int, 0)
	for index, flagElem := range flagsElems {
		flagString, ok := flagElem.(string) // Check if list element is a string
		if !ok {                            // Element isn't a string
			return flagsElems // Invalid list element format. Return original
		}

		result := ReFlag.FindStringSubmatch(flagString) // Check if element its a valid flag
		if len(result) >= 3 {
			flag := result[1]           // Get flag name
			existingFlags[flag] = index // Save latest apparition for the flag name
		}
	}

	finalFlagsElems := make([]any, 0, len(flagsElems))

	// Remove duplicated flags
	for index, flagElem := range flagsElems {
		flagString := flagElem.(string)
		result := ReFlag.FindStringSubmatch(flagString) // Check if element its a valid flag
		if len(result) >= 3 {
			flag := result[1]
			if !protectedFlags[flag] && existingFlags[flag] != index { // Check if flag its not protected and its not latest apparition
				continue // Remove duplicated flag
			}
		}
		finalFlagsElems = append(finalFlagsElems, flagElem) // Add latest apparition
	}

	// Convert from list form to string form if originally was a string
	if isString {
		finalFlagsString := ""
		for index, part := range finalFlagsElems {
			if index > 0 {
				finalFlagsString += "\n"
			}
			finalFlagsString += part.(string)
		}
		return finalFlagsString
	}

	return finalFlagsElems
}

/*
CleanDockerCompose
This functions is responsible for the process of cleaning a generated
docker compose script.
params :-
a. dockerComposePath string
Path of the docker compose file to clean
returns :-
a. error
Error if any
*/
func CleanDockerCompose(dockerComposePath string) error {
	// Get docker compose file data
	file, err := os.Open(dockerComposePath)
	if err != nil {
		return fmt.Errorf(configs.CleaningDCFileError, err)
	}

	raw, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf(configs.CleaningDCFileError, err)
	}

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf(configs.CleaningDCFileError, err)
	}

	err = file.Close()
	if err != nil {
		return fmt.Errorf(configs.CleaningDCFileError, err)
	}

	// Parse docker compose data
	dockerComposeData := yaml.MapSlice{}
	if err = yaml.Unmarshal(raw, &dockerComposeData); err != nil {
		return fmt.Errorf(configs.CleaningDCFileError, err)
	}

	// Construct cleaned docker compose data
	fixedDCD := yaml.MapSlice{}
	for _, section := range dockerComposeData { // Process docker compose data sections
		fixedSection := section
		if section.Key == "services" { // Clean services section
			services, ok := section.Value.(yaml.MapSlice)
			fixedServices := section.Value
			if ok {
				fixedServicesList := yaml.MapSlice{}
				for _, service := range services { // Process docker compose services
					serviceSections, ok := service.Value.(yaml.MapSlice)
					fixedService := service
					if ok {
						fixedServiceSections := yaml.MapSlice{}
						for _, serviceSection := range serviceSections { // Clean service
							fixedServiceSection := serviceSection
							if serviceSection.Key == "command" { // Process service commands
								fixedServiceSectionValue := cleanFlags(serviceSection.Value) // Remove duplicated flags
								fixedServiceSection = yaml.MapItem{
									Key:   "command",
									Value: fixedServiceSectionValue,
								}
							}
							fixedServiceSections = append(fixedServiceSections, fixedServiceSection)
						}
						fixedService = yaml.MapItem{
							Key:   service.Key,
							Value: fixedServiceSections,
						}
					}
					fixedServicesList = append(fixedServicesList, fixedService)
				}
				fixedServices = fixedServicesList
			}
			fixedSection = yaml.MapItem{
				Key:   section.Key,
				Value: fixedServices,
			}
		}
		fixedDCD = append(fixedDCD, fixedSection)
	}

	// Overwrite docker compose file with cleaned data
	fixed, err := yaml.Marshal(fixedDCD)
	if err != nil {
		return fmt.Errorf("error cleaning docker compose file: %v", err)
	}

	return os.WriteFile(dockerComposePath, fixed, info.Mode())
}

/*
CleanEnvFile
This functions is resposible for the process of cleaning a generated `.env`
file. It removes the duplicated env var in the file keeping only the latest
apparititon of it.
params :-
a. envFilePath string
Path of the generated `.env` file
returns :-
a. error
Error if any
*/
func CleanEnvFile(envFilePath string) error {
	// Get `.env` file data
	file, err := os.Open(envFilePath)
	if err != nil {
		return fmt.Errorf(configs.CleaningEnvFileError, err)
	}

	rawLines, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf(configs.CleaningEnvFileError, err)
	}
	lines := strings.Split(string(rawLines), "\n")

	info, err := file.Stat()
	file.Close()
	if err != nil {
		return fmt.Errorf(configs.CleaningEnvFileError, err)
	}

	ReENVVAR := regexp.MustCompile(`^ *(?P<VAR>[a-zA-Z0-9_]+) *= *(?P<VAL>.+) *$`) // Variable regex

	// Find duplicated vars
	existingVars := make(map[string]int, 0)
	for index, line := range lines {
		if line == "" { // Ignore empty lines
			continue
		}
		result := ReENVVAR.FindStringSubmatch(line) // Check line its a valid variable
		if len(result) >= 3 {
			envVar := result[1]            // Get var name
			existingVars[(envVar)] = index // Save latest apparition for the var name
		}
	}

	cleanedLines := make([]string, 0, len(lines))
	for index, line := range lines {
		if line == "" { // Ignore empty lines
			continue
		}
		result := ReENVVAR.FindStringSubmatch(line) // Check line its a valid variable
		if len(result) >= 3 {
			envVar := result[1]
			if existingVars[envVar] != index { // Check if its not latest apparition
				continue // Remove duplicates
			}
		}
		cleanedLines = append(cleanedLines, line) // Add latest apparition
	}

	// Overwrite env file with cleaned data
	cleanedText := strings.Join(cleanedLines, "\n")
	err = os.WriteFile(envFilePath, []byte(cleanedText), info.Mode())
	if err != nil {
		return fmt.Errorf(configs.CleaningEnvFileError, err)
	}

	return nil
}

/*
CleanGenerated
This functions handles the process of cleaning the generation results files
params :-
a. gr GenerationResults
The generations results to be cleaned
returns:-
a. error
Error if any
*/
func CleanGenerated(genPath string) error {
	err := CleanEnvFile(filepath.Join(genPath, configs.DefaultEnvFileName))
	if err != nil {
		return err
	}

	return CleanDockerCompose(filepath.Join(genPath, configs.DefaultDockerComposeScriptName))
}
