// 配置文件库
// 北京实易时代科技有限公司
// 2014-10-22 V0.1.1
// 增加float64类型
package config

import (
	"errors"
	"github.com/BurntSushi/toml"
	"strings"
)

type Config struct {
	ConfFile  string
	Version   string
	PathLevel int
	Item      map[string]interface{}
}

func NewConfig(file string, level int) (*Config, error) {
	var tmp interface{}
	if _, err := toml.DecodeFile(file, &tmp); err != nil {
		return nil, err
	}

	c := new(Config)
	c.ConfFile = file
	c.Version = "0.1.1"
	c.PathLevel = level
	c.Item = make(map[string]interface{})
	c.loadConfig(tmp, []string{})

	return c, nil
}

func (c *Config) ReloadConfig() error {
	var tmp interface{}
	if _, err := toml.DecodeFile(c.ConfFile, &tmp); err != nil {
		return err
	}

	c.Item = make(map[string]interface{})
	c.loadConfig(tmp, []string{})
	return nil
}

func (c *Config) loadConfig(tree interface{}, path []string) {
	if c.PathLevel > 0 && len(path) >= c.PathLevel {
		return
	}
	for key, value := range tree.(map[string]interface{}) {
		fullPath := append(path, key)
		pathKey := strings.Join(fullPath, ".")
		switch orig := value.(type) {
		case map[string]interface{}:
			c.loadConfig(orig, fullPath)
		//case []map[string]interface{}:

		//for i, v := range orig {
		//config[pathKey]

		//}
		default:
			//case string:
			c.Item[pathKey] = orig
			//case int:
			//    config[pathKey] = orig

			// case []map[string]interface{}:
			//     typed := make([]map[string]interface{}, len(orig))
			//     for i, v := range orig {
			//         typed[i] = translate(v).(map[string]interface{})
			//     }
			//     return typed
			// case []interface{}:
			//     typed := make([]interface{}, len(orig))
			//     for i, v := range orig {
			//         typed[i] = translate(v)
			//     }

			//     // We don't really need to tag arrays, but let's be future proof.
			//     // (If TOML ever supports tuples, we'll need this.)
			//     return tag("array", typed)
			// case time.Time:
			// 	return tag("datetime", orig.Format("2006-01-02T15:04:05Z"))
			// case bool:
			// 	return tag("bool", fmt.Sprintf("%v", orig))
			// case int64:
			// 	return tag("integer", fmt.Sprintf("%d", orig))
			// case float64:
			// 	return tag("float", fmt.Sprintf("%v", orig))
			// case string:
			// 	return tag("string", orig)
		}

	}
}

func (c *Config) Int(key string, def int) (int, error) {
	value, ok := c.Item[key]
	if !ok {
		return def, nil
	}
	switch v := value.(type) {
	case int:
		return v, nil
	case int64:
		return int(v), nil
	}
	return def, errors.New("Type Not Match Int")
}

func (c *Config) Int64(key string, def int64) (int64, error) {
	value, ok := c.Item[key]
	if !ok {
		return def, nil
	}
	switch v := value.(type) {
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	}
	return def, errors.New("Type Not Match Int64")
}

func (c *Config) Float64(key string, def float64) (float64, error) {
	value, ok := c.Item[key]
	if !ok {
		return def, nil
	}
	switch v := value.(type) {
	//case float:
	//	return float64(v), nil
	case float64:
		return v, nil
	}
	return def, errors.New("Type Not Match Float64")
}

func (c *Config) String(key string, def string) (string, error) {
	value, ok := c.Item[key]
	if !ok {
		return def, nil
	}

	switch str := value.(type) {
	case string:
		return str, nil
	}
	return def, errors.New("Type Not Match String")
}

func (c *Config) Bool(key string, def bool) (bool, error) {
	value, ok := c.Item[key]
	if !ok {
		return def, nil
	}
	switch v := value.(type) {
	case bool:
		return v, nil
	}
	return def, errors.New("Type Not Match Bool")
}

func (c *Config) Array(key string) ([]interface{}, error) {
	value, ok := c.Item[key]
	if !ok {
		return nil, nil
	}
	switch v := value.(type) {
	case []interface{}:
		return v, nil
	}
	return nil, errors.New("Type Not Match Array")
}

func (c *Config) Map(key string) (map[string]interface{}, error) {
	value, ok := c.Item[key]
	if !ok {
		return nil, nil
	}
	switch v := value.(type) {
	case map[string]interface{}:
		return v, nil
	}
	return nil, errors.New("Type Not Match Map")
}
