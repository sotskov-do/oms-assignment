-- Table: public.building
--DROP TABLE public.building;
CREATE TABLE IF NOT EXISTS public.building (
	id serial PRIMARY KEY NOT NULL,
	"name" varchar UNIQUE NOT NULL,
	address text
);

-- Table: public.apartment
--DROP TABLE public.apartment;
CREATE TABLE IF NOT EXISTS public.apartment (
    id serial PRIMARY KEY NOT NULL,
    building_id integer NOT NULL,
    "number" varchar,
    "floor" integer,
    sq_meters integer,
    CONSTRAINT building_id FOREIGN KEY (building_id)
        REFERENCES public.building (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
);

INSERT INTO public.building (name, address) VALUES
('building_1', 'Eliyahu Meridor 79'),
('building_2', 'HaMishlatim 4'),
('building_3', NULL),
('building_4', 'Mifrats Shlomo 96');

INSERT INTO public.apartment (building_id, "number", "floor", sq_meters) VALUES
(1, '5', 2, 30),
(1, '6', 2, 45),
(2, '100', 12, 80),
(4, '42', 6, 40);
