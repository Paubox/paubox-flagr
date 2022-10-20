package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/paubox/paubox-flagr/pkg/entity"
	"github.com/paubox/paubox-flagr/pkg/util"
	"github.com/paubox/paubox-flagr/swagger_gen/restapi/operations/flag"
	"gorm.io/gorm"
)

func (c *crud) CreateFlag(params flag.CreateFlagParams) middleware.Responder {
	f := &entity.Flag{}
	if params.Body != nil {
		f.Description = util.SafeString(params.Body.Description)
		f.CreatedBy = getSubjectFromRequest(params.HTTPRequest)

		key, err := entity.CreateFlagKey(params.Body.Key)
		if err != nil {
			return flag.NewCreateFlagDefault(400).WithPayload(
				ErrorMessage("cannot create flag. %s", err))
		}
		f.Key = key
		f.Enabled = true
	}

	tx := getDB().Begin()

	if err := tx.Create(f).Error; err != nil {
		tx.Rollback()
		return flag.NewCreateFlagDefault(500).WithPayload(
			ErrorMessage("cannot create flag. %s", err))
	}

	if params.Body.Template == "feature_flag" {
		if err := LoadSimpleBooleanFlagTemplate(f, tx); err != nil {
			tx.Rollback()
			return flag.NewCreateFlagDefault(500).WithPayload(
				ErrorMessage("cannot create flag. %s", err))
		}
	} else if params.Body.Template != "" {
		return flag.NewCreateFlagDefault(400).WithPayload(
			ErrorMessage("unknown value for template: %s", params.Body.Template))
	}

	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return flag.NewCreateFlagDefault(500).WithPayload(ErrorMessage("%s", err))
	}

	resp := flag.NewCreateFlagOK()
	payload, err := e2rMapFlag(f)
	if err != nil {
		return flag.NewCreateFlagDefault(500).WithPayload(
			ErrorMessage("cannot map flag. %s", err))
	}
	resp.SetPayload(payload)

	entity.SaveFlagSnapshot(getDB(), f.ID, getSubjectFromRequest(params.HTTPRequest))

	return resp
}

// LoadSimpleBooleanFlagTemplate loads the simple boolean flag template into
// a new flag. It creates a single segment, variant ('on'), and distribution.
func LoadSimpleBooleanFlagTemplate(flag *entity.Flag, tx *gorm.DB) error {
	// Create our default segment
	s := &entity.Segment{}
	s.FlagID = flag.ID
	s.RolloutPercent = uint(100)
	s.Rank = entity.SegmentDefaultRank
	s.Description = "Default Slider"

	if err := tx.Create(s).Error; err != nil {
		return err
	}

	// .. and our default Variant
	v := &entity.Variant{}
	v.FlagID = flag.ID
	v.Key = "on"

	v2 := &entity.Variant{}
	v2.FlagID = flag.ID
	v2.Key = "off"

	if err := tx.Create(v).Error; err != nil {
		return err
	}

	if err := tx.Create(v2).Error; err != nil {
		return err
	}

	// .. and our default Distribution
	d := &entity.Distribution{}
	d.SegmentID = s.ID
	d.VariantID = v2.ID
	d.VariantKey = v2.Key
	d.Percent = uint(100)

	d2 := &entity.Distribution{}
	d2.SegmentID = s.ID
	d2.VariantID = v.ID
	d2.VariantKey = v.Key
	d2.Percent = uint(0)

	if err := tx.Create(d).Error; err != nil {
		return err
	}

	if err := tx.Create(d2).Error; err != nil {
		return err
	}

	s.Distributions = append(s.Distributions, *d)
	flag.Variants = append(flag.Variants, *v, *v2)
	flag.Segments = append(flag.Segments, *s)

	return nil
}
