package input

import (
	"os"
	"strconv"
	"strings"

	"github.com/natemarks/zoochecker/version"
	"github.com/rs/zerolog"
)

// ClusterNode represents a single Zookeeper node in the cluster
type ClusterNode struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Timeout int    `json:"timeout"`
}

// Cluster represents a Zookeeper cluster
type Cluster struct {
	Nodes []ClusterNode `json:"nodes"`
}

const (
	// DefaultPort is the default port for Zookeeper
	DefaultPort = 2181
	// DefaultTimeout is the default timeout for Zookeeper
	DefaultTimeout = 5
)

// ParseCluster parses command line parameters into a Cluster object
func ParseCluster() Cluster {
	var cluster Cluster

	args := os.Args[1:]
	for _, arg := range args {
		// Split each argument by ":" to separate host and port
		parts := strings.Split(arg, ":")
		host := parts[0]

		// Use default port if none is provided
		port := DefaultPort
		if len(parts) > 1 {
			parsedPort, err := strconv.Atoi(parts[1])
			if err == nil {
				port = parsedPort
			}
		}

		// Create ClusterNode and append to cluster
		node := ClusterNode{
			Host:    host,
			Port:    port,
			Timeout: DefaultTimeout, // Set default timeout
		}
		cluster.Nodes = append(cluster.Nodes, node)
	}

	return cluster
}

// GetLogger returns a logger for the application
func GetLogger() (log zerolog.Logger) {
	log = zerolog.New(os.Stdout).With().Str("version", version.Version).Timestamp().Logger()
	return log
}
