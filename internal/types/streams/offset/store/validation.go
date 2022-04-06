package store

import "github.com/flyingdice/proton-pack/internal/validation"

var ErrStoreMapIsNil = validation.NewError(
	"offset_store_map_is_nil",
	"the offset store map must be allocated",
)

var defaultChecks = []validation.Check[*Store]{
	checkStoreNotNil(),
}

// checkStoreNotNil validates store map is not nil.
func checkStoreNotNil() validation.Check[*Store] {
	return func(s *Store) *validation.Error {
		if s.store == nil {
			return ErrStoreMapIsNil
		}
		return nil
	}
}
