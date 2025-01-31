# IMAGEN CONTENEDOR

## CRITERIOS DE SELECCIÓN OBJETIVOS

Se emplearán los siguientes criterios para elegir una imagen específica en la ejecución de nuestros tests:

- Mantenimiento: la imagen se actualiza frecuentemente. Podemos verificar la situación actual del mantenimiento de cada imagen en su github o en las páginas oficiales de cada una.
- Incluye el lenguaje empleado: las imágenes que incluyan go tendrán preferencia ya que no se tendrá que instalar.
- Bajo coste de memoria: cuanto más pesada, más se tardará en descargar una imagen cada vez que se quiera utilizar y consumirá más recursos. Es por ello que, en caso de que los otros criterios no sean suficientes para elegir la herramienta, se tendrá en cuenta la que tenga un tamaño más reducido. Esto se puede comprobar desde el apartado "Tags" de cada herramienta en [docker hub](https://hub.docker.com/).

## ACLARACIONES

Ya que uno de los requisitos básicos será que la imagen tenga golang instalado por defecto la única imagen base que nos quedará será la de golang. Sin embargo, existen, dentro de esta imagen, varios "sabores" o versiones de la misma que compararemos a continuación.

## OPCIONES

### [Imagen de docker para golang](https://hub.docker.com/_/golang)

Es una imagen mantenida con constancia, con su último push realizado hace 7 días. El peso de su versión latest es de 289.26 MB. Al ser la imagen oficial de go, incluye el lenguaje instalado y tendrá preferencia.

Dentro de la misma imagen podremos encontrar alternativas o "flavours" (según lo especificado en docker hub):

- `golang:<version>` : es la imagen por defecto, tiene las características de Debian.

- `golang:<version>-alpine` : es una imagen basada en alpine, por lo que nos encontramos igualmente ante una bajada considerable del espacio ocupado (sobre los 5 MB).

- `golang:<version>-windowsservercore` : no nos sirve ya que es la imagen diseñada para funcionar en windows server.

### [Alpine](https://hub.docker.com/_/alpine)

Aunque con menos constancia que la anterior, la imagen de Alpine es actualizada a menudo y su último push fue realizado hace 21 días. Es una imagen bien conocida por su minimalismo y su reducido tamaño (3.41 MB la versión latest), pero esto hace que venga con las herramientas justas y necesarias y sin incluir el lenguaje golang.

### [BusyBox](https://hub.docker.com/_/busybox)

Último push realizado hace 16 días. Está actualizada. Su versión latest pesa 2.11 MB.

### [Debian](https://hub.docker.com/_/debian)

Está mantenido. Su última actualización se llevó acabo hace 15 días. Su versión latest pesa 47.17 MB.

## SELECCIÓN FINAL

Emplearemos la imagen golang con el "flavour" de alpine, ya que es la que mejor se acoge a los criterios de selección: buen mantenimiento, incluye golang y ocupa poco espacio de almacenamiento.
Para llevar a cabo la construcción de nuestro dockerfile utilizaremos la documentación oficial de [dockerdocs](https://docs.docker.com/guides/golang/run-tests/)
