
// eslint-disable-next-line no-unused-vars
class MirrorMedia {
  constructor(video, canvas) {
    this.id = video.id
    this.video = video
    this.canvas = canvas
    this.stream = null
  }

  async setupMedia() {
    const stream = this.canvas.captureStream()

    this.stream = stream

    this.video.srcObject = stream
    this.video.play()
  }

  getStream() {
    return this.stream
  }

  async onConnected() {}
  async onDisconnected() {}
}
