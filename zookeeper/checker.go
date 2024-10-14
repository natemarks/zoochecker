package zookeeper

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/natemarks/zoochecker/input"
)

// SendToClusterNode sends a string to the ClusterNode's host and port via TCP and returns the response
func SendToClusterNode(node input.ClusterNode, message string) (string, error) {
	// Combine host and port into an address
	address := net.JoinHostPort(node.Host, strconv.Itoa(node.Port))

	// Create a timeout duration
	timeoutDuration := time.Duration(node.Timeout) * time.Second

	// Establish a TCP connection with timeout
	conn, err := net.DialTimeout("tcp", address, timeoutDuration)
	if err != nil {
		return "", fmt.Errorf("failed to connect to %s: %w", address, err)
	}
	defer conn.Close()

	// Send the message
	_, err = conn.Write([]byte(message + "\n"))
	if err != nil {
		return "", fmt.Errorf("failed to send message: %w", err)
	}

	// Create a buffer to read the response
	buf := make([]byte, 4096) // Adjust buffer size as necessary
	n, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	// Return the response as a string
	return string(buf[:n]), nil
}

// NodeIsOk check node with ruok
func NodeIsOk(node input.ClusterNode) bool {
	result, err := SendToClusterNode(node, "ruok")
	if result != "imok" || err != nil {
		return false
	}
	return true
}

// Status of the zookeeper cluster
type Status struct {
	Mode      string `json:"mode"`
	Followers int    `json:"followers"`
}

func parseZookeeperStatus(input string) Status {
	lines := strings.Split(input, "\n") // Split input by newline
	statusMap := make(map[string]string)

	for _, line := range lines {
		// Split each line by whitespace
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			// First element is the key, second is the value
			statusMap[parts[0]] = parts[1]
		}
	}

	// Initialize the result struct
	result := Status{
		Mode:      "standalone", // Default value for mode
		Followers: 0,            // Default value for followers
	}

	// Check for 'zk_server_state' and set 'mode'
	if mode, exists := statusMap["zk_server_state"]; exists {
		result.Mode = mode
	}

	// Check for 'zk_synced_followers' and set 'followers'
	if followers, exists := statusMap["zk_synced_followers"]; exists {
		// Convert followers value to int
		followersInt, err := strconv.Atoi(followers)
		if err == nil {
			result.Followers = followersInt
		}
	}

	return result
}

// NodeStatus returns the status of a Zookeeper cluster
func NodeStatus(node input.ClusterNode) Status {
	result, err := SendToClusterNode(node, "mntr")
	if err != nil {
		return Status{}
	}
	return parseZookeeperStatus(result)
}

// ClusterStatus represents the status of a Zookeeper cluster
type ClusterStatus struct {
	Results         []Status `json:"results"`
	Followers       int      `json:"followers"`
	Leaders         int      `json:"leaders"`
	SyncedFollowers int      `json:"synced_followers"`
}

// AddNodeResult adds a node status to the ClusterStatus
func (c *ClusterStatus) AddNodeResult(status Status) {
	c.Results = append(c.Results, status)
	if status.Mode == "leader" {
		c.Leaders++
		c.SyncedFollowers = status.Followers
	} else if status.Mode == "follower" {
		c.Followers++
	}
}
