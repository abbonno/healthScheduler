# TEST RUNNER

## CRITERIOS DE SELECCIÓN OBJETIVOS

Se emplearán los siguientes criterios en la búsqueda y selección de un test runner adecuado para nuestro proyecto.

- Mantenimiento: se trata de una herramienta que se actualiza a menudo respecto a errores o nuevas necesidades. También se empleará la plataforma ["snyk Advisor"](https://snyk.io/advisor/golang) que mide el grado de "health" o salud de una herramienta dado su estado de mantenimiento.
- Herramienta propia del lenguaje: en caso de que el anterior factor no sea claramente decisivo, se optará por elegir de entre las herramientas la que sea propia del lenguaje.

## OPCIONES

### [Ginkgo](https://github.com/onsi/ginkgo)

En su repositorio es común observar commits y modificaciones por lo que se encuentra en mantenimiento constante. Se data de 109 commits el último mes, como prueba de ello. No existe información en snyk Advisor, por lo que no se puede juzgar mediante este parámetro.

### [Goconvey](https://github.com/smartystreets/goconvey)

No se encuentra en mantenimiento y los commits realizados son prácticamente nulos en los últimos siete años (6 el último mes). Las entradas en snyk Advisor, correspondientemente, le dan una baja puntiación de health (entre 40 y 50 sobre 100).

### [Goblin](https://github.com/franela/goblin)

Igualmente, no se ha seguido su mantenimiento en los últimos ocho años. Sus puntos de salud en snyk Advisor rondan los 40-60 sobre 100.

### [Go testing package](https://pkg.go.dev/testing)

Se trata del paquete de testing de la biblioteca estándar de golang por lo que es preferente.

## SELECCIÓN FINAL

Siguiendo los criterios elegidos anteriormente, escogeremos usar el paquete de la biblioteca estándar de golang: **Go testing package**.
