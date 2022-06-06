package repository

import (
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"

	"github.com/iSezam/go-delivery/package/common/errs"
	"github.com/iSezam/go-delivery/package/common/storage"
)

const PageWidth = 20

func NewUkrposhtaRepository(db *storage.PostgresDB) *UkrposhtaPostgresRepository {
	return &UkrposhtaPostgresRepository{db}
}

type UkrposhtaPostgresRepository struct {
	db *storage.PostgresDB
}

func (t *UkrposhtaPostgresRepository) LastUpdated(table string) (lastTime time.Time) {
	if !t.db.Connect() {
		return time.Now()
	}
	conn, err := t.db.Pool.Acquire(t.db.Ctx)
	if err != nil {
		fmt.Println(err)
		return time.Now()
	}
	defer conn.Release()

	var rows pgx.Rows
	rows, err = conn.Query(t.db.Ctx, fmt.Sprintf("SELECT updated_at FROM public.ukrposhta_%s ORDER BY 1 desc LIMIT 1", table))
	defer rows.Close()

	if err != nil && err != pgx.ErrNoRows {
		fmt.Println(err)
		return time.Now()
	}

	for rows.Next() {
		rows.Scan(&lastTime)
	}

	if rows.CommandTag().RowsAffected() == 0 {
		return time.Now()
	}
	return lastTime
}

func (t *UkrposhtaPostgresRepository) Deactivate(table string, edgeTime time.Time) error {
	if !t.db.Connect() {
		return errs.DBconnect
	}

	conn, err := t.db.Pool.Acquire(t.db.Ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Release()

	commandTag, err := conn.Exec(t.db.Ctx, fmt.Sprintf("UPDATE public.ukrposhta_%s SET is_deleted=true,updated_at=now(),ver=now() WHERE updated_at<$1 AND not is_deleted", table), edgeTime)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return errs.DBNoRows
	}
	return nil
}

func (t *UkrposhtaPostgresRepository) UpdateRegion(regionID, regionUA, regionEN, regionRU, regionKATOTTG, regionKOATUU string, isDeleted bool) error {
	if !t.db.Connect() {
		return errs.DBconnect
	}

	conn, err := t.db.Pool.Acquire(t.db.Ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Release()

	commandTag, err := conn.Exec(t.db.Ctx, "INSERT INTO public.ukrposhta_region as t (region_id,region_ua,region_en,region_ru,region_katottg,region_koatuu,is_deleted,updated_at,ver) VALUES($1,$2,$3,$4,$5,$6,$7,now(),now()) "+
		"ON CONFLICT (region_id) DO "+
		"UPDATE SET region_ua=$2,region_en=$3,region_ru=$4,region_katottg=$5,region_koatuu=$6,is_deleted=$7,updated_at=now(),"+
		"ver=CASE WHEN t.region_ua=$2 AND t.region_en=$3 AND t.region_ru=$4 AND t.region_katottg=$5 AND t.region_koatuu=$6 AND t.is_deleted=$7 THEN t.ver ELSE now() END", regionID, regionUA, regionEN, regionRU, regionKATOTTG, regionKOATUU, isDeleted)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return errs.DBNoRows
	}
	return nil
}

func (t *UkrposhtaPostgresRepository) GetRegion(regionID string) (regionUA, regionEN, regionRU, regionKATOTTG, regionKOATUU string, isDeleted bool) {
	// TODO GetRegion by ID
	return "", "", "", "", "", false
}

func (t *UkrposhtaPostgresRepository) GetAllRegions(f func(regionID, regionUA, regionEN, regionRU, regionKATOTTG, regionKOATUU string, isDeleted bool) (ok bool)) error {
	if !t.db.Connect() {
		return errs.DBconnect
	}

	conn, err := t.db.Pool.Acquire(t.db.Ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	var (
		regionID, regionUA, regionEN, regionRU, regionKATOTTG, regionKOATUU string
		isDeleted                                                           bool
		rows                                                                pgx.Rows
		page                                                                int
	)
	for {
		rows, err = conn.Query(t.db.Ctx, "SELECT region_id,region_ua,region_en,region_ru,region_katottg,region_koatuu,is_deleted FROM public.ukrposhta_region ORDER BY region_id LIMIT $1 OFFSET $2", PageWidth, PageWidth*page)

		if err != nil && err != pgx.ErrNoRows {
			rows.Close()
			return err
		}

		for rows.Next() {
			err = rows.Scan(&regionID, &regionUA, &regionEN, &regionRU, &regionKATOTTG, &regionKOATUU, &isDeleted)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if !f(regionID, regionUA, regionEN, regionRU, regionKATOTTG, regionKOATUU, isDeleted) {
				rows.Close()
				return nil
			}
		}
		rows.Close()

		if rows.CommandTag().RowsAffected() == PageWidth {
			page++
			continue
		}

		if page == 0 {
			return errs.DBNoRows
		}
		break
	}
	return nil
}

// TODO: full repo Loader->UkrposhtaRepository
