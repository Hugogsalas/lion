DELIMITER $$
create procedure PAEditorialLibro(
	in IDEditorial int,
    in IDLibro int
)
begin
	if IDEditorial!=0 and IDLibro!=0 then
		select 
		editorial.Nombre,
		Libro.Titulo,
		Libro.precio 
		from lioness.editorial,lioness.libro,lioness.editoriallibro
		where
        editorial.ID=IDEditorial and 
        Libro.ID=IDLibro and 
        editoriallibro.IDEditorial=IDEditorial and 
        editoriallibro.IDLibro=IDLibro;
	elseif IDEditorial!=0 then
		select 
        editorial.Nombre,
		Libro.Titulo,
		Libro.precio 
		from lioness.editorial,lioness.libro,lioness.editoriallibro
		where 
        editorial.ID=IDEditorial and
        editoriallibro.IDLibro=Libro.ID and
        editoriallibro.IDEditorial=IDEditorial;
    else 
		select 
		editorial.Nombre,
		Libro.Titulo,
		Libro.precio 
		from lioness.editorial,lioness.libro,lioness.editoriallibro
		where 
        Libro.ID=IDLibro and
        editoriallibro.IDEditorial=editorial.ID and
        editoriallibro.IDLibro=IDLibro;
	end if;
END$$
DELIMITER ;