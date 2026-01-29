package util

func Build(message, watermark string) string {
	return "📩 Pesan Anonim\n\n" +
		message +
		"\n\n— " + watermark
}
