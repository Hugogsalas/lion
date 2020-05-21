create database lioness;

use lioness;

create table Usuarios(
    ID int primary key auto_increment not null,
    Correo varchar(30) not null,
    Nombre varchar(15) not null,
    ApellidoPaterno varchar(15) not null,
    ApellidoMaterno varchar(15) not null,
    Sexo varchar(1) not null,
    Clave varchar(200) not null
);

create table Autor(
    ID int primary key auto_increment not null,
    Nombre varchar(30) not null,
    ApellidoPaterno varchar(30) not null,
    ApellidoMaterno varchar(30) not null
);

create table Editorial(
    ID int primary key auto_increment not null,
    Nombre varchar (50) not null
);

create table TiposExposicion(
    ID int primary key auto_increment not null,
    Descripcion varchar (50) not null
);

create table Exposicion(
    ID int primary key auto_increment not null,
    Titulo varchar(50) not null,
    Presentador varchar(20) not null,
    Duracion int not null,
    IDTipo int not null,
    FOREIGN KEY (IDTipo) REFERENCES TiposExposicion(ID) ON DELETE CASCADE
);

create table Itinerario(
    ID int primary key auto_increment not null,
    Dia Date not null
);

create table Sello(
    ID int primary key auto_increment not null,
    Descripcion varchar (50) not null
);

create table Libro(
    ID int primary key auto_increment not null,
    Titulo varchar(50) not null,
    IDAutor int not null,
    Precio float not null,
    FOREIGN KEY (IDAutor) REFERENCES Autor(ID) ON DELETE CASCADE
);

create table Sala(
    ID int primary key auto_increment not null,
    Nombre varchar(20) not null
);

create table Stan(
    ID int key auto_increment not null,
    IDEditorial int not null,
    Numero int not null,
    FOREIGN KEY (IDEditorial) REFERENCES Editorial(ID) ON DELETE CASCADE
);

create table TiposTalleres(
    ID int primary key auto_increment not null,
    Descripcion varchar (50) not null
);

create table Taller(
    ID int primary key auto_increment not null,
    Nombre varchar(20) not null,
    Enfoque varchar(20) not null,
    IDTipo int not null,
    FOREIGN KEY (IDTipo) REFERENCES TiposTalleres(ID) ON DELETE CASCADE
);

create table AutorLibro(
    primary key (IDAutor, IDLibro),
    IDAutor int not null,
    IDLibro int not null,
    FOREIGN KEY (IDAutor) REFERENCES Autor(ID) ON DELETE CASCADE,
    FOREIGN KEY (IDLibro) REFERENCES Libro(ID) ON DELETE CASCADE
);

create table EditorialLibro (
    primary key (IDEditorial, IDLibro),
    IDEditorial int not null,
    IDLibro int not null,
    FOREIGN KEY (IDEditorial) REFERENCES Editorial(ID) ON DELETE CASCADE,
    FOREIGN KEY (IDLibro) REFERENCES Libro(ID) ON DELETE CASCADE
);

create table ItinerarioExposicion (
    primary key (IDItenerario, IDExposicion),
    IDItenerario int not null,
    IDExposicion int not null,
    Horario Time not null,
    FOREIGN KEY (IDItenerario) REFERENCES Itinerario(ID) ON DELETE CASCADE,
    FOREIGN KEY (IDExposicion) REFERENCES Exposicion(ID) ON DELETE CASCADE
);

create table ItinerarioTaller (
    primary key (IDItenerario, IDTaller),
    IDItenerario int not null,
    IDTaller int not null,
    Horario Time not null,
    FOREIGN KEY (IDItenerario) REFERENCES Itinerario(ID) ON DELETE CASCADE,
    FOREIGN KEY (IDTaller) REFERENCES Taller(ID) ON DELETE CASCADE
);

create table SalaExposicion(
    primary key (IDSala, IDExposicion),
    IDSala int not null,
    IDExposicion int not null,
    FOREIGN KEY (IDSala) REFERENCES Sala(ID) ON DELETE CASCADE,
    FOREIGN KEY (IDExposicion) REFERENCES Exposicion(ID) ON DELETE CASCADE
);

create table SalaTaller(
    primary key (IDSala, IDTaller),
    IDSala int not null,
    IDTaller int not null,
    FOREIGN KEY (IDSala) REFERENCES Sala(ID) ON DELETE CASCADE,
    FOREIGN KEY (IDTaller) REFERENCES Taller(ID) ON DELETE CASCADE
);

create table SelloLibro(
    primary key (IDLibro, IDSello),
    IDLibro int not null,
    IDSello int not null,
    FOREIGN KEY (IDLibro) REFERENCES Libro(ID) ON DELETE CASCADE,
    FOREIGN KEY (IDSello) REFERENCES Sello(ID) ON DELETE CASCADE
);
