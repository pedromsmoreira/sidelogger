package loggers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateWithLoggerNameAsNullReturnNullLogger(t *testing.T) {
	// arrange
	name := "null"
	expectedName := "NullLogger"
	// act
	logger, err := CreateLogger(name)
	// assert
	assert.NotNil(t, err)
	assert.Equal(t, expectedName, logger.GetName())

}

func TestCreateWithLoggerNameAsEmptyStringReturnNullLogger(t *testing.T) {
	// arrange
	name := ""
	expectedName := "NullLogger"
	// act
	logger, err := CreateLogger(name)
	// assert
	assert.NotNil(t, err)
	assert.Equal(t, expectedName, logger.GetName())
}

func TestCreateWithLoggerNameAsZapReturnZapLogger(t *testing.T) {
	// arrange
	name := "uberzap"
	expectedName := "zap-logger"
	// act
	logger, err := CreateLogger(name)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, expectedName, logger.GetName())
}

func TestCreateWithLoggerNameAsLogrusReturnLogrusLogger(t *testing.T) {
	// arrange
	name := "logrus"
	expectedName := "logrus-logger"
	// act
	logger, err := CreateLogger(name)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, expectedName, logger.GetName())

}

func TestCreateWithLoggerNameAsApexReturnApexLogger(t *testing.T) {
	// arrange
	name := "apex"
	expectedName := "apex-logger"
	// act
	logger, err := CreateLogger(name)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, expectedName, logger.GetName())
}
