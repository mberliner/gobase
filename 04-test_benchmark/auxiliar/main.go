//Package auxiliar sirve para verificar Tests, Cobertura, Benchmarking y export de Funciones
//go test
//go test -bench .
//go test -cover
//go test -coverprofile c.out
//go tool cover -html=c.out
package auxiliar

import "strings"

//Cat recibe un slice de strings y los retorna separados por un separador
//su uso aqui es a fin de verificar la performance con benchmark y Tests de go
//Lleva mayusculas porque se exporta de este package hacia otros
func Cat(xs []string, sep string) string {
	r := xs[0]
	for _, v := range xs[1:] {
		r += sep
		r += v
	}
	return r
}

//Join recibe un slice de strings y los retorna separados por un separador
//su uso aqui es a fin de verificar la performance con benchmark y Tests de go
//Lleva mayusculas porque se exporta de este package hacia otros
func Join(xs []string, sep string) string {
	return strings.Join(xs, sep)
}
