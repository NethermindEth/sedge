package generate

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

func cleanFlags(rawFlags any) any {
	// Prepare raw flags
	isString := false
	flagsElems, ok := rawFlags.([]any)
	if !ok {
		flagsString, ok := rawFlags.(string)
		if ok {
			isString = true
			parts := strings.Split(flagsString, "\n")
			for _, part := range parts {
				if part != "" {
					flagsElems = append(flagsElems, part)
				}
			}
		} else {
			return rawFlags
		}
	}

	// Find existing duplicates
	ReFlag := regexp.MustCompile(`^ *--(?P<VAR>[a-zA-Z0-9_\-\.]+) *[= ]{1} *(?P<VAL>.+) *$`)
	existingFlags := make(map[string]int, 0)
	for index, flagElem := range flagsElems {
		flagString, ok := flagElem.(string)
		if !ok { // Element isn't a string
			return flagsElems
		}

		result := ReFlag.FindStringSubmatch(flagString)
		if result != nil && len(result) >= 3 {
			flag := result[1]
			existingFlags[flag] = index
		}
	}

	finalFlagsElems := make([]any, 0, len(flagsElems))

	for index, flagElem := range flagsElems {
		flagString := flagElem.(string)
		result := ReFlag.FindStringSubmatch(flagString)
		if result != nil && len(result) >= 3 {
			flag := result[1]
			if existingFlags[flag] != index {
				continue
			}
		}
		finalFlagsElems = append(finalFlagsElems, flagElem)
	}

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

func CleanDockerCompose(dockerComposePath string) error {
	file, err := os.Open(dockerComposePath)
	if err != nil {
		return fmt.Errorf("error cleaning docker compose file: %v", err)
	}

	raw, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error cleaning docker compose file: %v", err)
	}

	info, err := file.Stat()
	file.Close()
	if err != nil {
		return fmt.Errorf("error cleaning docker compose file: %v", err)
	}

	dockerComposeData := yaml.MapSlice{}
	if err = yaml.Unmarshal(raw, &dockerComposeData); err != nil {
		return fmt.Errorf("error cleaning docker compose file: %v", err)
	}

	fixedDCD := yaml.MapSlice{}
	for _, section := range dockerComposeData {
		fixedSection := section
		if section.Key == "services" {
			services, ok := section.Value.(yaml.MapSlice)
			fixedServices := section.Value
			if ok {
				fixedServicesList := yaml.MapSlice{}
				for _, service := range services {
					serviceSections, ok := service.Value.(yaml.MapSlice)
					fixedService := service
					if ok {
						fixedServiceSections := yaml.MapSlice{}
						for _, serviceSection := range serviceSections {
							fixedServiceSection := serviceSection
							if serviceSection.Key == "command" {
								fixedServiceSectionValue := cleanFlags(serviceSection.Value)
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

	fixed, err := yaml.Marshal(fixedDCD)
	if err != nil {
		return fmt.Errorf("error cleaning docker compose file: %v", err)
	}

	return ioutil.WriteFile(dockerComposePath, fixed, info.Mode())
}

func CleanEnvFile(envFilePath string) error {
	file, err := os.Open(envFilePath)
	if err != nil {
		return fmt.Errorf("error cleaning env file: %v", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	lines := make([]string, 0, 20)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	info, err := file.Stat()
	file.Close()
	if err != nil {
		return fmt.Errorf("error cleaning env file: %v", err)
	}

	var ReENVVAR = regexp.MustCompile(`^ *(?P<VAR>[a-zA-Z0-9_]+) *= *(?P<VAL>.+) *$`)

	existingVars := make(map[string]int, 0)
	for index, line := range lines {
		result := ReENVVAR.FindStringSubmatch(line)
		if result != nil && len(result) >= 3 {
			envVar := result[1]
			existingVars[(envVar)] = index
		}
	}

	cleanedLines := make([]string, 0, len(lines))
	for index, line := range lines {
		result := ReENVVAR.FindStringSubmatch(line)
		if result != nil && len(result) >= 3 {
			envVar := result[1]
			if existingVars[envVar] != index {
				continue
			}
		}
		cleanedLines = append(cleanedLines, line)
	}

	cleanedText := strings.Join(cleanedLines, "\n")
	err = os.WriteFile(envFilePath, []byte(cleanedText), info.Mode())
	if err != nil {
		return fmt.Errorf("error cleaning env file: %v", err)
	}

	return nil
}

func CleanGenerated(gr GenerationResults) error {
	err := CleanEnvFile(gr.EnvFilePath)
	if err != nil {
		return err
	}

	return CleanDockerCompose(gr.DockerComposePath)
}
