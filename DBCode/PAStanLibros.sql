DELIMITER $$
create procedure PAStanEditoriales(
	in ID int,
    in IDEditorial int,
    in Numero int
)
begin
	if ID!=0 and IDEditorial!=0 and Numero is not null then
		select 
		Stan.ID,
		Stan.Numero,
		Editorial.ID,
        Editorial.Nombre
		from 
        lioness.Stan,
        lioness.Editorial
		where
        Editorial.ID=IDEditorial and
        Stan.Numero=Numero and
        Stan.IDEditorial=IDEditorial and
        Stan.ID=ID;
	elseif ID!=0 and IDEditorial!=0 then
		select 
		Stan.ID,
		Stan.Numero,
		Editorial.ID,
		Editorial.Nombre
		from 
		lioness.Stan,
		lioness.Editorial
		where
        Editorial.ID=IDEditorial and
		Stan.ID=ID and 
		Stan.IDEditorial=IDEditorial;
    elseif ID!=0 and Numero is not null then
		select 
		Stan.ID,
		Stan.Numero,
		Editorial.ID,
        Editorial.Nombre
		from 
        lioness.Stan,
        lioness.Editorial
		where
        Editorial.ID=Stan.IDEditorial and
        Stan.Numero=Numero and
        Stan.ID=ID;
	elseif IDEditorial!=0 and Numero is not null then
		select 
		Stan.ID,
		Stan.Numero,
		Editorial.ID,
        Editorial.Nombre
		from 
        lioness.Stan,
        lioness.Editorial
		where
        Editorial.ID=IDEditorial and
        Stan.Numero=Numero and
        Stan.IDEditorial=IDEditorial;
	else
		select 
		Stan.ID,
		Stan.Numero,
		Editorial.ID,
        Editorial.Nombre
		from 
        lioness.Stan,
        lioness.Editorial
		where
		Editorial.ID=Stan.IDEditorial;
	end if;
END$$
DELIMITER ;