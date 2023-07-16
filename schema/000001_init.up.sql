CREATE TABLE
  sensor (
    id serial not null unique,
    codename varchar(255) not null,
    x int not null,
    y int not null,
    z int not null,
    data_output_rate int not null
  );

CREATE TABLE
  fish (
    id serial not null unique,
    name varchar(255) not null
  );

CREATE TABLE
  detected_fishes (
    id serial not null unique,
    sensor_id int not null,
    fish_id int not null,
    count int not null
  );

CREATE TABLE
  sensor_detection (
    id serial not null unique,
    sensor_id int references sensor (id) on delete cascade not null,
    detected_fishes_id int references detected_fishes (int) on delete cascade not null,
    temperature real not null,
    transparency int not null,
    detection_time time not null
  );
