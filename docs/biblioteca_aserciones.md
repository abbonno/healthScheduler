# BIBLIOTECA DE ASERCIONES

## CRITERIOS DE SELECCIÓN OBJETIVOS

Se emplearán los siguientes criterios en la búsqueda y selección de una de las bibliotecas de aserciones empleadas en nuestro proyecto.

- Mantenimiento: se actualiza frecuentemente. Esto se puede observar en los datos observados en el apartado "Insights" de los repositorios en github. También se empleará la plataforma ["snyk Advisor"](https://snyk.io/advisor/golang) que mide el grado de "health" o salud de una herramienta dado su estado de mantenimiento.
- Forma parte de la librería estándar: en caso de que el anterior factor no sea claramente decisivo, se optará por elegir de entre las herramientas la que forme parte de la librería estándar del lenguaje elegido, lo cual se considera buena práctica.

## OPCIONES

### [Gomega](https://github.com/onsi/gomega)

Gomega se apoya en Ginkgo para su uso, por lo que ambos proyectos se mantienen a la par (se actualiza frecuentemente). El número de commits realizado a su repositorio este mes es de 61, lo cual apoya lo anteriormente mencionado. Se coloca en la última revisión con una alta puntuación "health" de 96 sobre 100.

### [Testify](https://github.com/stretchr/testify)

Es actualizado con mucha constancia desde 2014. En el último mes se pueden observar 74 commits. No encontramos entradas sobre el proyecto original testify en snyk Advisor (existen entradas de baja puntuación sobre algunos forks de sus colaboradores, pero no sirven para juzgar el proyecto completo).

### [Paquete gotest.tools/assert](https://github.com/gotestyourself/gotest.tools)

No se actualiza a menudo. Se han realizado 23 commits este mes, lo cual es bastante inferior en comparación con el resto de herramientas. No existe una entrada para puntuar la herramienta en snyk Advisor, por lo que no se puede juzgar mediante este parámetro.

### [Go testing package](https://pkg.go.dev/testing)

Se trata del paquete de testing de la biblioteca estándar de golang por lo que es preferente.

## SELECCIÓN FINAL

Siguiendo los criterios se optará por emplear el paquete de la biblioteca estándar de golang: **Go testing package**. La propia herramienta de testing Golang ofrece la posibilidad de realizar validaciones, suficientes para realizar la función de las aserciones.
