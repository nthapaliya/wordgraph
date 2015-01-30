# WordGraph

A wip library implementing fast and compact word lookup data structures.

## Getting there, but for now:

- [Trie](http://en.wikipedia.org/wiki/Trie) is a simple and fast, but bulky word automaton
- [Dawg](http://en.wikipedia.org/wiki/Deterministic_acyclic_finite_state_automaton) is a compressed trie, removing and merging redundant states. As of now, creating a Dawg only accepts words in lexicographical order
- CDawg and MDawg. Compressed versions of the Dawg. MDawg has vast space reductions, but has about 10X slower lookups than Dawg and CDawg.

## Upcoming:

- [Gaddag](http://en.wikipedia.org/wiki/GADDAG) is a niche data-structure that can be used for scrabble game solvers

## Limitations:
- Trie and Dawg only accept lowercase ascii (a-z only). Its currently hardcoded like this for speed, but I'll work on support for all 8 bit chars.
- Dawg.CreateFromList() can only accept a list in lexicographical order. It will return a nil otherwise. Work on non-lexicographic addition ongoing.
- MDawg and CDawg are read only. You must add all the required words to a Dawg before compression.
