CREATE TABLE chat_rooms(
    id int primary key not null auto_increment,
    name varchar(255) not null,
    created_by int not null,
    isPrivate TINYINT,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp
);