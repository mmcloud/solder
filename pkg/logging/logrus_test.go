package logging_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"

	"github.com/mmcloud/solder/mocks" // Import the generated mocks
	"github.com/mmcloud/solder/pkg/constants"
	"github.com/mmcloud/solder/pkg/interfaces"
	"github.com/mmcloud/solder/pkg/logging"
)

var _ = Describe("LogrusLogger", func() {
	var (
		mockLogger interfaces.Logger
		config     *logging.LoggingConfig
		buffer     *bytes.Buffer
	)

	BeforeEach(func() {
		buffer = new(bytes.Buffer)

		config = &logging.LoggingConfig{
			Level:      constants.DebugLevel,
			LoggerKind: constants.LogrusLoggerKind,
		}

		// Create a mock LogrusLogger (assumes you've run mockery to generate the mock)
		mockLogrusLogger := new(mocks.Logger)

		// Set up expectations on the mock logrus logger
		mockLogrusLogger.On("Debug", mock.AnythingOfType("string")).Run(func(args mock.Arguments) {
			// Store the log message in the buffer for later verification
			buffer.WriteString(args.Get(0).(string) + "\n")
		}).Return()

		// Similar setup for Info, Warn, Error, Fatal, SetLevel, and GetLevel

		// Create a mock LoggerFactory that returns our mock logrus logger
		mockLoggerFactory := new(mocks.LoggerFactory)
		mockLoggerFactory.On("GetLogger").Return(mockLogrusLogger, nil)

		var err error
		mockLogger, err = mockLoggerFactory.GetLogger()
		Expect(err).NotTo(HaveOccurred())

		// Redirect output to buffer
		mockLogger.(*logging.LogrusLogger).Logger.Out = buffer
	})

	AfterEach(func() {
		buffer.Reset() // clear the buffer after each test
		mockLogger = nil
	})

	Describe("SetLevel", func() {
		It("should set the logging level correctly", func() {
			mockLogger.SetLevel(constants.DebugLevel)
			Expect(mockLogger.GetLevel()).To(Equal(constants.DebugLevel))
		})

		It("should handle invalid levels gracefully", func() {
			mockLogger.SetLevel(constants.Level("INVALID"))
			Expect(mockLogger.GetLevel()).To(Equal(constants.DefaultLevel)) // Default level
		})
	})

	Describe("Logging Methods", func() {
		var (
			mockLogger *logrus.Logger
			buffer     *bytes.Buffer
		)

		BeforeEach(func() {
			buffer = &bytes.Buffer{}
			mockLogger = &logrus.Logger{
				Out:       buffer,
				Formatter: new(logrus.TextFormatter),
				Hooks:     make(logrus.LevelHooks),
			}
			logger := logging.NewLogrusLogger(config) // Declare and define the logger variable
			logger.Logger = mockLogger
			logger.SetLevel(constants.DebugLevel) // Ensure the level is set to debug for capturing all logs
		})

		It("should log a message at the debug level", func() {
			mockLogger.Debug("This is a debug message")
			Expect(buffer.String()).To(ContainSubstring("This is a debug message"))
		})

		It("should log a message at the info level", func() {
			mockLogger.Info("This is an info message")
			Expect(buffer.String()).To(ContainSubstring("This is an info message"))
		})

		It("should log a message at the warn level", func() {
			mockLogger.Warn("This is a warning message")
			Expect(buffer.String()).To(ContainSubstring("This is a warning message"))
		})

		It("should log a message at the error level", func() {
			mockLogger.Error("This is an error message")
			Expect(buffer.String()).To(ContainSubstring("This is an error message"))
		})

		It("should log a message at the fatal level", func() {
			mockLogger.Fatal("This is a fatal message")
			Expect(buffer.String()).To(ContainSubstring("FATAL: This is a fatal message"))
		})
	})

	Describe("GetLevel", func() {
		It("should return the current logging level", func() {
			Expect(mockLogger.GetLevel()).To(Equal(constants.InfoLevel))
		})
	})

	Describe("GetLoggerKind", func() {
		It("should return the correct logger type", func() {
			Expect(mockLogger.GetLoggerKind()).To(Equal(constants.LogrusLoggerKind))
		})
	})

	Describe("Logging Levels", func() {
		var (
			mockLogger *logrus.Logger
			buffer     *bytes.Buffer
		)

		BeforeEach(func() {
			buffer = &bytes.Buffer{}
			mockLogger = &logrus.Logger{
				Out:       buffer,
				Formatter: new(logrus.TextFormatter),
				Hooks:     make(logrus.LevelHooks),
			}
			logger := logging.NewLogrusLogger(config) // Corrected function name
			logger.Logger = mockLogger
			logger.SetLevel(constants.InfoLevel)
		})

		It("should not log a debug message when the level is set to info", func() {
			mockLogger.Debug("This is a debug message")
			Expect(buffer.String()).NotTo(ContainSubstring("This is a debug message"))
		})
	})
})
