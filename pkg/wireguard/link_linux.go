// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

//go:build linux

package wireguard

import (
	"fmt"

	"github.com/jsimonetti/rtnetlink"
	"golang.org/x/sys/unix"
)

func createWireguardDevice(name string) (string, error) {
	c, err := rtnetlink.Dial(nil)
	if err != nil {
		return "", fmt.Errorf("error connecting to netlink: %w", err)
	}

	defer c.Close() //nolint:errcheck

	err = c.Link.New(&rtnetlink.LinkMessage{
		Family: unix.AF_UNSPEC,
		Type:   unix.ARPHRD_NONE,
		Attributes: &rtnetlink.LinkAttributes{
			Name: name,
			Type: 65534,
			Info: &rtnetlink.LinkInfo{
				Kind: linkKindWireguard,
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("error creating wireguard device: %w", err)
	}

	return name, nil
}

func deleteWireguardDevice(name string) error {
	c, err := rtnetlink.Dial(nil)
	if err != nil {
		return err
	}

	defer c.Close() //nolint:errcheck

	kind, err := c.Link.ListByKind(linkKindWireguard)
	if err != nil {
		return err
	}

	for _, linkMessage := range kind {
		if linkMessage.Attributes.Name == name {
			err = c.Link.Delete(linkMessage.Index)
			if err != nil {
				return err
			}

			break
		}
	}

	return nil
}
