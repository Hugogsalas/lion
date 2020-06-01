

DELIMITER $$
create procedure PASalaTaller(
	in IDSala int,
    in IDTaller int
)
begin
	if IDSala!=0 and IDTaller!=0 then
		select 
		Sala.ID,
        Sala.Nombre,
        Taller.ID,
		Taller.Nombre,
		Taller.Enfoque,
        Taller.Duracion,
        TiposTalleres.Descripcion
		from 
        lioness.Taller,
        lioness.Sala,
        lioness.SalaTaller,
        lioness.TiposTalleres
		where
        Sala.ID=IDSala and 
        Taller.ID=IDTaller and  
        SalaTaller.IDSala=IDSala and 
        SalaTaller.IDTaller=IDTaller and
        TiposTalleres.ID=Taller.IDTipo;
	elseif IDSala!=0 then
		select 
		Sala.ID,
        Sala.Nombre,
        Taller.ID,
		Taller.Nombre,
		Taller.Enfoque,
        Taller.Duracion,
        TiposTalleres.Descripcion
		from 
        lioness.Taller,
        lioness.Sala,
        lioness.SalaTaller,
        lioness.TiposTalleres
		where
        Sala.ID=IDSala and 
        SalaTaller.IDSala=IDSala and 
        SalaTaller.IDTaller=Taller.ID and
        TiposTalleres.ID=Taller.IDTipo;
    else 
		select 
		Sala.ID,
        Sala.Nombre,
        Taller.ID,
		Taller.Nombre,
		Taller.Enfoque,
        Taller.Duracion,
        TiposTalleres.Descripcion
		from 
        lioness.Taller,
        lioness.Sala,
        lioness.SalaTaller,
        lioness.TiposTalleres
		where
        Taller.ID=IDTaller and  
        SalaTaller.IDSala=Sala.ID and 
        SalaTaller.IDTaller=IDTaller and
        TiposTalleres.ID=Taller.IDTipo;
	end if;
END$$
DELIMITER ;