# TEST RUNNER

## CRITERIOS DE SELECCIÓN OBJETIVOS

Se emplearán los siguientes criterios en la búsqueda y selección de un test runner adecuado para nuestro proyecto.

- Mantenimiento: se trata de una herramienta que se actualiza a menudo respecto a errores o nuevas necesidades. También se empleará la plataforma ["snyk Advisor"](https://snyk.io/advisor/golang) que mide el grado de "health" o salud de una herramienta dado su estado de mantenimiento.
- Forma parte de la librería estándar: en caso de que el anterior factor no sea claramente decisivo, se optará por elegir de entre las herramientas la que forme parte de la librería estándar del lenguaje elegido, lo cual se considera buena práctica.

## OPCIONES

### [Ginkgo](https://github.com/onsi/ginkgo)

En su repositorio es común observar commits y modificaciones por lo que se encuentra en mantenimiento constante. Se data de 109 commits el último mes, como prueba de ello. No existe información en snyk Advisor, por lo que no se puede juzgar mediante este parámetro.

### [Goconvey](https://github.com/smartystreets/goconvey)

No se encuentra en mantenimiento y los commits realizados son prácticamente nulos en los últimos siete años (6 el último mes). Las entradas en snyk Advisor, correspondientemente, le dan una baja puntiación de health (entre 40 y 50 sobre 100).

### [Goblin](https://github.com/franela/goblin)

Igualmente, no se ha seguido su mantenimiento en los últimos ocho años. Sus puntos de salud en snyk Advisor rondan los 40-60 sobre 100.

### [Go testing package](https://pkg.go.dev/testing)

Podemos valorar en el [historial de la carpeta de la herramienta en el repositorio de github de golang](https://github.com/golang/go/commits/master/src/testing) que se realizan actualizaciones constantemente. Se trata del paquete de testing de la biblioteca estándar de golang por lo que es preferente.

## SELECCIÓN FINAL

Siguiendo los criterios elegidos anteriormente, escogeremos usar el paquete de la biblioteca estándar de golang: **Go testing package**.

## CLI

Al haber elegido **Go testing package**, por ende, estaremos eligiendo emplear como CLI la suborden "test", incluida en el mismo paquete.
