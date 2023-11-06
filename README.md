# MySQL Example for Go

This simple program shows how to use Go packages to
work with a MySQL-database. The example features a database
of papers with a single table. The code loads data from
the database with a simple SELECT query with a WHERE clause.

## Create Database and User

First you will need to create the database, the user and link them together.
I was running this on Linux:

```
sudo su
mysql -uroot
```

Inside the MySQL console:
```
CREATE DATABASE sqlgo;

CREATE USER 'sqlgo'@'localhost' IDENTIFIED BY 'S-qGoL3-4az';

GRANT ALL PRIVILEGES ON sqlgo.* TO 'sqlgo'@'localhost';
```

Test it by trying to log in (from the bash console):
```
mysql -usqlgo -p sqlgo
```

## Create Table

Before running the code, you will need to create the following table
and add some data into it. The status value of some of the rows should be
more than 0.

```
CREATE TABLE papers (
  id INTEGER NOT NULL AUTO_INCREMENT,
  title VARCHAR(128),
  content TEXT,
  status INTEGER,
  PRIMARY KEY (id)
);
```

## Insert Test Data

```
INSERT INTO papers (title, content, status) VALUES ('Клаун', 'Пока всем', 1);
INSERT INTO papers (title, content, status) VALUES ('Жимбо', 'Привет всем', 2);
INSERT INTO papers (title, content, status) VALUES ('Иргишц', 'О воробьях...', 2);
```
