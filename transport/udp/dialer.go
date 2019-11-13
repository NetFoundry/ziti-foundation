/*
	Copyright 2019 Netfoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package udp

import (
	"github.com/netfoundry/ziti-foundation/transport"
	"net"
)

// Dial attempts to dial a UDP endpoint and create a connection to it
func Dial(destination, name string) (transport.Connection, error) {
	socket, err := net.Dial("udp", destination)
	if err != nil {
		return nil, err
	}

	return &connection{
		detail: &transport.ConnectionDetail{
			Address: "udp:" + destination,
			InBound: false,
			Name:    name,
		},
		socket:  socket,
		fullBuf: make([]byte, MaxPacketSize),
	}, nil
}
