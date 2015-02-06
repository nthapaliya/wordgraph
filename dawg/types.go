package dawg

// Dawg is  directed acyclic word graph. Letters are stored as edges of letters,
// and prefixes and matching suffixes are merged. This provides fast lookup, at a
// fraction of the memory consumption of say, a dictionary. It also allows word-related
// operations like looking up words that start with a specific prefix, for example.
//
type Dawg struct {
	root     *State
	register map[string]*State
	count    int
}

// State holds the flags and outgoing edges to other states
//
type State struct {
	final    bool
	children *Child
	hash     string
	id       int
}

// Child is a list of outgoing edges
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
