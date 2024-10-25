# Milestones

## Milestone 0: Análisis del problema y primer acercamiento

Analizar las historias de usuario. Decidir las tecnologías que serán necesarias para comenzar a tratar el problema. Con respecto a la historia de usuario 1, definir las entidades y objetos comprendidos en ella.

+ **Mínimo viable**

- Se ha escogido un lenguaje y una estructura de directorios adecuados para afrontar el proyecto.
- Existe una estructura para cada entidad básica del problema (una clase, un objeto, una variable, una constante, un struct...).

## Milestone 1: Creación horario inicial

Asignación de los horarios de cada enfermero de forma básica, con tal de emplearlo como base para la lógica a aplicar sobre el mismo. En este caso, se toma nota de las historias de usuario 1 y 4 en las que se busca consultar esta información.

+ **Mínimo viable**

- Se emplea la información obtenida en el readme para determinar las reglas (variables, constantes...).
- Se define un algoritmo de asignación equilibrada de turnos, dado un número de trabajadores.
- El algoritmo cumple todos los aspectos de los contratos con cierta flexibilidad (pueden sobrepasar las horas anuales con tal de cumplir los turnos necesarios).

## Milestone 2: Condicionar turnos en el horario

Asignación de los horarios de cada enfermero siguiendo las condiciones de cada contrato de una manera más estricta y adaptándolo a un caso más real. Otro acercamiento hacia las historias de usuario 1 y 4.

+ **Mínimo viable**

- Se calcula la cantidad de personal necesaria para cumplir las horas y la ocupación.
- El algoritmo creado en el milestone 1 deberá ahora adaptarse al número de personal adecuado.
