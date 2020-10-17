/* events */
CREATE TABLE events (
    id text PRIMARY KEY,
    serial serial UNIQUE,
    type text NOT NULL,
    user_id text NOT NULL,
    game_id text,
    payload jsonb NOT NULL,
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
    difficulty_options jsonb NOT NULL,
    game_options jsonb NOT NULL,
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
    answer text,
    correct boolean NOT NULL DEFAULT FALSE,
    created_at timestamp without time zone NOT NULL,
    answered_at timestamp without time zone
);

CREATE INDEX proj_prompts_user_id_idx ON proj_prompts (user_id text_ops);

CREATE INDEX proj_prompts_game_id_idx ON proj_prompts (game_id text_ops);
