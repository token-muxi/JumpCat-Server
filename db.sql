CREATE TABLE `room` (
    `room` int NOT NULL,
    `p1` text NOT NULL,
    `p2` text,
    `p1_ready` bool NOT NULL,
    `p2_ready` bool NOT NULL,
    `map` json NOT NULL,
    PRIMARY KEY (`room`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
