package main

import (
	"fmt"
	"os"

	"github.com/natemarks/zoochecker/input"
	"github.com/natemarks/zoochecker/version"
	"github.com/natemarks/zoochecker/zookeeper"
)

func main() {
	// Parse command line arguments into a Cluster object
	cluster := input.ParseCluster()
	fmt.Println("zoochecker version: ", version.Version)
	log := input.GetLogger()
	log.Info().Msgf("Cluster: %+v", cluster)
	error := false
	clusterStatus := zookeeper.ClusterStatus{}
	for _, node := range cluster.Nodes {
		if zookeeper.NodeIsOk(node) {
			log.Info().Msgf("Node %s:%d is OK", node.Host, node.Port)
		} else {
			log.Error().Msgf("Node %s:%d is not OK", node.Host, node.Port)
			error = true
		}

		clusterStatus.AddNodeResult(zookeeper.NodeStatus(node))
	}
	if clusterStatus.Leaders != 1 {
		log.Error().Msgf("Expected 1 leader, got %d", clusterStatus.Leaders)
		error = true
	}
	if clusterStatus.Followers != len(cluster.Nodes)-1 {
		log.Error().Msgf("Expected %d followers, got %d",
			len(cluster.Nodes)-1, clusterStatus.Followers)
		error = true
	}
	if clusterStatus.SyncedFollowers != clusterStatus.Followers {
		log.Error().Msgf("Expected %d synced followers, got %d",
			clusterStatus.Followers, clusterStatus.SyncedFollowers)
		error = true
	}
	if error {
		os.Exit(1)
	}
	log.Info().Msg("Cluster is OK")
	os.Exit(0)
}
