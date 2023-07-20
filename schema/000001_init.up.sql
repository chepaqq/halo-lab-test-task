CREATE TABLE
  fish (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL UNIQUE
  );

CREATE TABLE
  sensor_group (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL UNIQUE
  );

CREATE TABLE
  coordinates (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    x INT NOT NULL,
    y INT NOT NULL,
    z INT NOT NULL
  );

CREATE TABLE
  sensor (
    index INT NOT NULL,
    group_id INT REFERENCES sensor_group (id) NOT NULL,
    PRIMARY KEY (index, group_id),
    coordinates_id INT REFERENCES coordinates (id) NOT NULL,
    data_output_rate INT NOT NULL,
    temperature REAL DEFAULT 0,
    transparency INT DEFAULT 0
  );

CREATE TABLE
  detected_fish (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    fish_id INT REFERENCES fish (id) NOT NULL,
    count INT NOT NULL,
    sensor_index INT NOT NULL,
    sensor_group_id INT NOT NULL,
    FOREIGN KEY (sensor_index, sensor_group_id) REFERENCES sensor (index, group_id),
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

CREATE TABLE
  sensor_history (
    id SERIAL PRIMARY KEY NOT NULL,
    index INT NOT NULL,
    group_id INT NOT NULL,
    temperature REAL NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    FOREIGN KEY (index, group_id) REFERENCES sensor (index, group_id)
  );
