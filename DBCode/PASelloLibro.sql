DELIMITER $$
create procedure PASelloLibro(
	in IDSello int,
    in IDLibro int
)
begin
	if IDSello!=0 and IDLibro!=0 then
		select 
		sello.ID,
		sello.Descripcion,
		Editorial.Nombre,
		Libro.ID,
		Libro.Titulo,
		Libro.precio 
		from 
        lioness.sello,
        lioness.libro,
        lioness.sellolibro,
		lioness.Editorial
		where
		Editorial.ID=sello.IDEditorial and
        sello.ID=IDSello and 
        Libro.ID=IDLibro and 
        sellolibro.IDSello=IDSello and 
        sellolibro.IDLibro=IDLibro;
	elseif IDSello!=0 then
		select 
		sello.ID,
        sello.Descripcion,
		Editorial.Nombre,
		Libro.ID,
		Libro.Titulo,
		Libro.precio 
		from 
        lioness.sello,
        lioness.libro,
        lioness.sellolibro,
		lioness.Editorial
		where 
		Editorial.ID=sello.IDEditorial and
        sello.ID=IDSello and
        sellolibro.IDLibro=Libro.ID and
        sellolibro.IDSello=IDSello;
    elseif IDLibro!=0 then
		select 
		sello.ID,
		sello.Descripcion,
		Editorial.Nombre,
		Libro.ID,
		Libro.Titulo,
		Libro.precio 
		from 
        lioness.Sello,
        lioness.libro,
		lioness.Editorial,
        lioness.sellolibro
		where 
		Editorial.ID=sello.IDEditorial and
        Libro.ID=IDLibro and
        sellolibro.IDSello=sello.ID and
        sellolibro.IDLibro=IDLibro;
	else 
		select 
		sello.ID,
		sello.Descripcion,
		Editorial.Nombre,
		Libro.ID,
		Libro.Titulo,
		Libro.precio 
		from 
        lioness.Sello,
        lioness.libro,
        lioness.sellolibro,
		lioness.Editorial
		where 
		Editorial.ID=sello.IDEditorial and
        sellolibro.IDSello=sello.ID and
        sellolibro.IDLibro=Libro.ID;
	end if;
END$$
DELIMITER ;