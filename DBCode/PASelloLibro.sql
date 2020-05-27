

DELIMITER $$
create procedure PASelloLibro(
	in IDSello int,
    in IDLibro int
)
begin
	if IDSello!=0 and IDLibro!=0 then
		select 
		sello.ID,
		Libro.ID,
		sello.Descripcion,
		Libro.Titulo,
		Libro.precio 
		from lioness.sello,lioness.libro,lioness.sellolibro
		where
        sello.ID=IDSello and 
        Libro.ID=IDLibro and 
        sellolibro.IDSello=IDSello and 
        sellolibro.IDLibro=IDLibro;
	elseif IDSello!=0 then
		select 
		sello.ID,
		Libro.ID,
        sello.Descripcion,
		Libro.Titulo,
		Libro.precio 
		from lioness.sello,lioness.libro,lioness.sellolibro
		where 
        Autor.ID=IDSello and
        sellolibro.IDLibro=Libro.ID and
        AutorLibro.IDSello=IDSello;
    else 
		select 
		sello.ID,
		Libro.ID,
		sello.Descripcion,
		Libro.Titulo,
		Libro.precio 
		from lioness.Sello,lioness.libro,lioness.sellolibro
		where 
        Libro.ID=IDLibro and
        sellolibro.IDSello=sello.ID and
        sellolibro.IDLibro=IDLibro;
	end if;
END$$
DELIMITER ;