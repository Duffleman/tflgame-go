/* events */
CREATE TABLE events (
    id text PRIMARY KEY,
    serial serial UNIQUE,
    type text NOT NULL,
    user_id text NOT NULL,
    game_id text,
    payload jsonb,
    created_at timestamp without time zone NOT NULL
);

CREATE INDEX events_user_id_idx ON events (user_id text_ops);

CREATE INDEX events_game_id_idx ON events (game_id text_ops);


/* users */
CREATE TABLE proj_users (
    id text PRIMARY KEY,
    handle text NOT NULL,
    numeric text NOT NULL,
    pin text,
    score integer NOT NULL DEFAULT 0,
    created_at timestamp without time zone NOT NULL
);

CREATE UNIQUE INDEX proj_users_handle_numeric_idx ON proj_users (handle text_ops, numeric text_ops);


/* games */
CREATE TABLE proj_games (
    id text PRIMARY KEY,
    user_id text NOT NULL REFERENCES proj_users (id),
    rounds int NOT NULL,
    include_random_spaces boolean NOT NULL,
    change_letter_order boolean NOT NULL,
    reveal_word_length boolean NOT NULL,
    lines jsonb,
    zones jsonb,
    score integer NOT NULL DEFAULT 0,
    created_at timestamp without time zone NOT NULL,
    finished_at timestamp without time zone
);

CREATE INDEX proj_games_user_id_idx ON proj_games (user_id text_ops);


/* prompts */
CREATE TABLE proj_prompts (
    id text PRIMARY KEY,
    user_id text NOT NULL REFERENCES proj_users (id),
    game_id text NOT NULL REFERENCES proj_games (id),
    prompt text NOT NULL,
    answer text NOT NULL,
	answer_given text,
    correct boolean NOT NULL DEFAULT FALSE,
    created_at timestamp without time zone NOT NULL,
    answered_at timestamp without time zone,
	hint_given_at timestamp without time zone
);

CREATE INDEX proj_prompts_user_id_idx ON proj_prompts (user_id text_ops);

CREATE INDEX proj_prompts_game_id_idx ON proj_prompts (game_id text_ops);


/* tfl_modes */
CREATE TABLE tfl_modes (
    id text PRIMARY KEY,
    name text NOT NULL
);


/* tfl_lines */
CREATE TABLE tfl_lines (
    id text PRIMARY KEY,
    name text NOT NULL,
    mode_name text NOT NULL,
    created_at timestamp without time zone NOT NULL,
    modified_at timestamp without time zone NOT NULL
);


/* tfl_stops */
CREATE TABLE tfl_stops (
    id text PRIMARY KEY,
    name text NOT NULL,
    short_name text NOT NULL,
    ics_code text NOT NULL,
    station_naptan text NOT NULL,
    status boolean NOT NULL,
    lat double precision NOT NULL,
    lon double precision NOT NULL
);


/* tfl_stops_zones */
CREATE TABLE tfl_stops_zones (
    stop_id text NOT NULL REFERENCES tfl_stops (id),
    zone text NOT NULL
);

CREATE UNIQUE INDEX tfl_stops_zones_stop_id_zone_idx ON tfl_stops_zones (stop_id text_ops, zone text_ops);


/* tfl_lines_stops */
CREATE TABLE tfl_lines_stops (
    line_id text NOT NULL REFERENCES tfl_lines (id),
    stop_id text NOT NULL REFERENCES tfl_stops (id),
    mode text NOT NULL
);

CREATE UNIQUE INDEX tfl_lines_stops_line_id_stop_id_idx ON tfl_lines_stops (line_id text_ops, stop_id text_ops);
