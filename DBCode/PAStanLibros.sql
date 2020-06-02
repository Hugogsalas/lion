

DELIMITER $$
create procedure PAStanEditoriales(
	in ID int,
    in IDEditorial int,
    in Numero int
)
begin
	if ID=0 and IDEditorial=0 and Numero is not null then
		select 
		Stan.ID,
		Stan.Numero,
		Editorial.ID,
        Editorial.Nombre
		from 
        lioness.Stan,
        lioness.Editorial
		where
        Stan.Numero=Numero and
        Stan.IDEditorial=Editorial.ID;
	elseif ID!=0 then
		if Numero is not null then
			select 
			Stan.ID,
			Stan.Numero,
			Editorial.ID,
			Editorial.Nombre
			from 
			lioness.Stan,
			lioness.Editorial
			where
            Stan.ID=ID and 
			Stan.Numero=Numero and
			Stan.IDEditorial=Editorial.ID;
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
            Stan.ID=ID and 
			Stan.IDEditorial=Editorial.ID;
		end if;
    else 
		if Numero is not null then
			select 
			Stan.ID,
			Stan.Numero,
			Editorial.ID,
			Editorial.Nombre
			from 
			lioness.Stan,
			lioness.Editorial
			where
            Stan.IDEditorial=IDEditorial and 
			Stan.Numero=Numero and
			Stan.IDEditorial=Editorial.ID;
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
            Stan.IDEditorial=IDEditorial and 
			Stan.IDEditorial=Editorial.ID;
		end if;
	end if;
END$$
DELIMITER ;