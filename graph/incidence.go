package graph

import "sync"

// TODO - regarding locks, the incidences mapping is likely going
// to be built once and updated very rarely, so the performance
// issues of a mutex during creation should be inconsequential.
// However, read locking during regular use could run into large
// performance costs for the rare case where the lock is actually
// needed. The likely best solution is to use a lock during the
// construction of the new incidences but leave it out of the
// structure once built.

// SyncMapper implements access to a sync.Map struct
type SyncMapper interface {
	Delete(key interface{})
	Load(key interface{}) (value interface{}, ok bool)
	Range(f func(key, value interface{}) bool)
	Store(key, value interface{})
	LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
	LoadAndDelete(key interface{}) (value interface{}, loaded bool)
}

func NewIncidenceMap() SyncMapper {
	return &incidenceMutex{internal: make(incidenceMap)}
}

// incidenceMutex is like a Go map[interface{}]interface{} but is safe for concurrent use
// by multiple goroutines without additional locking or coordination.
// Loads, stores, and deletes run in amortized constant time.
//
// The Map type is specialized. Most code should use a plain Go map instead,
// with separate locking or coordination, for better type safety and to make it
// easier to maintain other invariants along with the map content.
//
// The Map type is optimized for two common use cases: (1) when the entry for a given
// key is only ever written once but read many times, as in caches that only grow,
// or (2) when multiple goroutines read, write, and overwrite entries for disjoint
// sets of keys. In these two cases, use of a Map may significantly reduce lock
// contention compared to a Go map paired with a separate Mutex or RWMutex.
//
// The zero Map is empty and ready for use. A Map must not be copied after first use.
//
// Reference: Go standard library sync.Map (as of v1.16.4)
type incidenceMutex struct {
	sync.RWMutex
	internal incidenceMap
}

func (rm *incidenceMutex) Load(key *edge) (value []*node, ok bool) {
	rm.RLock()
	result, ok := rm.internal[key]
	rm.RUnlock()
	return result, ok
}

func (rm *incidenceMutex) Delete(key *edge) {
	rm.Lock()
	delete(rm.internal, key)
	rm.Unlock()
}

func (rm *incidenceMutex) Store(key *edge, value []*node) {
	rm.Lock()
	rm.internal[key] = value
	rm.Unlock()
}

// incidenceMap represents a mapping of nodes to edges
type incidenceMap map[*edge][]*node

func (i *incidenceMap) Len() int { return len(*i) }

// func (i *incidenceMap) Get(n *node) (e *edges) {

// 	for k, v := range *i {
// 		if v == n {
// 			*e = append(*e, n)
// 		}
// 	}
// }
