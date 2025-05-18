const activities = [
  {
    nombre: "taekwondo",
    descripcion: "Arte marcial coreana",
    horarios: [
      { dia: 2, "hora-inicio": "18:30", "hora-fin": "20:00" },
      { dia: 4, "hora-inicio": "18:30", "hora-fin": "20:00" }
    ]
  },
  {
    nombre: "zumba",
    descripcion: "ritmos latinos",
    horarios: [
      { dia: 1, "hora-inicio": "19:30", "hora-fin": "20:30" },
      { dia: 3, "hora-inicio": "19:30", "hora-fin": "20:30" }
    ]
  }
];

const diasSemana = ["Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"];


const Activities = () => {

  const isLoggedIn = localStorage.getItem("isLoggedIn") === "true";

  return (
      <div>
          <h2>Actividades</h2>
          {activities.map((actividad, index) => (
              <div className="activity-card" key={index}>
                  <h3>{ actividad.nombre }</h3>
                  <p> { actividad.descripcion } </p>
                  <ul>
                      { actividad.horarios.map((horario, i) => (
                          <li key={i}>{ diasSemana[horario.dia] } : { horario["hora-inicio"]} - {horario["hora-fin"]}</li>
                      ))}
                      
                  </ul>
                  {isLoggedIn && (
                        <button onClick={() => alert(`Inscripto en ${actividad.nombre}`)}>
                            Inscribirme
                        </button>
                    )}
              </div>
          ))}
          
      </div>
  )

}

export default Activities;