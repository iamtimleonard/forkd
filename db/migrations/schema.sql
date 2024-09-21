
--
-- Postgres SQL Schema dump automatic generated by geni
--


-- TABLES 

CREATE TABLE ingredient_tags (
 ingredient character varying (255) NOT NULL,
 tag character varying (255) NOT NULL
);

CREATE TABLE ingredients (
 name character varying (255) NOT NULL,
 description text 
);

CREATE TABLE measurement_units (
 description text ,
 name character varying (255) NOT NULL
);

CREATE TABLE measurement_units_tags (
 tag character varying (255) NOT NULL,
 measurement character varying (255) NOT NULL
);

CREATE TABLE recipe_comments (
 id bigint  NOT NULL,
 post_date timestamp without time zone  NOT NULL,
 recipe_id bigint  NOT NULL,
 author_id bigint  NOT NULL,
 content text  NOT NULL
);

CREATE TABLE recipe_ingredients (
 revision_id bigint  NOT NULL,
 unit character varying (255) NOT NULL,
 ingredient character varying (255) NOT NULL,
 quantity real  NOT NULL,
 id bigint  NOT NULL,
 comment text 
);

CREATE TABLE recipe_revision_delta (
 to_recipe_revision_id bigint  NOT NULL,
 from_recipe_revision_id bigint  NOT NULL,
 id bigint  NOT NULL
);

CREATE TABLE recipe_revisions (
 publish_date timestamp without time zone  NOT NULL,
 recipe_id bigint  NOT NULL,
 description text ,
 id bigint  NOT NULL
);

CREATE TABLE recipes (
 initial_publish_date timestamp without time zone  NOT NULL,
 author_id bigint  NOT NULL,
 slug character varying (75) NOT NULL,
 forked_from bigint ,
 id bigint  NOT NULL,
 description text 
);

CREATE TABLE schema_migrations (
 id character varying (255) NOT NULL
);

CREATE TABLE tags (
 description text ,
 name character varying (255) NOT NULL
);

CREATE TABLE users (
 join_date timestamp without time zone  NOT NULL,
 id bigint  NOT NULL,
 email character varying (255) NOT NULL,
 username character varying (50) NOT NULL
);

-- CONSTRAINTS 

ALTER TABLE ingredient_tags ADD CONSTRAINT fk_ingredient_tags_ingredient FOREIGN KEY (ingredient) REFERENCES ingredients(name);

ALTER TABLE ingredient_tags ADD CONSTRAINT fk_ingredient_tags_tag FOREIGN KEY (tag) REFERENCES tags(name);

ALTER TABLE ingredients ADD CONSTRAINT ingredients_pkey PRIMARY KEY (name);

ALTER TABLE measurement_units ADD CONSTRAINT measurement_units_pkey PRIMARY KEY (name);

ALTER TABLE measurement_units_tags ADD CONSTRAINT fk_measurement_tags_measurement FOREIGN KEY (measurement) REFERENCES measurement_units(name);

ALTER TABLE measurement_units_tags ADD CONSTRAINT fk_measurement_tags_tag FOREIGN KEY (tag) REFERENCES tags(name);

ALTER TABLE recipe_comments ADD CONSTRAINT fk_recipe_comment FOREIGN KEY (recipe_id) REFERENCES recipes(id);

ALTER TABLE recipe_comments ADD CONSTRAINT fk_recipe_comment_author FOREIGN KEY (author_id) REFERENCES users(id);

ALTER TABLE recipe_comments ADD CONSTRAINT recipe_comments_pkey PRIMARY KEY (id);

ALTER TABLE recipe_ingredients ADD CONSTRAINT fk_recipe_ingredient FOREIGN KEY (ingredient) REFERENCES ingredients(name);

ALTER TABLE recipe_ingredients ADD CONSTRAINT fk_recipe_ingredient_quantity FOREIGN KEY (unit) REFERENCES measurement_units(name);

ALTER TABLE recipe_ingredients ADD CONSTRAINT fk_recipe_revision_ingredients FOREIGN KEY (revision_id) REFERENCES recipe_revisions(id);

ALTER TABLE recipe_ingredients ADD CONSTRAINT recipe_ingredients_pkey PRIMARY KEY (id);

