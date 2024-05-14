CREATE TABLE IF NOT EXISTS Statuses(
    status varchar(30) primary key
);

CREATE TABLE IF NOT EXISTS Commands(
    id serial primary key,
    script text not null,
    status varchar(30),
    output text not null,

    FOREIGN KEY (status) REFERENCES Statuses(status)
        ON DELETE SET NULL ON UPDATE CASCADE
);