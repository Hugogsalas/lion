
DELIMITER $$
create procedure PAAutorLibro(
	in IDAutor int,
    in IDLibro int
)
begin
	if IDAutor!=0 and IDLibro!=0 then
		select 
		Autor.Nombre,
		Autor.ApellidoPaterno,
		Autor.ApellidoMaterno,
		Libro.Titulo,
		Libro.precio 
		from lioness.Autor,lioness.libro,lioness.AutorLibro
		where
        Autor.ID=IDAutor and 
        Libro.ID=IDLibro and 
        AutorLibro.IDAutor=IDAutor and 
        AutorLibro.IDLibro=IDLibro;
	elseif IDAutor!=0 then
		select 
		Autor.Nombre,
		Autor.ApellidoPaterno,
		Autor.ApellidoMaterno,
		Libro.Titulo,
		Libro.precio 
		from lioness.Autor,lioness.libro,lioness.AutorLibro
		where 
        Autor.ID=IDAutor and
        AutorLibro.IDLibro=Libro.ID and
        AutorLibro.IDAutor=IDAutor;
    else 
		select 
		Autor.Nombre,
		Autor.ApellidoPaterno,
		Autor.ApellidoMaterno,
		Libro.Titulo,
		Libro.precio 
		from lioness.Autor,lioness.libro,lioness.AutorLibro
		where 
        Libro.ID=IDLibro and
        AutorLibro.IDAutor=Autor.ID and
        AutorLibro.IDLibro=IDLibro;
	end if;
END$$
DELIMITER ;