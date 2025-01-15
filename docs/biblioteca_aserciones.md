# BIBLIOTECA DE ASERCIONES

## CRITERIOS DE SELECCIÓN OBJETIVOS

Se emplearán los siguientes criterios en la búsqueda y selección de una de las bibliotecas de aserciones empleadas en nuestro proyecto.

- Mantenimiento: se actualiza frecuentemente. Esto se puede observar en los datos observados en el apartado "Insights" de los repositorios en github. También se empleará la plataforma ["snyk Advisor"](https://snyk.io/advisor/golang) que mide el grado de "health" o salud de una herramienta dado su estado de mantenimiento.
- Estándar: en caso de que el anterior factor no sea claramente decisivo, se optará por elegir la herramienta que cumple el estándar del lenguaje elegido.

## OPCIONES

### [Gomega](https://github.com/onsi/gomega)

Gomega se apoya en Ginkgo para su uso, por lo que ambos proyectos se mantienen a la par (se actualiza frecuentemente). El número de commits realizado a su repositorio este mes es de 61, lo cual apoya lo anteriormente mencionado. Se coloca en la última revisión con una alta puntuación "health" de 96 sobre 100.

### [Testify](https://github.com/stretchr/testify)

Es actualizado con mucha constancia desde 2014. En el último mes se pueden observar 74 commits. No encontramos entradas sobre el proyecto original testify en snyk Advisor (existen entradas de baja puntuación sobre algunos forks de sus colaboradores, pero no sirven para juzgar el proyecto completo).

### [Paquete gotest.tools/assert](https://github.com/gotestyourself/gotest.tools)

Se trata de una librería moderna de aserciones entre las herramientas de gotest.tools, pero tiene un mantenimiento un poco inestable. Prueba de ello es que solo se han realizado 23 commits este mes en un intervalo de 3 días seguidos. No existe una entrada para puntuar la herramienta en snyk Advisor, por lo que no se puede juzgar mediante este parámetro.

## SELECCIÓN FINAL

Siguiendo los criterios se optará por emplear el estándar de Golang. La propia herramienta de testing estándar de Golang ofrece la posibilidad de realizar validaciones, suficientes para realizar la función de las aserciones.
