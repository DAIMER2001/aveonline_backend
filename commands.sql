-- DROP SCHEMA public;

CREATE SCHEMA public AUTHORIZATION postgres;

COMMENT ON SCHEMA public IS 'standard public schema';

-- DROP SEQUENCE public.invoices_invoice_id_seq;

CREATE SEQUENCE public.invoices_invoice_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;

-- Permissions

ALTER SEQUENCE public.invoices_invoice_id_seq OWNER TO daimer;
GRANT ALL ON SEQUENCE public.invoices_invoice_id_seq TO daimer;

-- DROP SEQUENCE public.medicines_medicine_id_seq;

CREATE SEQUENCE public.medicines_medicine_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;

-- Permissions

ALTER SEQUENCE public.medicines_medicine_id_seq OWNER TO daimer;
GRANT ALL ON SEQUENCE public.medicines_medicine_id_seq TO daimer;

-- DROP SEQUENCE public.promotions_promotion_id_seq;

CREATE SEQUENCE public.promotions_promotion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;

-- Permissions

ALTER SEQUENCE public.promotions_promotion_id_seq OWNER TO daimer;
GRANT ALL ON SEQUENCE public.promotions_promotion_id_seq TO daimer;
-- public.invoices definition

-- Drop table

-- DROP TABLE public.invoices;

CREATE TABLE public.invoices (
	invoice_id bigserial NOT NULL,
	date_payment timestamptz NULL,
	full_payment numeric(13, 2) NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT invoices_pkey PRIMARY KEY (invoice_id)
);

-- Permissions

ALTER TABLE public.invoices OWNER TO daimer;
GRANT ALL ON TABLE public.invoices TO daimer;


-- public.medicines definition

-- Drop table

-- DROP TABLE public.medicines;

CREATE TABLE public.medicines (
	medicine_id bigserial NOT NULL,
	"name" varchar(50) NULL,
	price numeric(13, 2) NULL,
	"location" varchar(50) NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT medicines_pkey PRIMARY KEY (medicine_id)
);

-- Permissions

ALTER TABLE public.medicines OWNER TO daimer;
GRANT ALL ON TABLE public.medicines TO daimer;


-- public.promotions definition

-- Drop table

-- DROP TABLE public.promotions;

CREATE TABLE public.promotions (
	promotion_id bigserial NOT NULL,
	description varchar(100) NULL,
	percentage numeric(5, 2) NULL,
	start_date timestamptz NULL,
	end_date timestamptz NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT promotions_pkey PRIMARY KEY (promotion_id)
);

-- Permissions

ALTER TABLE public.promotions OWNER TO daimer;
GRANT ALL ON TABLE public.promotions TO daimer;


-- public.invoice_medicines definition

-- Drop table

-- DROP TABLE public.invoice_medicines;

CREATE TABLE public.invoice_medicines (
	invoice_invoice_id int8 NOT NULL,
	medicine_medicine_id int8 NOT NULL,
	CONSTRAINT invoice_medicines_pkey PRIMARY KEY (invoice_invoice_id, medicine_medicine_id),
	CONSTRAINT fk_invoice_medicines_invoice FOREIGN KEY (invoice_invoice_id) REFERENCES public.invoices(invoice_id) ON DELETE SET NULL ON UPDATE CASCADE,
	CONSTRAINT fk_invoice_medicines_medicine FOREIGN KEY (medicine_medicine_id) REFERENCES public.medicines(medicine_id) ON DELETE SET NULL ON UPDATE CASCADE
);

-- Permissions

ALTER TABLE public.invoice_medicines OWNER TO daimer;
GRANT ALL ON TABLE public.invoice_medicines TO daimer;


-- public.invoice_promotions definition

-- Drop table

-- DROP TABLE public.invoice_promotions;

CREATE TABLE public.invoice_promotions (
	invoice_invoice_id int8 NOT NULL,
	promotion_promotion_id int8 NOT NULL,
	CONSTRAINT invoice_promotions_pkey PRIMARY KEY (invoice_invoice_id, promotion_promotion_id),
	CONSTRAINT fk_invoice_promotions_invoice FOREIGN KEY (invoice_invoice_id) REFERENCES public.invoices(invoice_id) ON DELETE SET NULL ON UPDATE CASCADE,
	CONSTRAINT fk_invoice_promotions_promotion FOREIGN KEY (promotion_promotion_id) REFERENCES public.promotions(promotion_id) ON DELETE SET NULL ON UPDATE CASCADE
);

-- Permissions

ALTER TABLE public.invoice_promotions OWNER TO daimer;
GRANT ALL ON TABLE public.invoice_promotions TO daimer;


-- public.promotion_medicines definition

-- Drop table

-- DROP TABLE public.promotion_medicines;

CREATE TABLE public.promotion_medicines (
	promotion_promotion_id int8 NOT NULL,
	medicine_medicine_id int8 NOT NULL,
	CONSTRAINT promotion_medicines_pkey PRIMARY KEY (promotion_promotion_id, medicine_medicine_id),
	CONSTRAINT fk_promotion_medicines_medicine FOREIGN KEY (medicine_medicine_id) REFERENCES public.medicines(medicine_id) ON DELETE SET NULL ON UPDATE CASCADE,
	CONSTRAINT fk_promotion_medicines_promotion FOREIGN KEY (promotion_promotion_id) REFERENCES public.promotions(promotion_id) ON DELETE SET NULL ON UPDATE CASCADE
);

-- Permissions

ALTER TABLE public.promotion_medicines OWNER TO daimer;
GRANT ALL ON TABLE public.promotion_medicines TO daimer;




-- Permissions

GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO public;