ALTER TABLE recipe_revision_delta ADD CONSTRAINT fk_recipe_revision_deltas_from FOREIGN KEY (to_recipe_revision_id) REFERENCES recipe_revisions(id);

ALTER TABLE recipe_revision_delta ADD CONSTRAINT fk_recipe_revision_deltas_to FOREIGN KEY (from_recipe_revision_id) REFERENCES recipe_revisions(id);

ALTER TABLE recipe_revision_delta ADD CONSTRAINT recipe_revision_delta_pkey PRIMARY KEY (id);

ALTER TABLE recipe_revisions ADD CONSTRAINT fk_recipe_revisions FOREIGN KEY (recipe_id) REFERENCES recipes(id);

ALTER TABLE recipe_revisions ADD CONSTRAINT recipe_revisions_pkey PRIMARY KEY (id);

ALTER TABLE recipes ADD CONSTRAINT fk_recipe_author FOREIGN KEY (author_id) REFERENCES users(id);

ALTER TABLE recipes ADD CONSTRAINT fk_recipe_fork FOREIGN KEY (forked_from) REFERENCES recipes(id);

ALTER TABLE recipes ADD CONSTRAINT recipes_pkey PRIMARY KEY (id);

ALTER TABLE recipes ADD CONSTRAINT recipes_slug_key UNIQUE (slug);

ALTER TABLE schema_migrations ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (id);

ALTER TABLE tags ADD CONSTRAINT tags_pkey PRIMARY KEY (name);

ALTER TABLE users ADD CONSTRAINT users_email_key UNIQUE (email);

ALTER TABLE users ADD CONSTRAINT users_pkey PRIMARY KEY (id);

ALTER TABLE users ADD CONSTRAINT users_username_key UNIQUE (username);

-- INDEXES 

CREATE UNIQUE INDEX ingredients_pkey ON public.ingredients USING btree (name)

CREATE UNIQUE INDEX measurement_units_pkey ON public.measurement_units USING btree (name)

CREATE UNIQUE INDEX recipe_comments_pkey ON public.recipe_comments USING btree (id)

CREATE UNIQUE INDEX recipe_ingredients_pkey ON public.recipe_ingredients USING btree (id)

CREATE UNIQUE INDEX recipe_revision_delta_pkey ON public.recipe_revision_delta USING btree (id)

CREATE UNIQUE INDEX recipe_revisions_pkey ON public.recipe_revisions USING btree (id)

CREATE UNIQUE INDEX recipes_pkey ON public.recipes USING btree (id)

CREATE UNIQUE INDEX recipes_slug_key ON public.recipes USING btree (slug)

CREATE UNIQUE INDEX schema_migrations_pkey ON public.schema_migrations USING btree (id)

CREATE UNIQUE INDEX tags_pkey ON public.tags USING btree (name)

CREATE UNIQUE INDEX users_email_key ON public.users USING btree (email)

CREATE UNIQUE INDEX users_pkey ON public.users USING btree (id)

CREATE UNIQUE INDEX users_username_key ON public.users USING btree (username)

-- SEQUENCES 

CREATE SEQUENCE recipe_comments_id_seq AS bigint START WITH 1 MINVALUE 1 MAXVALUE 9223372036854775807 INCREMENT BY 1 CYCLE;

CREATE SEQUENCE recipe_ingredients_id_seq AS bigint START WITH 1 MINVALUE 1 MAXVALUE 9223372036854775807 INCREMENT BY 1 CYCLE;

CREATE SEQUENCE recipe_revision_delta_id_seq AS bigint START WITH 1 MINVALUE 1 MAXVALUE 9223372036854775807 INCREMENT BY 1 CYCLE;

CREATE SEQUENCE recipe_revisions_id_seq AS bigint START WITH 1 MINVALUE 1 MAXVALUE 9223372036854775807 INCREMENT BY 1 CYCLE;

CREATE SEQUENCE recipes_id_seq AS bigint START WITH 1 MINVALUE 1 MAXVALUE 9223372036854775807 INCREMENT BY 1 CYCLE;

CREATE SEQUENCE users_id_seq AS bigint START WITH 1 MINVALUE 1 MAXVALUE 9223372036854775807 INCREMENT BY 1 CYCLE;

