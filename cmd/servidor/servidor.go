package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	base "base/pkg"
)

func main() {
	// Uso os.Args para validar argumentos
	argumentos := os.Args[1:]
	if len(argumentos) != 1 {
		fmt.Println("Error: Se debe pasar un solo argumento, debe ser un valor de 0 a 2")
		os.Exit(1)
	}

	puerto, ok := base.MapaNodos[argumentos[0]]
	if !ok {
		fmt.Println("Error: El argumento debe ser un valor de 0 a 2")
		os.Exit(1)
	}

	// Crear el servidor gRPC
	servicio := base.NuevoServidor(argumentos[0])
	servidorReal := grpc.NewServer()
	base.RegisterBaseServer(servidorReal, &servicio)

	// Iniciar el servidor gRPC en el puerto especificado
	direccion := "localhost:" + puerto
	listen, error := net.Listen("tcp", direccion)

	// Comprobacion de fallo al escuchar
	if error != nil {
		log.Fatalf("fallo al escuchar: %s: %v", direccion, error)
	}

	fmt.Println("Servidor: ", servicio.IdNodo, " - Inciando Puerto: ", listen.Addr().String())

	// Goroutine para manejar señales de interrupción
	go func() {
		// Esperar señales de interrupción
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan

		// Cerrar el servidor gRPC y salir
		fmt.Println("\nInterrupción detectada. Cerrando el servidor gRPC...")
		servidorReal.GracefulStop()
		os.Exit(0)
	}()

	if err := servidorReal.Serve(listen); err != nil {
		log.Fatalf("Fallo al servidor: %v", err)
	}
}
