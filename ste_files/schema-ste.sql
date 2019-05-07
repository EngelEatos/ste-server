-- NOTE: the code below contains the SQL for the selected object
-- as well for its dependencies and children (if applicable).
-- 
-- This feature is only a convinience in order to permit you to test
-- the whole object's SQL definition at once.
-- 
-- When exporting or generating the SQL for the whole database model
-- all objects will be placed at their original positions.


-- object: ste | type: SCHEMA --
-- DROP SCHEMA IF EXISTS ste CASCADE;
CREATE SCHEMA ste;
-- ddl-end --
ALTER SCHEMA ste OWNER TO postgres;
-- ddl-end --

-- object: ste.ste_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS ste.ste_seq CASCADE;
CREATE SEQUENCE ste.ste_seq
	INCREMENT BY 1
	MINVALUE 0
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
ALTER SEQUENCE ste.ste_seq OWNER TO postgres;
-- ddl-end --

-- object: ste.author_of_novel | type: TABLE --
-- DROP TABLE IF EXISTS ste.author_of_novel CASCADE;
CREATE TABLE ste.author_of_novel (
	author_id integer NOT NULL,
	novel_id integer NOT NULL,
	CONSTRAINT author_novel_pk PRIMARY KEY (author_id,novel_id)

);
-- ddl-end --
ALTER TABLE ste.author_of_novel OWNER TO postgres;
-- ddl-end --

-- object: ste.chapter_of_novel | type: TABLE --
-- DROP TABLE IF EXISTS ste.chapter_of_novel CASCADE;
CREATE TABLE ste.chapter_of_novel (
	novel_id integer NOT NULL,
	chapter_id integer NOT NULL,
	CONSTRAINT novel_chapter_pk PRIMARY KEY (novel_id,chapter_id)

);
-- ddl-end --
ALTER TABLE ste.chapter_of_novel OWNER TO postgres;
-- ddl-end --

