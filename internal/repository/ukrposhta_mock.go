package repository

func NewMockRepository() *UkrposhtaPostgresRepository {
	return &UkrposhtaPostgresRepository{}
}

type UkrposhtaMockRepository struct {
}

func (t *UkrposhtaMockRepository) UpdateRegion(regionID, regionUA, regionEN, regionRU, regionKATOTTG, regionKOATUU string, isDeleted bool) error {
	return nil
}

// TODO: full repo Loader->UkrposhtaRepository
