-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS Kategori
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    name character varying(100) NOT NULL,
    created_at timestamp(3) without time zone NOT NULL,
    created_by character varying(50) NOT NULL,
    modified_at timestamp(3) without time zone,
    modified_by character varying(50)
);
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name = 'Kategori_pkey' AND table_name = 'Kategori') THEN
        ALTER TABLE Kategori ADD CONSTRAINT "Kategori_pkey" PRIMARY KEY (id);
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS Buku
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
    modified_by character varying(50)
);

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name = 'Buku_pkey' AND table_name = 'Buku') THEN
        ALTER TABLE Buku ADD CONSTRAINT "Buku_pkey" PRIMARY KEY (id);
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name = 'buku_category_id_to_kategori_id' AND table_name = 'Buku') THEN
        ALTER TABLE Buku ADD CONSTRAINT buku_category_id_to_kategori_id FOREIGN KEY (category_id)
            REFERENCES public."Kategori" (id) MATCH SIMPLE
            ON UPDATE NO ACTION
            ON DELETE NO ACTION;
    END IF;
END $$;



-- +migrate StatementEnd