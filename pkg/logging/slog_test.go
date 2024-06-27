package logging_test

// import (
// 	. "github.com/mmcloud/solder/pkg/logging"
// 	. "github.com/onsi/ginkgo/v2"
// 	. "github.com/onsi/gomega"
// )

// var _ = Describe("SlogLogger", func() {
// 	var (
// 		logger *SlogLogger
// 		config LoggingConfig
// 		buffer *bytes.Buffer
// 	)

// 	BeforeEach(func() {
// 		buffer = &bytes.Buffer{}
// 		config = LoggingConfig{
// 			Level: "INFO",
// 		}
// 		logger = NewSlogLogger(config)
// 		logger.Logger = slog.New(slog.NewJSONHandler(buffer, &slog.HandlerOptions{Level: &slog.LevelVar{}}))
// 	})

// 	Describe("SetLevel", func() {
// 		It("should set the logging level correctly", func() {
// 			logger.SetLevel("DEBUG")
// 			Expect(logger.GetLevel()).To(Equal("DEBUG"))
// 		})

// 		It("should handle invalid levels gracefully", func() {
// 			logger.SetLevel("INVALID")
// 			Expect(logger.GetLevel()).To(Equal("DEBUG")) // Default level
// 		})
// 	})

// 	Describe("Logging Methods", func() {
// 		BeforeEach(func() {
// 			logger.SetLevel("DEBUG") // Ensure the level is set to debug for capturing all logs
// 		})

// 		It("should log a message at the debug level", func() {
// 			logger.Debug("This is a debug message")
// 			Expect(buffer.String()).To(ContainSubstring("This is a debug message"))
// 		})

// 		It("should log a message at the info level", func() {
// 			logger.Info("This is an info message")
// 			Expect(buffer.String()).To(ContainSubstring("This is an info message"))
// 		})

// 		It("should log a message at the warn level", func() {
// 			logger.Warn("This is a warning message")
// 			Expect(buffer.String()).To(ContainSubstring("This is a warning message"))
// 		})

// 		It("should log a message at the error level", func() {
// 			logger.Error("This is an error message")
// 			Expect(buffer.String()).To(ContainSubstring("This is an error message"))
// 		})

// 		It("should log a message at the fatal level", func() {
// 			defer func() {
// 				if r := recover(); r != nil {
// 					// Check the log output before the program exits
// 					Expect(buffer.String()).To(ContainSubstring("FATAL: This is a fatal message"))
// 				}
// 			}()
// 			logger.Fatal("This is a fatal message")
// 		})
// 	})

// 	Describe("GetLevel", func() {
// 		It("should return the current logging level", func() {
// 			Expect(logger.GetLevel()).To(Equal("INFO"))
// 		})
// 	})

// 	Describe("GetLoggerKind", func() {
// 		It("should return the correct logger type", func() {
// 			Expect(logger.GetLoggerKind()).To(Equal(SlogLoggerKind))
// 		})
// 	})

// 	Describe("Logging Levels", func() {
// 		BeforeEach(func() {
// 			logger.SetLevel("INFO")
// 		})

// 		It("should not log a debug message when the level is set to info", func() {
// 			logger.Debug("This is a debug message")
// 			Expect(buffer.String()).NotTo(ContainSubstring("This is a debug message"))
// 		})
// 	})
// })
