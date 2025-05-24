CREATE TABLE `room` (
    `room` int NOT NULL,
    `p1` text NOT NULL,
    `p2` text,
    `is_start` bool NOT NULL,
    `map` json NOT NULL,
    PRIMARY KEY (`room`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
