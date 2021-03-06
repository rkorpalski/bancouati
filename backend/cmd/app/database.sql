CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, password TEXT, email TEXT, nome TEXT, receivealert INTEGER, isAdmin INTEGER);
CREATE TABLE IF NOT EXISTS servant (id INTEGER PRIMARY KEY AUTOINCREMENT, nome TEXT, cargo TEXT, orgao TEXT, salario REAL, sentalert INTEGER, UNIQUE(nome, cargo, orgao) ON CONFLICT REPLACE);
CREATE TABLE IF NOT EXISTS alert (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, useremail TEXT, clientname TEXT, clientsalary REAL, sendDate TEXT);
CREATE TABLE IF NOT EXISTS client (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, isPotentialClient INTEGER, UNIQUE(name) ON CONFLICT REPLACE);

INSERT INTO user VALUES(1, 'admin', 'd42f0e5b5ae41ad2d552008ba647fbff63f66b18', 'admin@admin.com', "Admin", 1, 1);