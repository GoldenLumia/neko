package desktop

import (
	"goldenlumia/neko/internal/desktop/xevent"
	"goldenlumia/neko/internal/types"
)

func (manager *DesktopManagerCtx) GetCursorChangedChannel() chan uint64 {
	return xevent.CursorChangedChannel
}

func (manager *DesktopManagerCtx) GetClipboardUpdatedChannel() chan struct{} {
	return xevent.ClipboardUpdatedChannel
}

func (manager *DesktopManagerCtx) GetEventErrorChannel() chan types.DesktopErrorMessage {
	return xevent.EventErrorChannel
}
