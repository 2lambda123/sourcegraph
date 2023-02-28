package runner_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/sourcegraph/sourcegraph/enterprise/cmd/executor/internal/worker/command"
	"github.com/sourcegraph/sourcegraph/enterprise/cmd/executor/internal/worker/runner"
	"github.com/sourcegraph/sourcegraph/enterprise/internal/executor/types"
)

func TestDockerRunner_Setup(t *testing.T) {
	tests := []struct {
		name               string
		dockerAuthConfig   types.DockerAuthConfig
		expectedDockerAuth string
		expectedErr        error
	}{
		{
			name:             "Setup",
			dockerAuthConfig: types.DockerAuthConfig{},
		},
		{
			name: "Setup",
			dockerAuthConfig: types.DockerAuthConfig{
				Auths: map[string]types.DockerAuthConfigAuth{
					"index.docker.io": {
						Auth: []byte("foobar"),
					},
				},
			},
			expectedDockerAuth: `{"auths":{"index.docker.io":{"auth":"Zm9vYmFy"}}}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			options := command.DockerOptions{}
			dockerRunner := runner.NewDockerRunner(nil, nil, "", options, test.dockerAuthConfig)

			ctx := context.Background()
			err := dockerRunner.Setup(ctx)
			defer dockerRunner.Teardown(ctx)

			if test.expectedErr != nil {
				require.Error(t, err)
				assert.EqualError(t, err, test.expectedErr.Error())
			} else {
				require.NoError(t, err)
				fmt.Println(dockerRunner.TempDir())
				entries, err := os.ReadDir(dockerRunner.TempDir())
				require.NoError(t, err)
				if len(test.expectedDockerAuth) == 0 {
					require.Len(t, entries, 0)
				} else {
					require.Len(t, entries, 1)
					dockerAuthEntries, err := os.ReadDir(filepath.Join(dockerRunner.TempDir(), entries[0].Name()))
					require.NoError(t, err)
					require.Len(t, dockerAuthEntries, 1)
					f, err := os.ReadFile(filepath.Join(dockerRunner.TempDir(), entries[0].Name(), dockerAuthEntries[0].Name()))
					require.NoError(t, err)
					assert.Equal(t, test.expectedDockerAuth, string(f))
				}
			}
		})
	}
}

func TestDockerRunner_Teardown(t *testing.T) {
	dockerRunner := runner.NewDockerRunner(nil, nil, "", command.DockerOptions{}, types.DockerAuthConfig{})
	ctx := context.Background()
	err := dockerRunner.Setup(ctx)
	require.NoError(t, err)

	dir := dockerRunner.TempDir()

	_, err = os.Stat(dir)
	require.NoError(t, err)

	err = dockerRunner.Teardown(ctx)
	require.NoError(t, err)

	_, err = os.Stat(dir)
	require.Error(t, err)
	assert.True(t, os.IsNotExist(err))
}

func TestDockerRunner_Run(t *testing.T) {
	cmd := new(fakeCommand)
	logger := command.NewMockLogger()
	dir := "/some/dir"
	options := command.DockerOptions{
		ConfigPath:     "/docker/config",
		AddHostGateway: true,
		Resources: command.ResourceOptions{
			NumCPUs:   10,
			Memory:    "1G",
			DiskSpace: "10G",
		},
	}
	spec := runner.Spec{
		CommandSpec: command.Spec{
			Key:     "some-key",
			Command: []string{"echo", "hello"},
			Dir:     "/workingdir",
			Env:     []string{"FOO=bar"},
		},
		Image:      "alpine",
		ScriptPath: "/some/script",
	}

	dockerRunner := runner.NewDockerRunner(cmd, logger, dir, options, types.DockerAuthConfig{})

	expectedCommandSpec := command.Spec{
		Key: "some-key",
		Command: []string{
			"docker",
			"--config",
			"/docker/config",
			"run",
			"--rm",
			"--add-host=host.docker.internal:host-gateway",
			"--cpus",
			"10",
			"--memory",
			"1G",
			"-v",
			"/some/dir:/data",
			"-w",
			"/data/workingdir",
			"-e",
			"FOO=bar",
			"--entrypoint",
			"/bin/sh",
			"alpine",
			"/data/.sourcegraph-executor/some/script",
		},
	}
	cmd.
		On("Run", mock.Anything, logger, expectedCommandSpec).
		Return(nil)

	err := dockerRunner.Run(context.Background(), spec)
	require.NoError(t, err)
}
