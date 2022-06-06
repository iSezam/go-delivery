package loader

import (
	"context"
	"fmt"
	"github.com/iSezam/go-delivery/package/ukrposhta"
	"time"
)

const (
	OK                          = true
	regionSleepDuration         = 7 * 24 * time.Hour
	regionBeforeRetryDuration   = 24 * time.Hour
	districtSleepDuration       = 1 * 24 * time.Hour
	districtBeforeRetryDuration = 1 * time.Hour
)

type UkrposhtaRepository interface {
	LastUpdated(table string) (lastTime time.Time)
	Deactivate(table string, edgeTime time.Time) error

	UpdateRegion(regionID, regionUA, regionEN, regionRU, regionKATOTTG, regionKOATUU string, isDeleted bool) error
	GetRegion(regionID string) (regionUA, regionEN, regionRU, regionKATOTTG, regionKOATUU string, isDeleted bool)
	GetAllRegions(func(regionID, regionUA, regionEN, regionRU, regionKATOTTG, regionKOATUU string, isDeleted bool) (ok bool)) error
	// TODO: all models
}

type UkrposhtaLoader struct {
	repo         UkrposhtaRepository
	ctx          context.Context
	regionTime   time.Time
	districtTime time.Time
}

func NewUkrposhtaLoader(repo UkrposhtaRepository) UkrposhtaLoader {
	return UkrposhtaLoader{
		repo,
		context.Background(),
		repo.LastUpdated("region").Add(regionSleepDuration),
		repo.LastUpdated("district"),
		// TODO: start time for all models
	}
}

func (u *UkrposhtaLoader) Run() {
	// try districts, then remove
	u.getDistricts()

	// TODO
	//go u.ukrposhtaWorker()
}

func (u *UkrposhtaLoader) Stop() {
	u.ctx.Done()
}

func (u *UkrposhtaLoader) ukrposhtaWorker() {
	for {
		select {
		case <-u.ctx.Done():
			return
		default:

		}
		if u.regionTime.Before(time.Now()) {
			u.getRegions()
		}
		// TODO: another model here (district, city, street, warehouse)

		time.Sleep(time.Second)
	}
}

func (u *UkrposhtaLoader) getRegions() {
	regions, _, err := ukrposhta.Regions("", "")
	if err != nil {
		u.regionTime = time.Now().Add(regionBeforeRetryDuration)
		return
	}

	t := time.Now()
	for _, region := range regions {
		err = u.repo.UpdateRegion(region.RegionID, region.RegionUA, region.RegionEN, region.RegionRU, region.RegionKATOTTG, region.RegionKOATUU, false)
		if err != nil {

		}
	}
	u.repo.Deactivate("region", t)
	u.regionTime = time.Now().Add(regionSleepDuration)
}

func (u *UkrposhtaLoader) getDistricts() {
	allOK := true
	t := time.Now()
	u.repo.GetAllRegions(func(regionID, regionUA, regionEN, regionRU, regionKATOTTG, regionKOATUU string, isDeleted bool) bool {
		if isDeleted {
			return OK
		}
		districts, _, err := ukrposhta.Districts(regionID, "")
		if err != nil {
			u.districtTime = time.Now().Add(districtBeforeRetryDuration)
			allOK = false
			return !OK
		}

		for _, district := range districts {
			// TODO: upsert
			fmt.Println(district)
		}
		return OK
	})
	if allOK {
		u.repo.Deactivate("district", t)
		u.districtTime = time.Now().Add(districtSleepDuration)
	}
}

func (u *UkrposhtaLoader) getCities() {
	// TODO: get cities by districts
}

func (u *UkrposhtaLoader) getStreets() {
	// TODO: get streets by city
}

func (u *UkrposhtaLoader) getWarehouses() {
	// TODO: get warehouses by city
}
