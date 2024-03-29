package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	base "base/pkg"
)

func main() {
	conexion, _ := grpc.Dial(
		// dirección del servidor
		"localhost:12345",
		// indica que se debe conectar usando TCP sin SSL
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// bloquea el hilo hasta que la conexión se establezca
		grpc.WithBlock(),
	)

	// crea un nuevo cliente gRPC sobre la conexión
	cliente := base.NewBaseClient(conexion)
	ctx := context.Background()

	cliente.Guardar(ctx, &base.ParametroGuardar{Clave: "1", Valor: "AAA"})
	cliente.Guardar(ctx, &base.ParametroGuardar{Clave: "2", Valor: "BBB"})
	cliente.Guardar(ctx, &base.ParametroGuardar{Clave: "3", Valor: "CCC"})
	cliente.Guardar(ctx, &base.ParametroGuardar{Clave: "4", Valor: "DDD"})
	cliente.Guardar(ctx, &base.ParametroGuardar{Clave: "5", Valor: "EEE"})
	cliente.Guardar(ctx, &base.ParametroGuardar{Clave: "6", Valor: "FFF"})
	cliente.Guardar(ctx, &base.ParametroGuardar{Clave: "7", Valor: "GGG"})
	cliente.Guardar(ctx, &base.ParametroGuardar{Clave: "8", Valor: "HHH"})
	cliente.Guardar(ctx, &base.ParametroGuardar{Clave: "9", Valor: "III"})
	cliente.Guardar(ctx, &base.ParametroGuardar{Clave: "10", Valor: "JJJ"})
	resultado, _ := cliente.Obtener(ctx, &base.ParametroObtenerEliminar{Clave: "4"})
	fmt.Println(resultado)
	cliente.Eliminar(ctx, &base.ParametroObtenerEliminar{Clave: "4"})
	cliente.Eliminar(ctx, &base.ParametroObtenerEliminar{Clave: "5"})

	// Probando si existe Clave 4 que se elimino
	resultadoEliminado4, _ := cliente.Obtener(ctx, &base.ParametroObtenerEliminar{Clave: "4"})
	fmt.Println("Clave 4 ->", resultadoEliminado4)

	// Probando si existe Clave 5 que se elimino
	resultadoEliminado5, _ := cliente.Obtener(ctx, &base.ParametroObtenerEliminar{Clave: "5"})
	fmt.Println("Clave 5 ->", resultadoEliminado5)

}
