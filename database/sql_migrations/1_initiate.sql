-- +migrate Up

CREATE TABLE IF NOT EXISTS "Kategori"
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    name character varying(100) NOT NULL,
    created_at timestamp(3) without time zone NOT NULL,
    created_by character varying(50) NOT NULL,
    modified_at timestamp(3) without time zone,
    modified_by character varying(50),
    CONSTRAINT "Kategori_pkey" PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS "Buku"
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    title character varying(150) NOT NULL,
    description character varying(300) NOT NULL,
    image_url character varying(200) ,
    release_year integer,
    price integer,
    total_page integer,
    thickness character varying(20),
    category_id integer,
    created_at timestamp(3) without time zone NOT NULL,
    created_by character varying(50) NOT NULL,
    modified_at timestamp(3) without time zone,
    modified_by character varying(50),
    CONSTRAINT "Buku_pkey" PRIMARY KEY (id),
    CONSTRAINT buku_category_id_to_kategori_id FOREIGN KEY (category_id)
        REFERENCES "Kategori" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);
