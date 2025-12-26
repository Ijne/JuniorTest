CREATE TABLE IF NOT EXISTS public.subscriptions (
    id bigserial NOT NULL,
    service_name varchar NOT NULL,
    price integer NOT NULL,
    user_id uuid NOT NULL,
    start_date date NOT NULL,
    end_date date,
    CONSTRAINT subscriptions_pk PRIMARY KEY (id)
);