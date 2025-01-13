# TEST RUNNER

## CRITERIOS DE SELECCIÓN OBJETIVOS

Se emplearán los siguientes criterios en la búsqueda y selección de un test runner adecuado para nuestro proyecto.

- Mantenimiento: se trata de una herramienta que se actualiza a menudo respecto a errores o nuevas necesidades.
- Estándar: en caso de que el anterior factor no sea claramente decisivo, se optará por elegir la herramienta que cumple el estándar del lenguaje elegido.

## OPCIONES

### [Ginkgo](https://github.com/onsi/ginkgo)

En su repositorio es común observar commits y modificaciones por lo que se encuentra en mantenimiento constante.

### [Goconvey](https://github.com/smartystreets/goconvey)

No se encuentra en mantenimiento y los commits realizados son prácticamente nulos en los últimos siete años.

### [Goblin](https://github.com/franela/goblin)

Igualmente, no se ha seguido su mantenimiento en los últimos ocho años.

### [Go testing package](https://pkg.go.dev/testing)

Se trata del paquete estándar de golang. Tiene mantenimiento constante.

## SELECCIÓN FINAL

Siguiendo los criterios elegidos anteriormente, escogeremos usar el paquete estándar de go.
