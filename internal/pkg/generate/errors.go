package generate

import "errors"

// TemplateNotFoundError is returned when the template is not found
var TemplateNotFoundError = errors.New("template not found")

// TemplateParseError is returned when the template cannot be parsed
var TemplateParseError = errors.New("template parse error")

// VariableNotFoundError is returned when the variable is not found
var VariableNotFoundError = errors.New("variable not found")

// TemplateExecuteError is returned when the template cannot be executed
var TemplateExecuteError = errors.New("template execute error")

// CanNotCreateComposeFileError is returned when the compose file cannot be created
var CanNotCreateComposeFileError = errors.New("cannot create compose file")

// EmptyDataError is returned when the data is nil
var EmptyDataError = errors.New("data is nil")

// UnableToGetClientsInfoError is returned when the client information cannot be retrieved
var UnableToGetClientsInfoError = errors.New("unable to get clients information")

// ConsensusClientNotValidError is returned when the consensus client is not valid
var ConsensusClientNotValidError = errors.New("consensus client not valid")

// ExecutionClientNotValidError is returned when the execution client is not valid
var ExecutionClientNotValidError = errors.New("execution client not valid")

// ValidatorClientNotValidError is returned when the validator client is not valid
var ValidatorClientNotValidError = errors.New("validator client not valid")
