CREATE TABLE users (
   id int primary key not null auto_increment,
   username varchar(255) not null ,
   password varchar(255) not null,
   created_at TIMESTAMP default CURRENT_TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);