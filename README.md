# Gollama: Cliente Go para Ollama

**Gollama** es un cliente de Go diseñado para interactuar con el servicio Ollama, proporcionando acceso a modelos de lenguaje de gran tamaño (LLM) como **Llama 3.2**. El cliente permite enviar preguntas a modelos LLM y recibir respuestas generadas a través de la API de Ollama, todo configurado de manera flexible a través de un archivo `config.yaml`.

## Características

- Interactúa con modelos de lenguaje como **Llama 3.2**.
- Configuración a través de un archivo YAML.
- Reintentos automáticos con configuraciones de tiempo de espera.
- Manejador de errores personalizado para distintos tipos de fallos.
- Soporte completo para el servicio Ollama a través de su API REST.

## Requisitos

- Go 1.18 o superior.
- Un servicio de Ollama en ejecución (local o remoto).

# Explicación del Código

## Carga de la Configuración
Al iniciar el programa, se carga el archivo `config.yaml` utilizando la función `LoadConfig()`. Esta configuración incluye parámetros como la URL de la API de Ollama, el modelo a utilizar y los parámetros de reintento.

## Creación de la Solicitud
Se crea un objeto `Request` que incluye:
- **Model**: El modelo de lenguaje que se usará para generar la respuesta (configurado previamente en el archivo YAML).
- **Prompt**: El texto o la pregunta que se le envía al modelo.
- **Stream**: Indica si la respuesta se devolverá de manera continua o no.

## Generación de Respuesta
La función `Generate()` se encarga de enviar la solicitud a la API de Ollama. Si la solicitud es exitosa, se obtiene una respuesta que se imprime en la consola.

# Personalización de la Configuración

Puedes modificar el archivo `config.yaml` según tus necesidades para cambiar cómo interactúa el cliente con la API de Ollama. Las opciones más importantes son:

- **api_url**: La URL del servicio de Ollama (por defecto `http://localhost:11434`). Cambia esta URL si el servicio está ejecutándose en otro lugar.
- **timeout**: Tiempo de espera máximo para las solicitudes HTTP (en segundos).
- **retry_count**: Número de intentos que realiza el cliente en caso de fallos de conexión.
- **retry_wait_time**: Tiempo de espera entre reintentos (en segundos).
- **retry_max_wait_time**: Tiempo máximo de espera entre reintentos.
- **model**: El modelo LLM que deseas usar (por ejemplo, `llama3.2`).
- **content_type**: Tipo de contenido de las solicitudes HTTP (por defecto `application/json`).
- **api_endpoint**: El endpoint de la API de Ollama (por ejemplo, `/api/generate`).

# Manejo de Errores

El cliente Gollama maneja varios tipos de errores que pueden ocurrir durante la ejecución, y los mensajes de error se definen en el archivo `config.yaml`. Algunos de los errores comunes son:

- **model_empty**: Error cuando no se proporciona un modelo en la solicitud. Se lanza si el campo `Model` está vacío.
- **prompt_empty**: Error cuando no se proporciona un texto (prompt) para el modelo. Se lanza si el campo `Prompt` está vacío.
- **request_error**: Error cuando falla la solicitud HTTP al servicio Ollama.
- **response_error**: Error cuando ocurre un problema al procesar la respuesta del servicio Ollama.
- **status_code_error**: Error cuando la respuesta de la API no tiene un código de estado 200 (OK).

Si alguno de estos errores ocurre, el cliente genera un mensaje detallado según el tipo de error.

# Código de Conducta

Por favor, sigue las pautas y buenas prácticas de código al contribuir. Asegúrate de que tus cambios estén correctamente documentados y probados, y que sigan el estilo de código del proyecto.

# Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo `LICENSE` para más detalles.
