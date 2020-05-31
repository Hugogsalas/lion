DELIMITER $$
create procedure PAItinerarioTaller(
	in IDItinerario int,
    in IDTaller int,
    in Horario time 
)
begin
	if IDItinerario!=0 and IDTaller!=0 then
		if Horario is not null then
			select
			Itinerario.ID,
			Itinerario.Dia,
			Taller.ID,
			Taller.Nombre,
			Taller.Duracion,
			Taller.Enfoque,
			TiposTalleres.Descripcion,
            ItinerarioTaller.Horario
			from 
			lioness.Itinerario,
			lioness.Taller,
			lioness.TiposTalleres,
			lioness.ItinerarioTaller
			where
			Itinerario.ID=IDItinerario and 
			Taller.ID=IDTaller and 
			TiposTalleres.ID=Taller.IDTipo and
			ItinerarioTaller.IDItinerario=IDItinerario and 
			ItinerarioTaller.IDTaller=IDTaller and
            ItinerarioTaller.Horario=Horario;
        else
			select 
			Itinerario.ID,
			Itinerario.Dia,
			Taller.ID,
			Taller.Nombre,
			Taller.Duracion,
			Taller.Enfoque,
			TiposTalleres.Descripcion,
            ItinerarioTaller.Horario
			from 
			lioness.Itinerario,
			lioness.Taller,
			lioness.TiposTalleres,
			lioness.ItinerarioTaller
			where
			Itinerario.ID=IDItinerario and 
			Taller.ID=IDTaller and 
			TiposTalleres.ID=Taller.IDTipo and
			ItinerarioTaller.IDItinerario=IDItinerario and 
			ItinerarioTaller.IDTaller=IDTaller;
        end if;
	elseif IDTaller!=0 then
		if Horario is not null then
			select 
			Itinerario.ID,
			Itinerario.Dia,
			Taller.ID,
			Taller.Nombre,
			Taller.Duracion,
			Taller.Enfoque,
			TiposTalleres.Descripcion,
            ItinerarioTaller.Horario
			from 
			lioness.Itinerario,
			lioness.Taller,
			lioness.TiposTalleres,
			lioness.ItinerarioTaller
			where
			Taller.ID=IDTaller and 
			TiposTalleres.ID=Taller.IDTipo and
			ItinerarioTaller.IDTaller=Taller.ID and
			ItinerarioTaller.IDItinerario=IDItinerario and
            ItinerarioTaller.Horario=Horario;
		else 
			select 
			Itinerario.ID,
			Itinerario.Dia,
			Taller.ID,
			Taller.Nombre,
			Taller.Duracion,
			Taller.Enfoque,
			TiposTalleres.Descripcion,
            ItinerarioTaller.Horario
			from 
			lioness.Itinerario,
			lioness.Taller,
			lioness.TiposTalleres,
			lioness.ItinerarioTaller
			where
			Taller.ID=IDTaller and 
			TiposTalleres.ID=Taller.IDTipo and
			ItinerarioTaller.IDTaller=IDTaller and
			ItinerarioTaller.IDItinerario=Itinerario.ID ;
        end if;
	elseif IDItinerario!=0 then
		if Horario is not null then
			select 
			Itinerario.ID,
			Itinerario.Dia,
			Taller.ID,
			Taller.Nombre,
			Taller.Duracion,
			Taller.Enfoque,
			TiposTalleres.Descripcion,
            ItinerarioTaller.Horario
			from 
			lioness.Itinerario,
			lioness.Taller,
			lioness.TiposTalleres,
			lioness.ItinerarioTaller
			where
			Itinerario.ID=IDItinerario and 
			TiposTalleres.ID=Taller.IDTipo and
			ItinerarioTaller.IDTaller=IDTaller and
			ItinerarioTaller.IDItinerario=Itinerario.ID and
            ItinerarioTaller.Horario=Horario;
		else 
			select 
			Itinerario.ID,
			Itinerario.Dia,
			Taller.ID,
			Taller.Nombre,
			Taller.Duracion,
			Taller.Enfoque,
			TiposTalleres.Descripcion,
            ItinerarioTaller.Horario
			from 
			lioness.Itinerario,
			lioness.Taller,
			lioness.TiposTalleres,
			lioness.ItinerarioTaller
			where
			Itinerario.ID=IDItinerario and 
			TiposTalleres.ID=Taller.IDTipo and
			ItinerarioTaller.IDTaller=Taller.ID and
			ItinerarioTaller.IDItinerario=IDItinerario ;
        end if;
    else
		if Horario is not null then
			select 
			Itinerario.ID,
			Itinerario.Dia,
			Taller.ID,
			Taller.Nombre,
			Taller.Duracion,
			Taller.Enfoque,
			TiposTalleres.Descripcion,
            ItinerarioTaller.Horario
			from 
			lioness.Itinerario,
			lioness.Taller,
			lioness.TiposTalleres,
			lioness.ItinerarioTaller
			where
			TiposTalleres.ID=Taller.IDTipo and
            ItinerarioTaller.IDTaller=Taller.ID and
			ItinerarioTaller.IDItinerario=Itinerario.ID and
            ItinerarioTaller.Horario=Horario;
        else
			select 
			Itinerario.ID,
			Itinerario.Dia,
			Taller.ID,
			Taller.Nombre,
			Taller.Duracion,
			Taller.Enfoque,
			TiposTalleres.Descripcion,
            ItinerarioTaller.Horario
			from 
			lioness.Itinerario,
			lioness.Taller,
			lioness.TiposTalleres,
			lioness.ItinerarioTaller
			where
			TiposTalleres.ID=Taller.IDTipo and
            ItinerarioTaller.IDTaller=Taller.ID and
			ItinerarioTaller.IDItinerario=Itinerario.ID;
		end if;
	end if;
END$$
DELIMITER ;









