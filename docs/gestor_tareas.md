# GESTOR DE TAREAS

## CRITERIOS DE SELECCIÓN OBJETIVOS

Se emplearán los siguientes criterios en la selección de una de las herramientas exploradas para gestionar las tareas del proyecto.

- Actualizado: tendrán preferencia los gestores de dependencias que no hayan sido "abandonados" o no reciban modificaciones de forma habitual, por lo que tienen un buen mantenimiento.
- Estabilidad: de la misma manera, aunque no sean actualizados continuamente serán bien consideradas las herramientas que se encuentren en etapas de desarrollo muy estables y, por lo tanto, sin errores críticos.

## OPCIONES

### [Make](https://www.gnu.org/software/make/)

En primer lugar analizaremos quizás el más conocido y generalizado. Make es una herramienta que controla la generación de ejecutables y otros archivos no fuente de un programa a partir de un archivo llamado makefile. Este enumera cada uno de los archivos no fuente y cómo procesarlos a partir de otros. El archivo make emplea reglas específicas para construir los *target* a partir de archivos fuente (inputs) y especifica sus dependencias.

Como ventaja, make es un proyecto desarrollado desde los 70, con una madurez más que suficiente para confiar en su correcto funcionamiento en todo momento, a pesar de no tener tantas actualizaciones en la actualidad.

### [Mage](https://magefile.org/)

Como alternativa a Make, en muchos foros se recomienda Mage. Mage es un reemplazo para Make que utiliza Go puro para definir y ejecutar tareas. No requiere un archivo Makefile ni otra sintaxis externa.

Respecto a su mantenimiento, el último commit fue realizado hace dos semanas y se puede ver que el proyecto es actualizado a menudo.
Sin embargo, y aunque la popularidad de una herramienta no sea un criterio de selección válido, en este caso lo pequeño de su comunidad puede suponer un problema a largo plazo, puesto que se trata de una herramienta de código abierto. Esto puede deparar en un mantenimiento pobre.

### [Task](https://taskfile.dev/)

Desde el mismo foro se habla de Task. Este, igualmente, según su documentación, dice diferenciarse de make por su simplicidad.

Visto su github, task se actualiza constantemente desde hace, al menos, 7 años, con una comunidad bastante más activa que la de mage (aunque siga sin ser tan tradicionalmente empleada como make).

### [Ansible](https://github.com/ansible/ansible)

Ansible maneja "gestión de configuración, despliegue de aplicaciones, aprovisionamiento en la nube, ejecución de tareas, automatización en la red y orquestación de múltiples nodos" mediante archivos yaml en los que especificamos las tareas.
Es una herramienta con una gran comunidad detrás y actualizaciones constantes desde 2012.

### [Goyek](https://github.com/goyek/goyek)

Se trata de una librería de automatización de tareas en go que permite organizar flujos de trabajo mediante funciones de este lenguaje.
Como proyecto es bastante simple y con muchos menos colaboradores, por ende, con menos mantenimiento. Además tiene su origen en 2021.

## SELECCIÓN FINAL

Se ha elegido, finalmente, la herramienta **Task** por las siguientes razones.
Con respecto a **Mage**, se ha optado darle prioridad al criterio de mantenimiento porque el riesgo de que la herramienta caiga en abandono puede causar errores.
Con **Buck2** pasa al revés, que aún es muy nuevo y se debate con la estabilidad que pueda ofrecer.
Sobre **Goyek**, al ser un proyecto poco pulido y con poca esperanza en su futuro desarrollo se queda atrás del resto de herramientas.
Por el resto, entre **Make**, **Ansible** y **Task** todas son buenas opciones, por lo que se elige, a criterio propio, **Task**.
