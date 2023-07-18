CREATE TABLE
  fish (
    id serial primary key not null unique,
    name varchar(255) not null
  );

CREATE TABLE
  sensor (
    index serial not null,
    group varchar(40) not null,
    primary key (index, group),
    x int not null,
    y int not null,
    z int not null,
    data_output_rate int not null
  );

CREATE TABLE
  detected_fishes (
    id serial not null unique,
    sensor_index int not null,
    sensor_group varchar(40) not null,
    foreign key (sensor_index, sensor_group) references sensor (index, group) on delete cascade not null,
    fish_id int not null,
    count int not null
  );

CREATE TABLE
  sensor_detection (
    id serial primary key not null unique,
    sensor_index int not null,
    sensor_group varchar(40) not null,
    foreign key (sensor_index, sensor_group) references sensor (index, group) on delete cascade not null,
    detected_fishes_id int references detected_fishes (int) on delete cascade not null,
    temperature real not null,
    transparency int not null,
    detected_at time not null
  );
