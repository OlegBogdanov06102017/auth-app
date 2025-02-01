CREATE TABLE customers 
(
    id serial not null unique,
    phone varchar(250) not null unique,
    email varchar(250) not null unique
);

CREATE TABLE customers_lists
(
    id serial not null unique,
    userID varchar(250) not null unique,
    countryID serial not null unique
);
