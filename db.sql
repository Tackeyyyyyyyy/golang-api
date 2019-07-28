-- auto-generated definition
CREATE TABLE daily_weight
(
  id              INT AUTO_INCREMENT
    PRIMARY KEY,
  weight          FLOAT    NULL,
  measurement_day DATETIME NULL
)
  ENGINE = InnoDB;

INSERT INTO daily_weight.daily_weight (id, weight, measurement_day) VALUES (1, 65.56, '2019-07-20 10:36:06');
INSERT INTO daily_weight.daily_weight (id, weight, measurement_day) VALUES (2, 65.23, '2019-07-21 10:36:06');
INSERT INTO daily_weight.daily_weight (id, weight, measurement_day) VALUES (3, 64.52, '2019-07-23 10:36:06');
INSERT INTO daily_weight.daily_weight (id, weight, measurement_day) VALUES (4, 63.89, '2019-07-24 10:36:06');
INSERT INTO daily_weight.daily_weight (id, weight, measurement_day) VALUES (5, 65.23, '2019-07-25 10:36:06');
INSERT INTO daily_weight.daily_weight (id, weight, measurement_day) VALUES (6, 69.12, '2019-07-26 10:36:06');