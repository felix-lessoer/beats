// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package ibmmq

//Config contains all configuration objects
type Config struct {
	QueueManager       string           `config:"bindingQueueManager"`
	RemoteQueueManager []string         `config:"targetQueueManager"`
	QMgrStat           bool             `config:"queueManagerStatus"`
	PubSub             bool             `config:"pubSub"`
	ConnectionConfig   ConnectionConfig `config:"cc"`
}

//ConnectionConfig contains the configuration to connect to MQ
type ConnectionConfig struct {
	ClientMode bool   `config:"clientMode"`
	MqServer   string `config:"mqServer"`
	UserID     string `config:"user"`
	Password   string `config:"password"`
}

var (
	//DefaultConfig contains the default configuration for this module
	DefaultConfig = Config{
		PubSub:             false,
		QMgrStat:           true,
		RemoteQueueManager: []string{""},
		ConnectionConfig: ConnectionConfig{
			ClientMode: false,
			UserID:     "",
			Password:   "",
		},
	}
)