//Usar:
//go test
//go test -bench .
//go test -cover
//go test -coverprofile c.out
//go tool cover -html=c.out
package auxiliar

import (
	"fmt"
	"strings"
	"testing"
)

func TestCat(t *testing.T) {
	result := "Este deberia ser el resultado"
	xs := strings.Split(result, " ")
	s := Cat(xs, " ")
	if s != result {
		t.Error("got", s, "want", result)
	}
}

func TestJoin(t *testing.T) {
	result := "Este deberia ser el resultado"
	xs := strings.Split(result, " ")
	s := Join(xs, " ")
	if s != result {
		t.Error("got", s, "want", result)
	}
}

//Tiene sentido si va a ser paarte de un package
func ExampleCat() {
	s := "Este deberia ser el resultado"
	xs := strings.Split(s, " ")
	fmt.Println(Cat(xs, " "))
	// Output:
	// Este deberia ser el resultado
}

func ExampleJoin() {
	s := "Este deberia ser el resultado"
	xs := strings.Split(s, " ")
	fmt.Println(Join(xs, " "))
	// Output:
	// Este deberia ser el resultado
}

const s = `La noche del catorce de marzo de 1939, en un departamento de la Zeltnergasse de Praga, Jaromir Hladík, autor de la inconclusa tragedia Los enemigos, de una Vindicación de la eternidad y de un examen de las indirectas fuentes judías de Jakob Boehme, soñó con un largo ajedrez. No lo disputaban dos individuos sino dos familias ilustres; la partida había sido entablada hace muchos siglos; nadie era capaz de nombrar el olvidado premio, pero se murmuraba que era enorme y quizá infinito; las piezas y el tablero estaban en una torre secreta; Jaromir (en el sueño) era el primogénito de una de las familias hostiles; en los relojes resonaba la hora de la impostergable jugada; el soñador corría por las arenas de un desierto lluvioso y no lograba recordar las figuras ni las leyes del ajedrez. En ese punto, se despertó. Cesaron los estruendos de la lluvia y de los terribles relojes. Un ruido acompasado y unánime, cortado por algunas voces de mando, subía de la Zeltnergasse. Era el amanecer, las blindadas vanguardias del Tercer Reich entraban en Praga.
El diecinueve, las autoridades recibieron una denuncia; el mismo diecinueve, al atardecer, Jaromir Hladík fue arrestado. Lo condujeron a un cuartel aséptico y blanco, en la ribera opuesta del Moldau. No pudo levantar uno solo de los cargos de la Gestapo: su apellido materno era Jaroslavski, su sangre era judía, su estudio sobre Boehme era judaizante, su firma delataba el censo final de una protesta contra el Anschluss. En 1928, había traducido el Sepher Yezirah para la editorial Hermann Barsdorf; el efusivo catálogo de esa casa había exagerado comercialmente el renombre del traductor; ese catálogo fue hojeado por Julius Rothe, uno de los jefes en cuyas manos estaba la suerte de Hladík. No hay hombre que, fuera de su especialidad, no sea crédulo; dos o tres adjetivos en letra gótica bastaron para que Julius Rothe admitiera la preeminencia de Hladík y dispusiera que lo condenaran a muerte, pour encourager les autres. Se fijó el día veintinueve de marzo, a las nueve a.m. Esa demora (cuya importancia apreciará después el lector) se debía al deseo administrativo de obrar impersonal y pausadamente, como los vegetales y los planetas.
El primer sentimiento de Hladík fue de mero terror. Pensó que no lo hubieran arredrado la horca, la decapitación o el degüello, pero que morir fusilado era intolerable. En vano se redijo que el acto puro y general de morir era lo temible, no las circunstancias concretas. No se cansaba de imaginar esas circunstancias: absurdamente procuraba agotar todas las variaciones. Anticipaba infinitamente el proceso, desde el insomne amanecer hasta la misteriosa descarga. Antes del día prefijado por Julius Rothe, murió centenares de muertes, en patios cuyas formas y cuyos ángulos fatigaban la geometría, ametrallado por soldados variables, en número cambiante, que a veces lo ultimaban desde lejos; otras, desde muy cerca. Afrontaba con verdadero temor (quizá con verdadero coraje) esas ejecuciones imaginarias; cada simulacro duraba unos pocos segundos; cerrado el círculo, Jaromir interminablemente volvía a las trémulas vísperas de su muerte. Luego reflexionó que la realidad no suele coincidir con las previsiones; con lógica perversa infirió que prever un detalle circunstancial es impedir que éste suceda. Fiel a esa débil magia, inventaba, para que no sucedieran, rasgos atroces; naturalmente, acabó por temer que esos rasgos fueran proféticos. Miserable en la noche, procuraba afirmarse de algún modo en la sustancia fugitiva del tiempo. Sabía que éste se precipitaba hacia el alba del día veintinueve; razonaba en voz alta: Ahora estoy en la noche del veintidós; mientras dure esta noche (y seis noches más) soy invulnerable, inmortal. Pensaba que las noches de sueño eran piletas hondas y oscuras en las que podía sumergirse. A veces anhelaba con impaciencia la definitiva descarga, que lo redimiría, mal o bien, de su vana tarea de imaginar. El veintiocho, cuando el último ocaso reverberaba en los altos barrotes, lo desvió de esas consideraciones abyectas la imagen de su drama Los enemigos.
Hladík había rebasado los cuarenta años. Fuera de algunas amistades y de muchas costumbres, el problemático ejercicio de la literatura constituía su vida; como todo escritor, medía las virtudes de los otros por lo ejecutado por ellos y pedía que los otros lo midieran por lo que vislumbraba o planeaba. Todos los libros que había dado a la estampa le infundían un complejo arrepentimiento. En sus exámenes de la obra de Boehme, de Abnesra y de Flood, había intervenido esencialmente la mera aplicación; en su traducción del Sepher Yezirah, la negligencia, la fatiga y la conjetura. Juzgaba menos deficiente, tal vez, la Vindicación de la eternidad: el primer volumen historia las diversas eternidades que han ideado los hombres, desde el inmóvil Ser de Parménides hasta el pasado modificable de Hinton; el segundo niega (con Francis Bradley) que todos los hechos del universo integran una serie temporal. Arguye que no es infinita la cifra de las posibles experiencias del hombre y que basta una sola “repetición” para demostrar que el tiempo es una falacia… Desdichadamente, no son menos falaces los argumentos que demuestran esa falacia; Hladík solía recorrerlos con cierta desdeñosa perplejidad. También había redactado una serie de poemas expresionistas; éstos, para confusión del poeta, figuraron en una antología de 1924 y no hubo antología posterior que no los heredara. De todo ese pasado equívoco y lánguido quería redimirse Hladík con el drama en verso Los enemigos. (Hladík preconizaba el verso, porque impide que los espectadores olviden la irrealidad, que es condición del arte.)
Este drama observaba las unidades de tiempo, de lugar y de acción; transcurría en Hradcany, en la biblioteca del barón de Roemerstadt, en una de las últimas tardes del siglo diecinueve. En la primera escena del primer acto, un desconocido visita a Roemerstadt. (Un reloj da las siete, una vehemencia de último sol exalta los cristales, el aire trae una arrebatada y reconocible música húngara.) A esta visita siguen otras; Roemerstadt no conoce las personas que lo importunan, pero tiene la incómoda impresión de haberlos visto ya, tal vez en un sueño. Todos exageradamente lo halagan, pero es notorio -primero para los espectadores del drama, luego para el mismo barón- que son enemigos secretos, conjurados para perderlo. Roemerstadt logra detener o burlar sus complejas intrigas; en el diálogo, aluden a su novia, Julia de Weidenau, y a un tal Jaroslav Kubin, que alguna vez la importunó con su amor. Éste, ahora, se ha enloquecido y cree ser Roemerstadt… Los peligros arrecian; Roemerstadt, al cabo del segundo acto, se ve en la obligación de matar a un conspirador. Empieza el tercer acto, el último. Crecen gradualmente las incoherencias: vuelven actores que parecían descartados ya de la trama; vuelve, por un instante, el hombre matado por Roemerstadt. Alguien hace notar que no ha atardecido: el reloj da las siete, en los altos cristales reverbera el sol occidental, el aire trae la arrebatada música húngara. Aparece el primer interlocutor y repite las palabras que pronunció en la primera escena del primer acto. Roemerstadt le habla sin asombro; el espectador entiende que Roemerstadt es el miserable Jaroslav Kubin. El drama no ha ocurrido: es el delirio circular que interminablemente vive y revive Kubin.
Nunca se había preguntado Hladík si esa tragicomedia de errores era baladí o admirable, rigurosa o casual. En el argumento que he bosquejado intuía la invención más apta para disimular sus defectos y para ejercitar sus felicidades, la posibilidad de rescatar (de manera simbólica) lo fundamental de su vida. Había terminado ya el primer acto y alguna escena del tercero; el carácter métrico de la obra le permitía examinarla continuamente, rectificando los hexámetros, sin el manuscrito a la vista. Pensó que aun le faltaban dos actos y que muy pronto iba a morir. Habló con Dios en la oscuridad. Si de algún modo existo, si no soy una de tus repeticiones y erratas, existo como autor de Los enemigos. Para llevar a término ese drama, que puede justificarme y justificarte, requiero un año más. Otórgame esos días, Tú de Quien son los siglos y el tiempo. Era la última noche, la más atroz, pero diez minutos después el sueño lo anegó como un agua oscura.
Hacia el alba, soñó que se había ocultado en una de las naves de la biblioteca del Clementinum. Un bibliotecario de gafas negras le preguntó: ¿Qué busca? Hladík le replicó: Busco a Dios. El bibliotecario le dijo: Dios está en una de las letras de una de las páginas de uno de los cuatrocientos mil tomos del Clementinum. Mis padres y los padres de mis padres han buscado esa letra; yo me he quedado ciego, buscándola. Se quitó las gafas y Hladík vio los ojos, que estaban muertos. Un lector entró a devolver un atlas. Este atlas es inútil, dijo, y se lo dio a Hladík. Éste lo abrió al azar. Vio un mapa de la India, vertiginoso. Bruscamente seguro, tocó una de las mínimas letras. Una voz ubicua le dijo: El tiempo de tu labor ha sido otorgado. Aquí Hladík se despertó.
Recordó que los sueños de los hombres pertenecen a Dios y que Maimónides ha escrito que son divinas las palabras de un sueño, cuando son distintas y claras y no se puede ver quien las dijo. Se vistió; dos soldados entraron en la celda y le ordenaron que los siguiera.
Del otro lado de la puerta, Hladík había previsto un laberinto de galerías, escaleras y pabellones. La realidad fue menos rica: bajaron a un traspatio por una sola escalera de fierro. Varios soldados -alguno de uniforme desabrochado- revisaban una motocicleta y la discutían. El sargento miró el reloj: eran las ocho y cuarenta y cuatro minutos. Había que esperar que dieran las nueve. Hladík, más insignificante que desdichado, se sentó en un montón de leña. Advirtió que los ojos de los soldados rehuían los suyos. Para aliviar la espera, el sargento le entregó un cigarrillo. Hladík no fumaba; lo aceptó por cortesía o por humildad. Al encenderlo, vio que le temblaban las manos. El día se nubló; los soldados hablaban en voz baja como si él ya estuviera muerto. Vanamente, procuró recordar a la mujer cuyo símbolo era Julia de Weidenau…
El piquete se formó, se cuadró. Hladík, de pie contra la pared del cuartel, esperó la descarga. Alguien temió que la pared quedara maculada de sangre; entonces le ordenaron al reo que avanzara unos pasos. Hladík, absurdamente, recordó las vacilaciones preliminares de los fotógrafos. Una pesada gota de lluvia rozó una de las sienes de Hladík y rodó lentamente por su mejilla; el sargento vociferó la orden final.
El universo físico se detuvo
Las armas convergían sobre Hladík, pero los hombres que iban a matarlo estaban inmóviles. El brazo del sargento eternizaba un ademán inconcluso. En una baldosa del patio una abeja proyectaba una sombra fija. El viento había cesado, como en un cuadro. Hladík ensayó un grito, una sílaba, la torsión de una mano. Comprendió que estaba paralizado. No le llegaba ni el más tenue rumor del impedido mundo. Pensó estoy en el infierno, estoy muerto. Pensó estoy loco. Pensó el tiempo se ha detenido. Luego reflexionó que en tal caso, también se hubiera detenido su pensamiento. Quiso ponerlo a prueba: repitió (sin mover los labios) la misteriosa cuarta égloga de Virgilio. Imaginó que los ya remotos soldados compartían su angustia: anheló comunicarse con ellos. Le asombró no sentir ninguna fatiga, ni siquiera el vértigo de su larga inmovilidad. Durmió, al cabo de un plazo indeterminado. Al despertar, el mundo seguía inmóvil y sordo. En su mejilla perduraba la gota de agua; en el patio, la sombra de la abeja; el humo del cigarrillo que había tirado no acababa nunca de dispersarse. Otro “día” pasó, antes que Hladík entendiera.
Un año entero había solicitado de Dios para terminar su labor: un año le otorgaba su omnipotencia. Dios operaba para él un milagro secreto: lo mataría el plomo alemán, en la hora determinada, pero en su mente un año transcurría entre la orden y la ejecución de la orden. De la perplejidad pasó al estupor, del estupor a la resignación, de la resignación a la súbita gratitud.
No disponía de otro documento que la memoria; el aprendizaje de cada hexámetro que agregaba le impuso un afortunado rigor que no sospechan quienes aventuran y olvidan párrafos interinos y vagos. No trabajó para la posteridad ni aun para Dios, de cuyas preferencias literarias poco sabía. Minucioso, inmóvil, secreto, urdió en el tiempo su alto laberinto invisible. Rehizo el tercer acto dos veces. Borró algún símbolo demasiado evidente: las repetidas campanadas, la música. Ninguna circunstancia lo importunaba. Omitió, abrevió, amplificó; en algún caso, optó por la versión primitiva. Llegó a querer el patio, el cuartel; uno de los rostros que lo enfrentaban modificó su concepción del carácter de Roemerstadt. Descubrió que las arduas cacofonías que alarmaron tanto a Flaubert son meras supersticiones visuales: debilidades y molestias de la palabra escrita, no de la palabra sonora… Dio término a su drama: no le faltaba ya resolver sino un solo epíteto. Lo encontró; la gota de agua resbaló en su mejilla. Inició un grito enloquecido, movió la cara, la cuádruple descarga lo derribó.
Jaromir Hladík murió el veintinueve de marzo, a las nueve y dos minutos de la mañana.`

var xs []string

func BenchmarkCat(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cat(xs, " ")
	}
}

func BenchmarkJoin(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(xs, " ")
	}
}
