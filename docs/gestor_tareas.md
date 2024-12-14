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

### [Buck2](https://buck2.build/)

    Buck2 es la build tool desarrollada y usada por Meta. Está escrita en Rust y en su página principal remarcan el uso de ejecución remota y construcción incremental guardando en caché las instrucciones anteriormente compiladas.
    Es una herramientam muy reciente, con solo dos años de desarrollo, aunque parte de buck1 que comenzó a ser desarrollada en 2014. Su comunidad es activa y recibe actualizaciones a diario.
    La propia página oficial admite como desventaja su estabilidad: "Buck2 está en desarrollo activo, lo que significa que el riesgo de regresión es correspondientemente mayor. Puede haber problemas, pero se solucionarán lo más rápido posible (y se aprenderán lecciones) a través del proceso de revisión SEV de Meta."

### [Rake](https://ruby.github.io/rake/)

    Como otras herramientas a las que solo se cambia una letra a partir de "Make" (mage, rake, cake...), es un Make-like pero implementado en Ruby.
    Su comunidad no es muy activa en comparación con las herramientas anteriores y esto sería justificable por el hecho de no tener que lidiar con muchos issues debido a su desarrollo anterior, habiendo comenzado en 2008.

### [Ninja](https://ninja-build.org/)

    En el caso de Ninja, se trata de una herramienta programada a nivel de ensamblador, consiguiendo mucha más velocidad. Es más minimalista pero más optimizado que, por ejemplo,Make.
    Es un proyecto bastante actualizado desde 2012 y, aunque sus cambios se hayan frenado, sigue teniendo un desarrollo activo a día de hoy.

## SELECCIÓN FINAL

Se ha elegido, finalmente, la herramienta **Task** por las siguientes razones.
Con respecto a **Mage**, se ha optado darle prioridad al criterio de mantenimiento porque el riesgo de que la herramienta caiga en abandono puede causar errores. Algo parecido ocurre con **Rake** que, a pesar de tener cierta madurez, parece un proyecto algo inactivo.
Con **Buck2** pasa al revés, que aún es muy nuevo y se debate con la estabilidad que pueda ofrecer.
Por el resto, entre **Make**, **Ninja** y **Task** todas son buenas opciones, por lo que se elige, a criterio propio, **Task**.
