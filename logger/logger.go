package logger

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Función para crear un logger con rotación
func Create(basePath, pagina, categoria string) *logrus.Logger {
	// Construye la ruta según los parámetros
	var logPath string
	if categoria != "" {
		logPath = filepath.Join(basePath, pagina, categoria, "log.log")
	} else {
		logPath = filepath.Join(basePath, pagina, "general.log")
	}

	// Crea los directorios si no existen
	if err := os.MkdirAll(filepath.Dir(logPath), os.ModePerm); err != nil {
		logrus.Errorf("Error creando directorios: %v", err)
	}

	// Configura `lumberjack` para manejar la rotación
	rotator := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    100,   // Tamaño máximo del archivo en MB
		MaxAge:     0,     // No eliminar por antigüedad
		MaxBackups: 0,     // Sin límite de copias de respaldo
		Compress:   false, // No comprimir archivos antiguos
	}

	// Configura el logger de `logrus`
	logger := logrus.New()
	logger.SetOutput(rotator)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel) // Asegura que registre mensajes Debug

	return logger
}
