# BIBLIOTECA DE ASERCIONES

## CRITERIOS DE SELECCIÓN OBJETIVOS

Se emplearán los siguientes criterios en la búsqueda y selección de una de las bibliotecas de aserciones empleadas en nuestro proyecto.

- Mantenimiento: se actualiza frecuentemente.
- Estándar: en caso de que el anterior factor no sea claramente decisivo, se optará por elegir la herramienta que cumple el estándar del lenguaje elegido.

## OPCIONES

### [Gomega](https://github.com/onsi/gomega)

Gomega se apoya en Ginkgo para su uso, por lo que ambos proyectos se mantienen a la par (se actualiza frecuentemente).

### [Testify](https://github.com/stretchr/testify)

Es actualizado con mucha constancia desde 2014.

### [Paquete gotest.tools/assert](https://github.com/gotestyourself/gotest.tools)

Se trata de una librería moderna de aserciones entre las herramientas de gotest.tools, pero tiene un mantenimiento un poco inestable.

### Estándar (con el paquete testing)

La propia herramienta de testing estándar de Golang ofrece la posibilidad de realizar validaciones.

## SELECCIÓN FINAL

Siguiendo los criterios se optará por emplear las aserciones que ofrece el estándar de Golang.
