CREATE TABLE IF NOT EXISTS books (
    id int auto_increment primary key,
    title varchar(255) not null,
    authors varchar(255),
    description text,
    user_id int not null,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=INNODB;