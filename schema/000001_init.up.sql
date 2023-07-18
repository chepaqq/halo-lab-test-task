CREATE TABLE
  fish (
    id serial primary key not null unique,
    name varchar(255) not null
  );

CREATE TABLE
  sensor (
    index serial not null,
    group_name varchar(40) not null,
    primary key (index, group_name),
    x int not null,
    y int not null,
    z int not null,
    data_output_rate int not null
  );

CREATE TABLE
  detected_fishes (
    id serial primary key not null unique,
    fish_id int references fish (id) not null,
    count int not null
  );

CREATE TABLE
  sensor_detection (
    id serial primary key not null unique,
    sensor_index int not null,
    sensor_group varchar(40) not null,
    foreign key (sensor_index, sensor_group) references sensor (index, group_name),
    detected_fishes_id int references detected_fishes (id) not null,
    temperature real not null,
    transparency int not null,
    detected_at time not null
  );

INSERT INTO
  fish (name)
VALUES
  ('Atlantic Cod'),
  ('Sailfish'),
  ('Tuna'),
  ('Salmon'),
  ('Mackerel'),
  ('Grouper');

INSERT INTO
  sensor (index, group_name, x, y, z, data_output_rate)
VALUES
  (1, 'alpha', 10, 20, 30, 5),
  (2, 'alpha', 15, 25, 35, 10),
  (1, 'beta', 20, 30, 40, 7),
  (2, 'beta', 25, 35, 45, 12),
  (3, 'beta', 30, 40, 50, 8);

INSERT INTO
  detected_fishes (fish_id, count)
VALUES
  (1, 12),
  (2, 4),
  (1, 8),
  (3, 6),
  (5, 10),
  (1, 15),
  (4, 5),
  (2, 3),
  (6, 7);

INSERT INTO
  sensor_detection (
    sensor_index,
    sensor_group,
    detected_fishes_id,
    temperature,
    transparency,
    detected_at
  )
VALUES
  (1, 'alpha', 1, 25.5, 80, '09:00:00'),
  (1, 'alpha', 2, 29.8, 75, '09:05:00'),
  (2, 'alpha', 3, 27.3, 85, '09:10:00'),
  (2, 'alpha', 4, 26.7, 70, '09:15:00'),
  (2, 'alpha', 5, 28.1, 77, '09:20:00'),
  (1, 'beta', 6, 24.6, 82, '09:25:00'),
  (1, 'beta', 7, 26.8, 78, '09:30:00'),
  (2, 'beta', 8, 25.2, 88, '09:35:00'),
  (3, 'beta', 9, 23.9, 79, '09:40:00');
