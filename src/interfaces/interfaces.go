package interfaces

import "github.com/voodooEntity/gits/src/transport"

type PluginInterface interface {
	Execute(transport.TransportEntity, string, string) ([]transport.TransportEntity, error)
	GetConfig() transport.TransportEntity
	New() PluginInterface
}
