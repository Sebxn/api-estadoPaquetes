package routes

import (
	"backend/api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigureRoutes(r *mux.Router) {
	// allowedOrigins := []string{"http://facturacion.lumonidy.studio", "http://localhost:3000"}

	// c := middleware.CorsMiddleware(allowedOrigins)
	// r.Use(c)

	r.Handle("/api/facturacion", http.HandlerFunc(handlers.HomeHandler))

	r.Handle("/api/facturacion/actualizar_estado", http.HandlerFunc(handlers.ActualizarEstadoReserva))
	// Agrega más configuraciones de rutas aquí si es necesario
}
