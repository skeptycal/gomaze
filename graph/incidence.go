package graph

import "sync"

// NOTE - regarding locks, the incidences mapping is likely going
// to be built once and updated very rarely, so the performance
// issues of a mutex during creation should be inconsequential.
// However, read locking during regular use could run into large
// performance costs. The likely best solution is to use a lock
// during the construction of the new incidences but leave it out
// of the structure once built.

// NewIncidenceMap returns a new empty sync.Map ready for use.
func NewIncidenceMap() SyncMapper {
	return &sync.Map{}
}

// SyncMapper implements access to a sync.Map struct that is
// optimized for keys that are mostly written only once and
// where multiple goroutines require concurrent access.
//
//  sync.Map information:
//
// sync.Map is like a Go map[interface{}]interface{} but is safe
// for concurrent use by multiple goroutines without additional
// locking or coordination.
//
// Loads, stores, and deletes run in amortized constant time.
//
// The sync.Map type is specialized. Most code should use a plain
// Go map instead, with separate locking or coordination, for better
// type safety and to make it easier to maintain other invariants
// along with the map content.
//
// The Map type is optimized for two common use cases:
//
// (1) when the entry for a given key is only ever written once
// but read many times, as in caches that only grow, or
//
// (2) when multiple goroutines read, write, and overwrite entries
// for disjoint sets of keys. In these two cases, use of a Map may
// significantly reduce lock contention compared to a Go map paired
// with a separate Mutex or RWMutex.
//
// The zero Map is empty and ready for use. A Map must not be copied
// after first use.
type SyncMapper interface {

	// Load returns the value stored in the map for a key, or nil if no
	// value is present.
	// The ok result indicates whether value was found in the map.
	Load(key interface{}) (value interface{}, ok bool)

	// Store sets the value for a key.
	Store(key, value interface{})

	// LoadOrStore returns the existing value for the key if present.
	// Otherwise, it stores and returns the given value.
	// The loaded result is true if the value was loaded, false if stored.
	LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)

	// LoadAndDelete deletes the value for a key, returning the previous
	// value, if any.
	// The loaded result reports whether the key was present.
	LoadAndDelete(key interface{}) (value interface{}, loaded bool)

	// Delete deletes the value for a key.
	Delete(key interface{})

	// Range calls f sequentially for each key and value present in
	// the map. If f returns false, range stops the iteration.
	//
	// Range does not necessarily correspond to any consistent
	// snapshot of the Map's contents: no key will be visited more
	// than once, but if the value for any key is stored or deleted
	// concurrently, Range may reflect any mapping for that key
	// from any point during the Range call.
	//
	// Range may be O(N) with the number of elements in the map
	// even if f returns false after a constant number of calls.
	Range(f func(key, value interface{}) bool)
}
