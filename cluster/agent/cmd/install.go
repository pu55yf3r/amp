package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/appcelerator/amp/cluster/agent/admin"
	"github.com/appcelerator/amp/cluster/agent/pkg/docker"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/compose/convert"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/term"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/subfuzion/stack/stack"
)

const (
	targetSingle  = "single"
	targetCluster = "cluster"
)

func NewInstallCommand() *cobra.Command {
	installCmd := &cobra.Command{
		Use:   "install",
		Short: "Set up amp services in swarm environment",
		RunE:  install,
	}
	return installCmd
}

func install(cmd *cobra.Command, args []string) error {
	stdin, stdout, stderr := term.StdStreams()
	dockerCli := docker.NewDockerCli(stdin, stdout, stderr)

	namespace := "amp"
	if len(args) > 0 {
		namespace = args[0]
	}

	target := targetSingle // TODO: Add a parameter or detect the number of swarm nodes
	files, err := getStackFiles("./stacks", target)
	if err != nil {
		return err
	}

	for _, f := range files {
		log.Println(f)
		if strings.Contains(f, "test") {
			err := deployTest(dockerCli, f, "test", 60 /* timeout in seconds */)
			stack.Remove(dockerCli, stack.RemoveOptions{Namespaces: []string{"test"}})
			if err != nil {
				return err
			}
		} else {
			err := deploy(dockerCli, f, namespace)
			if err != nil {
				return err
			}
			time.Sleep(10 * time.Second)
		}
	}
	return nil
}

// returns sorted list of yaml file pathnames
func getStackFiles(path string, target string) ([]string, error) {
	if path == "" {
		path = "./stacks"
	}
	if target == "" {
		target = targetSingle
	}

	// a bit more work but we can't just use filepath.Glob
	// since we need to match both *.yml and *.yaml
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	stackfiles := []string{}
	for _, f := range files {
		name := f.Name()
		// not compiling regex since only expecting less than a dozen stackfiles
		matched, err := regexp.MatchString("\\.ya?ml$", name)
		if err != nil {
			log.Println(err)
		} else if matched {
			if strings.Contains(name, targetSingle) && target != targetSingle {
				continue
			}
			if strings.Contains(name, targetCluster) && target != targetCluster {
				continue
			}
			stackfiles = append(stackfiles, filepath.Join(path, name))
		}
	}
	return stackfiles, nil
}

func deploy(d *command.DockerCli, stackfile string, namespace string) error {
	if namespace == "" {
		// use the stackfile basename as the default stack namespace
		namespace = filepath.Base(stackfile)
		namespace = strings.TrimSuffix(namespace, filepath.Ext(namespace))
	}

	opts := stack.DeployOptions{
		Namespace:        namespace,
		Composefile:      stackfile,
		ResolveImage:     stack.ResolveImageNever,
		SendRegistryAuth: false,
		Prune:            false,
	}
	err := stack.Deploy(d, opts)
	return err
}

func deployTest(d *command.DockerCli, stackfile string, namespace string, timeout int) error {
	// Deploy the test stack
	err := deploy(d, stackfile, namespace)
	if err != nil {
		return err
	}

	// Create a docker client
	c, err := client.NewClient(admin.DefaultURL, admin.DefaultVersion, nil, nil)
	if err != nil {
		return err
	}

	for i := 0; i < timeout; i++ {
		time.Sleep(time.Second)

		// List stack tasks
		options := types.TaskListOptions{Filters: filters.NewArgs()}
		options.Filters.Add("label", convert.LabelNamespace+"="+namespace)
		tasks, err := stack.ListTasks(context.Background(), c, options)
		if err != nil {
			return err
		}

		// Assert we have at least one task
		if len(tasks) == 0 {
			continue
		}

		// Assert we have only one task
		if len(tasks) != 1 {
			return fmt.Errorf("too many tasks for test: %d", len(tasks))
		}

		task := tasks[0]
		//log.Println("task.Status.State", task.Status.State)
		//log.Println("task.DesiredState", task.DesiredState)
		//log.Println("task.Status.Err", task.Status.Err)

		// If the task has an error, the test has failed
		if task.Status.Err != "" {
			return fmt.Errorf("test failed: %s", task.Status.Err)
		}

		// Check that the task has the expected status (complete, shutdown and no error)
		if task.Status.State == swarm.TaskStateComplete && task.DesiredState == swarm.TaskStateShutdown {
			log.Println("Test successful")
			return nil
		}
	}
	return errors.New("test timed out")
}