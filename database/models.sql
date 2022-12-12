-- drop table users cascade;
-- drop table profiles cascade;
-- drop table posts cascade;
-- drop table comments cascade;
-- drop table categories cascade;
-- drop table postscategorias cascade;



create table profiles (
  id serial not null unique,
  name varchar(20),
  lastname varchar (30),
  age int,
  address varchar (30),
  img varchar (10),
  primary key(id)
);

create table users (
  id serial not null unique,
  email varchar(30),
  password varchar(100),
  idprofiles int,
  primary key(id),
  CONSTRAINT fk_users_profiles FOREIGN KEY(idprofiles) REFERENCES profiles(id)
);
create table posts(
    id serial not null unique,
    description varchar (100),
    title varchar (20),
    idusers int,
    primary key(id),
    CONSTRAINT fk_posts_users foreign key (idusers)references users(id)
);
create table comments(
    id serial not null unique,
    description varchar (100),
    idusers int,
    idposts int,
    primary key(id),
    constraint fk_comments_users foreign key(idusers)references users(id),
    constraint fk_comments_posts foreign key (idposts)references posts(id)
);

create table categories(
    id serial not null unique,
    category varchar(20),
    primary key(id)
);
create table postscategorias(
    id serial not null unique,
    idposts int,
    idcategories int,
    primary key(id),
    constraint fk_postscategories_categories foreign key (idcategories)references categories(id),
    constraint fk_postscategories_posts foreign key (idposts)references posts(id)
);


-- -models the inner join
-- select * from users u
-- inner join profiles p
-- on u.id =p.idpro; 

-- select * from users;
-- select * from users inner join profiles on users.id = profiles.idpro;


-- select p.idpro,p.name,p.lastname,p.age,p.address,p.img,u.*from users as
-- u inner join profiles as p on p.idpro =u.idprofiles 

-- select p.idpro,p.name,p.lastname,p.age,p.address,p.img,
-- u.id,u.email,u.password,u.idprofiles from users as u inner join profiles as p on p.idpro=u.idprofiles  

-- select u.id,u.email,u.password,u.idprofiles,p.*from profiles as p inner join users as u on u.id=p.idpro