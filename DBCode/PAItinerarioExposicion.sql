
DELIMITER $$
create procedure PAItinerarioExposicion(
	in IDItinerario int,
    in IDExposicion int,
    in Horario time 
)
begin
	if IDItinerario!=0 and IDExposicion!=0 then
		if Horario is not null then
			select
			Itinerario.ID,
			Itinerario.Dia,
			Exposicion.ID,
			Exposicion.Presentador,
			Exposicion.Duracion,
			Exposicion.Titulo,
			TiposExposicion.Descripcion,
            ItinerarioExposicion.Horario
			from 
			lioness.Itinerario,
			lioness.Exposicion,
			lioness.TiposExposicion,
			lioness.ItinerarioExposicion
			where
			Itinerario.ID=IDItinerario and 
			Exposicion.ID=IDExposicion and 
			TiposExposicion.ID=Exposicion.IDTipo and
			ItinerarioExposicion.IDItinerario=IDItinerario and 
			ItinerarioExposicion.IDExposicion=IDExposicion and
            ItinerarioExposicion.Horario=Horario;
        else
			select 
			Itinerario.ID,
			Itinerario.Dia,
			Exposicion.ID,
			Exposicion.Presentador,
			Exposicion.Duracion,
			Exposicion.Titulo,
			TiposExposicion.Descripcion,
            ItinerarioExposicion.Horario
			from 
			lioness.Itinerario,
			lioness.Exposicion,
			lioness.TiposExposicion,
			lioness.ItinerarioExposicion
			where
			Itinerario.ID=IDItinerario and 
			Exposicion.ID=IDExposicion and 
			TiposExposicion.ID=Exposicion.IDTipo and
			ItinerarioExposicion.IDItinerario=IDItinerario and 
			ItinerarioExposicion.IDExposicion=IDExposicion;
        end if;
	elseif IDExposicion!=0 then
		if Horario is not null then
			select 
			Itinerario.ID,
			Itinerario.Dia,
			Exposicion.ID,
			Exposicion.Presentador,
			Exposicion.Duracion,
			Exposicion.Titulo,
			TiposExposicion.Descripcion,
            ItinerarioExposicion.Horario
			from 
			lioness.Itinerario,
			lioness.Exposicion,
			lioness.TiposExposicion,
			lioness.ItinerarioExposicion
			where
			Exposicion.ID=IDExposicion and 
			TiposExposicion.ID=Exposicion.IDTipo and
			ItinerarioExposicion.IDExposicion=Exposicion.ID and
			ItinerarioExposicion.IDItinerario=IDItinerario and
            ItinerarioExposicion.Horario=Horario;
		else 
			select 
			Itinerario.ID,
			Itinerario.Dia,
			Exposicion.ID,
			Exposicion.Presentador,
			Exposicion.Duracion,
			Exposicion.Titulo,
			TiposExposicion.Descripcion,
            ItinerarioExposicion.Horario
			from 
			lioness.Itinerario,
			lioness.Exposicion,
			lioness.TiposExposicion,
			lioness.ItinerarioExposicion
			where
			Exposicion.ID=IDExposicion and 
			TiposExposicion.ID=Exposicion.IDTipo and
			ItinerarioExposicion.IDExposicion=IDExposicion and
			ItinerarioExposicion.IDItinerario=Itinerario.ID ;
        end if;
	elseif IDItinerario!=0 then
		if Horario is not null then
			select 
			Itinerario.ID,
			Itinerario.Dia,
			Exposicion.ID,
			Exposicion.Presentador,
			Exposicion.Duracion,
			Exposicion.Titulo,
			TiposExposicion.Descripcion,
            ItinerarioExposicion.Horario
			from 
			lioness.Itinerario,
			lioness.Exposicion,
			lioness.TiposExposicion,
			lioness.ItinerarioExposicion
			where
			Itinerario.ID=IDItinerario and 
			TiposExposicion.ID=Exposicion.IDTipo and
			ItinerarioExposicion.IDExposicion=IDExposicion and
			ItinerarioExposicion.IDItinerario=Itinerario.ID and
            ItinerarioExposicion.Horario=Horario;
		else 
			select 
			Itinerario.ID,
			Itinerario.Dia,
			Exposicion.ID,
			Exposicion.Presentador,
			Exposicion.Duracion,
			Exposicion.Titulo,
			TiposExposicion.Descripcion,
            ItinerarioExposicion.Horario
			from 
			lioness.Itinerario,
			lioness.Exposicion,
			lioness.TiposExposicion,
			lioness.ItinerarioExposicion
			where
			Itinerario.ID=IDItinerario and 
			TiposExposicion.ID=Exposicion.IDTipo and
			ItinerarioExposicion.IDExposicion=Exposicion.ID and
			ItinerarioExposicion.IDItinerario=IDItinerario ;
        end if;
    else
		if Horario is not null then
			select 
			Itinerario.ID,
			Itinerario.Dia,
			Exposicion.ID,
			Exposicion.Presentador,
			Exposicion.Duracion,
			Exposicion.Titulo,
			TiposExposicion.Descripcion,
            ItinerarioExposicion.Horario
			from 
			lioness.Itinerario,
			lioness.Exposicion,
			lioness.TiposExposicion,
			lioness.ItinerarioExposicion
			where
			TiposExposicion.ID=Exposicion.IDTipo and
            ItinerarioExposicion.IDExposicion=Exposicion.ID and
			ItinerarioExposicion.IDItinerario=Itinerario.ID and
            ItinerarioExposicion.Horario=Horario;
        else
			select 
			Itinerario.ID,
			Itinerario.Dia,
			Exposicion.ID,
			Exposicion.Presentador,
			Exposicion.Duracion,
			Exposicion.Titulo,
			TiposExposicion.Descripcion,
            ItinerarioExposicion.Horario
			from 
			lioness.Itinerario,
			lioness.Exposicion,
			lioness.TiposExposicion,
			lioness.ItinerarioExposicion
			where
			TiposExposicion.ID=Exposicion.IDTipo and
            ItinerarioExposicion.IDExposicion=Exposicion.ID and
			ItinerarioExposicion.IDItinerario=Itinerario.ID;
		end if;
	end if;
END$$
DELIMITER ;