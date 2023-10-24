Grado en Ingeniería Informática

Asignatura: Gestión del Conocimiento en las Organizaciones

Tema: Sistemas de Recomendación
Introducción y Propósito del Programa:

Dentro del vasto campo de la Ingeniería Informática y, en particular, en el estudio de la gestión del conocimiento en las organizaciones, se ha diseñado un programa específico centrado en los sistemas de recomendación. Este software tiene como objetivo primordial procesar información basada en determinadas métricas y aplicar técnicas de predicción para emitir recomendaciones precisas.

Para lograr una eficaz adaptabilidad y versatilidad, el programa permite a los usuarios definir ciertos parámetros de ejecución que influirán directamente en los resultados obtenidos.

Parámetros de Configuración y Uso:

name:
Descripción: Este parámetro permite al usuario especificar la ruta exacta del archivo de datos que desea procesar.
Valor Predeterminado: En caso de no proporcionar una ruta específica, el programa automáticamente opta por tomar la información del archivo situado en ./tabla.txt.
metric:
Descripción: Es fundamental para decidir qué fórmula matemática o métrica se empleará en el cálculo de las recomendaciones. Se han incluido tres opciones para que el usuario elija.
Opciones Disponibles:
CP (Correlación de Pearson): Una medida estadística que indica el grado de relación lineal entre dos variables. Si no se define otra métrica, esta será seleccionada automáticamente.
DC (Distancia coseno): Mide la similitud entre dos vectores en un espacio multidimensional.
DE (Distancia Euclídea): Calcula la distancia directa entre dos puntos en un espacio.
neighbors:
Descripción: Determina cuántos vecinos cercanos (otros datos o entradas) se considerarán al emitir una recomendación.
Valor Predeterminado: Si no se especifica un número, el programa considerará 3 vecinos por defecto.
prediction:
Descripción: Define el algoritmo o método de predicción a implementar en el proceso de recomendación.
Opciones Disponibles:
PS (Predicción simple): Se basa en una estimación directa sin ajustes.
DM (Diferencia con la media): Ajusta la predicción teniendo en cuenta la media de las calificaciones o datos.
Instrucciones Detalladas de Ejecución:

Ejecución por Defecto:
Al invocar el programa sin especificar ningún parámetro:
bash
Copy code
go run main.go
El sistema automáticamente empleará el archivo ./tabla.txt, la métrica "CP", considerará 3 vecinos y aplicará la predicción "PS".
Ejecución Personalizada:
Si un usuario desea personalizar su ejecución, puede especificar cada parámetro de la siguiente manera:
bash
Copy code
go run main.go -name tabla.txt -metric CP -neighbors 3 -prediction DM
Acceso a la Ayuda:
El programa también provee una función de ayuda, para orientar a los usuarios en su uso. Para acceder a esta función, simplemente deben ejecutar:
bash
Copy code
go run main.go -h
Con la intención de proporcionar una herramienta eficaz y adaptable a las necesidades de cada organización, este programa busca ser un pilar en el ámbito de los sistemas de recomendación y la gestión del conocimiento.
