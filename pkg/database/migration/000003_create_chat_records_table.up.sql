CREATE TABLE chat_records(
    id int primary key not null auto_increment,
    message TEXT not null,
    created_by int not null,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp
);