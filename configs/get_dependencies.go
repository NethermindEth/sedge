package configs

import "github.com/spf13/viper"

/*
GetDependencies :
This function is responsible for giving the dependencies needed for 1Click setup

params :-
None

returns :-
a. []string
List of dependencies needed for 1Click setup
*/
func GetDependencies() []string {
	return viper.GetStringSlice("dependencies")
}
