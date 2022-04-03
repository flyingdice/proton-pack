package topic_partition

import (
	"github.com/flyingdice/proton-pack/internal/validation"
)

var defaultChecks = []validation.Check[TopicPartition]{
	checkFieldsAreValid(),
}

// checkFieldsAreValid validates topic and partition are valid.
func checkFieldsAreValid() validation.Check[TopicPartition] {
	return func(tp TopicPartition) *validation.Error {
		if err := validation.RunCheckers(tp.Topic, tp.Partition); err != nil {
			return validation.NewError(
				"topic_partition_must_have_valid_fields",
				err.Error(),
			)
		}
		return nil
	}
}
