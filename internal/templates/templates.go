package templates

/*
GetRawTemplates :
Create a map of supported clients docker-compose.yml raw templates and
environment variables raw templates.

params :-
a. tmpType string
Kind of the template to be generated. Supported values are "execution", "consensus", "validator" and "env".

returns :-
a. map[string]string
Map of supported clients raw templates of template kind provided
*/
func GetRawTemplates(tmpType string) map[string]string {
	var tmps map[string]string

	switch tmpType {
	case "execution":
		tmps = map[string]string{
			"nethermind": Nethermind,
		}
	case "consensus":
		tmps = map[string]string{
			"lighthouse": LighthouseConsensus,
			"lodestar":   LodestarConsensus,
			"prysm":      PrysmConsensus,
			"teku":       TekuConsensus,
		}
	case "validator":
		tmps = map[string]string{
			"lighthouse": LighthouseValidator,
			"lodestar":   LodestarValidator,
			"prysm":      PrysmValidator,
			"teku":       TekuValidator,
		}
	case "env":
		tmps = map[string]string{
			"nethermind":           NethermindEnv,
			"lighthouse_consensus": LighthouseConsensusEnv,
			"lighthouse_validator": LighthouseValidatorEnv,
			"lodestar_consensus":   LodestarConsensusEnv,
			"lodestar_validator":   LodestarValidatorEnv,
			"prysm_consensus":      PrysmConsensusEnv,
			"prysm_validator":      PrysmValidatorEnv,
			"teku_consensus":       TekuConsensusEnv,
			"teku_validator":       TekuValidatorEnv,
		}
	}

	return tmps
}
