

DELIMITER $$
create procedure PASalaExposicion(
	in IDSala int,
    in IDExposicion int
)
begin
	if IDSala!=0 and IDExposicion!=0 then
		select 
		Sala.ID,
		Sala.Nombre,
        Exposicion.ID,
		Exposicion.Presentador,
		Exposicion.Titulo,
        Exposicion.Duracion,
        TiposExposicion.Descripcion
		from 
        lioness.Sala,
        lioness.Exposicion,
        lioness.SalaExposicion,
        lioness.TiposExposicion
		where
        Sala.ID=IDSala and 
        Exposicion.ID=IDExposicion and 
        TiposExposicion.ID=Exposicion.IDTipo and
        SalaExposicion.IDSala=IDSala and 
        SalaExposicion.IDExposicion=IDExposicion;
	elseif IDSala!=0 then
		select 
		Sala.ID,
		Sala.Nombre,
        Exposicion.ID,
		Exposicion.Presentador,
		Exposicion.Titulo,
        Exposicion.Duracion,
        TiposExposicion.Descripcion
		from 
        lioness.Sala,
        lioness.Exposicion,
        lioness.SalaExposicion,
        lioness.TiposExposicion
		where
        Sala.ID=IDSala and 
        TiposExposicion.ID=Exposicion.IDTipo and
        SalaExposicion.IDSala=IDSala and 
        SalaExposicion.IDExposicion=Exposicion.ID;
    else 
		select 
		Sala.ID,
		Sala.Nombre,
        Exposicion.ID,
		Exposicion.Presentador,
		Exposicion.Titulo,
        Exposicion.Duracion,
        TiposExposicion.Descripcion
		from 
        lioness.Sala,
        lioness.Exposicion,
        lioness.SalaExposicion,
        lioness.TiposExposicion
		where
        Exposicion.ID=IDExposicion and 
        TiposExposicion.ID=Exposicion.IDTipo and
        SalaExposicion.IDSala=Sala.ID and 
        SalaExposicion.IDExposicion=IDExposicion;
	end if;
END
$$DELIMITER ;