CREATE TABLE USER(
Id_user int not NULL AUTO_INCREMENT PRIMARY KEY,
Nama_user VARCHAR(100) not NULL,
email VARCHAR(100) not NULL,
phone VARCHAR(15) not NULL,
alamat TEXT,
foto_profil TEXT,
kata_sandi VARCHAR(20),
create_at datetime DEFAULT current_timestamp(),
update_at datetime DEFAULT current_timestamp()


);

CREATE TABLE TopUp(
Id_TopUp int not NULL AUTO_INCREMENT PRIMARY KEY,
Jumlah_Topup decimal,
users_Id int,
Stat ENUM("BERHASIL","GAGAL"),
create_at datetime DEFAULT current_timestamp(),
update_at datetime on update current_timestamp(),

CONSTRAINT FK_topupUsers FOREIGN KEY(users_id) REFERENCES USER(Id_user)
);

CREATE TABLE Transfers(
Id_TF int AUTO_INCREMENT PRIMARY KEY,
pengirim_id int,
phone VARCHAR(15) not NULL,
Jumlah_TF decimal,
penerima_id int,  
create_at datetime DEFAULT current_timestamp(),
update_at datetime on update current_timestamp(),

CONSTRAINT FK_Tfpenerima FOREIGN KEY(penerima_id) REFERENCES Penerima(id_user),
CONSTRAINT FK_Tfpengirim FOREIGN KEY(pengirim_id) REFERENCES Pengirim(id_user)
);




CREATE TABLE Saldo(
id_users int AUTO_INCREMENT PRIMARY KEY,
Jumlah_saldo decimal,
create_at datetime DEFAULT current_timestamp(),
update_at datetime on update current_timestamp(),

CONSTRAINT FK_saldo FOREIGN KEY(id_users) REFERENCES USER(Id_user)
);



