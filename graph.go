package event

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Logging is flag to log event executions
var Logging = true

type node struct {
	event Event
	next  *node
}

// graph is a event graph of events
type graph struct {
	mu    sync.RWMutex
	graph map[string]*node
}

// addEdge adds a new edge to the graph
func (g *graph) addEdge(src, dst Event) {
	g.mu.Lock()
	defer g.mu.Unlock()
	node := &node{
		event: dst,
		next:  g.graph[src.Name],
	}
	g.graph[src.Name] = node
}

// runChain runs events chain from given event
func (g *graph) runChain(e Event) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	g.bfs(e)
}

func (g *graph) bfs(e Event) {
	var logOutput []string

	// assume max events chain < 5
	// todo queue
	q := make([]Event, 1, 4)
	q[0] = e

	visited := make(map[string]struct{}, len(g.graph))
	visited[e.Name] = struct{}{}

	start := time.Now().Format(time.RFC3339)
	if Logging {
		logOutput = []string{}
	}

	for len(q) > 0 {
		v := q[0]
		q = q[1:]

		visited[v.Name] = struct{}{}
		v.Exec()

		if Logging {
			logOutput = append(logOutput, v.String())
		}

		ptr := g.graph[v.Name]
		for ptr != nil {
			if _, ok := visited[ptr.event.Name]; !ok {
				q = append(q, ptr.event)
			}
			ptr = ptr.next
		}
	}

	if Logging {
		fmt.Println(start, "[events]", strings.Join(logOutput, " -> "))
	}
}
