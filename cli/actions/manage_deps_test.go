package actions_test

import (
	"errors"
	"runtime"
	"testing"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	sedge_mocks "github.com/NethermindEth/sedge/mocks"
	"github.com/NethermindEth/sedge/test"
	"github.com/golang/mock/gomock"
)

func TestManageDependencies(t *testing.T) {
	// Skip test on windows
	if runtime.GOOS == "windows" {
		t.Skip("Skipping test on windows")
	}

	tests := []struct {
		name         string
		options      actions.ManageDependenciesOptions
		handleResult error
		pendingDeps  []string
		wantErr      bool
	}{
		{
			name:         "Install, existing dependencies",
			options:      actions.ManageDependenciesOptions{Dependencies: configs.Dependencies, Install: true},
			handleResult: nil,
			pendingDeps:  []string{},
			wantErr:      false,
		},
		{
			name:         "Install, non existing dependencies, supported",
			options:      actions.ManageDependenciesOptions{Dependencies: []string{"good_dep_1", "good_dep_2"}, Install: true},
			handleResult: nil,
			pendingDeps:  []string{"good_dep_1", "good_dep_2"},
			wantErr:      false,
		},
		{
			name:         "Install, existing and non existing dependencies, supported",
			options:      actions.ManageDependenciesOptions{Dependencies: []string{"docker", "good_dep_1"}, Install: true},
			handleResult: nil,
			pendingDeps:  []string{"good_dep_1"},
			wantErr:      false,
		},
		{
			name:         "Install, bad dependencies, error",
			options:      actions.ManageDependenciesOptions{Dependencies: []string{"bad_dep_1", "bad_dep_2"}, Install: true},
			handleResult: errors.New("error"),
			pendingDeps:  []string{"bad_dep_1", "bad_dep_2"},
			wantErr:      true,
		},
		{
			name:         "Prompt, existing dependencies",
			options:      actions.ManageDependenciesOptions{Dependencies: configs.Dependencies, Install: false},
			handleResult: nil,
			pendingDeps:  []string{},
			wantErr:      false,
		},
		{
			name:         "Prompt, non existing dependencies, supported",
			options:      actions.ManageDependenciesOptions{Dependencies: []string{"good_dep_1", "good_dep_2"}, Install: false},
			handleResult: nil,
			pendingDeps:  []string{"good_dep_1", "good_dep_2"},
			wantErr:      false,
		},
		{
			name:         "Prompt, existing and non existing dependencies, supported",
			options:      actions.ManageDependenciesOptions{Dependencies: []string{"docker", "good_dep_1"}, Install: false},
			handleResult: nil,
			pendingDeps:  []string{"good_dep_1"},
			wantErr:      false,
		},
		{
			name:         "Prompt, bad dependencies, error",
			options:      actions.ManageDependenciesOptions{Dependencies: []string{"bad_dep_1", "bad_dep_2"}, Install: false},
			handleResult: errors.New("error"),
			pendingDeps:  []string{"bad_dep_1", "bad_dep_2"},
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			handlers := sedge_mocks.NewMockDependenciesHandlers(ctrl)
			if len(tt.pendingDeps) > 0 && tt.options.Install {
				handlers.EXPECT().InstallDependencies(nil, tt.pendingDeps).DoAndReturn(func(_ commands.CommandRunner, pending []string) error {
					for _, dep := range pending {
						if tt.handleResult == nil {
							test.CreateFakeDep(t, dep)
						}
					}
					return tt.handleResult
				}).Times(1)
			} else if len(tt.pendingDeps) > 0 && !tt.options.Install {
				handlers.EXPECT().InstallOrShowInstructions(nil, tt.pendingDeps).DoAndReturn(func(_ commands.CommandRunner, pending []string) error {
					for _, dep := range pending {
						if tt.handleResult == nil {
							test.CreateFakeDep(t, dep)
						}
					}
					return tt.handleResult
				}).Times(1)
			}
			sedgeActions := actions.NewSedgeActions(actions.SedgeActionsOptions{
				DepsHandlers: handlers,
			})
			if err := sedgeActions.ManageDependencies(tt.options); (err != nil) != tt.wantErr {
				t.Errorf("ManageDependencies() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.handleResult == nil {
				for _, dep := range tt.pendingDeps {
					test.DeleteFakeDep(dep)
				}
			}
		})
	}
}
