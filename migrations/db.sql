create table 
	users (
		id serial primary key,
		email varchar(20) default '',
		password varchar(100) default '',
		created_at varchar(255) default '',
		updated_at varchar(255) default '',
		deleted_at varchar(255) default ''
	);

INSERT INTO users (email, user_password, role_id) VALUES ('admin@admin.com', '$2a$10$DMphGc0NQ1MJZCD6tyNeBOOrpP6REzj/t.iCwr9HCNwbZ4TN7xE8S', 1); 

create table user_roles(
	user_id integer references users(id) on delete cascade on update cascade,
	role_id integer,
	primary key(user_id , role_id)
);
