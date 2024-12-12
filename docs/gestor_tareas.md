# GESTOR DE TAREAS

## CRITERIOS DE SELECCIÓN OBJETIVOS

    Se emplearán los siguientes criterios en la selección de una de las herramientas exploradas para gestionar las tareas del proyecto.

    - 1. Se valorarán positivamente las herramientas que no añadan complejidad al proyecto (es preferible usar una específica, evitando aprender nuevos lenguajes, pero si es más complicado que escribir un comando sencillo en otro lenguaje se elegirá uno genérico).
    - 2. Actualizado: tendrán preferencia los gestores de dependencias que no hayan sido "abandonados" o no reciban modificaciones de forma habitual, por lo que tienen un buen mantenimiento. De la misma manera, aunque no sean actualizados continuamente serán bien consideradas las herramientas que se encuentren en etapas de desarrollo muy estables y, por lo tanto, sin errores críticos.
    - 3. Herramienta concisa: sirve para lo que sirve, evitando lo especificado en los errores comunes *"Usar herramientas que hacen más cosas, incluso mucho más cosas, que verificar la sintaxis. Se trata simplemente de ver si el código compila o no compila, no si sigue las mejores prácticas y demás."*
    - 4. Peso: si la herramienta necesita muchas dependencias externas u ocupa demasiado espacio en su instalación se verá reflejado negativamente. Además, es probable que, si añade mucho contenido no cumpla el criterio anterior.

## OPCIONES

### Go build

    Puede ser sencillo confundir en los foros y fuentes de información *task managers* (gestor de tareas) y *build tools* (herramientas de construcción). Esto se debe a que existen muchas herramientas de construcción que llegan a cumplir tareas de testing, generación de documentación y más, a parte llevar a cabo su compilación. Pero muchas otras solo cumplen esta última, por lo que no nos sirven. Es por esto que no se tiene en cuenta **go build** que tendría más peso que el resto por ser la oficial de golang.

### Bazle

    En el caso de **Bazle**, la situación es contraria a la de go build. Es una herramienta desarrollada por google, por lo que se podría asumir que le va bien a golang. Pero el caso es que está diseñada para escenarios más complejos y dedicada a software en múltiples lenguajes como java, c++, python, rust... No tiene lugar considerarla porque su uso es demasiado amplio y su funcionalidad adicional no es necesaria en este proyecto.

### Make

    En primer lugar analizaremos quizás el más conocido y generalizado. Make es una herramienta que controla la generación de ejecutables y otros archivos no fuente de un programa a partir de un archivo llamado makefile. Este enumera cada uno de los archivos no fuente y cómo procesarlos a partir de otros. El archivo make emplea reglas específicas para construir los *target* a partir de archivos fuente (inputs) y especifica sus dependencias.

VENTAJAS: es un proyecto desarrollado desde los 70, con una madurez más que suficiente para confiar en su correcto funcionamiento en todo momento, a pesar de no tener tantas actualizaciones en la actualidad.
A pesar de tener que comprender un nuevo formato, make es especialmente conciso en la escritura de sus reglas, sobre todo para las necesidades del proyecto.
Con respecto a su peso, no es necesario más que descargar la herramienta y crear un archivo make.

DESVENTAJAS: puede aportar complejidad al proyecto, ya que hay que aprender una nueva sintáxis totalmente independiente de golang.
Le faltan ciertas utilidades más modernas como detectar errores mientras se escribe, lo que al no conocerla puede dar lugar a muchos tropiezos.

### Mage

    Como alternativa a Make, en muchos foros se recomienda Mage. Mage es un reemplazo para Make que utiliza Go puro para definir y ejecutar tareas. No requiere un archivo Makefile ni otra sintaxis externa.

VENTAJAS: el último commit fue realizado hace dos semanas y se puede ver que el proyecto es actualizado a menudo.
Es específica de golang.
No tiene dependencias fuera de la librería estándar de go y funciona a partir de go 1.7.

DESVENTAJAS: aunque la popularidad de una herramienta no sea un criterio de selección válido, en este caso lo pequeño de su comunidad puede suponer un problema a largo plazo, puesto que se trata de una herramienta de código abierto. Esto puede deparar en un mantenimiento pobre.

### Task

    Desde el mismo foro se habla de Task. Este, igualmente, según su documentación, dice diferenciarse de make por su simplicidad.

VENTAJAS: visto su github, task se actualiza constantemente desde hace, al menos, 7 años, con una comunidad bastante más activa que la de mage (aunque siga sin ser tan tradicionalmente empleada como make).
Su instalación es muy ligera y no añade más que un binario.

DESVENTAJAS: aunque su código está escrito (principalmente) en go, las reglas que se emplean en su uso son de tipo YAML (aún así, este está diseñado para cierta naturalidad del lenguaje)

## SELECCIÓN FINAL

Se ha elegido, finalmente, la herramienta **Task** por las siguientes razones.
En comparativa con **Make**, la forma de programar con **Task** es bastante más intuitiva y cómoda, al emplear un lenguaje más natural.
Con respecto a **Mage**, se ha optado darle prioridad al criterio de mantenimiento porque el riesgo de que la herramienta caiga en abandono puede causar errores.

Por el resto de criterios, las opciones se mantienen bastante igualadas, por lo que no son de suficiente peso como para desviar la atención de los mencionados.
