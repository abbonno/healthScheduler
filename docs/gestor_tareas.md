# GESTOR DE TAREAS

## CRITERIOS DE SELECCIÓN OBJETIVOS

    Se emplearán los siguientes criterios en la selección de una de las herramientas exploradas para gestionar las tareas del proyecto.

    - Actualizado: tendrán preferencia los gestores de dependencias que no hayan sido "abandonados" o no reciban modificaciones de forma habitual, por lo que tienen un buen mantenimiento. De la misma manera, aunque no sean actualizados continuamente serán bien consideradas las herramientas que se encuentren en etapas de desarrollo muy estables y, por lo tanto, sin errores críticos.

## OPCIONES

### Make

    En primer lugar analizaremos quizás el más conocido y generalizado. Make es una herramienta que controla la generación de ejecutables y otros archivos no fuente de un programa a partir de un archivo llamado makefile. Este enumera cada uno de los archivos no fuente y cómo procesarlos a partir de otros. El archivo make emplea reglas específicas para construir los *target* a partir de archivos fuente (inputs) y especifica sus dependencias.

    Como ventaja, make es un proyecto desarrollado desde los 70, con una madurez más que suficiente para confiar en su correcto funcionamiento en todo momento, a pesar de no tener tantas actualizaciones en la actualidad.

### Mage

    Como alternativa a Make, en muchos foros se recomienda Mage. Mage es un reemplazo para Make que utiliza Go puro para definir y ejecutar tareas. No requiere un archivo Makefile ni otra sintaxis externa.

    Respecto a su mantenimiento, el último commit fue realizado hace dos semanas y se puede ver que el proyecto es actualizado a menudo.
    Sin embargo, y aunque la popularidad de una herramienta no sea un criterio de selección válido, en este caso lo pequeño de su comunidad puede suponer un problema a largo plazo, puesto que se trata de una herramienta de código abierto. Esto puede deparar en un mantenimiento pobre.

### Task

    Desde el mismo foro se habla de Task. Este, igualmente, según su documentación, dice diferenciarse de make por su simplicidad.

    Visto su github, task se actualiza constantemente desde hace, al menos, 7 años, con una comunidad bastante más activa que la de mage (aunque siga sin ser tan tradicionalmente empleada como make).

## SELECCIÓN FINAL

Se ha elegido, finalmente, la herramienta **Task** por las siguientes razones.
Con respecto a **Mage**, se ha optado darle prioridad al criterio de mantenimiento porque el riesgo de que la herramienta caiga en abandono puede causar errores.
Por el resto, entre **Make** y **Task** ambas son buenas opciones, por lo que se elige, a criterio propio, **Task**.
