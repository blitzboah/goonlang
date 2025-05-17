-- +goose Up
-- +goose StatementBegin
create table if not exists workout_entries (
  id bigserial primary key,
  workout_id bigint not null references workouts(id) on delete cascade,
  exercise varchar(255) not null,
  sets integer not null,
  reps integer,
  duration integer,
  weight decimal(5,2),
  notes text,
  order_index integer not null,
  created_at timestamp with time zone default current_timestamp,
  constraint valid_workout_entry check(
    (reps is not null or duration is not null) and
    (reps is null or duration is null)
  )
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table workout_entries;
-- +goose StatementEnd
