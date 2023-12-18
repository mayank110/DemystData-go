package config


// AppConfig represents the configuration for the payment processor service
type AppConfig struct {
	TodoURL    string
	DefaultTodo int
}

var c *AppConfig

func initConfig() {
	/** 
		This function can be extened to write the logic of 
		1. Env specific config.
		2. File specific config. 
		3. Fetch Secrets from Vault.
	**/

	c = &AppConfig{
		TodoURL: "https://jsonplaceholder.typicode.com/todos/",
		DefaultTodo: 20,
	}

}

// GetConfig returns a serialized instance of the app config object
func GetConfig() *AppConfig {
	if c == nil {
		initConfig()
	}
	return c
}