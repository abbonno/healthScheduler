# healthScheduler

Proyecto de planificación de horarios de trabajo entre personal sanitario

## Problemática a resolver

El problema tiene lugar en un hospital. Los enfermeros tienen que cumplir ciertos horarios de trabajo divididos en los siguientes turnos de trabajo: mañana, tarde y noche; y en distintas áreas del hospital. Esta distribución es conocida de antemano, puesto que se turna siguiendo siempre el mismo orden. Sin embargo, pueden surgir ciertas circunstancias en las que esta organización sufra alteraciones: baja de un empleado, vacaciones, días de asuntos propios, cambios de turno... Es por eso que se debe encontrar la manera de generar un cronograma (schedule) que reparta óptimamente turnos teniendo en cuenta las diversas condiciones de cada contrato y repartiendo de forma equilibrada sustituciones, así como los nuevos horarios para enfermeros que reemprendan su labor.

## Factores a tener en cuenta

### Áreas

    Intermedia 1, 2, 3 y 4. Urgencias 1 y 2. Planta 1 y 2.
    En total hacen falta 8 enfermeros y 2 sustitutos al día.
    El número de empleados es variable y hay cierta flexibilidad en las horas a completar.

### Turnos

    Mañana: 8:00 - 15:00
    Tarde: 15:00 - 22:00
    Noche: 22:00 - 8:00

### Tipos de contrato

    Rotatorio con noches: 1483 horas/año.
    Rotatorio sin noches: 1514 horas/año.
    Jornada reducida (con noches/sin noches): 2/3 del horario correspondiente (de los anteriores).
    "Mochilas": sustituyen el 1/3 de horas que no hacen los de jornada reducida.
    Sustitutos: contrato variable.

### Días libres (variables según antigüedad)

    Asuntos propios: 6 días.
    Vacaciones: 1 mes (30-31 días) o 2 quincenas o por días sueltos (22 días).
    *Bajas: 12 meses + 6 meses.

*No computan como días libres pero es necesario tenerlos en cuenta

## Documentos

[Ir a la carpeta docs](./docs/)
Configuración

- [Configuración SSH](.\docs\configuracion\claveSSH.png)
- [Configuración perfil 1](.\docs\configuracion\conf1.png)
- [Configuración perfil 2](.\docs\configuracion\conf2.png)

Documentación

- ["User journeys"](.\docs\documentacion\user_journeys.md)
- [Historias de Usuario](.\docs\documentacion\HU.md)
- [Milestones](.\docs\documentacion\milestones.md)
- [Gestor Dependencias](./docs/gestor_dependencias.md).
- [Gestor Tareas](./docs/gestor_tareas.md).
- [Biblioteca de aserciones](.\docs\biblioteca_aserciones.md)
- [Test runner](.\docs\test_runner.md)
  
Tareas

- "task check": comprueba la sintaxis de los fuentes existentes
- "task test": ejecuta los tests (aún por crear)
