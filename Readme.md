# Mejora del proyecto del servidor

## Microservicios en GO

En esta aplicación, el modelo de negocio del equipo es el desarrollo de una tienda online. Partiendo de esta base el microservicio son productos que los usuarios pueden comprar. De esta manera se ha refactorizado i/o creado nueva funcionalidad

- Common: con aspectos de configuración globales al microservicio así como utilidades para la creación de sesiones de base de datos
- Models: en donde se definirán los modelos utilizados por el microservicio
- Routers: Para la definición de las rutas o endpoints que publicará el 
microservicio
- Data: En donde se incluyen las funciones que son ejecutadas para obtener la información de respuesta de los endpoints del microservicio
- Controllers: Que incluirá todo lo necesario para que, partiendo de la información de respuesta de las utilidades de la carpeta Data, construirá las respuestas finales de los endpoints del microservicio.

### main.go:
![Carpeta common](./imgs/main.png)

- El main, primeramente al conectarse con la base de datos hará una migración de las entidades de la base de datos. Estas entidades estarán definidas en la carpeta de los modelos del proyecto. Una vez hecha la conexión y las migraciones, se iniciarán las rutas con sus respectivos controladores. Para finalizar se desplegará el servidor en el puerto 8080 de nuestra localhost.

### Common:
![Carpeta common](./imgs/common.png)

- **database.go:** Aquí estará la configuración de la connexión de la base de datos.

- **utils.go:** El interior de este archivo es para la utilización de herramientas como el display de errores o generar tokens.

### Models:
![Carpeta models](./imgs/models.png)

- **model.go:** Definición de todas las entidades del proyecto, en este caso es el de productos y el de todos sus datos relacionados como el autor y los comentarios.

![Carpeta models](./imgs/models2.png)

### Routers:

![Carpeta routers](./imgs/routers.png)

- **router.go:** Inicia las rutas llamando al enrutador de los productos

![Carpeta routers](./imgs/routers2.png)


- **products.go:** Inicia las rutas de los productos, en este caso solo hay dos la de obtener todos los productos (/products) o la de obtener solo un producto (/product). Cada una de estas con su respectivo controlador que manejará la peticióny devolverá el resultado.

![Carpeta routers](./imgs/routers1.png)

### Data:

![Carpeta data](./imgs/data.png)

- **repositories.go:** Tiene las operaciones con la base de datos para obtener lo resultados de los distintos productos. Estas funciones serán llamadas por los controladores para buscar los datos en la base de datos y crear la respuesta para el usuario que ha hecho la petición.

![Carpeta data](./imgs/data1.png)

### Controllers:

![Carpeta controllers](./imgs/controllers.png)

- **productController.go:** Tiene los metodos que controlarán, validarán y responderán las peticiones de los usuarios. En este caso son dos: 

![Carpeta controllers](./imgs/controller1.png)
![Carpeta controllers](./imgs/controller2.png)

- **resourses.go:** Este archivo es una esopecie de "serializer" pero solo define el tipo de respuesta que tendrá.

![Carpeta controllers](./imgs/controller3.png)

## Extensión de servicios

En este apartado intentaremos crear una plantilla para reutilizar código de nuestro Docker Compose, para ello si utilizamos una versión de docker inferior a la 3 podremos utilizar "**extends**" que básicamente se trata crear un servicio base del cual extenderán los servicios que tengan estas mismas características en común. 

Para versiones de docker compose superiores o igual a 3 podemos utilizar los **campos de extensión**, que se trata de crear una especificación común, que se puede reutilizar para todos los servicios.

La versión de nuestro docker compose es la 3 por lo que hemos optado por los campos de extensión.

Para ello estos campos deben de estar ubicados en la raíz del Compose y su nombre debe empezar por la secuencia `x-` y posteriormente debemos escribir el nombre por el cual llamaremos a este campo, precedido por el signo `&` (**&default-command**).

![Creación del campo de extensión](./imgs/fieldsExtension1.png)

Para utilizar este campo debemos utilizar el singo ´*´ sucedido de el nombre dado anteriormente (*default-command).

Para reutilizar la especificación command tenemos que utilizar una variable de entorno, por que según el servicio debe llamar a un microservicio u otro.

![Llamada del campo de extensión para app](./imgs/fieldsExtension2.png)

En esta imagen podemos observar como utilizamos el comando y reescribimos el valor de la variable de entrono con el valor **App**.

![Llamada del campo de extensión para products](./imgs/fieldsExtension3.png)

En esta imagen observamos utilizamos el comando con el valor **products**.