-- object: ste.source | type: TABLE --
-- DROP TABLE IF EXISTS ste.source CASCADE;
CREATE TABLE ste.source (
	id serial NOT NULL,
	url text,
	name text,
	CONSTRAINT sources_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE ste.source OWNER TO postgres;
-- ddl-end --

-- object: ste.chapter_queue | type: TABLE --
-- DROP TABLE IF EXISTS ste.chapter_queue CASCADE;
CREATE TABLE ste.chapter_queue (
	novel_id integer NOT NULL,
	chapter_id integer NOT NULL,
	"queuedAt" date,
	finished bool DEFAULT false,
	"finishedAt" date,
	CONSTRAINT chapter_queue_pk PRIMARY KEY (novel_id,chapter_id)

);
-- ddl-end --
ALTER TABLE ste.chapter_queue OWNER TO postgres;
-- ddl-end --

-- object: ste.novel_queue | type: TABLE --
-- DROP TABLE IF EXISTS ste.novel_queue CASCADE;
CREATE TABLE ste.novel_queue (
	novel_id integer NOT NULL,
	url text,
	"sourceId" integer,
	"queuedAt" date NOT NULL,
	finished bool DEFAULT false,
	"finishedAt" date,
	CONSTRAINT novel_queue_pk PRIMARY KEY (novel_id,"queuedAt")

);
-- ddl-end --
ALTER TABLE ste.novel_queue OWNER TO postgres;
-- ddl-end --

-- object: ste.author | type: TABLE --
-- DROP TABLE IF EXISTS ste.author CASCADE;
CREATE TABLE ste.author (
	id serial NOT NULL,
	name text,
	url text,
	CONSTRAINT author_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE ste.author OWNER TO postgres;
-- ddl-end --

-- object: ste.genre_of_novel | type: TABLE --
-- DROP TABLE IF EXISTS ste.genre_of_novel CASCADE;
CREATE TABLE ste.genre_of_novel (
	novel_id integer NOT NULL,
	genre_id integer NOT NULL,
	CONSTRAINT nu_novel_genre_pk PRIMARY KEY (novel_id,genre_id)

);
-- ddl-end --
ALTER TABLE ste.genre_of_novel OWNER TO postgres;
-- ddl-end --

-- object: ste.genre | type: TABLE --
-- DROP TABLE IF EXISTS ste.genre CASCADE;
CREATE TABLE ste.genre (
	id integer NOT NULL,
	name text,
	url text,
	CONSTRAINT nu_genres_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE ste.genre OWNER TO postgres;
-- ddl-end --

-- object: ste.tag | type: TABLE --
-- DROP TABLE IF EXISTS ste.tag CASCADE;
CREATE TABLE ste.tag (
	id integer NOT NULL,
	name text,
	url text,
	CONSTRAINT nu_tags_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE ste.tag OWNER TO postgres;
-- ddl-end --

-- object: ste.novel | type: TABLE --
-- DROP TABLE IF EXISTS ste.novel CASCADE;
CREATE TABLE ste.novel (
	id serial NOT NULL,
	title text NOT NULL,
	chaptercount integer,
	novel_id_str text,
	type integer,
	description text,
	language_id integer,
	year integer,
	status integer,
	licensed bool,
	completly_translated bool,
	cover_id integer,
	source_id integer NOT NULL,
	"updatedAt" date,
	"fetchedAt" date,
	CONSTRAINT nu_novel_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE ste.novel OWNER TO postgres;
-- ddl-end --

-- object: ste.tag_of_novel | type: TABLE --
-- DROP TABLE IF EXISTS ste.tag_of_novel CASCADE;
CREATE TABLE ste.tag_of_novel (
	novel_id integer NOT NULL,
	tag_id integer NOT NULL,
	CONSTRAINT nu_novel_tags_pk PRIMARY KEY (novel_id,tag_id)

);
-- ddl-end --
ALTER TABLE ste.tag_of_novel OWNER TO postgres;
-- ddl-end --

-- object: ste.cover | type: TABLE --
-- DROP TABLE IF EXISTS ste.cover CASCADE;
CREATE TABLE ste.cover (
	id serial NOT NULL,
	url text NOT NULL,
	downloaded bool DEFAULT false,
	path text,
	CONSTRAINT cover_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE ste.cover OWNER TO postgres;
-- ddl-end --

-- object: ste.language | type: TABLE --
-- DROP TABLE IF EXISTS ste.language CASCADE;
CREATE TABLE ste.language (
	id integer NOT NULL,
	name text,
	url text,
	CONSTRAINT "Language_pk" PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE ste.language OWNER TO postgres;
-- ddl-end --

-- object: ste.recommendation | type: TABLE --
-- DROP TABLE IF EXISTS ste.recommendation CASCADE;
CREATE TABLE ste.recommendation (
	novel_id integer NOT NULL,
	recommended_novel_id integer NOT NULL,
	CONSTRAINT "Recommendation_pk" PRIMARY KEY (novel_id,recommended_novel_id)

);
-- ddl-end --
ALTER TABLE ste.recommendation OWNER TO postgres;
-- ddl-end --

-- object: ste.chapter | type: TABLE --
-- DROP TABLE IF EXISTS ste.chapter CASCADE;
CREATE TABLE ste.chapter (
	id integer NOT NULL DEFAULT nextval('ste.ste_seq'::regclass),
	title text,
	url text,
	idx integer,
	downloaded bool,
	path text,
	CONSTRAINT chapters_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE ste.chapter OWNER TO postgres;
-- ddl-end --

-- object: "FK_author_of_novel_author" | type: CONSTRAINT --
-- ALTER TABLE ste.author_of_novel DROP CONSTRAINT IF EXISTS "FK_author_of_novel_author" CASCADE;
ALTER TABLE ste.author_of_novel ADD CONSTRAINT "FK_author_of_novel_author" FOREIGN KEY (author_id)
REFERENCES ste.author (id) MATCH FULL
ON DELETE CASCADE ON UPDATE CASCADE;
-- ddl-end --

-- object: "FK_author_of_novel_novel" | type: CONSTRAINT --
-- ALTER TABLE ste.author_of_novel DROP CONSTRAINT IF EXISTS "FK_author_of_novel_novel" CASCADE;
ALTER TABLE ste.author_of_novel ADD CONSTRAINT "FK_author_of_novel_novel" FOREIGN KEY (novel_id)
REFERENCES ste.novel (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "FK_chapter_of_novel_chapter" | type: CONSTRAINT --
-- ALTER TABLE ste.chapter_of_novel DROP CONSTRAINT IF EXISTS "FK_chapter_of_novel_chapter" CASCADE;
ALTER TABLE ste.chapter_of_novel ADD CONSTRAINT "FK_chapter_of_novel_chapter" FOREIGN KEY (chapter_id)
REFERENCES ste.chapter (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE CASCADE;
-- ddl-end --

-- object: "FK_chapter_of_novel_novel" | type: CONSTRAINT --
-- ALTER TABLE ste.chapter_of_novel DROP CONSTRAINT IF EXISTS "FK_chapter_of_novel_novel" CASCADE;
ALTER TABLE ste.chapter_of_novel ADD CONSTRAINT "FK_chapter_of_novel_novel" FOREIGN KEY (novel_id)
REFERENCES ste.novel (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "FK_chapter_queue_chapter" | type: CONSTRAINT --
-- ALTER TABLE ste.chapter_queue DROP CONSTRAINT IF EXISTS "FK_chapter_queue_chapter" CASCADE;
ALTER TABLE ste.chapter_queue ADD CONSTRAINT "FK_chapter_queue_chapter" FOREIGN KEY (chapter_id)
REFERENCES ste.chapter (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "FK_chapter_queue_novel" | type: CONSTRAINT --
-- ALTER TABLE ste.chapter_queue DROP CONSTRAINT IF EXISTS "FK_chapter_queue_novel" CASCADE;
ALTER TABLE ste.chapter_queue ADD CONSTRAINT "FK_chapter_queue_novel" FOREIGN KEY (novel_id)
REFERENCES ste.novel (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "FK_novel_queue_novel" | type: CONSTRAINT --
-- ALTER TABLE ste.novel_queue DROP CONSTRAINT IF EXISTS "FK_novel_queue_novel" CASCADE;
ALTER TABLE ste.novel_queue ADD CONSTRAINT "FK_novel_queue_novel" FOREIGN KEY (novel_id)
REFERENCES ste.novel (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "FK_genre_of_novel_novel" | type: CONSTRAINT --
-- ALTER TABLE ste.genre_of_novel DROP CONSTRAINT IF EXISTS "FK_genre_of_novel_novel" CASCADE;
ALTER TABLE ste.genre_of_novel ADD CONSTRAINT "FK_genre_of_novel_novel" FOREIGN KEY (novel_id)
REFERENCES ste.novel (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "FK_genre_of_novel_genre" | type: CONSTRAINT --
-- ALTER TABLE ste.genre_of_novel DROP CONSTRAINT IF EXISTS "FK_genre_of_novel_genre" CASCADE;
ALTER TABLE ste.genre_of_novel ADD CONSTRAINT "FK_genre_of_novel_genre" FOREIGN KEY (genre_id)
REFERENCES ste.genre (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "FK_novel_chapter" | type: CONSTRAINT --
-- ALTER TABLE ste.novel DROP CONSTRAINT IF EXISTS "FK_novel_chapter" CASCADE;
ALTER TABLE ste.novel ADD CONSTRAINT "FK_novel_chapter" FOREIGN KEY (cover_id)
REFERENCES ste.cover (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "FK_novel_source" | type: CONSTRAINT --
-- ALTER TABLE ste.novel DROP CONSTRAINT IF EXISTS "FK_novel_source" CASCADE;
ALTER TABLE ste.novel ADD CONSTRAINT "FK_novel_source" FOREIGN KEY (source_id)
REFERENCES ste.source (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "FK_novel_language" | type: CONSTRAINT --
-- ALTER TABLE ste.novel DROP CONSTRAINT IF EXISTS "FK_novel_language" CASCADE;
ALTER TABLE ste.novel ADD CONSTRAINT "FK_novel_language" FOREIGN KEY (language_id)
REFERENCES ste.language (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "FK_tag_of_novel_novel" | type: CONSTRAINT --
-- ALTER TABLE ste.tag_of_novel DROP CONSTRAINT IF EXISTS "FK_tag_of_novel_novel" CASCADE;
ALTER TABLE ste.tag_of_novel ADD CONSTRAINT "FK_tag_of_novel_novel" FOREIGN KEY (novel_id)
REFERENCES ste.novel (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "FK_tag_of_novel_tag" | type: CONSTRAINT --
-- ALTER TABLE ste.tag_of_novel DROP CONSTRAINT IF EXISTS "FK_tag_of_novel_tag" CASCADE;
ALTER TABLE ste.tag_of_novel ADD CONSTRAINT "FK_tag_of_novel_tag" FOREIGN KEY (tag_id)
REFERENCES ste.tag (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "FK_recommendation_novel" | type: CONSTRAINT --
-- ALTER TABLE ste.recommendation DROP CONSTRAINT IF EXISTS "FK_recommendation_novel" CASCADE;
ALTER TABLE ste.recommendation ADD CONSTRAINT "FK_recommendation_novel" FOREIGN KEY (novel_id)
REFERENCES ste.novel (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "FK_recommendation_novel2" | type: CONSTRAINT --
-- ALTER TABLE ste.recommendation DROP CONSTRAINT IF EXISTS "FK_recommendation_novel2" CASCADE;
ALTER TABLE ste.recommendation ADD CONSTRAINT "FK_recommendation_novel2" FOREIGN KEY (recommended_novel_id)
REFERENCES ste.novel (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

