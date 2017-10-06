package e2r

import (
	"github.com/checkr/flagr/pkg/entity"
	"github.com/checkr/flagr/pkg/util"
	"github.com/checkr/flagr/swagger_gen/models"
)

// MapFlag maps flag
func MapFlag(e *entity.Flag) *models.Flag {
	r := &models.Flag{}
	r.ID = int64(e.ID)
	r.Description = util.StringPtr(e.Description)
	return r
}

// MapFlags maps flags
func MapFlags(e []entity.Flag) []*models.Flag {
	ret := make([]*models.Flag, len(e), len(e))
	for i, f := range e {
		ret[i] = MapFlag(&f)
	}
	return ret
}

// MapSegment maps segment
func MapSegment(e *entity.Segment) *models.Segment {
	r := &models.Segment{}
	r.ID = int64(e.ID)
	r.Description = util.StringPtr(e.Description)
	r.Rank = util.Int32Ptr(int32(e.Rank))
	r.RolloutPercent = util.Int32Ptr(int32(e.RolloutPercent))
	return r
}

// MapSegments maps segments
func MapSegments(e []entity.Segment) []*models.Segment {
	ret := make([]*models.Segment, len(e), len(e))
	for i, s := range e {
		ret[i] = MapSegment(&s)
	}
	return ret
}

// MapConstraint maps constraint
func MapConstraint(e *entity.Constraint) *models.Constraint {
	r := &models.Constraint{}
	r.ID = int64(e.ID)
	r.Property = util.StringPtr(e.Property)
	r.Operator = util.StringPtr(e.Operator)
	r.Value = util.StringPtr(e.Value)
	return r
}

// MapConstraints maps constraints
func MapConstraints(e []entity.Constraint) []*models.Constraint {
	ret := make([]*models.Constraint, len(e), len(e))
	for i, c := range e {
		ret[i] = MapConstraint(&c)
	}
	return ret
}

// MapDistribution maps to a distribution
func MapDistribution(e *entity.Distribution) *models.Distribution {
	r := &models.Distribution{
		Bitmap:     e.Bitmap,
		ID:         int64(e.ID),
		Percent:    util.Int32Ptr(int32(e.Percent)),
		VariantID:  util.Int64Ptr(int64(e.VariantID)),
		VariantKey: util.StringPtr(e.VariantKey),
	}
	return r
}

// MapDistributions maps distribution
func MapDistributions(e []entity.Distribution) []*models.Distribution {
	ret := make([]*models.Distribution, len(e), len(e))
	for i, d := range e {
		ret[i] = MapDistribution(&d)
	}
	return ret
}