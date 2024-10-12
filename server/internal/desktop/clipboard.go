package desktop

import "goldenlumia/neko/internal/desktop/clipboard"

func (manager *DesktopManagerCtx) ReadClipboard() string {
	return clipboard.Read()
}

func (manager *DesktopManagerCtx) WriteClipboard(data string) {
	clipboard.Write(data)
}
