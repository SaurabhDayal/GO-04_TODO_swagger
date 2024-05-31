Drop table if exists todos;


CREATE TABLE IF NOT EXISTS todos
(
    id          uuid                     default gen_random_uuid() not null primary key,
    title       TEXT                                               NOT NULL,
    description TEXT                                               NOT NULL,
    pending_at  TIMESTAMP WITH TIME ZONE,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    archived_at TIMESTAMP WITH TIME ZONE
);