package discord

/*
	#cgo LDFLAGS: -L../../lib -ldiscord_partner_sdk
	#include "../../lib/cdiscord.h"
	#include <stdlib.h>
*/
import "C"

type DiscordClient struct {
}

func NewDiscordClient() *DiscordClient {
	return &DiscordClient{}
}
