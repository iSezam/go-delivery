create table if not exists ukrposhta_region (
    region_id varchar(36) primary key,
    region_ua varchar,
    region_en varchar,
    region_ru varchar,
    region_katottg varchar,
    region_koatuu varchar,
    is_deleted bool,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    ver TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
