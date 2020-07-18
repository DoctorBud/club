// eslint-disable-next-line no-unused-vars
class Peering {
  constructor(elmID, stream, signals) {
    this.videosElm = document.getElementById(elmID)
    this.stream = stream
    this.signals = signals
    this.offer = null
    this.peers = {}
  }

  async onJoin(join) {
    let peer = this.getPeer(join.peerId)

    this.stream.getTracks().forEach((track) => {
      peer.addTrack(track, this.stream)
    })

    const offer = await peer.createOffer({
      offerToReceiveVideo: 1,
      offerToReceiveAudio: 1,
    })

    await peer.setLocalDescription(offer)

    return offer
  }

  async onLeave(leave) {
    let peer = this.removePeer(leave.peerId)
    if (!peer) {
      return
    }

    let videosElm = document.getElementById("videos")
    let videoElm = document.getElementById(leave.peerId)
    videosElm.removeChild(videoElm)

    peer.close()
  }

  async onOffer(offer) {
    let peer = this.getPeer(offer.peerId)
    peer.setRemoteDescription(offer.offer)

    this.stream.getTracks().forEach((track) => {
      peer.addTrack(track, this.stream)
    })

    const answer = await peer.createAnswer()
    peer.setLocalDescription(answer)

    return answer
  }

  async onAnswer(answer) {
    let peer = this.getPeer(answer.peerId)
    await peer.setRemoteDescription(answer.answer)
  }

  async onICECandidate(candidate) {
    let peer = this.getPeer(candidate.peerId)
    peer.addIceCandidate(candidate.candidate)
  }

  newPeer(peerId) {
    const config = {
      iceServers: [{
        urls: "stun:stun1.l.google.com:19302"
      }]
    }

    let peer = new RTCPeerConnection(config)
    let video = document.createElement("video")
    video.id = peerId
    video.autoplay = true
    // video.controls = true
    // video.muted = true // https://developer.mozilla.org/en-US/docs/Web/Media/Autoplay_guide
    // video.poster = 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z/C/HgAGgwJ/lK3Q6wAAAABJRU5ErkJggg=='

    this.videosElm.appendChild(video)

    // setTimeout(() => {
    //   video.play()
    // }, 1000)

    peer.addEventListener('icecandidate', ({ candidate }) => {
      if (candidate) {
        this.signals.sendICECandidate(peerId, candidate)
      }
    })

    peer.addEventListener("track", (track) => {
      video.srcObject = track.streams[0]
    })

    return peer
  }

  getPeer(peerId) {
    if (this.peers[peerId] === undefined) {
      this.peers[peerId] = this.newPeer(peerId)
    }

    return this.peers[peerId]
  }

  removePeer(peerId) {
    if (this.peers[peerId] === undefined) {
      return null
    }

    let peer = this.peers[peerId]
    delete this.peers[peerId]

    return peer
  }

  onConnected() { }
  onDisconnected() { }
}

