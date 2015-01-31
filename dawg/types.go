package dawg

// Dawg ...
//
type Dawg struct {
	root     *State
	register map[string]*State
}

// State ...
//
type State struct {
	final    bool
	children *Child
	hash     string
}

// Child collection of states
//
type Child [26]*State

// Root ...
//
func (dg Dawg) Root() *State { return dg.root }

// Register ...
func (dg Dawg) Register() map[string]*State { return dg.register }

// Final ...
func (st State) Final() bool { return st.final }

// Children ...
func (st State) Children() *Child { return st.children }

// Hash ...
//
func (st State) Hash() string { return st.hash }
