package event

type Event struct {
	Name string
	Func func()
}

// Exec executes event Func
func (e Event) Exec() {
	if e.Func != nil {
		e.Func()
	}
}

// String returns event name
// If event contains a func it marks as "*"
func (e Event) String() string {
	if e.Func == nil {
		return e.Name
	}
	return "*" + e.Name
}

// Manager contains the graph of events
// "event1" -> "event2" -> "event3"
// "event3" -> "event4"
// "event4" -> nil
type Manager struct {
	graph *graph
}

func NewManager() *Manager {
	return &Manager{graph: &graph{
		graph: make(map[string]*node),
	}}
}

// Bind binds events src -> dst
func (s *Manager) Bind(src, dst Event) {
	s.graph.addEdge(src, dst)
}

// Initiate runs a chain of events from given e
func (s *Manager) Initiate(e Event) {
	s.graph.runChain(e)
}
