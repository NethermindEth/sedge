package configs

// All the strings that are needed for debugging, info and error logging.
const (
	InstallingDependenciesError = "something went wrong while installing dependencies. %s"
	ShowingInstructionsError    = "something went wrong while showing the instructions for installing %s"
	ScriptPathError             = "failed to get path for instructions file. Error: %s"
	ReadingInstructionError     = "failed to read instructions from file %s"
	OSNotSupportedError         = "installation not supported for %s"
	ProvideClientsError         = "Please provide both execution client and consensus client"
	IncorrectClientError        = "incorrect %s client name %s. Please provide correct client name. Use 'listClient' command to see the list of supported clients"
	NoClientsFoundError         = "no %s clients found. Please check your configuration file"
	ClientNotSpecifiedError     = "please enter %s"
	CreatingFileError           = "failed to create file %s. Error: %s"
	OpeningFileError            = "failed to open file %s. Error: %s"
	ClosingFileError            = "failed to close file %s"
	GeneratingScriptsError      = "generating docker-compose files for execution client %s, consensus client %s and validator client %s failed. Error: %s"
	ClientNotSupportedError     = "client %s is not supported. Please use 'listClient' command to see the list of supported clients"
	GetRawTemplatesError        = "failed to get raw templates for %s"
	LoadingTemplatesError       = "error loading templates: %s"
	PrintingFileError           = "something went wrong printing file %s. Error: %s"
	DependenciesPending         = "pending dependencies: %s"
)
